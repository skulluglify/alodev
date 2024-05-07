package tests

import (
	"database/sql"
	"dvwk/models"
	"github.com/google/uuid"
	"reflect"
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
		t.Logf("Field: %v, Tag: %v\n", field, tag)
		parts := ParseTags(tag, "tags")
		t.Logf("Parts: %v\n", parts)
	}
}

func TestParseBodyWithStrfmtReflect(t *testing.T) {
	body := &models.Categories{
		Model: &models.Model{
			Id:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdateAt:  time.Now().UTC(),
			DeletedAt: sql.Null[time.Time]{
				Valid: false,
				V:     time.Time{},
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

	if model, ok := typ.FieldByName("Model"); ok {
		t.Logf("Model: %v\n", model)
		modelType := BypassPointerTypeReflect(model.Type)

		// Logging
		t.Logf("Model Name: %s\n", modelType.Name())
		t.Logf("Model Type: %v\n", modelType)
	}
}
