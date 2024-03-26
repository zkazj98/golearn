package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"reflect"
	"sort"
	"sync"
)

const (
	CategoryOne = iota
	CategoryTwo
	CategoryThree
)

const (
	a         = iota * 42
	b float64 = iota * 42
	c         = iota * 42
)

func main() {

	/*var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = ""
	}
	fmt.Println(s)*/
	/*_, b, c := resNum()
	fmt.Println(b)
	fmt.Println(c)*/
	/*zhizhen()*/
	/*v := 1
	z := &v
	fmt.Println(incr(&v))
	incr(&v) // side effect: v is now 2
	fmt.Println(*z)*/ // 通过指针间接赋值
	/*var x, y, z int
	x, y, z = 2, 3, 4
	fmt.Println(x, y, z)*/
	/*str := [...]string{1: "ass", 2: "sdasdassd"}
	sp := str[1:3]
	for i, v := range sp {
		fmt.Println(i)
		fmt.Println(v)
	}*/
	/*	fmt.Println(str)
		var a conv.Celsius = 192.6
		z := conv.CToF(a)
		fmt.Println(z)*/
	/*arr := [...]int{1, 2, 3, 4, 5, 6, 78, 9, 78}
	reverse(arr[:])
	fmt.Println(arr)
	fmt.Printf("arr= %p", &arr)
	reverse(arr[:2])
	fmt.Println(arr)
	fmt.Printf("arr= %p", &arr)
	reverse(arr[2:])
	fmt.Println(arr)
	fmt.Printf("arr= %p", &arr)
	fmt.Println()
	arr1 := arr[2:]
	fmt.Printf("arr1= %p", &arr1)*/
	/*z := make([]int, 3, 4)
	fmt.Println(z)
	fmt.Println(z[:])
	str := "dhashdiaus"
	var arr []rune
	for _, v := range str {
		arr = append(arr, v)
	}
	fmt.Printf("%q\n", arr)*/

	/*	h := map[string]int{
			"1": 90,
		}
		h["dhasu"] = 312

		for k, v := range h {
			fmt.Println(k)
			fmt.Println(v)
		}*/
	/*	a := [...]int{1, 2, 3, 4, 6}
		fmt.Println()
		fmt.Printf("%p", &a)

		b := a[0:2]
		fmt.Println()
		fmt.Printf("%p", &b)

		c := make([]int, 10)
		fmt.Println()
		fmt.Printf("%p", &c)

		fmt.Println()
		copy(c, b)
		fmt.Println(len(c))
		for _, v := range c {
			fmt.Println(v)
		}

		ma := map[string]int{"a": 2}
		for k, v := range ma {
			fmt.Println(k)
			fmt.Println(v)
		}

		mapp := make(map[int]string)
		mapp[1] = "dasd"
		fmt.Println(mapp[1])

		s := struct {
			name string
			age  int
		}{"312", 12}
		fmt.Println(s)

		type User struct {
			name string
			age  int
			sex  bool
		}

		type Student struct {
			user   *User
			number int
		}

		user := &User{name: "zk", age: 12, sex: true}

		fmt.Println(*user)

		student := Student{user: user, number: 3123123}

		fmt.Println(student)

		res := do(add, 1, 2)
		fmt.Println(res)

		var sum = func(a, b int) int {
			return a + b
		}
		fmt.Println(sum)

		defer func() {
			println("dasdasdasd")
		}()

		defer func() {
			println("dasdasdasdaaaaaaa")
		}()

		/*	defer func() {
				if err := recover(); err != nil {
					fmt.Println(err)
				}
			}()

			defer func() {
				panic("cuo le")
			}()*/
	/*err := Errorf("cuole")
	if err == nil {
		fmt.Println(err)
	}*/
	/*var a *int
	a = new(int)
	*a = 100
	fmt.Println(*a)

	var b map[string]int
	b = make(map[string]int)
	b["沙河娜扎"] = 100
	fmt.Println(b)*/
	/*type person struct {
		name string
		city string
		age  int8
	}

	type test struct {
		a, b, c, d int8
	}

	n := &test{
		1, 2, 3, 4,
	}

	z := test{
		1, 2, 3, 4,
	}

	fmt.Printf("n.a=%p \n", &n.a)
	fmt.Printf("n.b=%p \n", &n.b)
	fmt.Printf("n.c=%p \n", &n.c)
	fmt.Printf("n.d=%p \n", &n.d)

	fmt.Printf("z.a=%p \n", &z.a)
	fmt.Printf("z.b=%p \n", &z.b)
	fmt.Printf("z.c=%p \n", &z.c)
	fmt.Printf("z.d=%p \n", &z.d)

	p6 := &person{
		name: "小王子",
		city: "北京",
		age:  18,
	}

	p5 := person{
		name: "小王子",
		city: "北京",
		age:  18,
	}

	fmt.Println(*p6)
	fmt.Println(p5)*/
	//Person 结构体

	//NewPerson 构造函数
	/*arg := NewPerson("zk", 26)
	arg.Setage(100)
	fmt.Println(*arg)

	arg1 := NewPerson("zk1", 25)
	arg1.Dream(12)
	fmt.Println(*arg1)*/
	/*an := &Animal{
		name: "asas",
	}
	d := &Dog{
		Feet:   12,
		animal: *an,
	}
	fmt.Printf("an p :%p", &an)
	fmt.Printf("d p :%p", &d.animal)*/
	/*c := &Class{
		Title:    "101",
		Students: make([]Student, 0, 200),
	}
	for i := 0; i < 10; i++ {
		stu := Student{
			Name:   fmt.Sprintf("stu%02d", i),
			Gender: "男",
			ID:     i,
		}
		c.Students = append(c.Students, stu)
	}
	//JSON序列化：结构体-->JSON格式的字符串
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}
	fmt.Printf("json:%s\n", data)

	//JSON反序列化：JSON格式的字符串-->结构体
	str := `{"Title":"101","Students":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu04"},{"ID":5,"Gender":"男","Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Name":"stu08"},{"ID":9,"Gender":"男","Name":"stu09"}]}`
	c1 := &Class{}
	err = json.Unmarshal([]byte(str), c1)
	if err != nil {
		fmt.Println("json unmarshal failed!")
		return
	}
	fmt.Printf("%#v\n", c1)*/
	/*s1 := &Student{ID: 13}
	data, err := json.Marshal(s1)
	if err != nil {
		fmt.Println("json marshal failed!")
		return
	}
	fmt.Printf("json str:%s\n", data)

	hello.SayHi()*/

	/*b := new(Bird)
	MakeHungry(b)*/
	/*	var x Payer
		x = Zhifubao{}
		x.pay(100)
		x = &Weixin{}
		x.pay(100)*/

	/*a := &reverseIntSlice{}
	aa := a.Len()
	a.Less(1, 2)
	fmt.Println(aa)*/

	/*var m Mover

	m = &Daoge{}
	v, ok := m.(*Daoge)
	fmt.Println(v)
	fmt.Println(ok)

	str := "dsad"
	justifyType(str)

	fmt.Printf("%T", m)*/
	/*var a int64 = 3
	reflectSetValue2(&a)*/

	/*	var a *int
		//指针为空那么就返回true 有值返回false
		fmt.Println("var a *int IsNil:", reflect.ValueOf(a).IsNil())
		//参数无效返回false 有效返回true
		fmt.Println("nil IsValid:", reflect.ValueOf(nil).IsValid())

		reflect.TypeOf(a)

		// 实例化一个匿名结构体
		b := struct{}{}
		// 尝试从结构体中查找"abc"字段
		fmt.Println("不存在的结构体成员:", reflect.ValueOf(b).FieldByName("abc").IsValid())
		// 尝试从结构体中查找"abc"方法
		fmt.Println("不存在的结构体方法:", reflect.ValueOf(b).MethodByName("abc").IsValid())
		// map
		c := map[string]int{}
		// 尝试从map中查找一个不存在的键
		fmt.Println("map中不存在的键：", reflect.ValueOf(c).MapIndex(reflect.ValueOf("娜扎")).IsValid())*/
	/*wg.Add(1)
	s := &School{
		name: "ja",
		person: Person{
			name: "zhangsan",
			age:  42,
		},
	}
	go hello()
	size := reflect.TypeOf(*s).NumField()
	fmt.Println("字段长 :", size)
	time.Sleep(time.Second)
	wg.Wait()*/
	ch := producer()
	consumer(ch)

}

