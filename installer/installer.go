package installer

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/Eldius/tools-setup-go/logger"
	"github.com/Eldius/tools-setup-go/specs"
)

/*
Install will install the spec passad as parameter
*/
func Install(specName string) {
	for i, s := range specs.ListSpecs() {
		fmt.Println(i, " => ", s.Name, ": ", s.Version)
	}
	//destFile := downloadFile(spec.URL, specName)
	//extractedFiles, err := UnpackFile(destFile)
	//if err != nil {
	//	panic(err.Error())
	//}

	//MoveToBinFolder(extractedFiles)
}

func downloadFile(url string, packageName string) string {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Failed to download package from ", url)
		panic(err.Error())
	}
	defer resp.Body.Close()

	logger.DebugInterface(resp.Header)

	filepath := fmt.Sprintf("/tmp/%s%s", packageName, getFileExtension(resp.Header))

	logger.Info("- download from ", url, " to ", filepath)

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		panic(err.Error())
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		panic(err.Error())
	}
	return filepath
}

func getFileExtension(header http.Header) string {
	switch header["Content-Type"][0] {
	case "application/zip":
		return ".zip"
	default:
		return ""
	}
}
