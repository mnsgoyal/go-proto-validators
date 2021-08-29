// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: examples/nested.proto

package validator_examples

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/maanasasubrahmanyam-sd/go-proto-validators"
	regexp "regexp"
	github_com_maanasasubrahmanyam-sd_go_proto_validators "github.com/maanasasubrahmanyam-sd/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *InnerMessage) Validate() error {
	if !(this.SomeInteger > 0) {
		return github_com_maanasasubrahmanyam-sd_go_proto_validators.FieldError("SomeInteger", fmt.Errorf(`value '%v' must be greater than '0'`, this.SomeInteger))
	}
	if !(this.SomeInteger < 100) {
		return github_com_maanasasubrahmanyam-sd_go_proto_validators.FieldError("SomeInteger", fmt.Errorf(`value '%v' must be less than '100'`, this.SomeInteger))
	}
	return nil
}

var _regex_OuterMessage_ImportantString = regexp.MustCompile(`^[a-z]{2,5}$`)

func (this *OuterMessage) Validate() error {
	if !_regex_OuterMessage_ImportantString.MatchString(this.ImportantString) {
		return github_com_maanasasubrahmanyam-sd_go_proto_validators.FieldError("ImportantString", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-z]{2,5}$"`, this.ImportantString))
	}
	if nil == this.Inner {
		return github_com_maanasasubrahmanyam-sd_go_proto_validators.FieldError("Inner", fmt.Errorf("message must exist"))
	}
	if this.Inner != nil {
		if err := github_com_maanasasubrahmanyam-sd_go_proto_validators.CallValidatorIfExists(this.Inner); err != nil {
			return github_com_maanasasubrahmanyam-sd_go_proto_validators.FieldError("Inner", err)
		}
	}
	return nil
}