func producer() <-chan int {
	ch := make(chan int, 2)
	ch <- 100
	ch <- 200
	close(ch)
	return ch
}

func consumer(ch <-chan int) {
	sum := 0
	for i := range ch {
		sum += i
	}
	fmt.Println(sum)
}

func f2(ch chan int) {
	/*for {
		v, ok := <-ch
		if !ok {
			fmt.Println("通道关了")
			break
		}
		fmt.Printf("v:%#v ok:%#v \n", v, ok)
	}*/
	for v := range ch {
		fmt.Printf("v:%#v \n", v)
	}
}

var wg sync.WaitGroup

func hello() {
	fmt.Println("go fun")
	wg.Done()
}

type School struct {
	name   string
	person Person
}

func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	b := v.IsNil()
	fmt.Println(b)
	if v.Kind() == reflect.Int64 {
		v.SetInt(100)
	}
}

func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	b := v.IsNil()
	fmt.Println(b)
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(100)
	}
}

func reflectType(a interface{}) {
	v := reflect.TypeOf(a)

	fmt.Printf("type Type:%v  Kind:%v \n", v.Name(), v.Kind())
}

func reflectValue(i interface{}) {
	v := reflect.ValueOf(i)
	k := v.Kind()

	switch k {
	case reflect.Int64:
		// v.Int()从反射中获取整型的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Float32:
		// v.Float()从反射中获取浮点型的原始值，然后通过float32()强制类型转换
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		// v.Float()从反射中获取浮点型的原始值，然后通过float64()强制类型转换
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))

	}
}

