package common

import (
	"reflect"
	"log"
	"fmt"
)

//todo 转换ids
func ConvertIds(){

}

//todo 转换id 为key 的map
func ConvertMapByIds(){

}




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

//todo targets传入一个数组， 通过指针数组传回值
// 可以实现转换， 但是数组无法通配, 使用范围很局限
// 配合ToSlice  可以转换为[]interface{}  本质上只有interface{}可以接收，例如dao.insertAll
// 调用方接收后需要手动转换为slice
// targets 传入一个非指针对象
func Converts(objects interface{}, targets ...interface{}) interface{} {
	if reflect.TypeOf(objects).Kind() != reflect.Slice {
		panic(NewServiceException(20002))
	}
	targetsLen := len(targets)
	ovs := reflect.ValueOf(objects)
	if ovs.Len() <= 0 {
		return targets[targetsLen:]
	}
	if len(targets)<= 0 {
		panic(NewServiceException(20005))
	}

	targetTemp := targets[0]

	oType := ovs.Index(0).Type()
	fieldNum := ovs.Index(0).Type().NumField()

	tType := reflect.TypeOf(targetTemp)
	fmt.Println(tType)
	//// 获取 字段数

	for i:=0; i< ovs.Len();i++ {
		object := ovs.Index(i)
		target := reflect.New(tType)
		for i:=0; i<fieldNum; i++ {
			field := object.Field(i)
			fieldName := oType.Field(i).Name
			fmt.Println(fieldName)
			value := field.Interface()
			// 判断是否有这个field
			if  tsf, ok:= tType.FieldByName(fieldName); ok {
				// 判断字段类型是否一样
				kind :=tsf.Type.Kind()
				if kind != field.Kind() {
					fmt.Println(kind)
					fmt.Println(field.Kind())
					log.Println("convert error，field type not onsistent in"+fieldName+ ",convert object:" + oType.String() +"=>>>" + tType.String())
					panic(NewServiceError(10000))
				}
				structFieldValue := target.Elem().FieldByName(fieldName) //结构体单个属性值
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
		targets = append(targets, target.Elem().Interface())

	}
	fmt.Println(len(targets))
	return targets[targetsLen:]
}


func ToSlice(arr interface{}) []interface{} {
	v := reflect.ValueOf(arr)
	if v.Kind() != reflect.Slice {
		panic("toslice arr not slice")
	}
	l := v.Len()
	ret := make([]interface{}, l)
	for i := 0; i < l; i++ {
		ret[i] = v.Index(i).Interface()
	}
	return ret
}


