# 🪝 Servicex

[![Go Doc](_icons/godoc.svg)](https://pkg.go.dev/github.com/FishGoddess/servicex)
[![License](_icons/license.svg)](https://opensource.org/licenses/MIT)
[![License](_icons/coverage.svg)](_icons/coverage.svg)
![Test](https://github.com/FishGoddess/servicex/actions/workflows/check.yml/badge.svg)

**Servicex** 是一些服务生态的集合包，开箱即用。

[Read me in English](./README.en.md)

### 📝 功能特性

* 常用拦截器，比如超时、panic 保护、耗时统计等等
* 链路 id 使用 context 透传
* 随机字符串生成，UUID 生成，自增序列获取
* 栈信息获取，自动设置 max proc 数值
* 带自动解析的 Duration 类型，快速获取时间的 Clock

_历史版本的特性请查看 [HISTORY.md](./HISTORY.md)。未来版本的新特性和计划请查看 [FUTURE.md](./FUTURE.md)。_

### 🔧 使用方式

> $ go get -u github.com/FishGoddess/servicex

### 👥 贡献者

如果您觉得 servicex 缺少您需要的功能，请不要犹豫，马上参与进来，发起一个 _**issue**_。

### 📦 使用 servicex 的项目

| 项目     | 作者         | 描述        | 链接                                                                                         |
|--------|------------|-----------|--------------------------------------------------------------------------------------------|
| postar | avino-plan | 开箱即用的邮件服务 | [Github](https://github.com/avino-plan/postar) / [码云](https://gitee.com/avino-plan/postar) |

最后，我想感谢 JetBrains 公司的 **free JetBrains Open Source license(s)**，因为 servicex 是用该计划下的 Idea /
GoLand 完成开发的。

<a href="https://www.jetbrains.com/?from=servicex" target="_blank"><img src="./_icons/jetbrains.png" width="250"/></a>