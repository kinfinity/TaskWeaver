/*
Copyright © 2023 kokou.egbewatt@gmail.com

	TASKWEAVER CLIENT -> WEAVER
*/
package main

//

import (
	"fmt"
	"os"
	"taskweaver/pkg/weaver"
)

func main() {
	// if args < 1 print usage
	if len(os.Args) == 1 {
		PrintUsage()
		os.Exit(1)
	}
	// Create new weaver
	weaverInstance := weaver.GetWeaverInstance()
	if err := weaverInstance.Run(); err != nil {
		weaverInstance.Logger().Fatalf("%s", err)
	}
}

func PrintUsage() {
	fmt.Println(
		`USAGE:

	
	`)
}
