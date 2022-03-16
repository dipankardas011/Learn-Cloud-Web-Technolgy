package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
//func (ii *IPAddr) String() string {
//	str := fmt.Sprintf("%v.%v.%v.%v", ii[0], ii[1], ii[2], ii[3])
//	return str
//}

func (ii IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ii[0], ii[1], ii[2], ii[3])
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)  // here it required a value called variable function to accept
		fmt.Printf("%v: %v\n", name, &ip) // here as we are using &ip we need pointer function to accept
	}
}
