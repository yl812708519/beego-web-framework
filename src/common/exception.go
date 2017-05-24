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
)

var (
	ExceptionConfig config.Configer
)


type ServiceException struct {

	Code int
	Message string
}

type ServiceError struct {
	Code int
	Message string
}

func NewServiceException(exceptionCode int) ServiceException {
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
}




