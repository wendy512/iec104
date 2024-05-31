# iec104

[![License](https://img.shields.io/badge/license-Apache--2.0-green.svg)](https://www.apache.org/licenses/LICENSE-2.0.html)

English | [中文](README_zh_CN.md)

IEC104 library

## Overview
iec104 is an open source (Apache-2.0 license) implementation of the IEC 104 client and server library that implements the MMS, GOOSE and SV protocols.
It can be used to implement IEC 61850 compliant clients and PCs on embedded systems and PCs running Linux, Windows Server application.

## Features

The library support the following IEC 104 protocol features:

* client/server for CS 104 TCP/IP communication
* Support for much application layer(except file object) message types
* Support for buffered and unbuffered reports
* Data access service (get data, set data)
* all data set services (get values, set values, browse)
* dynamic data set services (create and delete)
* TLS support

## How to use
```shell  
go get -u github.com/wendy512/iec104
```

## License
iec104 is based on the [Apache License 2.0](./LICENSE) agreement, and iec104 relies on some third-party components whose open source agreement is also Apache License 2.0.
## Contact

- Email：<taowenwuit@gmail.com>