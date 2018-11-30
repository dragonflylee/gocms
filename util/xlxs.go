package util

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"

	"github.com/tealeg/xlsx"
)

var (
	xlsxType = map[reflect.Type]func(v interface{}) string{
		reflect.TypeOf(time.Time{}): func(v interface{}) string {
			return v.(time.Time).Format(formatDateTime)
		},
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
			v := f.Type
			if v.Kind() == reflect.Ptr {
				v = v.Elem()
			}
			if _, exist = xlsxType[v]; exist {
				row.AddCell().SetString(f.Name)
			} else if v.Kind() == reflect.Struct {
				xlsxField(row, v)
			} else {
				row.AddCell().SetString(f.Name)
			}
		} else if name, _ := parseXlsxTag(tags); "-" != name {
			row.AddCell().SetString(name)
		}
	}
}

func xlsxCell(row *xlsx.Row, objT reflect.Type, objV reflect.Value) {
	for i := 0; i < objT.NumField(); i++ {
		fieldV := objV.Field(i)
		fieldT := objT.Field(i)
		if fieldT.Anonymous && fieldT.Type.Kind() == reflect.Struct {
			xlsxCell(row, fieldT.Type, fieldV)
			continue
		}
		name, opts := parseXlsxTag(fieldT.Tag.Get("xlsx"))
		if name == "-" {
			continue
		}
		cell := row.AddCell()
		if !fieldV.IsValid() {
			cell.SetString("-")
			continue
		}
		fieldV = reflect.Indirect(fieldV)
		if enum, exist := opts["enum"]; exist {
			list := strings.Split(enum, ",")
			if index := int(fieldV.Int()); index < len(list) {
				cell.SetString(list[index])
			}
		} else if !fieldV.CanInterface() {
			log.Printf("tag `%s` bad field `%s` valid `%v`", name, fieldV.Type(), fieldV.IsValid())
		} else if m, exist := xlsxType[fieldV.Type()]; exist {
			cell.SetString(m(fieldV.Interface()))
		} else if s, ok := fieldV.Interface().(fmt.Stringer); ok {
			cell.SetString(s.String())
		} else {
			cell.SetValue(fieldV.Interface())
		}
	}
}

// Excel 导出Excel
func Excel(w http.ResponseWriter, data map[string]interface{}, format string, a ...interface{}) error {
	file := xlsx.NewFile()

	for name, obj := range data {
		sheet, err := file.AddSheet(name)
		if err != nil {
			return err
		}
		v := reflect.Indirect(reflect.ValueOf(obj))
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
			xlsxCell(row, t, reflect.Indirect(v.Index(i)))
		}
	}
	w.Header().Set("Content-Type", "application/vnd.ms-excel")
	w.Header().Set("Cache-Control", "must-revalidate, post-check=0, pre-check=0")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	w.Header().Set("Content-Disposition", "attachment; filename="+url.PathEscape(fmt.Sprintf(format, a...)))
	return file.Write(w)
}
