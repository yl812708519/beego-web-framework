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
		fieldName := oType.Field(i).Name
		value := field.Interface()
		tValue := reflect.ValueOf(target)
		// 判断是否有这个field
		if  _, ok:= tType.Elem().FieldByName(fieldName); ok {
			// 判断字段类型是否一样
			kind := tValue.Elem().FieldByName(fieldName).Kind()
			if kind != field.Kind() {
				log.Println("convert error，field type not onsistent in"+fieldName+ "convert object:" + oType.String() +"=>>>" + tType.String())
				panic(NewServiceError(10000))
			}
			structValue := reflect.ValueOf(target).Elem()        //结构体属性值
			structFieldValue := structValue.FieldByName(fieldName) //结构体单个属性值
			if !structFieldValue.CanSet() {
				panic(NewServiceError(10000))
				log.Println("can` not set field")
			}
			switch kind {
				case reflect.String:
					structFieldValue.SetString(value.(string))
				case reflect.Bool:
					structFieldValue.SetBool(value.(bool))
				case reflect.Int64:
					structFieldValue.SetInt(value.(int64))
				case reflect.Int:
					structFieldValue.SetInt(int64(value.(int)))
			}
		}
	}

	return target
}


// convert 时对象类型会丢失， 变为*interface{}
// 导致无法获取field， 数组获取或者
// 报错导致前段无法联调， 晚上再调
//func Converts(targetTemp interface{}, objects ...interface{}) interface{} {
//
//
//	if len(objects) <= 0 {
//		panic(NewServiceException(10000))
//	}
//	results := []interface{}{}
//	oType := reflect.TypeOf(objects[0])
//
//	fieldNum := reflect.ValueOf(objects[0]).NumField()
//	// 获取 字段数
//
//	for _, object := range objects {
//		target := targetTemp
//		oValue := reflect.ValueOf(object)
//		// 目标类型
//		tType := reflect.TypeOf(&target)
//		for i:=0; i<fieldNum; i++{
//			field := oValue.Field(i)
//			fieldName := oType.Field(i).Name
//			value := field.Interface()
//			tValue := reflect.ValueOf(&target)
//			// 判断是否有这个field
//			if  _, ok:= tType.Elem().FieldByName(fieldName); ok {
//				// 判断字段类型是否一样
//				kind := tValue.Elem().FieldByName(fieldName).Kind()
//				if kind != field.Kind() {
//					log.Println("convert error，field type not onsistent in"+fieldName+ "convert object:" + oType.String() +"=>>>" + tType.String())
//					panic(NewServiceError(10000))
//				}
//				structValue := reflect.ValueOf(&target).Elem()        //结构体属性值
//				structFieldValue := structValue.FieldByName(fieldName) //结构体单个属性值
//				if !structFieldValue.CanSet() {
//					panic(NewServiceError(10000))
//					log.Println("can` not set field")
//				}
//				switch kind {
//				case reflect.String:
//					structFieldValue.SetString(value.(string))
//				case reflect.Bool:
//					structFieldValue.SetBool(value.(bool))
//				case reflect.Int64:
//					structFieldValue.SetInt(value.(int64))
//				case reflect.Int:
//					structFieldValue.SetInt(int64(value.(int)))
//				}
//			}
//			results = append(results, target)
//		}
//	}
//
//	return results
//}








