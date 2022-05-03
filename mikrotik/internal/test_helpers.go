package internal

import (
	"fmt"
)

var ipCounter uint = 0xC0A80001      // 192 = C0, 168 = A8, 0 = 00, 1 = 01
var ipRangeCounter uint = 0xAC100001 // 172 = AC, 16 = 10, 0 = 00, 1 = 01
var macCounter = 0
var dnsCounter = 0

func GetNewIpAddr() string {
	ipCounter++
	return formatIPv4(ipCounter)
}

func GetNewIpAddrRange(count uint) string {
	var ipRangeStart = ipRangeCounter + 1
	ipRangeCounter = ipRangeCounter + count
	return fmt.Sprintf("%s-%s", formatIPv4(ipRangeStart), formatIPv4(ipRangeCounter))
}

func GetNewMacAddr() string {
	macCounter++

	if macCounter > 255 {
		macCounter = 1
	}

	return fmt.Sprintf("01:23:45:67:89:%02x", macCounter)
}

func GetNewDnsName() string {
	dnsCounter++
	return fmt.Sprintf("dns-%02d.terraform", dnsCounter)
}

func formatIPv4(ipAddr uint) string {
	return fmt.Sprintf("%d.%d.%d.%d", (ipAddr>>24)&0xFF, (ipAddr>>16)&0xFF, (ipAddr>>8)&0xFF, ipAddr&0xFF)
}
