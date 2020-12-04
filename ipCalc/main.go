package ipCalc

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/brotherpowers/ipsubnet"
)

func calculateIp() {
	for {
		var ip string
		fmt.Println("Enter IP address")
		fmt.Scanln(&ip)
		if ip == "-1" {
			break
		}

		var mask string
		fmt.Println("Enter network mask address")
		fmt.Scanln(&mask)

		ips := strings.SplitN(ip, ".", 4)
		masks := strings.SplitN(mask, ".", 4)

		var binIpNum string
		var binMaskNum string
		for i := 0; i < 4; i++ {
			n, _ := strconv.ParseInt(ips[i], 10, 64)
			bin := strconv.FormatInt(n, 2)
			initialLen := len(bin)
			for i := 0; i < 8-initialLen; i++ {
				bin += "0"
			}
			binIpNum = binIpNum + bin

			m, _ := strconv.ParseInt(masks[i], 10, 64)
			bin = strconv.FormatInt(m, 2)
			initialLen = len(bin)
			for i := 0; i < 8-initialLen; i++ {
				bin += "0"
			}
			binMaskNum = binMaskNum + bin
		}

		numberOfFreeDigits := 0
		for i := len(binMaskNum) - 1; i >= 0; i-- {
			if binMaskNum[i] == '1' {
				numberOfFreeDigits = len(binMaskNum) - i - 1
				break
			}
		}

		//networkNumber := ipNum & maskNum

		sub := ipsubnet.SubnetCalculator(ip, 32-numberOfFreeDigits)

		fmt.Println("-----------------------------")
		fmt.Println("IP:   " + sub.GetIPAddressBinary())
		fmt.Println("Mask: " + sub.GetSubnetMaskBinary())
		if sub.GetIPAddressBinary()[0] == 0 {
			fmt.Println("Network class: A")
		}
		if sub.GetIPAddressBinary()[0] == 1 && sub.GetIPAddressBinary()[1] == 0 {
			fmt.Println("Network class: B")
		}
		if sub.GetIPAddressBinary()[0] == 1 && sub.GetIPAddressBinary()[1] == 1 && sub.GetIPAddressBinary()[2] == 0 {
			fmt.Println("Network class: C")
		}

		fmt.Println()
		fmt.Println("Network: " + sub.GetNetworkPortion())
		fmt.Println("Host: " + sub.GetHostPortion())

		fmt.Println()
		fmt.Println("Number of usable addresses: " + strconv.Itoa(sub.GetNumberAddressableHosts()))
		machineNumber, _ := strconv.ParseInt(sub.GetHostPortionBinary(), 2, 64)
		fmt.Println("Current machine address: " + strconv.Itoa(int(machineNumber)))

		fmt.Println("-----------------------------------------------------")
		fmt.Println()
	}
}

func main() {
	calculateIp()
}
