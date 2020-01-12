package connect

import (
	"fmt"
	"log"
	"reflect"

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

type Technology struct {
	Name      string `json:"Name,omitempty"`
	Type      string `json:"Type,omitempty"`
	Powered   bool   `json:"Powered"`
	Connected bool   `json:"Connected"`
	Tethering bool   `json:"Tethering"`
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

func (c *ConnectionManager) GetTechnologyInfo(technology string) (*Technology, error) {
	info := Technology{}
	bo := c.GetManagerInterface()
	res := new([]interface{})
	err := bo.Call("net.connman.Manager.GetTechnologies", 0).Store(res)
	if err != nil {
		return nil, err
	}
	for _, item := range *res {
		if reflect.ValueOf(item).Kind() == reflect.Slice {
			path := fmt.Sprintf("/net/connman/technology/%s", technology)
			v := reflect.ValueOf(item)
			if string(v.Index(0).Interface().(dbus.ObjectPath)) == path {
				Decode(v.Index(1).Interface(), &info)
				break
			}
		}
	}
	return &info, nil
}
