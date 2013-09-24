package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Host struct {
	Ip     string
	Mac    string
	Subnet string
}

func checkLine(inputLine string) (string, bool) {
	inputLine = strings.TrimSpace(inputLine)

	if !strings.HasPrefix(inputLine, "#") {
		if len(inputLine) > 0 {
			return inputLine, true
		}
	}

	return "", false
}

func openFile(fileName string) *os.File {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	/*
		defer func() {
			if err := file.Close(); err != nil {
				panic(err)
			}
		}()
	*/

	return file
}

func parseFile(file *os.File, subnet string) map[string]Host {
	hostMap := make(map[string]Host)
	var checkedLine string
	var goodLine bool

	//opening file reader buffer
	fileLine := bufio.NewScanner(file)

	//fmt.Println("***")
	/*
	 * Big Ugly Parsing loop, surprisingly only Big O(n^2)
	 */
	for fileLine.Scan() {
		if err := fileLine.Err(); err != nil {
			panic(err)
		}
		//fmt.Println(strings.TrimSpace(fileLine.Text()))
		//check that line has information in it
		checkedLine, goodLine = checkLine(fileLine.Text())
		if goodLine {
			lineFields := strings.Fields(checkedLine)
			if lineFields[0] == "host" {
				hostName := lineFields[1]
				var hostIp string
				var hostMac string
				var hostSubnet string
				//hostMap[hostName] = Host{"10.10.0.1", "0:00:00:00:00:00"}
				//fmt.Println("Host:", hostName)
				fileLine.Scan()
				checkedLine, goodLine = checkLine(fileLine.Text())
				for !strings.Contains(checkedLine, "}") {
					if goodLine {
						hostFields := strings.Fields(checkedLine)
						switch hostFields[0] {
						case "hardware":
							hostMac = hostFields[2]
							hostMac = hostMac[:len(hostMac)-1]
							tempMac := strings.Split(hostMac, ":")
							hostMac = strings.Join(tempMac, "")
							//fmt.Println("Mac:", hostMac)
						case "fixed-address":
							hostIp = hostFields[1]
							hostIp = hostIp[:len(hostIp)-1]
							hostSubnet = subnet
							//hostSubnet = "0.0.0.0"
							//fmt.Println("Ip:", hostIp)
						}
					}
					fileLine.Scan()
					checkedLine, goodLine = checkLine(fileLine.Text())
				}
				hostMap[hostName] = Host{hostIp, hostMac, hostSubnet}
				//fmt.Println("***")
			}
		}
	}

	return hostMap
}

/*
 * Format the output to my liking
 */
func printMap(hostMap map[string]Host, dhcpServer string) {
	//fmt.Println("***")
	for key, value := range hostMap {
		//fmt.Println("Host:", key)
		//fmt.Println("Ip:", value.Ip)
		//fmt.Println("Mac:", value.Mac)
		//fmt.Println("Sub:", value.Subnet)
		//fmt.Println("***")
		//Dhcp Server 10.0.0.20 Scope 10.0.10.0 Add reservedip 10.0.10.21 001122334455 "Server1" "" "DHCP"
		fmt.Printf("dhcp server %s Scope %s Add reservedip %s %s \"%s\" \"%s\" \"DHCP\"\n", dhcpServer, value.Subnet, value.Ip, value.Mac, key, key)
	}
}

func main() {
	file := openFile(os.Args[1])
	subnet := os.Args[2]
	dhcpServer := os.Args[3]
	hostMap := parseFile(file, subnet)
	printMap(hostMap, dhcpServer)
	file.Close()
}
