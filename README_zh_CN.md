# iec61850

[![License](https://img.shields.io/badge/license-Apache--2.0-green.svg)](https://www.apache.org/licenses/LICENSE-2.0.html)

中文 | [English](README.md)

IEC104库

## 概述
IEC104 是实现 MMS、GOOSE 和 SV 协议的 IEC 104 客户端和服务器库的开源 (Apache-2.0 license) 实现，它可用于在运行 Linux、Windows 的嵌入式系统和 PC 上实施符合 IEC 61850 的客户端和服务器应用程序。

## 功能特性
该库支持以下IEC 104协议功能：

* CS 104 TCP/IP 通信的客户端/服务器
* 支持多种应用层（文件对象除外）消息类型
* 支持缓冲和非缓冲报告
* 数据访问服务（获取数据、设置数据）
* 所有数据集服务（获取值、设置值、浏览）
* 动态数据集服务（创建和删除）
* TLS 支持

## 如何使用
```shell  
go get -u github.com/wendy512/iec104
```

## 开源许可
iec104 基于 [Apache License 2.0](./LICENSE) 协议，iec104 依赖了一些第三方组件，它们的开源协议也为 Apache License 2.0。

## 联系方式

- 邮箱：<taowenwuit@gmail.com>