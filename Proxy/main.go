package main

import "proxy/objects"

func main() {
	driver := &objects.Driver{Age: 15}

	carProxy := objects.NewCarProxy(driver)

	carProxy.Drive()
}
