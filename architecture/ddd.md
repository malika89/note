DDD：Domain Driven Design 领域驱动设计

根据业务和语义边境，将一个或者多个聚合划定在一个限界上下文内，形成领域模型
多用于微服务和中台设计

### how it works
+ 1.事件风暴(业务场景和外部依赖)
+ 2.领域事件分析
+ 3.提取领域对象(紧密关联的实体)
+ 4.领域对象和代码模型映射(业务架构绑定系统架构)
+ 5.代码逻辑

### 常见三种架构模式
   + 分层架构-一般企业项目和微服务采用
     + 用户接口层
     + 应用层
     + 领域层
     + 基础层
     
   + 整洁架构
     + 领域模型
     + 领域服务
     + 应用服务
     + 基础设施、界面
     
   + 六边形架构：应用通过端口与界面进行交互
     + 核心业务逻辑(应用程序和领域模型)
     + 外部资源(app 、数据库资源)
  

### 分层
每层仅与其下方的层发生耦合

* 用户接口层
`
包含前端用户界面、web服务、
`
* 应用层
`
应用服务
`
* 领域层
`
聚合 领域服务(核心域、通用域、支撑域)
`
* 基础层
`
网关 数据库 缓存 总线
`

### 防腐层
微服务间集成时采用DDD中的防腐层（Anti-Corruption Layer, ACL 

#### 作用：
`
一个上下文通过一些适配和转换与另一个上下文交互，适用于不同应用间转换；可确保应用程序的设计不受限于对外部子系统的依赖
 + 1）在架构层面，通过引入防腐层有效隔离限界上下文之间的耦合；
 + 2）防腐层同时还可以扮演适配器、调停者、外观等角色；
 + 3）防腐层往往属于下游限界上下文，用以隔绝上游限界上下文可能发生的变化；
`

#### 使用场景
+ 旧版单体应用迁移到新版微服务系统，但是迁移计划发生在多个阶段，新旧系统之间的集成需要维护
+ 两个或更多不同的子系统（或限界上下文）具有不同的领域模型，需要对外部上下文的方法进行一次转义。




-----


### 1. 事件风暴
常用分析方法：
  用例分析
  场景分析
  旅程分析
  
### 实体与值对象
`
代码模型中实体的形式是实体类，这个类拥有实体的属性和方法(业务逻辑)。
实体类通常采用充血模型，业务逻辑在方法中实现。跨实体的领域逻辑在领域服务中实现

 DO 实体：有唯一ID;具有业务属性和业务行为
 值对象：属性集合，对实力的状态和特征进行补充描述
`

### 聚合与聚合根
`
聚合多个实体组成
聚合(跨聚合的业务逻辑在应用服务实现，同聚合的在领域服务实现)

`

### 其他概念
#### 1.领域事件

#### 2.仓储模式

#### 3.事件总线
