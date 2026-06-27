# 快手开放平台（快手小店）GO SDK

对接快手电商开放平台（快手小店）的 Go SDK。

所有凭证（appKey / appSecret / signSecret / msgSecret）均通过**传参**传入，不在 SDK 内写死，方便复用。

## 平台链接

- [快手电商开放平台](https://open.kwaixiaodian.com/)
- [APP授权说明](https://open.kwaixiaodian.com/zone/new/docs/dev?pageSign=e1d9e229332f4f233a04b44833a5dfe71614263940720)
- [消息服务说明](https://open.kwaixiaodian.com/zone/new/docs/dev?pageSign=120a69028f0e9ea69145644317ae8bcd1614263915925)
- [返回码说明](https://open.kwaixiaodian.com/zone/new/docs/dev?pageSign=5de448fecaddd4c58104e8aea442695b1614263998209)

## 安装命令

```sh
go get github.com/zsmhub/ks-sdk
```

## sdk调用示例

**去 ./demo 文件夹查看完整示例！**

[点击查看完整demo](https://github.com/zsmhub/ks-sdk/-/tree/master/demo)

## 注意

1. 一个订单只会有一个商品。客户一次性购买N个商品时快手会自动拆成N个订单。
2. 如果你想要的接口或回调消息找不到定义？你可以自己添加，so easy!