func justifyType(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string，value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("unsupport type！")
	}
}

type Mover interface {
	Move()
}

type Daoge struct {
}

func (d *Daoge) Move() {
	fmt.Println("狗叫")
}

// reverseIntSlice 是一个自定义类型，它嵌入了一个 sort.IntSlice。
type reverseIntSlice struct {
	arr [5]int
	sort.Interface
}

func (r reverseIntSlice) Len() int {
	return len(r.arr)
}

type Payer interface {
	pay(int64)
}

type Zhifubao struct {
}

func (z Zhifubao) pay(a int64) {
	fmt.Printf("支付宝支付了:%.2f \n", float64(a))
}

func (z *Weixin) pay(a int64) {
	fmt.Printf("微信支付了:%.2f \n", float64(a))
}

type Weixin struct {
}

func Zhifu(p Payer) {
	p.pay(100)
}

type Singer interface {
	Sing()
}

func MakeHungry(s Singer) {
	s.Sing()
}

type Bird struct {
}

func (b Bird) Sing() {
	fmt.Println("鸟在唱歌")
}

//Student 学生
type Student struct {
	ID     int `json:"id"`
	Gender string
	Name   string
}

//Class 班级
type Class struct {
	Title    string
	Students []Student
}

type Animal struct {
	name string
}

func (a Animal) move() {
	fmt.Printf("%s会动", a.name)
}

type Dog struct {
	Feet   int8
	animal Animal
}

func (d Dog) wang() {
	fmt.Printf("%s会汪汪叫", d.animal)
}

type Person struct {
	name string
	age  int8
}

func NewPerson(name string, age int8) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}
func (p Person) Dream(newAge int8) {
	p.age = newAge
}

func (p *Person) Setage(newAge int8) {
	p.age = newAge
}

func fun1() {
	fmt.Println("fun1")
}

func fun2() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("recover fun2")
			fmt.Println(err)
		}
	}()
	panic("panic fun2")
}

func fun3() {
	fmt.Println("fun3")
}

func add(a, b int) int {
	return a + b
}

func Errorf(format string, a ...interface{}) error {
	return errors.New(fmt.Sprintf(format, a))
}

func sub(a, b int) int {
	return a - b
}

type Op func(int, int) int

func do(z Op, a int, b int) int {
	return z(a, b)
}

func Uint8FromInt(n int) (uint8, error) {
	if 0 <= n && n <= math.MaxUint8 { // conversion is safe
		return uint8(n), nil
	}
	return 0, fmt.Errorf("%d is out of the uint8 range", n)
}

func zhizhen() {
	x := 1
	p := &x
	fmt.Println(x)
	fmt.Println(p)
	*p = 2
	fmt.Println(p)
	incr(&x)
}

func incr(p *int) int {
	*p++ // 非常重要：只是增加p指向的变量的值，并不改变p指针！！！
	return *p
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func resNum() int {
	a := 1
	return a
}

func findChong() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}
