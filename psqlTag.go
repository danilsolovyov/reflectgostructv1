package reflectgostructv1

import (
	"fmt"
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
		tag := reflect.TypeOf(s).Elem().Field(i).Tag.Get("psql")
		if !strings.Contains(tag, "IDENTITY") {
			fields[strings.Split(tag, " ")[0]] = fmt.Sprint(r.Elem().Field(i).Interface())
		}
	}
		
	var tags string
	var values string
	i = 1
	for k, v := range fields {
		tags += k
		values += "'"+v+"'"
		if i < len(fields) {
			tags += ", "
			values += ", "
		}
		i++
	}
	return tags, values
}

func GetPsqlTagsNames (s interface{}) string {
	r := reflect.ValueOf(s)
	numfield := r.Elem().NumField()
	if r.Kind() != reflect.Ptr {
		log.Fatal("Wrong type struct")
	}
	var i int
	var tags string
	for i = 0; i < numfield; i++ {
		tag := reflect.TypeOf(s).Elem().Field(i).Tag.Get("psql")
		tags += strings.Split(tag, " ")[0]
		if (i+1 < numfield) {
			tags += ", "
		}
	}
	return tags
}