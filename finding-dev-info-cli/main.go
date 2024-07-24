package main

import (
	"bufio"
	"encoding/json"
	"finding-dev-info-cli/interfaces"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	deviceChannel := make(chan []interfaces.Device)
	go read_file(deviceChannel)

	devices := <-deviceChannel // Receive devices slice from the channel

	for {
		fmt.Println("\nWhat would you like to search? (e.g. pacemaker, ICD, Azure) | type exit to quit\n")
		in := bufio.NewReader(os.Stdin)
		s, err := in.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		s = strings.TrimSpace(s)
		s = strings.ToUpper(s)
		if s == "EXIT" {
			break
		} else {
			found := false
			for idx, device := range devices {
				if strings.Index(strings.ToUpper(device.DeviceName), s) == 0 || strings.Contains(strings.ToUpper(device.DeviceType), s) || strings.Contains(strings.ToUpper(device.DeviceModelNumber), s) {
					if idx == 0 {
						fmt.Printf("---------------\n")
					}
					fmt.Printf("Device Type : %+v\n", device.DeviceType)
					fmt.Printf("Device Name : %+v\n", device.DeviceName)
					fmt.Printf("Model Number: %+v\n", device.DeviceModelNumber)
					fmt.Printf("---------------\n")
					found = true
				}
			}
			if !found {
				fmt.Println("No device found")
			}

		}
	}

}

// func read_file() []interfaces.Device {
func read_file(deviceChannel chan<- []interfaces.Device) {

	path := flag.String("path", "data\\device-data.json", "The path for device data file")

	flag.Parse()

	f, err := os.Open(*path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	// write code to parse the device-data.json file into a map
	decoder := json.NewDecoder(r)

	var devices []interfaces.Device
	err = decoder.Decode(&devices)
	if err != nil {
		log.Fatal(err)
	}

	// Now devices is a slice of Device structs populated from the JSON file
	// for _, device := range devices {
	//     fmt.Printf("%+v\n", device)
	// }
	//return devices

	// Assuming the rest of read_file implementation is correct and populates a devices slice
	// Send devices slice back through the channel
	deviceChannel <- devices
}
