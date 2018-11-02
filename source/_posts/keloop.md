---
title: keloop 迭代
---

[keloop](https://o2o.keloop.cn/) 第三次迭代

### 平台首页

- 1 [平台首页](https://pro.modao.cc/app/ad7875830e7833f2bdcd0f58c3d7356ab80d03f4#screen=sfae74949f962739b316e59)

> * 1.1 基于之前的页面配置之上 添加一个管理分类兼容之前的页面配置
> * 1.2 由于第一版页面配置和平台首页是一对一的关系 需要调整下表结构
> * 1.3 现在迭代的版本需要实现一个平台能有多个平台首页(一对多的关系)
> * 1.4 首先需要在 **module** 模块 数据表中 添加一个 ``module_type_id`` 加一个分类id
> * 1.5 在 **module_type** 表中记得把 模块排序一起存上

### 店铺设置

- 2 [店铺设置](https://pro.modao.cc/app/ad7875830e7833f2bdcd0f58c3d7356ab80d03f4#screen=s97A0FEAD9A1533448292648)

> * 2.1 需要添加的字段(店铺大图,店铺业务,支付方式分为平台支付,和店铺自己配置支付)
> * 2.2 店铺相册需要 一个分类表 ,一个类型详情(把平台id和分类id做联合索引)
> * 2.3 如果 店铺业务存放成一个二进制数组 查询列表比较麻烦看到时候是吧分开还是聚合
> * 2.4 其他信息共用之前的数据 
> * 2.5 tips:[店铺类型](https://note.youdao.com/share/?id=5a15c447684d6bb6952d706c2e02e026&type=note#/)固定成常量 数据库存储一个店铺类型就可以了
> * 2.6 店铺配置中的 外送业务 (外卖),开启团购服务必须购买了市场应用才行
> * 2.7 是否预定之前的好像没有配置 如果没有在shop_info中添加一个添加一个是否预定的状态