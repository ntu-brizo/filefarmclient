/*
Package filefarmclient ...

*/
package filefarmclient

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/ntu-brizo/filefarmclient/config"
)

// UploadArg describes the structure of requesting arguments of `Upload` API, including accesing key and file content.
type UploadArg struct {
	key     string
	content *[]byte
}

// DownloadArg describes the structure of requesting arguments of `Download` API, including accesing key.
type DownloadArg struct {
	key string
}

// randomFarmer returns one random farmer record by parsing config file
func randomFarmer() (farmer config.FarmerRecord, err error) {
	var conf config.Config
	conf, err = config.Parse()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	rand.Seed(time.Now().UnixNano())
	farmer = conf.Farmers[rand.Intn(len(conf.Farmers)-1)]
	return
}

// Hello says hello; just for testing
func Hello() (greet string, err error) {
	greet = "hello"
	return
}

// Upload uploads a file with key and content to FileFarm.
func Upload(arg UploadArg) (err error) {
	return
}

// Download downloads file content from FileFarm given its key.
func Download(arg DownloadArg) (content *[]byte, err error) {
	return
}
