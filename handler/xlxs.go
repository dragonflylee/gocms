package handler

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/tealeg/xlsx"
)

var (
	typeField = map[reflect.Type]bool{
		reflect.TypeOf(time.Time{}): true,
	}
)

type xlsxTag map[string]string

func parseXlsxTag(tags string) (name string, opts xlsxTag) {
	opts = make(xlsxTag)
	for _, value := range strings.Split(tags, ";") {
		v := strings.Split(value, ":")
		k := strings.TrimSpace(strings.ToLower(v[0]))
		if len(v) >= 2 {
			opts[k] = strings.Join(v[1:], ":")
		} else {
			name = k
		}
	}
	return name, opts
}

func xlsxField(row *xlsx.Row, t reflect.Type) {
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if tags, exist := f.Tag.Lookup("xlsx"); !exist {
			if _, exist = typeField[f.Type]; exist {
				row.AddCell().SetString(f.Name)
			} else if f.Type.Kind() == reflect.Struct {
				xlsxField(row, f.Type)
			} else if f.Type.Kind() == reflect.Ptr {
				xlsxField(row, f.Type.Elem())
			} else {
				row.AddCell().SetString(f.Name)
			}
		} else if name, _ := parseXlsxTag(tags); "-" != name {
			row.AddCell().SetString(name)
		}
	}
}

func xlsxCeil(row *xlsx.Row, v reflect.Value) {
	t := v.Type()
	for j := 0; j < v.NumField(); j++ {
		f := reflect.Indirect(v.Field(j))
		if tags, exist := t.Field(j).Tag.Lookup("xlsx"); !exist {
			if !f.IsValid() {
				row.AddCell().SetString("-")
			} else if _, exist := typeField[f.Type()]; exist {
				row.AddCell().SetValue(f.Interface())
			} else if f.Kind() == reflect.Struct {
				xlsxCeil(row, f)
			} else if f.CanInterface() {
				row.AddCell().SetValue(f.Interface())
			} else {
				log.Printf("bad field `%s` valid `%v`", f.Type(), f.IsValid())
			}
		} else if name, opts := parseXlsxTag(tags); "-" != name {
			if !f.IsValid() {
				row.AddCell().SetString("-")
			} else {
				if _, exist := typeField[f.Type()]; !exist && f.Kind() == reflect.Struct {
					if field, exist := opts["field"]; !exist {
						f = f.Field(0)
					} else if t, ok := f.Type().FieldByName(field); ok {
						f = f.FieldByIndex(t.Index)
					} else {
						f = f.Field(0)
					}
				}
				if enum, exist := opts["enum"]; exist {
					ceil := row.AddCell()
					list := strings.Split(enum, ",")
					if index := int(f.Int()); index < len(list) {
						ceil.SetString(list[index])
					}
				} else if f.CanInterface() {
					row.AddCell().SetValue(f.Interface())
				} else {
					log.Printf("tag `%s` bad field `%s` valid `%v`", name, f.Type(), f.IsValid())
				}
			}
		}
	}
}

// exportExcel 导出Excel
func exportExcel(data map[string]interface{}, w http.ResponseWriter) error {
	filename := fmt.Sprintf("%s-%d.xlsx", time.Now().Format(dateFormate), time.Now().Unix())
	file := xlsx.NewFile()

	for name, page := range data {
		sheet, err := file.AddSheet(name)
		if err != nil {
			return err
		}
		v := reflect.Indirect(reflect.ValueOf(page))
		if v.Kind() != reflect.Slice {
			return fmt.Errorf("expect slice but type `%s` found", v.Type())
		}
		if v.Len() <= 0 {
			return fmt.Errorf("slice `%s` is empty", v.Type())
		}
		t := reflect.Indirect(v.Index(0)).Type()
		if t.Kind() != reflect.Struct {
			return fmt.Errorf("expect struct but type `%s` found", v.Type())
		}
		// 填写表头
		row := sheet.AddRow()
		xlsxField(row, t)
		// 填写数据
		for i := 0; i < v.Len(); i++ {
			row = sheet.AddRow()
			xlsxCeil(row, reflect.Indirect(v.Index(i)))
		}
	}
	w.Header().Set("Content-Type", "application/vnd.ms-excel")
	w.Header().Set("Cache-Control", "must-revalidate, post-check=0, pre-check=0")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	return file.Write(w)
}
