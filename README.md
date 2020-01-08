<h1 align="center">Welcome to go-connect 👋 </h1>
<h2 align="center">🚧 Work in Progress</h2>
<p>
  <a href="https://github.com/tsirysndr/go-connect/commits/master">
    <img src="https://img.shields.io/github/last-commit/tsirysndr/go-connect.svg" target="_blank" />
  </a>
  <img alt="GitHub code size in bytes" src="https://img.shields.io/github/languages/code-size/tsirysndr/go-connect">
  <img alt="GitHub go.mod Go version" src="https://img.shields.io/github/go-mod/go-version/tsirysndr/go-connect">
  <a href="https://github.com/tsirysndr/go-connect/graphs/contributors">
    <img alt="GitHub contributors" src="https://img.shields.io/github/contributors/tsirysndr/go-connect">
  </a>
  <a href="https://github.com/tsirysndr/go-connect/blob/master/LICENSE">
    <img alt="License: MIT" src="https://img.shields.io/badge/license-MIT-green.svg" target="_blank" />
  </a>
</p>

> Go library for configuring Linux connection (requires connman & dbus)

## 🚚 Install

```sh
go get github.com/tsirysndr/go-connect
```

## 🚀 Usage

Import the package into your project.

```Go
import "github.com/tsirynsdr/go-connect"
```

Construct a new connection manager, then use the various services on the client to access different parts of the Connman dbus API. For example:

```Go
c := connect.NewConnectionManager()
c.Wifi.Scan()
res, _ := c.Wifi.GetNetworks()
r, _ := json.Marshal(res)
fmt.Println(string(r))
```

## Author

👤 **Tsiry Sandratraina**

* Twitter: [@tsiry_sndr](https://twitter.com/tsiry_sndr)
* Github: [@tsirysndr](https://github.com/tsirysndr)

## Show your support

Give a ⭐️ if this project helped you!