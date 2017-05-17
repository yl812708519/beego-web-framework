
# 项目描述

## 项目分为三层，view 托管于框架不在此中
* controller :  做请求数据的拼装过滤等工作， 尽量不涉及业务逻辑
* service： 服务层， 处理业务逻辑， 返回DTO对象给controller层
* dao: 数据层， 项目中使用models文件夹，（跟随beego创建规则）， 负责操作数据， 目前这层返回的直接是model实体， 并未包装

## 错误封装
项目中定义异常(serviceException)、 错误(serviceError), 业务中的逻辑错误使用serviceException通知controller， panic(&serviceException) 不会中断进程， 框架会捕获panic， 如果panic的参数可以转换为 serviceException, 会将http status 设为400 并返回相应的error message json。 serviceError目前未做其他的封装与使用


## config
配置同beego, 增加orm data source url 的相关配置



## orm
model 定义的init 需要 调用orm.Register(new(model))， 否则无法使用高级查询
orm 部分未经过测试。。。。


# 项目依赖
项目依赖使用golang 自带vendor管理
* 编译之前请设置  export GO15VENDOREXPERIMENT=1



就这样。。。想到再补。。
