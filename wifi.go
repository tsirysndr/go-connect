package connect

import (
	"encoding/json"
	"reflect"
)

type WifiService service

type Services struct {
	Type                     string            `json:"Type,omitempty"`
	Security                 string            `json:"Security,omitempty"`
	State                    string            `json:"State,omitempty"`
	Strength                 int               `json:"Strength,omitempty"`
	Favorite                 bool              `json:"Favorite,omitempty"`
	Immutable                bool              `json:"Immutable,omitempty"`
	AutoConnect              bool              `json:"AutoConnect,omitempty"`
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

func (s *WifiService) IsAvailable() {
}

func (s *WifiService) IsEnabled() {
}

func (s *WifiService) ToggleState() {
}

func (s *WifiService) Scan() error {
	bo := s.manager.GetTechnologyInterface("wifi")
	res := new([]interface{})
	return bo.Call("net.connman.Technology.Scan", 0).Store(res)
}

func (s *WifiService) GetNetworks() ([]Services, error) {
	networks := []Services{}
	bo := s.manager.GetManagerInterface()
	res := new([]interface{})
	err := bo.Call("net.connman.Manager.GetServices", 0).Store(res)
	if err != nil {
		return networks, err
	}
	for _, item := range *res {
		if reflect.ValueOf(item).Kind() == reflect.Slice {
			v := reflect.ValueOf(item)
			data := Services{}
			Decode(v.Index(1).Interface(), &data)
			networks = append(networks, data)
		}
	}

	return networks, nil
}

func (s *WifiService) Connect() {
}

func (s *WifiService) Disconnect() {
}

func (s *WifiService) Remove() {
}

func (s *WifiService) GetConnectedWifi() {
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
