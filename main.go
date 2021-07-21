package main

import (
	"github.com/atlas-dns/apis"
	"github.com/atlas-dns/common"
)

func main() {
	apis.SetupEnvironment()
	common.SetUpLogging()
	apis.InitialiseApplication()
}
