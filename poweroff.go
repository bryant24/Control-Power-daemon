package main

import (
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println("远程关闭目标机")
	arg := []string{"-m", "\\\\192.168.1.41", "-s", "-t", "3"}
	cmd := exec.Command("shutdown", arg...)
	d, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)

		return
	}
	fmt.Println(string(d))
	return
}
