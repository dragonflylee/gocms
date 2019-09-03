package util

import (
	"encoding"
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

// ParseForm will parse form values to struct via tag.
// Support for anonymous struct.
func parseFormToStruct(form url.Values, objT reflect.Type, objV reflect.Value) error {
	for i := 0; i < objT.NumField(); i++ {
		fieldV := objV.Field(i)
		if !fieldV.CanSet() {
			continue
		}

		fieldT := objT.Field(i)
		if fieldT.Anonymous && fieldT.Type.Kind() == reflect.Struct {
			err := parseFormToStruct(form, fieldT.Type, fieldV)
			if err != nil {
				return err
			}
			continue
		}

		tags := strings.Split(fieldT.Tag.Get("form"), ",")
		var tag string
		if len(tags) == 0 || len(tags[0]) == 0 {
			tag = strings.ToLower(fieldT.Name)
		} else if tags[0] == "-" {
			continue
		} else {
			tag = tags[0]
		}
		value := form.Get(tag)
		if len(value) == 0 {
			continue
		}

		if fieldV.Kind() == reflect.Ptr {
			if fieldV.IsNil() {
				fieldV.Set(reflect.New(fieldT.Type.Elem()))
			}
			fieldV = fieldV.Elem()
		}
		switch fieldV.Kind() {
		case reflect.Bool:
			b, err := strconv.ParseBool(value)
			if err != nil {
				return fmt.Errorf("parse `%s` failed: %v", tag, err)
			}
			fieldV.SetBool(b)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			x, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return fmt.Errorf("parse `%s` failed: %v", tag, err)
			}
			fieldV.SetInt(x)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			x, err := strconv.ParseUint(value, 10, 64)
			if err != nil {
				return fmt.Errorf("parse `%s` failed: %v", tag, err)
			}
			fieldV.SetUint(x)
		case reflect.Float32, reflect.Float64:
			x, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return fmt.Errorf("parse `%s` failed: %v", tag, err)
			}
			fieldV.SetFloat(x)
		case reflect.Interface:
			fieldV.Set(reflect.ValueOf(value))
		case reflect.String:
			fieldV.SetString(value)
		case reflect.Struct:
			if v, ok := fieldV.Addr().Interface().(encoding.TextUnmarshaler); ok {
				if err := v.UnmarshalText([]byte(value)); err != nil {
					return fmt.Errorf("parse `%s` failed: %v", tag, err)
				}
			} else {
				return fmt.Errorf("parse `%s` failed `%s`", tag, fieldV.Type())
			}
		case reflect.Slice, reflect.Array:
			v, exist := form[tag]
			if !exist || len(v) == 0 {
				continue
			}
			el := fieldT.Type.Elem()
			switch el.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				fieldV.Set(reflect.MakeSlice(reflect.SliceOf(el), len(v), len(v)))
				for i := 0; i < len(v); i++ {
					x, err := strconv.ParseInt(v[i], 10, 64)
					if err != nil {
						return fmt.Errorf("parse `%s` failed: %v", tag, err)
					}
					fieldV.Index(i).SetInt(x)
				}
			case reflect.String:
				fieldV.Set(reflect.ValueOf(v))
			}
		}
	}
	return nil
}

// ParseForm will parse form values to struct via tag.
func ParseForm(form url.Values, obj interface{}) error {
	objT := reflect.TypeOf(obj)
	objV := reflect.ValueOf(obj)
	if objT.Kind() != reflect.Ptr || objT.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("%v must be struct pointer", obj)
	}
	objT = objT.Elem()
	objV = objV.Elem()
	return parseFormToStruct(form, objT, objV)
}
