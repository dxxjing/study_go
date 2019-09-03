package iniConfig

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func Marshal(data interface{}) (result []byte,err error){
	//判断传入的是否为结构体
	typeInfo := reflect.TypeOf(data)
	if typeInfo.Kind() != reflect.Struct {
		err = errors.New("please pass struct")
		return
	}
	var conf []string
	//遍历Config中的字段
	valueInfo := reflect.ValueOf(data)
	for i := 0; i < typeInfo.NumField(); i++ {
		//取出key val
		sectionKey := typeInfo.Field(i)
		sectionVal := valueInfo.Field(i)

		fieldType := sectionKey.Type
		if fieldType.Kind() != reflect.Struct {
			continue
		}
		tagVal := sectionKey.Tag.Get("ini")
		if len(tagVal) == 0 {
			tagVal = sectionKey.Name
		}

		section := fmt.Sprintf("\n[%s]\n", tagVal)
		conf = append(conf, section)
		//遍历ServerConfig 或 MysqlConfig中的字段
		for j := 0; j < fieldType.NumField(); j++ {
			keyField := fieldType.Field(j)
			fieldTagVal := keyField.Tag.Get("ini")
			if len(fieldTagVal) == 0 {
				fieldTagVal = keyField.Name
			}

			valField := sectionVal.Field(j)
			item := fmt.Sprintf("%s=%v\n", fieldTagVal, valField.Interface())
			conf = append(conf, item)
		}
	}
	for _, val := range conf {
		byteVal := []byte(val)
		result = append(result, byteVal...)
	}
	return
}

func UnMarshal(iniData []byte,result interface{}) (err error) {
	iniSlice := strings.Split(string(iniData),"\n")
	vType := reflect.TypeOf(result)
	if vType.Kind() != reflect.Ptr {
		err = errors.New("param data not point")
		return
	}
	typeStruct := vType.Elem()
	if typeStruct.Kind() != reflect.Struct {
		return errors.New("param data not struct")
	}
	var lastFieldName string
	for index,line := range iniSlice {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		if line[0] == ';' || line[0] == '#'{
			continue
		}
		if line[0] == '[' {
			lastFieldName,err = parseSection(line,typeStruct)
			if err != nil {
				return
			}
			continue
		}
		err = parseItem(lastFieldName,line,result)
		if err != nil {
			err = fmt.Errorf("%v lineno:%d", err, index+1)
			return
		}
	}
	return
}

//ip=10.238.2.2
//遍历ServerConfig 或 MysqlConfig 中的字段 并给字段设置值
func parseItem(lastFieldName,line string,result interface{}) (err error){
	//ini文件中的行 没有=号 退出
	index := strings.Index(line,"=")
	if index == -1 {
		err = fmt.Errorf("sytax error, line:%s", line)
		return
	}
	//取出等号两边的key val
	key := strings.TrimSpace(line[0:index])
	val := strings.TrimSpace(line[index+1:])
	if len(key) == 0 || len(val) == 0 {
		err = fmt.Errorf("sytax error, line:%s", line)
		return
	}
	//取出Config 中的SrvConf 或 SqlConf字段
	resultVal := reflect.ValueOf(result)
	sectionVal := resultVal.Elem().FieldByName(lastFieldName)
	//判断SrvConf 或 SqlConf 是否为结构体
	sectionType := sectionVal.Type()
	if sectionType.Kind() != reflect.Struct {
		err = fmt.Errorf("field:%s must be struct", lastFieldName)
		return
	}
	keyFieldName := ""
	//遍历ServerConfig 或 MysqlConfig中的字段
	for i := 0; i < sectionType.NumField(); i++ {
		field := sectionType.Field(i)
		if field.Tag.Get("ini") == key {
			keyFieldName = field.Name
			break
		}
	}
	if len(keyFieldName) == 0 {
		return
	}

	fieldValue := sectionVal.FieldByName(keyFieldName)
	if fieldValue == reflect.ValueOf(nil) {
		return
	}
	//判断ServerConfig 或 MysqlConfig中的字段类型 并通过反射设置值
	switch fieldValue.Type().Kind() {
	case reflect.String:
		fieldValue.SetString(val)
	case reflect.Int8, reflect.Int16, reflect.Int, reflect.Int32, reflect.Int64:
		intVal, errRet := strconv.ParseInt(val, 10, 64)
		if errRet != nil {
			err = errRet
			return
		}
		fieldValue.SetInt(intVal)

	case reflect.Uint8, reflect.Uint16, reflect.Uint, reflect.Uint32, reflect.Uint64:
		intVal, errRet := strconv.ParseUint(val, 10, 64)
		if errRet != nil {
			err = errRet
			return
		}
		fieldValue.SetUint(intVal)
	case reflect.Float32, reflect.Float64:
		floatVal, errRet := strconv.ParseFloat(val, 64)
		if errRet != nil {
			return
		}

		fieldValue.SetFloat(floatVal)

	default:
		err = fmt.Errorf("unsupport type:%v", fieldValue.Type().Kind())
	}
	return
}


// [server] => SrvConf
func parseSection(line string,typeInfo reflect.Type) (fieldName string,err error) {
	if line[len(line)-1] != ']' {
		err = fmt.Errorf("syntax error, invalid section:%s", line)
		return
	}
	//去除中括号以及两边空格 并判断剩余字符串长度
	sectionName := strings.TrimSpace(line[1:len(line)-1])
	if len(line) == 0 {
		err = fmt.Errorf("syntax error, invalid section:%s", line)
		return
	}
	//遍历Config结构体中的字段
	for i := 0; i < typeInfo.NumField(); i++ {
		field := typeInfo.Field(i)
		if field.Tag.Get("ini") == sectionName {
			fieldName = field.Name
			break
		}
	}
	return
}
