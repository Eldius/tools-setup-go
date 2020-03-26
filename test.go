package main

import (
	"fmt"

	"github.com/Eldius/tools-setup-go/config"
	"github.com/spf13/viper"
	yaml "gopkg.in/yaml.v2"
)

func main() {
	terraform := config.ToolSpec{
		SetupType:  "download-and-unpack-binary",
		URL:        "https://releases.hashicorp.com/terraform/0.12.24/terraform_0.12.24_linux_amd64.zip",
		Version:    "0.12.24",
		VersionCmd: "torraform version",
	}
	packer := config.ToolSpec{
		SetupType:  "download-and-unpack-binary",
		URL:        "https://releases.hashicorp.com/packer/1.5.4/packer_1.5.4_linux_amd64.zip",
		Version:    "1.5.4",
		VersionCmd: "packer version",
	}

	specsMap := make(map[string]config.ToolSpec, 2)
	specsMap["terraform"] = terraform
	specsMap["packer"] = packer
	specs := config.SetupSpecConfig{
		Specs:     specsMap,
		BinFolder: "~/.tools-setup-go/bin",
		LogFile:   "~/.tools-setup-go/execution.log",
		Verbose:   true,
	}
	bs, err := yaml.Marshal(specs)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("---")
	fmt.Println(string(bs))
	fmt.Println("---")
	viper.SetDefault(".", specs)
	viper.WriteConfigAs("samples/config-test.yml")
}
