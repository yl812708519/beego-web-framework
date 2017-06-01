package common

import (
	"encoding/json"
     _ "github.com/astaxie/beego/config/yaml"
	"github.com/astaxie/beego/config"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"log"
	"github.com/astaxie/beego/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

var (
	ExceptionConfig config.Configer
)

const ValidExceptionCode int = 20000
var validErrMsgMap = map[string] string{
	"Required":     "不能为空",
	"Min":          "最小值是 %d",
	"Max":          "最大值是 %d",
	"Range":        "取值范围是 %d 到 %d",
	"MinSize":      "长度过短，最短长度为 %d",
	"MaxSize":      "输入过长，最大长度为 %d",
	"Length":       "指定长度为 %d",
	"Numeric":      "必须是一个数字字符串",
	"Match":        "必须匹配 %s",
	"NoMatch":      "必须不匹配 %s",
	"Email":        "必须是一个邮箱地址",
	"IP":           "必须是一个IP地址",
	"Base64":       "必须是一个base64字符串",
	"Mobile":       "必须是一个手机号码",
	"Tel":          "必须是一个电话号码",
	"Phone":        "必须是一个手机或电话号码",
	"ZipCode":      "必须是一个邮政编码",
}

type ServiceException struct {

	Code int
	Message string
}

type ServiceError struct {
	Code int
	Message string
}

func NewServiceException(exceptionCode int, msg ...string) ServiceException {

	if exceptionCode == ValidExceptionCode && len(msg) > 0 {
		return ServiceException{exceptionCode, msg[0]}
	}

	ecStr:= strconv.Itoa(exceptionCode)
	message := ExceptionConfig.String(ecStr)
	return ServiceException{exceptionCode, message}

}

func (e ServiceException) Json() []byte {

	content, err := json.Marshal(e)
	if err != nil {
		panic("convert error")
	}
	return content
}

func (e ServiceException) Error() string{
	content, _ := json.Marshal(e)
	return string(content)
}

func NewServiceError(exceptionCode int) ServiceError {
	ecStr:= strconv.Itoa(exceptionCode)
	message := ExceptionConfig.String(ecStr)
	return ServiceError{exceptionCode, message}
}

func (e ServiceError) Json() []byte {

	content, err := json.Marshal(e)
	if err != nil {
		panic("convert error")
	}
	return content
}

func (e ServiceError) Error() string{
	content, _ := json.Marshal(e)
	return string(content)
}




func init() {
	// 读取 错误的配置
	var configPath string
	workPath, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		fmt.Println(workPath)
		panic("can`t parse current path")
	}
	configPath = filepath.Join(workPath, "conf", "exception.conf")
	if !utils.FileExists(configPath) {
		log.Panicln(configPath)
		configPath = filepath.Join(beego.AppPath, "conf", "app.conf")
		if !utils.FileExists(configPath) {
			log.Panicln(configPath)
			log.Panicln("can`t found error config path")
		}
	}
	c, err := config.NewConfig("yaml", configPath)
	if err!= nil {
		log.Panicln("config parse error: " + err.Error())
	}
	ExceptionConfig = c

	// 修改valid 模块的报错
	validation.SetDefaultMessage(validErrMsgMap)

}




