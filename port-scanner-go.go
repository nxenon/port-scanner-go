package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)


// set global vars
var targetIP string
var portsList []string
var timeoutTCP time.Duration


func main() {

	parseArgs()

	startScanning()

}

func parseArgs(){

	parser := argparse.NewParser("port-scanner-go", "Start Port Scanner")

	ip_arg := parser.String("", "ip", &argparse.Options{Required: true, Help: "Target IP"})

	ports_arg := parser.String("", "port", &argparse.Options{Required: true,
		Help: "Ports Range e.g 80 or 1-1024 or 80,22,23"})

	timeout := parser.Int("", "timeout", &argparse.Options{Required: true,
		Help: "Timeout in Millisecond --> [Default : 500 ms]", Default: 500})

	parser.Parse(os.Args)

	if *ip_arg == ""{
		println("Enter Target IP")
		println(os.Args[0] + " -h for Help")
		os.Exit(0)

	}

	if *ports_arg == ""{
		println("Enter Port")
		println(os.Args[0] + " -h for Help")
		os.Exit(0)

	}

	targetIP = *ip_arg

	timeoutTCP = time.Duration(*timeout) * time.Millisecond

	ports_list := getPortsList(*ports_arg)

	portsList = ports_list

}

func getPortsList(
	port_var string,
	) (
	[]string,
	){

	// if port argument is like : 22,80,23
	if strings.Contains(port_var, ","){
		ports_list := strings.Split(port_var, ",")

		for p:= range ports_list{
			_, err_c := strconv.Atoi(ports_list[p])
			if err_c != nil{
				fmt.Printf("Invalid Port : %s\n", ports_list[p])
				os.Exit(0)

			}
		}

		return ports_list

	} else if strings.Contains(port_var, "-"){
		// if port argument is like : 1-1024

		port_min_and_max := strings.Split(port_var, "-")

		port_min, err := strconv.Atoi(port_min_and_max[0])
		if err != nil{
			println("Invalid Port Min Range : " + port_min_and_max[0])
			os.Exit(1)

		}

		port_max, err2 := strconv.Atoi(port_min_and_max[1])
		if err2 != nil{
			println("Invalid Port Max Range : " + port_min_and_max[1])
			os.Exit(1)

		}

		var ports_temp_list []string

		for p_min := port_min; p_min <= port_max; p_min++ {
			port_str := strconv.Itoa(p_min)
			ports_temp_list = append(ports_temp_list, port_str)

		}

		return ports_temp_list

	}

	// if port is single number like : 80
	_, err := strconv.Atoi(port_var) // check if port is correct (int)
	if err != nil{
		println("Invalid Port : " + port_var)
		os.Exit(1)

	}

	return []string{port_var}

}

func startScanning(){

	for p:= range portsList{

		go scanPort(portsList[p])

	}

	// wait 1 sec more than timeout for finishing go routines
	time.Sleep(timeoutTCP + (1000 * time.Millisecond))

}

func scanPort(
	port string,
	){

	d := net.Dialer{Timeout: timeoutTCP}
	_, err := d.Dial("tcp", targetIP + ":" + port)
	if err != nil {
		if add_err, ok := err.(*net.AddrError); ok {
			if add_err.Timeout(){
				return

			}
		} else if add_err, ok := err.(*net.OpError); ok {

			// handle lacked sufficient buffer space error

			if strings.TrimSpace(add_err.Err.Error()) == "bind: An operation on a socket could not be performed because " +
				"the system lacked sufficient buffer space or because a queue was full." {

				time.Sleep(timeoutTCP + (3000 * time.Millisecond))

				_, err_ae := d.Dial("tcp", targetIP + ":" + port)

				if err_ae != nil {
					if add_err, ok := err.(*net.AddrError); ok {
						if add_err.Timeout() {
							return

						}
					}
				}
			}

		} else{
			println(err.Error())
			os.Exit(1)

		}

		return

	}

	fmt.Printf("[+] Port %s/TCP is open\n", port)
	
}