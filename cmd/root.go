package cmd

import (
	"github.com/benben/serialbeat/beater"
	
	"github.com/elastic/beats/libbeat/cmd/instance"
	cmd "github.com/elastic/beats/libbeat/cmd"
)

// Name of this beat
var Name = "serialbeat"

// RootCmd to handle beats cli
var RootCmd = cmd.GenRootCmdWithSettings(beater.New, instance.Settings{
	Name: Name,
})
