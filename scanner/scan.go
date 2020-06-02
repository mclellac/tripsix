package scanner

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/briandowns/spinner"
	"golang.org/x/sync/semaphore"
)

// PortScanner struct
type PortScanner struct {
	IP   string
	Lock *semaphore.Weighted
}

// Ulimit gets the maximum number of open files allowed by the OS.
func Ulimit() int64 {
	out, err := exec.Command("bash", "-c", "ulimit -n").Output()
	if err != nil {
		panic(err)
	}

	s := strings.TrimSpace(string(out))

	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}

	return i
}

// ScanPort scans the target IP/port to determine if it's open or closed.
func ScanPort(ip string, port int, timeout time.Duration) {
	target := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.DialTimeout("tcp", target, timeout)

	// fmt.Sprintf("%s:%d", ip, port)

	if err != nil {
		if strings.Contains(err.Error(), "too many open files") {
			time.Sleep(timeout)
			ScanPort(ip, port, timeout)
		}
		return
	}

	conn.Close()
	fmt.Printf("%-18d %-7s %20s\n", port, "open", DescribePort(port))
}

// Start creates a channel to begin the port scan.
func (ps *PortScanner) Start(first, last int, timeout time.Duration) {
	wg := sync.WaitGroup{}
	defer wg.Wait()

	banner := `
████████╗██████╗ ██╗██████╗ ███████╗██╗██╗  ██╗
╚══██╔══╝██╔══██╗██║██╔══██╗██╔════╝██║╚██╗██╔╝
   ██║   ██████╔╝██║██████╔╝███████╗██║ ╚███╔╝
   ██║   ██╔══██╗██║██╔═══╝ ╚════██║██║ ██╔██╗
   ██║   ██║  ██║██║██║     ███████║██║██╔╝ ██╗
   ╚═╝   ╚═╝  ╚═╝╚═╝╚═╝     ╚══════╝╚═╝╚═╝  ╚═╝


`
	sep := strings.Repeat("-", 47)

	fmt.Print(banner)

	fmt.Printf("Scanning %v:%v-%v\n\n", ps.IP, first, last)

	fmt.Printf("%-18s %-7s %20s\n", "PORT", "STATE", "SERVICE")
	fmt.Println(sep)

	s := spinner.New(spinner.CharSets[25], 100*time.Millisecond)
	s.Suffix = " Scanning ports..."
	s.Writer = os.Stderr

	if err := s.Color("red", "bold"); err != nil {
		log.Fatalln(err)
	}

	s.Start()

	for port := first; port <= last; port++ {
		err := ps.Lock.Acquire(context.TODO(), 1)
		if err != nil {
			fmt.Print(err)
		}

		wg.Add(1)
		go func(port int) {
			defer ps.Lock.Release(1)
			defer wg.Done()
			ScanPort(ps.IP, port, timeout)
		}(port)
	}

	s.Stop()
	println("")
}
