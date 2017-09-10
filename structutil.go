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
