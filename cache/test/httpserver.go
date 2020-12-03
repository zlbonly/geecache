package main

import (
	"flag"
	"fmt"
	"log"
	"reflect"
	"strings"
	"sync"
	"time"
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
	name    string
	age     int
	address *string
	id      *int
)

func init() {

	// 通过传入变量地址的方式，绑定命令行参数到string变量

	flag.StringVar(&name, // 第一个参数，存放值的参数地址
		"name", // 第二个参数，命令行参数的名称
		"匿名",   // 第三个参数，命令行不输入时的默认值
		"您的姓名") // 第四个参数，该参数的描述信息，help命令时的显示

	flag.IntVar(&age,
		"age",
		-1,
		"您的年龄")

	address = flag.String("address", "未知", "您的住址")

	id = flag.Int("id", -1, "身份ID")

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

	fmt.Printf("%s您好，您的年龄：%d，您的住址:%s,您的ID:%d\n\n", name, age, *address, *id)

	var wg sync.WaitGroup

	typ := reflect.TypeOf(&wg)

	for i := 0; i < typ.NumMethod(); i++ {
		method := typ.Method(i)
		argv := make([]string, 0, method.Type.NumIn())

		returns := make([]string, 0, method.Type.NumOut())

		for j := 1; j < method.Type.NumIn(); j++ {
			argv = append(argv, method.Type.In(j).Name())
		}

		for j := 0; j < method.Type.NumOut(); j++ {
			returns = append(returns, method.Type.Out(j).Name())
		}

		log.Printf("func (w*%s),%s(%s) %s",
			typ.Elem().Name(),
			method.Name,
			strings.Join(argv, ","),
			strings.Join(returns, ","))
	}

	var bc int = 1024

	value := reflect.ValueOf(bc).Int()

	fmt.Println(value)

	// 可以从一个任何非接口类型的值创建一个 reflect.Type 值
	fmt.Println(reflect.TypeOf(bc)) // 获取变量bc的反射值对象

	fmt.Println(reflect.TypeOf(bc).Name())
	// 声明一个空结构体
	type cat struct {
		age   int
		name  string
		score float64
	}

	s := cat{1, "zlb", 13.0}
	typeOfCat := reflect.TypeOf(s)
	fmt.Println(typeOfCat)        // 获取结构体实例的反射类型对象
	fmt.Println(typeOfCat.Name()) // 反射类型对象的名称
	fmt.Println(typeOfCat.Kind()) // 反射类型对象的种类
	fmt.Println(typeOfCat.NumField())

	typeOfCatValue := reflect.ValueOf(s)
	for i := 0; i < typeOfCat.NumField(); i++ {
		key := typeOfCat.Field(i)
		value := typeOfCatValue.Field(i)
		fmt.Printf("第%d个字段是：%s:%v = %v \n", i+1, key.Name, key.Type, value)
	}

	/**

		场景：假设业务中需要调用服务接口A，要求超市时间为5秒，

	time.After()表示time.Duration长的时候后返回一条time.Time类型的通道消息。那么，基于这个函数，就相当于实现了定时器，且是无阻塞的。
	*/
	ch := make(chan string)
	go func() {
		fmt.Println("go func start .....")
		time.Sleep(time.Second * 2)
		ch <- "result"
	}()

	log.SetFlags(1)
	log.Printf("1111111111")
	select {
	case res := <-ch:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout")
	}

}
