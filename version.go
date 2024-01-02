package main

import "fmt"

var appVersion = "v0.0.0"
var appBuild = "0000000"

func printver() {
	fmt.Printf("GoRevoke %s\n", appVersion)
	fmt.Printf("Build Number: %s\n", appBuild)
}
