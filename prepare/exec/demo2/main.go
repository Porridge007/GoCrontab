package main

import (
	"os/exec"
	"fmt"
)

func main(){
	var(
		cmd *exec.Cmd
		output []byte
		err error
	)

	//生成cmd
	cmd = exec.Command("/bin/bash","-c","sleep 5;ls -l")
	//执行命令，捕获子进程输出（pipe）
	if output,err = cmd.CombinedOutput();err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Println(output)
}
