package reflectgostructv1

import (
	"log"
	"reflect"
	"strings"
)

func PsqlTagToSql(s interface{}) string {
	r := reflect.ValueOf(s)
	numfield := r.Elem().NumField()
	if r.Kind() != reflect.Ptr {
		log.Fatal("Wrong type struct")
	}
	var result string
	for i := 0; i < numfield; i++ {
		result += reflect.TypeOf(s).Elem().Field(i).Tag.Get("psql")
		if i+1 < numfield {
			result += ", \n"
		}
	}
	return result
}

func GetPsqlTagsAndValues(s interface{}) (string, string) {
	r := reflect.ValueOf(s)
	numfield := r.Elem().NumField()
	if r.Kind() != reflect.Ptr {
		log.Fatal("Wrong type struct")
	}
	var fields = make(map[string]string)
	var i int
	for i = 0; i < numfield; i++ {
		if !r.Elem().Field(i).IsZero() {
			tag := reflect.TypeOf(s).Elem().Field(i).Tag.Get("psql")
			fields[strings.Split(tag, " ")[0]] = r.Elem().Field(i).Interface().(string)
		}
	}
	var tags string
	var values string
	i = 0
	for k, v := range fields {
		tags += k
		values += "'"+v+"'"
		if i+1 < len(fields) {
			tags += ", "
			values += ", "
		}
		i++
	}
	return tags, values
}