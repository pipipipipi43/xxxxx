// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: exception.proto

package pb

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/protobuf/types/descriptorpb"
	_ "google.golang.org/protobuf/types/known/structpb"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *GetExceptionsRequest) Validate() error {
	if this.ScopeID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ScopeID", fmt.Errorf(`value '%v' must not be an empty string`, this.ScopeID))
	}
	return nil
}
func (this *GetExceptionsResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *GetExceptionEventIdsRequest) Validate() error {
	if this.ExceptionID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ExceptionID", fmt.Errorf(`value '%v' must not be an empty string`, this.ExceptionID))
	}
	if this.ScopeID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ScopeID", fmt.Errorf(`value '%v' must not be an empty string`, this.ScopeID))
	}
	return nil
}
func (this *GetExceptionEventIdsResponse) Validate() error {
	return nil
}
func (this *GetExceptionEventRequest) Validate() error {
	if this.ExceptionEventID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ExceptionEventID", fmt.Errorf(`value '%v' must not be an empty string`, this.ExceptionEventID))
	}
	if this.ScopeID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ScopeID", fmt.Errorf(`value '%v' must not be an empty string`, this.ScopeID))
	}
	return nil
}
func (this *GetExceptionEventResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *Exception) Validate() error {
	return nil
}
func (this *Stacks) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *ExceptionEvent) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	// Validation of proto3 map<> fields is unsupported.
	// Validation of proto3 map<> fields is unsupported.
	for _, item := range this.Stacks {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Stacks", err)
			}
		}
	}
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
