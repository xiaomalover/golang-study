package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	for index, value := range os.Args {
		fmt.Printf("args[%d]:%s\n", index, value)
	}

	//flag.type()
	name := flag.String("name", "张三", "姓名")
	age := flag.Int("age", 18, "年龄")
	married := flag.Bool("married", false, "婚否")
	delay := flag.Duration("delay", 0, "时间间隔")
	//解析命令行参数
	flag.Parse()
	fmt.Println(*name, *age, *married, *delay)
	//返回命令行参数后的其他参数
	fmt.Println(flag.Args())
	//返回命令行参数后的其他参数个数
	fmt.Println(flag.NArg())
	//返回使用的命令行参数个数
	fmt.Println(flag.NFlag())
	/**
		运行命令
		Run  go run command_args.go  -name=找人 -age=1 -married=true -delay=22s foo bar

		返回结果
		args[0]:C:\Users\Dell\AppData\Local\Temp\go-build779834302\b001\exe\command_args.exe
		args[1]:-name=找人
		args[2]:-age=1
		args[3]:-married=true
		args[4]:-delay=22s
		args[5]:foo
		args[6]:bar

		找人 1 true 22s
		[foo bar]
		2
		4
	*/

	//上面的等价于下面
	//flag.typeVar()
	/*var name1 string
	var age1 int
	var married1 bool
	var delay1 time.Duration

	flag.StringVar(&name1, "name1", "赵六", "姓名1")
	flag.IntVar(&age1, "age1", 10, "年龄1")
	flag.BoolVar(&married1, "married1", true, "婚否1")
	flag.DurationVar(&delay1, "delay1", time.Second * 2, "时间间隔1")
	flag.Parse()
	fmt.Println(name1, age1, married1, delay1)*/
}
