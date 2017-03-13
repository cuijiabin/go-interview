package main

import "fmt"
import "strings"
import "jvm-book/ch03/classfile"
import "jvm-book/ch03/classpath"

type Cmd struct {
	helpFlag    bool
	versionFlag bool
	cpOption    string
	XjreOption  string
	class       string
	args        []string
}

func main() {
	cmd := &Cmd{XjreOption: "D:\\Program Files\\Java\\jdk1.7.067\\jre", class: "java.lang.String"}
	startJVM(cmd)
}

func startJVM(cmd *Cmd) {
	//解析类路径
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	className := strings.Replace(cmd.class, ".", "/", -1)
	fmt.Println(cmd.class)
	//重点代码 分为两步 1.读取 2.转化
	cf := loadClass(className, cp)

	//打印类信息 Ctrl+Q 查看方法说明
	printClassInfo(cf)
}

func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
	classData, _, err := cp.ReadClass(className)
	//字节码数字
	//fmt.Println("class data:%v\n", classData)
	if err != nil {
		panic(err)
	}

	//转化字节码
	cf, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}

	return cf
}

func printClassInfo(cf *classfile.ClassFile) {
	fmt.Printf("version: %v.%v\n", cf.MajorVersion(), cf.MinorVersion())
	//for idx, v := range cf.ConstantPool() {
	//	//类型判断
	//	if ini, ok := v.(classfile.ConstantIntegerInfo); ok {
	//		fmt.Println(idx, ini.Value())
	//	}
	//
	//}

	fmt.Printf("constants count: %v\n", len(cf.ConstantPool()))
	fmt.Printf("access flags: 0x%x\n", cf.AccessFlags())
	fmt.Printf("this class: %v\n", cf.ClassName())
	fmt.Printf("super class: %v\n", cf.SuperClassName())
	fmt.Printf("interfaces: %v\n", cf.InterfaceNames())
	//fmt.Printf("fields count: %v\n", len(cf.Fields()))
	//for _, f := range cf.Fields() {
	//	fmt.Printf("  %s\n", f.Name())
	//}
	//fmt.Printf("methods count: %v\n", len(cf.Methods()))
	//for _, m := range cf.Methods() {
	//	fmt.Printf("  %s\n", m.Name())
	//}
}
