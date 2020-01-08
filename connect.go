package connect

import (
	"log"

	"github.com/tsirysndr/dbus"
)

const (
	CONNMAN_OBJECT_PATH = "net.connman"
	ETHERNET_PATH       = "/net/connman/service/ethernet"
	WIFI_PATH           = "/net/connman/service/wifi"
)

type ConnectionManager struct {
	base      *dbus.Conn
	common    service
	Bluetooth *BluetoothService
	Ethernet  *EthernetService
	Wifi      *WifiService
	Tethering *TetheringService
}

type service struct {
	manager *ConnectionManager
}

func NewConnectionManager() *ConnectionManager {
	c := &ConnectionManager{}
	base, err := dbus.SystemBus()
	c.base = base
	c.common.manager = c
	c.Bluetooth = (*BluetoothService)(&c.common)
	c.Ethernet = (*EthernetService)(&c.common)
	c.Wifi = (*WifiService)(&c.common)
	c.Tethering = (*TetheringService)(&c.common)
	if err != nil {
		log.Fatal(err)
	}
	return c
}

func (c *ConnectionManager) GetTechnologyInterface(path dbus.ObjectPath) dbus.BusObject {
	return c.base.Object(CONNMAN_OBJECT_PATH, "/net/connman/technology/"+path)
}

func (c *ConnectionManager) GetManagerInterface() dbus.BusObject {
	return c.base.Object(CONNMAN_OBJECT_PATH, "/")
}
