# GoFrame

[![Go Doc](https://godoc.org/github.com/gogf/gf?status.svg)](https://godoc.org/github.com/gogf/gf/g#pkg-subdirectories)
[![Build Status](https://travis-ci.org/gogf/gf.svg?branch=master)](https://travis-ci.org/gogf/gf)
[![Go Report](https://goreportcard.com/badge/github.com/gogf/gf?v=1)](https://goreportcard.com/report/github.com/gogf/gf)
[![Code Coverage](https://codecov.io/gh/gogf/gf/branch/master/graph/badge.svg)](https://codecov.io/gh/gogf/gf/branch/master)
[![Production Ready](https://img.shields.io/badge/production-ready-blue.svg)](https://github.com/gogf/gf)
[![License](https://img.shields.io/github/license/gogf/gf.svg?style=flat)](https://github.com/gogf/gf)

English | [简体中文](README_ZH.MD)

`GF(GoFrame)` is a modular, full-featured and production-ready application development framework of golang. Providing a series of core components and dozens of practical modules, such as: memcache, configure, validator, logging, array/queue/set/map containers, timer/timing tasks, file/memory lock, object pool, database ORM, etc. Supporting web server integrated with router, cookie, session, middleware, logger, template, https, hooks, rewrites and many more features.


# Installation
```
go get -u github.com/gogf/gf
```
suggested using `go.mod`:
```
require github.com/gogf/gf latest
```

# Limitation
```
golang version >= 1.10
```

# Documentation

* [APIDoc](https://godoc.org/github.com/gogf/gf)
* [中文文档](https://goframe.org)

# Architecture
<div align=center>
<img src="https://goframe.org/images/arch.png?v=10"/>
</div>

# Quick Start

```go
package main

import (
    "github.com/gogf/gf/frame/g"
    "github.com/gogf/gf/net/ghttp"
)

func main() {
    s := g.Server()
    s.BindHandler("/", func(r *ghttp.Request) {
        r.Response.Write("Hello World")
    })
    s.Run()
}
```

[More Features...](https://goframe.org/start/index)


# License

`GF` is licensed under the [MIT License](LICENSE), 100% free and open-source, forever.

# Donators

We currently accept donation by Alipay/WechatPay, please note your github/gitee account in your payment bill. If you like `GF`, why not [buy developer a cup of coffee](DONATOR.MD)?

# Thanks
<a href="https://www.jetbrains.com/?from=GoFrame"><img src="https://goframe.org/images/jetbrains.png" width="100" alt="JetBrains"/></a>


<!--
# Sponsor
We appreciate any kind of sponsorship for `GF` development. If you've got some interested, please contact john@goframe.org.
-->




