package hello

import "fmt"

func SayHi(name string) {
	fmt.Printf("你好%s，我是yangkunzhou。很高兴认识你~\n", name)
}

func ShowVersion() {
	fmt.Printf("当前版本号：V%s", "0.0.2")
}
