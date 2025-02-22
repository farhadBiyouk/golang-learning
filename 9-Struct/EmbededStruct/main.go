package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	Name   string
	Color  string
	Length int
	Width  int
	Weight int
	Price  int
	Brand  string
	MadeIn string
}

type ElectronicProduct struct {
	Product
	Ram                    int
	Cpu                    string
	ScreenSize             float32
	OperatingSystem        string
	OperatingSystemVersion string
}

type Mobile struct {
	ElectronicProduct
	SimcardCapacity int
	SimcardType     string
	NetworkType     string
	CameraCount     int
}

type Loptop struct {
	ElectronicProduct
	UsbPortCount int
	KeyboardType string
	HasCdRom     bool
}

func main() {

	mobile := Mobile{}
	mobile.Name = "Samsung M4150"
	mobile.Brand = "Samsung"
	mobile.Color = "Blue"
	mobile.Ram = 4
	mobile.OperatingSystem = "Android"
	mobile.OperatingSystemVersion = "11"

	mobile.CameraCount = 2
	mobile.NetworkType = "5G"
	mobile.SimcardType = "Nano"
	mobile.SimcardCapacity = 3

	laptop := Loptop{}

	laptop.Name = "MSI GX 4570"
	laptop.Brand = "MSI"
	laptop.Color = "Black"
	laptop.Ram = 16
	laptop.OperatingSystem = "Ubuntu"
	laptop.OperatingSystemVersion = "22.04"

	laptop.HasCdRom = false
	laptop.UsbPortCount = 3
	laptop.KeyboardType = "Light"

	laptopJson, _ := json.Marshal(laptop)
	mobileJson, _ := json.Marshal(mobile)

	fmt.Println(string(laptopJson))
	fmt.Println(string(mobileJson))

}
