// Package structutil provides functions to easily work woth structs

package structutil

import (
	"reflect"
)

// StructToMap : returns a map of names to values for struct passed .
func StructToMap(v interface{}) *map[string]interface{} {

	//modelRefType := v.Type()
	structmap := make(map[string]interface{})

	elements := reflect.ValueOf(v).Elem() //element values
	typeofT := elements.Type()            // type of Element - name of element
	for i := 0; i < elements.NumField(); i++ {
		f := elements.Field(i)
		elementName := typeofT.Field(i).Name
		//typeOfStructElement := f.Type()
		valueOfElement := f.Interface()
		structmap[elementName] = valueOfElement
	}
	return &structmap
}

//CopyStruct : Copies src struct fields to dest struct where both struct have elements with same names but different types
func CopyStruct(src interface{}, dest interface{}) {

	srcFields := reflect.ValueOf(src).Elem()
	newStructValue := reflect.ValueOf(dest)

	for count := 0; count < srcFields.NumField(); count++ {
		srcField := srcFields.Field(count)
		fieldname := srcFields.Type().Field(count).Name
		newField := newStructValue.Elem().FieldByName(fieldname)
		newField.SetString(fmt.Sprintf("%v", srcField.Interface()))
	}
}
