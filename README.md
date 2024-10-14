# iec104

[![License](https://img.shields.io/badge/license-Apache--2.0-green.svg)](https://www.apache.org/licenses/LICENSE-2.0.html)

English | [中文](README_zh_CN.md)

IEC104 library

## Overview
This project implements a client for the IEC 60870-5-104 protocol (commonly referred to as IEC 104) using the Go programming language. 
IEC 104 is a widely used protocol in the electrical and industrial automation sectors, enabling reliable and efficient communication for remote control and data acquisition.

## Features

The library support the following IEC 104 protocol features:

* TCP/IP Based Communication
  </br>Utilizes standard TCP/IP protocols for communication, ensuring compatibility with a wide range of network infrastructures.
* Multiple Information Types
  </br>Supports transmission of various information types including single point information, double point information, measured values (normalized, scaled, short floating point), integrated totals, and commands (single, double, set point).
* Real-time Data Exchange
  </br>Provides real-time data exchange capabilities, essential for monitoring and controlling industrial processes and electrical systems.
* Time Synchronization
  </br>Supports time synchronization commands to ensure that all connected devices maintain accurate and synchronized time.
* Event-driven Communication
  </br>Supports event-driven data transmission, allowing for efficient communication by only sending updates when changes occur.
* Quality and Priority Indicators
  </br>Includes quality and priority indicators for transmitted data, ensuring that the integrity and importance of the data are maintained.
* Automatic Reconnection
  </br>Implements automatic reconnection mechanisms to handle network disruptions, ensuring continuous and reliable communication.

## How to use
```shell  
go get -u github.com/wendy512/iec104
```

- [Client reads and writes values](tests/clien_test.go)

## License
iec104 is based on the [Apache License 2.0](./LICENSE) agreement.
## Contact

- Email：<wendy512@yeah.net>
