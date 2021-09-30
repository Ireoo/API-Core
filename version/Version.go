package version

import (
	"flag"
	"fmt"

	Info "github.com/Ireoo/API-Core/info"
)

var (
	ver bool
)

func init() {
	flag.BoolVar(&ver, "v", false, "版本信息")
	flag.Parse()

	if ver {
		//fmt.Printf(`API-Core version: %s`, version)
		fmt.Printf("API-Core version: %s\nbuild time: %s\n", Info.Version, Info.BuildTime)
		return
	}

	fmt.Printf("API-Core version: %s\nbuild time: %s\n", Info.Version, Info.BuildTime)
	fmt.Println("")
	fmt.Println("")
}
