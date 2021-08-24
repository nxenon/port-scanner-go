# Port-Scanner-GO
[![Tool Category](https://badgen.net/badge/Tool/Port%20Scanner/black)](https://github.com/nxenon/port-scanner-go)
[![APP Version](https://badgen.net/badge/Version/Beta/red)](https://github.com/nxenon/port-scanner-go)
[![Go Version](https://badgen.net/badge/Go/1.13/blue)](https://golang.org/doc/go1.13)
[![License](https://badgen.net/badge/License/GPLv2/purple)](https://github.com/nxenon/port-scanner-go/blob/master/LICENSE)


Simple TCP port scanner in golang.

Installation & Build
----
    You have to have GO version 1.13
    run:
    go build port-scanner-go.go

Run
----

    single port:
        ./port-scanner-go --ip 192.168.1.1 --port 80
    ports range:
        ./port-scanner-go --ip 192.168.1.1 --port 1-1024
    specific ports:
        ./port-scanner-go --ip 192.168.1.1 --port 80,443,22


Help
----
    usage: port-scanner-go [-h|--help] --ip "<value>" --port "<value>" --timeout <integer>
    
                           Start Port Scanner
    
    Arguments:
    
    -h  --help     Print help information
    --ip       Target IP
    --port     Ports Range e.g 80 or 1-1024 or 80,22,23
    --timeout  Timeout in Millisecond --> [Default : 500 ms]