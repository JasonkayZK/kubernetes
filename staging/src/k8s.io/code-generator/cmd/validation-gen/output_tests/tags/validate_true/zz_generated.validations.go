//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by validation-gen. DO NOT EDIT.

package validatetrue

import (
	context "context"
	fmt "fmt"

	operation "k8s.io/apimachinery/pkg/api/operation"
	safe "k8s.io/apimachinery/pkg/api/safe"
	validate "k8s.io/apimachinery/pkg/api/validate"
	field "k8s.io/apimachinery/pkg/util/validation/field"
	testscheme "k8s.io/code-generator/cmd/validation-gen/testscheme"
)

func init() { localSchemeBuilder.Register(RegisterValidations) }

// RegisterValidations adds validation functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterValidations(scheme *testscheme.Scheme) error {
	scheme.AddValidationFunc((*Struct)(nil), func(ctx context.Context, op operation.Operation, obj, oldObj interface{}, subresources ...string) field.ErrorList {
		if len(subresources) == 0 {
			return Validate_Struct(ctx, op, nil /* fldPath */, obj.(*Struct), safe.Cast[*Struct](oldObj))
		}
		return field.ErrorList{field.InternalError(nil, fmt.Errorf("no validation found for %T, subresources: %v", obj, subresources))}
	})
	return nil
}

func Validate_Struct(ctx context.Context, op operation.Operation, fldPath *field.Path, obj, oldObj *Struct) (errs field.ErrorList) {
	// field Struct.TypeMeta has no validation

	// field Struct.StringField
	errs = append(errs,
		func(fldPath *field.Path, obj, oldObj *string) (errs field.ErrorList) {
			errs = append(errs, validate.FixedResult(ctx, op, fldPath, obj, oldObj, true, "field Struct.StringField")...)
			return
		}(fldPath.Child("stringField"), &obj.StringField, safe.Field(oldObj, func(oldObj *Struct) *string { return &oldObj.StringField }))...)

	return errs
}
