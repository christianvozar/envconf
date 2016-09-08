// MIT Licensed.
// Christian R. Vozar <christian@rogueethic.com>
// Fabriqué en Nouvelle Orléans ⚜

// Inspired heavily by the work of Kelsey Hightower's envconfig
// https://github.com/kelseyhightower/envconfig

// Package envconf implements utility funtions for parsing envronment variables
// utilized for application settings.
// Factor 3 in the 12Factor application. http://12factor.net/config
package envconf

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// ErrInvalidSpecification indicates that a specification is of the wrong type.
var ErrInvalidSpecification = errors.New("invalid specification must be a struct")

// A ParseError occurs when an environment variable cannot be converted to
// the type required by a struct field during assignment.
type ParseError struct {
	KeyName   string
	FieldName string
	TypeName  string
	Value     string
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("envconfig.Process: assigning %[1]s to %[2]s: converting '%[3]s' to type %[4]s", e.KeyName, e.FieldName, e.Value, e.TypeName)
}

// Parse populates a struct from environment variables and returns the number of
// successfully parsed variables.
func Parse(envPrefix string, spec interface{}) (parseCount int, err error) {
	s := reflect.ValueOf(spec).Elem()
	if s.Kind() != reflect.Struct {
		return 0, ErrInvalidSpecification
	}
	typeOfSpec := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		if f.CanSet() {
			alt := typeOfSpec.Field(i).Tag.Get("vudou")
			fieldName := typeOfSpec.Field(i).Name
			if alt != "" {
				fieldName = alt
			}
			key := strings.ToUpper(fmt.Sprintf("%s_%s", envPrefix, fieldName))
			value := os.Getenv(key)
			if value == "" && alt != "" {
				key := strings.ToUpper(fieldName)
				value = os.Getenv(key)
			}

			def := typeOfSpec.Field(i).Tag.Get("default")
			if def != "" && value == "" {
				value = def
			}

			req := typeOfSpec.Field(i).Tag.Get("required")
			if value == "" {
				if req == "true" {
					return 0, fmt.Errorf("required key %s missing value", key)
				}
				continue
			}

			switch f.Kind() {
			case reflect.Slice:
				s := strings.Split(value, ",")
				f.Set(reflect.AppendSlice(f, reflect.ValueOf(s)))
			case reflect.String:
				f.SetString(value)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				intValue, err := strconv.ParseInt(value, 0, f.Type().Bits())
				if err != nil {
					return 0, &ParseError{
						KeyName:   key,
						FieldName: fieldName,
						TypeName:  f.Type().String(),
						Value:     value,
					}
				}
				f.SetInt(intValue)
			case reflect.Bool:
				boolValue, err := strconv.ParseBool(value)
				if err != nil {
					return 0, &ParseError{
						KeyName:   key,
						FieldName: fieldName,
						TypeName:  f.Type().String(),
						Value:     value,
					}
				}
				f.SetBool(boolValue)
			case reflect.Float32, reflect.Float64:
				floatValue, err := strconv.ParseFloat(value, f.Type().Bits())
				if err != nil {
					return 0, &ParseError{
						KeyName:   key,
						FieldName: fieldName,
						TypeName:  f.Type().String(),
						Value:     value,
					}
				}
				f.SetFloat(floatValue)
			}
		}
	}
	return parseCount, nil
}

func MustParse(envPrefix string, spec interface{}) {
	if _, err := Parse(envPrefix, spec); err != nil {
		panic(err)
	}
}
