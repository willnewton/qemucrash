package main

import (
	"fmt"
	"os/exec"
)

func doPrints() {
	i := 0
	for {
		fmt.Printf("This is print %d\n", i)
		c := exec.Command("ls", "/")
		_, err := c.CombinedOutput()
		if err != nil {
			fmt.Printf("error from command: %s\n", err.Error())
		}
		i++
	}
}

func main() {
	go doPrints()
	doPrints()
}
