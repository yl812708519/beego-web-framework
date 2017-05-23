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
				panic(NewServiceError(10001))
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
					structFieldValue.SetInt(value.(int64))
				case reflect.Int:
					structFieldValue.SetInt(int64(value.(int)))
			}
		}
	}

	return target
}

// error
//func Converts(objects interface{}, target interface{}) []interface{} {
//
//	result := []interface{}{}
//	tChan := make(chan interface{})
//	innerConverts(objects, target, tChan)
//	for {
//		if r, ok := <-tChan; ok {
//			result = append(result, r)
//		} else {
//			break
//		}
//	}
//	return result
//}
//
//func innerConverts(objects interface{}, target interface{}, channel chan interface{}){
//
//	switch reflect.TypeOf(objects).Kind() {
//		case reflect.Slice:
//			oValue := reflect.ValueOf(objects)
//			for i:=0; i<oValue.Len(); i++ {
//				tType := reflect.TypeOf(target)
//				tc := reflect.New(tType)
//				tc = Convert(oValue.Index(i).Interface(), tc)
//				channel <- tc
//
//			}
//			close(channel)
//	default:
//		panic(NewServiceError(20003))
//		log.Println("error type, need slice")
//	}
//}
//




