package common

import (
	"reflect"
	"log"
)

func Convert(o interface{}, target interface{}) interface{} {

	// 获取 字段数
	oType := reflect.TypeOf(o)
	oValue := reflect.ValueOf(o)
	fieldNum := oValue.NumField()

	// 目标类型
	tType := reflect.TypeOf(target)

	for i:=0; i<fieldNum; i++{
		field := oValue.Field(i)
		value := field.Interface()
		tValue := reflect.ValueOf(target)
		fieldName := oType.Field(i).Name
		// 判断是否有这个field
		if  _, ok:= tType.Elem().FieldByName(fieldName); ok {
			// 判断字段类型是否一样
			kind := tValue.Elem().FieldByName(fieldName).Kind()
			if kind != field.Kind() {
				panic(NewServiceError(10001))
				log.Println("convert error，field type not onsistent " + oType.String() +"=>>>" + tType.String() +"."+ fieldName)
			}
			structValue := reflect.ValueOf(target).Elem()        //结构体属性值
			structFieldValue := structValue.FieldByName(fieldName) //结构体单个属性值
			if !structFieldValue.CanSet() {
				panic(NewServiceError(10001))
				log.Println("can` not set field")
			}
			switch kind {
				case reflect.String:
					structFieldValue.SetString(value.(string))
				case reflect.Bool:
					structFieldValue.SetBool(value.(bool))
				case reflect.Int64:
				case reflect.Int:
					structFieldValue.SetInt(int64(value.(int)))
			}
		}
	}

	return target
}








