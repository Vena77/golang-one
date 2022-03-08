package main

import (
	"errors"
	"log"
	"reflect"
	"strconv"
)

type In struct {
	ID     int
	Name   string
	Chislo int
}

var mapIn = map[string]interface{}{
	"ID":     4,
	"Name":   "Vena",
	"Chislo": 77,
}

func SetData(s interface{}) error {
	//t1 := reflect.TypeOf(s)
	//t2 := reflect.Inderect(reflect.ValueOf(s))
	t2 := reflect.ValueOf(s)
	if t2.Kind() == reflect.Ptr {
		t2 = t2.Elem()
	}

	if t2.Kind() != reflect.Struct {
		return errors.New("it's not struct")
	}
	for i := 0; i < t2.NumField(); i++ {
		typeField := t2.Type().Field(i)

		if typeField.Type.Kind() == reflect.Struct {
			log.Printf("nested field: %v", typeField.Name)
			continue
		}

		fldv := mapIn[typeField.Name]
		ff := t2.FieldByName(typeField.Name)
		if fldvs, ok := fldv.(string); ok && ff.Type().AssignableTo(reflect.TypeOf(int(0))) {
			fInt, _ := strconv.Atoi(fldvs)
			ff.Set(reflect.ValueOf(fInt))
		} else {
			ff.Set(reflect.ValueOf(fldv))
		}
		/*	switch typeField.Type.Kind() {
			case reflect.Int:
				ff.SetInt(2)
			case reflect.String:
				ff.SetString("Slovo")
			case reflect.Int32:
				fmt.Println("int32")
			default:
				fmt.Println("неизвестно")
			}*/
		log.Printf("\tname=%s, type=%s, value=%v, tag=`%s`\n",
			typeField.Name,
			typeField.Type,
			t2.Field(i),
			typeField.Tag,
		)
	}
	return nil
}

func main() {
	resp := In{}
	SetData(&resp)
}
