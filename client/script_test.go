package client

import (
	"errors"
	"reflect"
	"strings"
	"testing"
)

var scriptSource string = ":put testing"
var scriptName string = "testing"
var scriptOwner string = "owner"
var scriptPolicies []string = []string{
	"ftp",
	"reboot",
	"read",
	"write",
	"policy",
	"test",
	"password",
	"sniff",
	"sensitive",
	"romon",
}
var scriptDontReqPerms = true

func TestCreateScriptAndDeleteScript(t *testing.T) {
	c := NewClient(GetConfigFromEnv())

	expectedScript := Script{
		Name:                   scriptName,
		Owner:                  scriptOwner,
		Source:                 scriptSource,
		PolicyString:           strings.Join(scriptPolicies, ","),
		DontRequirePermissions: scriptDontReqPerms,
	}
	script, err := c.CreateScript(
		scriptName,
		scriptOwner,
		scriptSource,
		scriptPolicies,
		scriptDontReqPerms,
	)

	if err != nil {
		t.Errorf("Error creating a script with: %v", err)
	}

	expectedScript.Id = script.Id

	if !reflect.DeepEqual(*script, expectedScript) {
		t.Errorf("The script does not match what we expected. actual: %v expected: %v", script, expectedScript)
	}

	err = c.DeleteScript(scriptName)

	if err != nil {
		t.Errorf("Error deleting a script with: %v", err)
	}
}

func TestFindScript_onNonExistantScript(t *testing.T) {
	c := NewClient(GetConfigFromEnv())

	name := "script-not-found"
	_, err := c.FindScript(name)

	if err == nil || errors.Is(err, &NotFound{}) {
		t.Errorf("client should have received error %T but got %T", &NotFound{}, err)
	}
}
