// --Summary:
//
//	Write a program to display server status.
//
// --Requirements:
// * Create a function to print server status, including:
//   - Number of servers
//   - Number of servers for each status (Online, Offline, Maintenance, Retired)
//
// * Store the existing slice of servers in a map
// * Default all of the servers to `Online`
// * Perform the following status changes and display server info:
package main

import "fmt"

const (
	Online = iota
	Offline
	Maintenance
	Retired
)

var statusText = map[int]string{
	Online:      "Online",
	Offline:     "Offline",
	Maintenance: "Maintenance",
	Retired:     "Retired",
}

var p = fmt.Printf

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}
	status := map[string]int{}

	for i := 0; i < len(servers); i++ {
		status[servers[i]] = Online
	}

	DisplayAllServers := func() {
		for _, server := range servers {
			p("server \"%s\" is %v\n", server, statusText[status[server]])
		}
	}

	//  - display server info
	//  - change `darkstar` to `Retired`
	//  - change `aiur` to `Offline`
	//  - display server info
	//  - change all servers to `Maintenance`
	//  - display server info

	status["darkstar"] = Retired
	status["aiue"] = Offline

	DisplayAllServers()

	for i := 0; i < len(servers); i++ {
		status[servers[i]] = Maintenance
	}

	DisplayAllServers()
}
