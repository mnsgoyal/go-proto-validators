package plugin

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gogo/protobuf/gogoproto"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
	validator "github.com/mnsgoyal/go-proto-validators"
)

const alphaPattern = "^[a-zA-Z]+$"
const defaultPattern = "^[a-zA-Z0-9]+$"


type plugin struct {
	*generator.Generator
	generator.PluginImports
	regexPkg      generator.Single
	fmtPkg        generator.Single
	validatorPkg  generator.Single
	useGogoImport bool
}

func NewPlugin(useGogoImport bool) generator.Plugin {
	return &plugin{useGogoImport: useGogoImport}
}

func (p *plugin) Name() string {
	return "validator"
}

func (p *plugin) Init(g *generator.Generator) {
	p.Generator = g
}

func (p *plugin) Generate(file *generator.FileDescriptor) {
	/*if !p.useGogoImport {
		vanity.TurnOffGogoImport(file.FileDescriptorProto)
	}*/
	p.PluginImports = generator.NewPluginImports(p.Generator)
	p.regexPkg = p.NewImport("regexp")
	p.fmtPkg = p.NewImport("fmt")
	p.validatorPkg = p.NewImport("github.com/mnsgoyal/go-proto-validators")

	for _, msg := range file.Messages() {
		if msg.DescriptorProto.GetOptions().GetMapEntry() {
			continue
		}
		p.generateRegexVars(file, msg)
		if gogoproto.IsProto3(file.FileDescriptorProto) {
			p.generateProto3Message(file, msg)
		} /*else {
			p.generateProto2Message(file, msg)
		}*/
	}
}

func getFieldValidatorIfAny(field *descriptor.FieldDescriptorProto) *validator.FieldValidator {
	if field.Options != nil {
		v, err := proto.GetExtension(field.Options, validator.E_Field)
		if err == nil && v.(*validator.FieldValidator) != nil {
			return (v.(*validator.FieldValidator))
		}
	}
	return nil
}


func (p *plugin) generateRegexVars(file *generator.FileDescriptor, message *generator.Descriptor) {
	ccTypeName := generator.CamelCaseSlice(message.TypeName())
	for _, field := range message.Field {
		validator := getFieldValidatorIfAny(field)
		if validator != nil {
			fieldName := p.GetOneOfFieldName(message, field)
			if validator.Alpha == nil {
				fmt.Fprintf(os.Stderr, "WARNING: regex and uuid validator is set for field %v.%v, is null.", ccTypeName, fieldName)
			} else if validator.Alpha != nil && *validator.Alpha {
				p.P(`var `, p.regexName(ccTypeName, fieldName), ` = `, p.regexPkg.Use(), `.MustCompile(`, "`", alphaPattern, "`", `)`)
			}else{
				p.P(`var `, p.regexName(ccTypeName, fieldName), ` = `, p.regexPkg.Use(), `.MustCompile(`, "`", defaultPattern, "`", `)`)
			}
		}
	}
}

func (p *plugin) GetFieldName(message *generator.Descriptor, field *descriptor.FieldDescriptorProto) string {
	fieldName := p.Generator.GetFieldName(message, field)
	if p.useGogoImport {
		return fieldName
	}
	if gogoproto.IsEmbed(field) {
		fieldName = generator.CamelCase(*field.Name)
	}
	return fieldName
}

func (p *plugin) generateProto3Message(file *generator.FileDescriptor, message *generator.Descriptor) {
	ccTypeName := generator.CamelCaseSlice(message.TypeName())
	p.P(`func (this *`, ccTypeName, `) Validate() error {`)
	p.In()

	for _, field := range message.Field {
		fieldValidator := getFieldValidatorIfAny(field)
		if fieldValidator == nil && !field.IsMessage() {
			continue
		}
		isOneOf := field.OneofIndex != nil
		fieldName := p.GetOneOfFieldName(message, field)
		variableName := "this." + fieldName
		//repeated := field.IsRepeated()
		// Golang's proto3 has no concept of unset primitive fields
		//nullable := (gogoproto.IsNullable(field) || !gogoproto.ImportsGoGoProto(file.FileDescriptorProto)) && field.IsMessage() && !(p.useGogoImport && gogoproto.IsEmbed(field))
		if p.fieldIsProto3Map(file, message, field) {
			p.P(`// Validation of proto3 map<> fields is unsupported.`)
			continue
		}
		if isOneOf {
			p.In()
			oneOfName := p.GetFieldName(message, field)
			oneOfType := p.OneOfTypeName(message, field)
			// if x, ok := m.GetType().(*OneOfMessage3_OneInt); ok {
			p.P(`if oneOfNester, ok := this.Get` + oneOfName + `().(* ` + oneOfType + `); ok {`)
			variableName = "oneOfNester." + p.GetOneOfFieldName(message, field)
		}
		if field.IsString() {
			p.generateAlphaValidator(variableName, ccTypeName, fieldName, fieldValidator)
		}
	}
	p.P(`return nil`)
	p.Out()
	p.P(`}`)
}


func (p *plugin) generateAlphaValidator(variableName string, ccTypeName string, fieldName string, fv *validator.FieldValidator) {
	if fv.Alpha != nil  {
		p.P(`if !`, p.regexName(ccTypeName, fieldName), `.MatchString(`, variableName, `) {`)
		p.In()
		errorStr := "be a string conforming to default regex " + strconv.Quote(defaultPattern)
		if *fv.Alpha {
			errorStr = "be a string conforming to alpha regex " + strconv.Quote(alphaPattern)
		}
		p.P(`return `, p.validatorPkg.Use(), `.FieldError("`, fieldName, `",`, p.fmtPkg.Use(), ".Errorf(`", errorStr, "`))")
		p.Out()
		p.P(`}`)
	}
}


func (p *plugin) fieldIsProto3Map(file *generator.FileDescriptor, message *generator.Descriptor, field *descriptor.FieldDescriptorProto) bool {
	if field.GetType() != descriptor.FieldDescriptorProto_TYPE_MESSAGE || !field.IsRepeated() {
		return false
	}
	typeName := field.GetTypeName()
	var msg *descriptor.DescriptorProto
	if strings.HasPrefix(typeName, ".") {
		// Fully qualified case, look up in global map, must work or fail badly.
		msg = p.ObjectNamed(field.GetTypeName()).(*generator.Descriptor).DescriptorProto
	} else {
		// Nested, relative case.
		msg = file.GetNestedMessage(message.DescriptorProto, field.GetTypeName())
	}
	return msg.GetOptions().GetMapEntry()
}

func (p *plugin) regexName(ccTypeName string, fieldName string) string {
	return "_regex_" + ccTypeName + "_" + fieldName
}
