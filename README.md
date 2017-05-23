
# 项目描述

## 项目分为三层，view 托管于框架不在此中
* controller :  做请求数据的拼装过滤等工作， 尽量不涉及业务逻辑
* service： 服务层， 处理业务逻辑， 返回DTO对象给controller层
* dao: 数据层， 项目中使用models文件夹，（跟随beego创建规则）， 负责操作数据， 目前这层返回的直接是model实体， 并未包装


##baseController
这里封装了部分C层常用的方法， 有些与框架的基类重名， 重名方法使用小写字母， 反正是同包使用

## 参数处理
框架提供的传入参数处理机制貌似不完善， 无法parse 嵌套对象(无法测试尚不明了)
使用解析json的形式可以达到要求，性能应该不是最优解
form-data、x-www-form-urlencode 可以使用this.Ctx.Input() 接收
application/json 只能使用 this.Ctx.Input.RequestBody()获取并解析


## 错误封装
项目中定义异常(serviceException)、 错误(serviceError)  
业务中的逻辑错误使用panic(&serviceException) 不会中断进程， 框架会捕获panic  
如果panic的参数可以转换为 serviceException, 会将http status 设为400 并返回相应的error message json。 serviceError目前未做其他的封装与使用


## config
配置同beego, 增加orm data source url 的相关配置  


## orm
model 定义的init 需要 调用orm.Register(new(model))， 否则无法使用高级查询  
model 定义， 最好写上对应的column  
model 不使用自动建表功能  
model 中bool类型在表中定义为tinyint(1) 0:false   1:true  orm会转换
表明使用复数名词， model不固定最好使用单数名词， 实现tableName 方法返回真正的表名  
orm 部分未经过测试。。。。  


# 项目依赖
项目依赖使用golang 自带vendor管理
> 编译之前请设置  export GO15VENDOREXPERIMENT=1



就这样。。。想到再补。。
