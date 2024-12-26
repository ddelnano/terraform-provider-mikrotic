package client

import (
	"log"

	"github.com/ddelnano/terraform-provider-mikrotik/client/types"
)

type SystemResources struct {
	Uptime  types.MikrotikDuration `mikrotik:"uptime,readonly"`
	Version string                 `mikrotik:"version,readonly"`
}

func (d *SystemResources) ActionToCommand(action Action) string {
	return map[Action]string{
		Find: "/system/resource/print",
	}[action]
}

func (client Mikrotik) GetSystemResources() (*SystemResources, error) {
	c, err := client.getMikrotikClient()
	if err != nil {
		return nil, err
	}
	sysResources := &SystemResources{}
	cmd := Marshal(sysResources.ActionToCommand(Find), sysResources)

	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	r, err := c.RunArgs(cmd)
	if err != nil {
		return nil, err
	}

	err = Unmarshal(*r, sysResources)
	return sysResources, err
}
