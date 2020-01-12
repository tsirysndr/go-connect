package connect

import (
	"github.com/tsirysndr/dbus"
)

type TetheringService service

func (s *TetheringService) IsTetheringEthernet() (bool, error) {
	info, err := s.manager.GetTechnologyInfo("ethernet")
	if err != nil {
		return false, err
	}
	return info.Tethering, nil
}

func (s *TetheringService) IsTetheringWifi() (bool, error) {
	info, err := s.manager.GetTechnologyInfo("wifi")
	if err != nil {
		return false, err
	}
	return info.Tethering, nil
}

func (s *TetheringService) GetTetheringSSID() {
}

func (s *TetheringService) GetTetheringPassphrase() {
}

func (s *TetheringService) EnableTethering(technology dbus.ObjectPath, SSID, passphrase string) error {
	bo := s.manager.GetTechnologyInterface(technology)
	if len(SSID) > 0 {
		call := bo.Call("net.connman.Technology.SetProperty", 0, "TetheringIdentifier", dbus.MakeVariant(SSID))
		if call.Err != nil {
			return call.Err
		}
	}
	if len(passphrase) > 0 {
		call := bo.Call("net.connman.Technology.SetProperty", 0, "TetheringPassphrase", dbus.MakeVariant(passphrase))
		if call.Err != nil {
			return call.Err
		}
	}
	call := bo.Call("net.connman.Technology.SetProperty", 0, "Tethering", dbus.MakeVariant(true))
	return call.Err
}

func (s *TetheringService) DisableTethering(technology dbus.ObjectPath) error {
	bo := s.manager.GetTechnologyInterface(technology)
	call := bo.Call("net.connman.Technology.SetProperty", 0, "Tethering", dbus.MakeVariant(false))
	return call.Err
}
