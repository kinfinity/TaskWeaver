package utils

import (
    "fmt"
    "reflect"

)

func PrettyPrintStruct(s interface{}) {
    v := reflect.ValueOf(s)
    t := v.Type()

    fmt.Println(t.Name() + " {")
    for i := 0; i < v.NumField(); i++ {
        field := v.Field(i)
        fieldName := t.Field(i).Name
        fieldValue := field.Interface()
        fmt.Printf("    %s: %v\n", fieldName, fieldValue)
        // fmt.Printf("%s\n", reflect.ValueOf(fieldValue).Type())
    }
    fmt.Println("}")

}