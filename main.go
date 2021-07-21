package main

import "github.com/atlas-dns/common"

func main() {
	common.SetupEnvironment()
	common.SetUpLogging()
	common.InitialiseApplication()
}
