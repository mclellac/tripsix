package main

// PortState type
type PortState uint8

const (
	PortUnknown PortState = iota
	PortOpen
	PortClosed
	PortFiltered
)

var DefaultPorts []int

func init() {

	for port := range knownPorts {
		DefaultPorts = append(DefaultPorts, port)
	}
}

// DescribePort returns a service string if it matches a known
// port, otherwise it returns an empty string.
func DescribePort(port int) string {
	if s, ok := knownPorts[port]; ok {
		return s
	}

	return ""
}
