package tests

import (
	"database/sql"
	"dvwk/models"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"reflect"
	"strings"
	"testing"
	"time"
)

type StrFmtTags struct {
	Limit int `tags:"name:limit;type:integer"`
}

func NewStrFmtTags(limit int) *StrFmtTags {
	return &StrFmtTags{
		Limit: limit,
	}
}

func ParseTags(tag reflect.StructTag, name string) map[string]string {
	tags := tag.Get(name)
	parts := make(map[string]string)
	var key, value string
	for i := 0; i < len(tags); i++ {
		if tags[i] == ';' {
			parts[key] = value
			key, value = "", ""
			continue
		}
		if tags[i] == ':' {
			key = value
			value = ""
			continue
		}
		value += string(tags[i])
	}
	parts[key] = value
	return parts
}

func BypassPointerTypeReflect(p reflect.Type) reflect.Type {
	if p.Kind() == reflect.Ptr {
		return BypassPointerTypeReflect(p.Elem())
	}
	return p
}

func BypassPointerValueReflect(p reflect.Value) reflect.Value {
	if p.Kind() == reflect.Ptr {
		return BypassPointerValueReflect(p.Elem())
	}
	return p
}

func TestStrfmtReflect(t *testing.T) {
	t.Log("Test: String Formatter Reflection")

	strFmtTags := NewStrFmtTags(51)

	val := reflect.ValueOf(strFmtTags)
	typ := BypassPointerTypeReflect(reflect.TypeOf(strFmtTags))

	// Logging
	t.Logf("Value: %v, Type: %v\n", val, typ)

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		tag := field.Tag
		t.Logf("CatchField: %v, Tag: %v\n", field, tag)
		parts := ParseTags(tag, "tags")
		t.Logf("Parts: %v\n", parts)
	}
}

func ValueContainPolymorphicOfBaseTypeReflect(val reflect.Value, baseType reflect.Type) bool {
	val = BypassPointerValueReflect(val)
	baseType = BypassPointerTypeReflect(baseType)
	if !val.IsValid() {
		return false
	}
	if val.Kind() == reflect.Struct {
		for i := 0; i < val.NumField(); i++ {
			valField := val.Field(i)
			valFieldType := BypassPointerTypeReflect(valField.Type())
			//fmt.Printf("CatchField Name: %s, Base Name: %s\n", valFieldType.Name(), baseType.Name())
			if valFieldType.Kind() == baseType.Kind() && valFieldType.Name() == baseType.Name() {
				return true
			} else if ValueContainPolymorphicOfBaseTypeReflect(valField, baseType) {
				return true
			}
		}
	}
	return false
}

func ValueContainPolymorphicOfBaseType(val any, baseType any) bool {
	return ValueContainPolymorphicOfBaseTypeReflect(reflect.ValueOf(val), reflect.TypeOf(baseType))
}

type CatchField struct {
	structField *reflect.StructField
	value       *reflect.Value
	valid       bool
}

func NewCatchField(field *reflect.StructField, value *reflect.Value) *CatchField {
	valid := false
	if field != nil && value != nil {
		if value.IsValid() {
			valid = true
		}
	}
	return &CatchField{
		structField: field,
		value:       value,
		valid:       valid,
	}
}

func (catchField *CatchField) Tag(key string) string {
	return catchField.structField.Tag.Get(key)
}

func (catchField *CatchField) Name() string {
	return catchField.structField.Name
}

func (catchField *CatchField) Type() reflect.Type {
	return catchField.structField.Type
}

func (catchField *CatchField) Value() reflect.Value {
	return *catchField.value
}

func (catchField *CatchField) Valid() bool {
	return catchField.valid
}

func DeepFieldSearchByNameValueReflect(target reflect.Value, name string) (*CatchField, error) {
	val := BypassPointerValueReflect(target)
	typ := BypassPointerTypeReflect(val.Type())
	if !val.IsValid() {
		return NewCatchField(nil, nil), errors.New("invalid value")
	}
	if val.Kind() != reflect.Struct {
		return NewCatchField(nil, nil), errors.New("value is not struct")
	}
	baseName := name
	pointing := strings.Index(name, ".")
	if 0 <= pointing {
		baseName = name[:pointing]
		truncIdx := pointing + 1
		if truncIdx < len(name) {
			name = name[truncIdx:]
		} else {
			name = ""
		}
	}
	baseEnd := len(baseName) - 1
	lastChar := baseName[baseEnd:]
	nullable := lastChar == "?" // nullish coalescing operator
	throwing := lastChar == "!" // panic error message
	if nullable || throwing {
		baseName = baseName[:baseEnd]
	}
	for i := 0; i < val.NumField(); i++ {
		valField := val.Field(i)
		typField := typ.Field(i)
		if strings.EqualFold(typField.Name, baseName) {
			if name != baseName+lastChar && name != "" {
				return DeepFieldSearchByNameValueReflect(valField, name)
			}
			return NewCatchField(&typField, &valField), nil
		}
	}
	if nullable {
		return NewCatchField(nil, nil), nil
	}
	return NewCatchField(nil, nil), errors.New(fmt.Sprintf("name %s not found", name))
}

func DeepFieldSearchByName(target any, name string) (any, error) {
	catchField, err := DeepFieldSearchByNameValueReflect(reflect.ValueOf(target), name)
	if catchField.Valid() {
		return catchField.Value().Interface(), err
	}
	return nil, errors.New("invalid value")
}

func TestParseBodyWithStrfmtReflect(t *testing.T) {
	body := &models.Categories{
		Model: &models.Model{
			Id:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdateAt:  time.Now().UTC(),
			DeletedAt: sql.Null[time.Time]{
				V:     time.Time{},
				Valid: false,
			},
		},
		Name:   "Business",
		Kind:   "Article",
		UserId: uuid.New(),
	}

	t.Logf("Body: %v\n", body)

	val := reflect.ValueOf(body)
	typ := BypassPointerTypeReflect(reflect.TypeOf(body))

	// Logging
	t.Logf("Value: %v, Type: %v\n", val, typ)

	polymorphic := ValueContainPolymorphicOfBaseType(body, models.Model{})
	t.Logf("Polymorphic: %v\n", polymorphic)

	fmt.Println(DeepFieldSearchByName(body, "model?.id!"))

	catchField, _ := DeepFieldSearchByNameValueReflect(val, "model?.id!")
	println(catchField)
	strfmt := catchField.Tag("strfmt")
	fmt.Println(strfmt)
}
