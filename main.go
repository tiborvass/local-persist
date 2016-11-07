package main

import (
	"fmt"
	"os"

	"github.com/docker/go-plugins-helpers/volume"
)

func main() {
	driver := newLocalPersistDriver(os.Args[1])

	handler := volume.NewHandler(driver)
	fmt.Println(handler.ServeUnix("root", "plugin"))
}
