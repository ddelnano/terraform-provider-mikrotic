package client

import (
	"log"

	"github.com/go-routeros/routeros"
)

type IpAddress struct {
	Id        string `mikrotik:".id"`
	Address   string `mikrotik:"address"`
	Comment   string `mikrotik:"comment"`
	Disabled  bool   `mikrotik:"disabled"`
	Interface string `mikrotik:"interface"`
	Network   string `mikrotik:"network"`
}

var ipAddressWrapper *resourceWrapper = &resourceWrapper{
	idField:               "id",
	mikrotikClientGetFunc: nil,
	actionsMap: map[string]string{
		"add":    "/ip/address/add",
		"find":   "/ip/address/print",
		"update": "/ip/address/set",
		"delete": "/ip/address/remove",
	},
	targetStruct:          &IpAddress{},
	addIDExtractorFunc:    func(r *routeros.Reply) string { return r.Done.Map["ret"] },
	recordIDExtractorFunc: func(r interface{}) string { return r.(*IpAddress).Id },
}

func (client Mikrotik) AddIpAddress(addr *IpAddress) (*IpAddress, error) {
	ipAddressWrapper.WithMikrotikClientGetter(client.getMikrotikClient)

	r, err := ipAddressWrapper.Add(addr)
	if err != nil {
		return nil, err
	}
	return r.(*IpAddress), nil
}

func (client Mikrotik) ListIpAddress() ([]IpAddress, error) {
	c, err := client.getMikrotikClient()

	if err != nil {
		return nil, err
	}
	cmd := []string{"/ip/address/print"}
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	r, err := c.RunArgs(cmd)

	if err != nil {
		return nil, err
	}
	log.Printf("[DEBUG] found ip address: %v", r)

	ipaddr := []IpAddress{}

	err = Unmarshal(*r, &ipaddr)

	if err != nil {
		return nil, err
	}

	return ipaddr, nil
}

func (client Mikrotik) FindIpAddress(id string) (*IpAddress, error) {
	ipAddressWrapper.WithMikrotikClientGetter(client.getMikrotikClient)
	ipaddr, err := ipAddressWrapper.Find(id)
	if err != nil {
		return nil, err
	}

	return ipaddr.(*IpAddress), nil
}

func (client Mikrotik) UpdateIpAddress(addr *IpAddress) (*IpAddress, error) {
	ipAddressWrapper.WithMikrotikClientGetter(client.getMikrotikClient)
	ipaddr, err := ipAddressWrapper.Update(addr)
	if err != nil {
		return nil, err
	}

	return ipaddr.(*IpAddress), nil
}

func (client Mikrotik) DeleteIpAddress(id string) error {
	return ipAddressWrapper.
		WithMikrotikClientGetter(client.getMikrotikClient).
		Delete(id)
}
