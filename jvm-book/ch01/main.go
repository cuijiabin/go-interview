package main

import "fmt"

/*
1.根据命令行参数直接生成cmd对象
2.根据系统的输入参数判断条件
3.版本信息需要打印版本
4.帮助信息需要显示帮助的信息
5.启动jvm
*/
func main() {
	cmd := parseCmd()

	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	fmt.Printf("classpath:%s class:%s args:%v\n",
		cmd.cpOption, cmd.class, cmd.args)
}

/*
第一章总结 编写了一个简化版的命令工具
*/
