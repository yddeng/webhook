package client

import (
	"fmt"
	"os/exec"
)

func Command(path string) {
	cmd := exec.Command("/bin/bash", "-c", path)

	//err := cmd.Run()
	//if err != nil {
	//	fmt.Println("Execute Command failed:" + err.Error())
	//}

	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Execute Shell:%s failed with error:%s \n", path, err.Error())
		return
	}
	fmt.Printf("Execute Shell: %s finished with output:\n %s ", path, string(output))
}
