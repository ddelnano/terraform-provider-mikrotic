package client

import (
	"github.com/go-routeros/routeros"
)

type InterfaceWireguard struct {
	Name       string `mikrotik:"name"`
	Comment    string `mikrotik:"comment"`
	Disabled   bool   `mikrotik:"disabled"`
	ListenPort int    `mikrotik:"listen-port"`
	Mtu        int    `mikrotik:"mtu"`
	PrivateKey string `mikrotik:"private-key"`
	PublicKey  string `mikrotik:"public-key"` //read only property
	Running    bool   `mikrotik:"running"`    //read only property
}

func (i *InterfaceWireguard) ActionToCommand(action Action) string {
	return map[Action]string{
		Add:    "/interface/wireguard/add",
		Find:   "/interface/wireguard/print",
		List:   "/interface/wireguard/print",
		Update: "/interface/wireguard/set",
		Delete: "/interface/wireguard/remove",
	}[action]
}

func (i *InterfaceWireguard) IDField() string {
	return "name"
}

func (i *InterfaceWireguard) ID() string {
	return i.Name
}

func (i *InterfaceWireguard) SetID(id string) {
	i.Name = id
}

func (i *InterfaceWireguard) AfterAddHook(r *routeros.Reply) {
	i.Name = r.Done.Map["ret"]
}

func (i *InterfaceWireguard) FindField() string {
	return "name"
}

func (i *InterfaceWireguard) FindFieldValue() string {
	return i.Name
}

func (i *InterfaceWireguard) DeleteField() string {
	return "name"
}

func (i *InterfaceWireguard) DeleteFieldValue() string {
	return i.Name
}

func (client Mikrotik) AddInterfaceWireguard(i *InterfaceWireguard) (*InterfaceWireguard, error) {
	res, err := client.Add(i)
	if err != nil {
		return nil, err
	}

	return res.(*InterfaceWireguard), nil
}

func (client Mikrotik) FindInterfaceWireguard(name string) (*InterfaceWireguard, error) {
	res, err := client.Find(&InterfaceWireguard{Name: name})
	if err != nil {
		return nil, err
	}

	return res.(*InterfaceWireguard), nil
}

func (client Mikrotik) UpdateInterfaceWireguard(i *InterfaceWireguard) (*InterfaceWireguard, error) {
	res, err := client.Update(i)
	if err != nil {
		return nil, err
	}

	return res.(*InterfaceWireguard), nil
}

func (client Mikrotik) DeleteInterfaceWireguard(name string) error {
	return client.Delete(&InterfaceWireguard{Name: name})
}
