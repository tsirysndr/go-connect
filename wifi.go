package connect

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"reflect"
	"time"

	"github.com/tsirysndr/dbus"
)

type WifiService service

type Service struct {
	Path                     dbus.ObjectPath
	Type                     string            `json:"Type,omitempty"`
	Security                 string            `json:"Security,omitempty"`
	State                    string            `json:"State,omitempty"`
	Strength                 int               `json:"Strength"`
	Favorite                 bool              `json:"Favorite"`
	Immutable                bool              `json:"Immutable"`
	AutoConnect              bool              `json:"AutoConnect"`
	Name                     string            `json:"Name,omitempty"`
	Ethernet                 Ethernet          `json:"Ethernet,omitempty"`
	IPv4                     IPv4              `json:"IPv4,omitempty"`
	IPv4Configuration        IPv4Configuration `json:"IPv4.Configuration,omitempty"`
	IPv6                     interface{}       `json:"IPv6,omitempty"`
	IPv6Configuration        IPv6Configuration `json:"IPv6.Configuration,omitempty"`
	Nameservers              []string          `json:"Nameservers,omitempty"`
	NameserversConfiguration []string          `json:"Nameservers.Configuration,omitempty"`
	Timeservers              []string          `json:"Timeservers,omitempty"`
	TimeserversConfiguration []string          `json:"Timeservers.Configuration,omitempty"`
	Domains                  []string          `json:"Domains,omitempty"`
	DomainsConfiguration     []string          `json:"Domains.Configuration,omitempty"`
	Proxy                    Proxy             `json:"Proxy,omitempty"`
	ProxyConfiguration       interface{}       `json:"Proxy.Configuration,omitempty"`
	MDNS                     bool              `json:"mDNS"`
	MDNSConfiguration        bool              `json:"mDNS.Configuration"`
	Provider                 interface{}       `json:"Provider,omitempty"`
}

type Ethernet struct {
	Address   string `json:"Address,omitempty"`
	Interface string `json:"Interface,omitempty"`
	MTU       int    `json:"MTU,omitempty"`
	Method    string `json:"Method,omitempty"`
}

type IPv4 struct {
	Address string `json:"Address,omitempty"`
	Gateway string `json:"Gateway,omitempty"`
	Method  string `json:"Method,omitempty"`
	Netmask string `json:"Netmask,omitempty"`
}

type IPv4Configuration struct {
	Method string `json:"Method,omitempty"`
}

type IPv6Configuration struct {
	Method string `json:"Method,omitempty"`
}

type Proxy struct {
	Method string `json:"Method,omitempty"`
}

func (s *WifiService) IsAvailable() bool {
	info, _ := s.manager.GetTechnologyInfo("wifi")
	return info.Type != ""
}

func (s *WifiService) IsEnabled() bool {
	info, _ := s.manager.GetTechnologyInfo("wifi")
	if info.Type == "" {
		return false
	}
	return info.Powered
}

func (s *WifiService) ToggleState(state bool) error {
	bo := s.manager.GetTechnologyInterface("wifi")
	call := bo.Call("net.connman.Technology.SetProperty", 0, "Powered", dbus.MakeVariant(state))
	return call.Err
}

func (s *WifiService) Scan() error {
	bo := s.manager.GetTechnologyInterface("wifi")
	res := new([]interface{})
	return bo.Call("net.connman.Technology.Scan", 0).Store(res)
}

func (s *WifiService) GetNetworks() ([]Service, error) {
	networks := []Service{}
	bo := s.manager.GetManagerInterface()
	res := new([]interface{})
	err := bo.Call("net.connman.Manager.GetServices", 0).Store(res)
	if err != nil {
		return networks, err
	}
	for _, item := range *res {
		if reflect.ValueOf(item).Kind() == reflect.Slice {
			v := reflect.ValueOf(item)
			data := Service{}
			Decode(v.Index(1).Interface(), &data)
			data.Path = v.Index(0).Interface().(dbus.ObjectPath)
			networks = append(networks, data)
		}
	}

	return networks, nil
}

func (s *WifiService) Connect(password, SSID, path string) (bool, error) {
	log.Println("Starting Wireless Agent")
	keyfile := "/tmp/preseed_data"
	content := fmt.Sprintf("%s\n%s", password, SSID)

	if err := ioutil.WriteFile(keyfile, []byte(content), 0644); err != nil {
		return false, err
	}

	const agent = "/tmp/wireless_agent.py"
	code := []byte(WIRELESS_AGENT)

	if err := ioutil.WriteFile(agent, code, 0644); err != nil {
		return false, err
	}

	if err := exec.Command("chmod", "a+x", agent).Run(); err != nil {
		return false, err
	}

	log.Println(agent, "fromfile")

	cmd := exec.Command(agent, "fromfile")

	if err := cmd.Start(); err != nil {
		return false, err
	}

	connected := 1
	connectionAttempts := 20
	bo := s.manager.GetServiceInterface(dbus.ObjectPath(path))

	for connected != 0 && connected < (connectionAttempts+1) {
		call := bo.Call("net.connman.Service.Connect", 0)
		connected = 0
		if call.Err != nil {
			log.Println(call.Err)
			connected += 1
			time.Sleep(1 * time.Second)
			log.Println("Connection agent not started yet, waiting a second")
			if call.Err.Error() != "Not registered" {
				break
			}
		} else {
			break
		}
	}

	log.Println("Connection to ", path, " : ", connected == 0)
	if err := cmd.Process.Kill(); err != nil {
		return connected == 0, err
	}
	if err := exec.Command("rm", agent, keyfile).Run(); err != nil {
		return connected == 0, err
	}

	return connected == 0, nil
}

func (s *WifiService) Disconnect(path string) error {
	bo := s.manager.GetServiceInterface(dbus.ObjectPath(path))
	call := bo.Call("net.connman.Service.Disconnect", 0)
	return call.Err
}

func (s *WifiService) Remove(path string) error {
	bo := s.manager.GetServiceInterface(dbus.ObjectPath(path))
	call := bo.Call("net.connman.Service.Remove", 0)
	return call.Err
}

func (s *WifiService) GetConnectedWifi() *Service {
	networks, _ := s.GetNetworks()
	for _, item := range networks {
		if item.State == "online" {
			return &item
		}
	}
	return nil
}

func (s *WifiService) HasNetworkConnection() {
}

func Decode(src interface{}, dest interface{}) error {
	data, err := json.Marshal(src)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, dest)
}
