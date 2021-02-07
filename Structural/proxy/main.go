package main

func main() {
	driver := &Driver{Age: 15}

	carProxy := NewCarProxy(driver)

	carProxy.Drive()
}
