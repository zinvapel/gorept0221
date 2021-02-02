package main

import "fmt"

type Address struct {
	Street string
	Building string
}

type Vendor struct {
	Name string
}

type Order struct {
	Id string
	DeliveryAddress *Address
	Vendor Vendor
}

func main() {
	o := Order{}
	fmt.Printf("'%p' '%s'\n", o.DeliveryAddress, o.Vendor.Name)
}
