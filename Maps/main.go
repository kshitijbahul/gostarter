package main

import (
	"fmt"
)

func main() {
	colours := map[string]string{
		"red":   "#ff000",
		"green": "#ff745",
	}
	//var colours map[string]string
	//colours := make(map[string]string)
	colours["white"] = "##0000"
	colours["10"] = "##mkasdasm"

	//fmt.Println(colours)
	parseMap(colours)
	delete(colours, "10")
	fmt.Println(colours)
}

func parseMap(colour map[string]string) {
	for key, value := range colour {
		fmt.Println(key, value)
	}
}
