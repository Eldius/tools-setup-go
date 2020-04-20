package specs

import (
	"fmt"
	"io/ioutil"

	"github.com/Eldius/tools-setup-go/logger"
	"gopkg.in/yaml.v3"
)

/*
ImportSpecs is an unit of tool spec
*/
func ImportSpecs(files []string) []ToolSpec {
	fmt.Println("import called with arguments: ", files)

	var specs []ToolSpec
	for _, f := range files {
		yamlFile, err := ioutil.ReadFile(f)
		if err != nil {
			panic(err.Error())
		}
		var specsList map[string]ToolSpec
		if err := yaml.Unmarshal(yamlFile, &specsList); err != nil {
			panic(err.Error())
		}
		for i, s := range specsList {
			logger.Debug(fmt.Sprintf("parsing %s", i))
			s.Name = i
			specs = append(specs, s)
		}
	}
	PersistSpecs(specs)

	for i, s := range ListSpecs() {
		fmt.Println(i, " => ", s.Name, ": ", s.Version)
	}

	return specs
}
