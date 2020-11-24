package main

import(
	"fmt"
	"flag"
)


// 相同底层类型的变量之间是可以转换的

/*接口型函数，指的是用函数实现接口，
这样在调用的时候就会非常简便，我称这种函数，为接口型函数，这种方式使用于只有一个函数的接口。
对于只有一个方法的接口，每次都要实现这个接口，
过于繁琐，所以通常通过接口型函数，来达到简化的目的，直接上代码
*/
type introduce interface {
	say(name string, age int)
}

type introduceFunc func(name string, age int)

func (f introduceFunc) say(name string, age int) {
	f(name, age)
}

func chineseSay(name string, age int) {
	fmt.Printf("大家好%s,今年%d岁\n", name, age)
}

func englishSay(name string, age int) {
	fmt.Printf("hi erveryone My name is %s,and I'm %d years old\n", name, age)
}
func each(m map[string]int, h introduce) {
	if m != nil && len(m) > 0 {
		for name, age := range m {
			h.say(name, age)
		}
	}
}


var (
	name string
	age int
	address *string
	id *int
)

func init()  {

	// 通过传入变量地址的方式，绑定命令行参数到string变量


	flag.StringVar(&name, // 第一个参数，存放值的参数地址
		"name", // 第二个参数，命令行参数的名称
		"匿名", // 第三个参数，命令行不输入时的默认值
		"您的姓名") // 第四个参数，该参数的描述信息，help命令时的显示


	flag.IntVar(&age,
		"age",
		-1,
		"您的年龄")

	address = flag.String("address","未知","您的住址")

	id = flag.Int("id",-1,"身份ID")



}
func main() {
	students := map[string]int{"hu": 12, "len": 13, "jim": 31}

	var handler introduceFunc = chineseSay
	each(students, handler)

	handler = englishSay
	each(students, handler)

	fmt.Println("----------------------")

	var a int = 2
	a = a << 10

	fmt.Print("左移的结果为")
	fmt.Println(a)

	var b int = 1024
	b = b >> 10
	fmt.Print("右移的结果")
	fmt.Println(b)


	flag.Parse()


	fmt.Printf("%s您好，您的年龄：%d，您的住址:%s,您的ID:%d\n\n",name,age,*address,*id)



}
