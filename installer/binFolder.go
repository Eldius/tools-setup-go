package installer

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/Eldius/tools-setup-go/config"
)

/*
MoveToBinFolder moves files to bin folder
*/
func MoveToBinFolder(fileSet []string) {
	binFolder := getBinFolder()
	for _, f := range fileSet {
		newpath := filepath.Join(binFolder, filepath.Base(f))
		fmt.Println("from ", f, " to ", newpath)
		err := os.Rename(f, newpath)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func getBinFolder() string {
	cfg := config.LoadSetupSpecsConfig()
	binFolderPath := cfg.ExpandedBinFolder()
	err := os.MkdirAll(binFolderPath, os.ModePerm)
	if err != nil {
		println("Error creating config folder...")
		println("-->", binFolderPath)
		println("-->", cfg.BinFolder)
		println(err.Error())
	}
	return binFolderPath
}
