package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


// Host stores information pertaining to a single network computer.
// This information includes the machines IP address, MAC address, 
// and subnet.
type Host struct {
	Ip     string
	Mac    string
	Subnet string
}

// checkLine takes in a single line from the input file and checks
// that it is not a comment; a line starting with the hash (#) mark.
// If the line is not a comment line then it is returned with a true,
// otherwise an empty string is returned with a false.
func checkLine(inputLine string) (string, bool) {
	inputLine = strings.TrimSpace(inputLine)

	if !strings.HasPrefix(inputLine, "#") {
		if len(inputLine) > 0 {
			return inputLine, true
		}
	}

	return "", false
}

// openFile takes in a string filename and attempts to open it as
// a path to a file.
// The file pointer is returned when it is opened correctly.
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

// parseFile takes in a file pointer and a user defined string describing
// the subnet. There is not enough information in the input file to correctly
// infer the subnet so we have to have the user input that manually. 
// The input file is parsed to find all the host configurations, this information
// is stored in a map and returned once the file has been fully parsed.
func parseFile(file *os.File, subnet string) map[string]Host {
	hostMap := make(map[string]Host)
	var checkedLine string
	var goodLine bool

	//opening file reader buffer
	fileLine := bufio.NewScanner(file)

	//fmt.Println("***")
	/*
	 * Big Ugly Parsing loop, surprisingly only Big O(n)
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

// printMap takes in the map returned from our parseFile function and
// walked through each key,value pair printing it out.
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
