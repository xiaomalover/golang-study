package main

import (
	"fmt"
	"math"
)

func main() {
	var a int = 0b1100
	var b int = 0xC
	var c int = 0o14
	var d float32 = math.MaxFloat32
	var e float64 = math.MaxFloat64
	fmt.Printf("%d\n%d\n%d\n%f\n%f", a, b, c, d, e)
	var myfloat float32 = 10000018
	fmt.Println("\n",myfloat)
	myfloat++
	fmt.Println(myfloat)

	test := [...]int{5:33, 3:22}
	test[2] = 11
	fmt.Println(test)

	arr1 := [...]int{1,2,3}
	arr2 := [3]int{1,2,3}
	fmt.Printf("arr1:%T\narr2:%T\n", arr1, arr2)

	type intarr3 [3]int
	arr3 := intarr3{33,44,55}
	fmt.Printf("%v, %T", arr3, arr3)

	sli := [6]int{1,22,3,4,5,6}
	sli_item := sli[2:5]
	fmt.Printf("\n haha %v,%T", sli_item, sli_item)

	make_arr := make([]int, 2)
	make_arr1 := make([]int, 2, 10)

	fmt.Printf("\n%v, %v", len(make_arr), cap(make_arr1))
	var aa []int
	fmt.Println( aa)

	var map1 map[string]int = map[string]int{"English": 1, "Chinese": 99}
	map1["history"] = 100
	fmt.Println(map1)
	if data, exist := map1["Chinese"]; exist {
		fmt.Println("Chinese", data)
	} else {
		fmt.Println("Chinese not exist")
	}

	for subject, score := range map1 {
		fmt.Println("subject:", subject, ";score:", score)
	}

	for _, item := range sli {
		fmt.Println(item)
	}

	scores := make(map[string]int)
	fmt.Println(scores["Chinese"])

	name := "hello"
	ptr1 := &name
	fmt.Println(*ptr1)

	name1 := new(int)
	*name1 = 123
	fmt.Println(*name1)

	aint := 1     // 定义普通变量
	ptr := &aint  // 定义指针变量
	fmt.Println("普通变量存储的是：", aint)
	fmt.Println("普通变量存储的是：", *ptr)
	fmt.Println("指针变量存储的是：", &aint)
	fmt.Println("指针变量存储的是：", ptr)

	var ptr2 *int
	ac := 44
	ptr2 = &ac
	fmt.Println(ptr2)
	fmt.Println(*ptr2)




	ax := [3]int{89, 90, 91}
	modify(ax[:])
	fmt.Println(ax)
	bx := [3]int{89, 90, 91}
	modifyPtr(&bx)
	fmt.Println(bx)

	xiaomin := Profile{name: "xiaoming", age: 33, gender: "男"}
	_, _ = xiaomin.increaseAge()
	xiaomin.FmtProfile()

	myc := Company{"腾讯", "深圳南山"}
	mys := Staff{"张三",33, "男", myc}
	mys.selfDescription()

	fmt.Println(sum(1,2,3,4,5,6))
	MyPrintf(11, "ddd", "ccc", true, myc)

	vlist(ax[:], func(v int) {
		fmt.Println(v)
	})

	/*phone := Phone{1999, "我是电话", 10}
	pc := Computer{8888, "我是电脑", 20}
	goods := []Goods{phone, pc}*/
}

func vlist(params []int, f func(int))  {
	for _, value := range params {
		f(value)
	}
}

func sum(params ...int) int {
	total := 0
	for _, value := range params {
		total += value
	}
	return total
}

func MyPrintf(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an int value.")
		case string:
			fmt.Println(arg, "is a string value.")
		case int64:
			fmt.Println(arg, "is an int64 value.")
		default:
			fmt.Println(arg, "is an unknown type.")
		}
	}
}

func modify(sls []int) {
	sls[0] = 99
}

func modifyPtr(arr *[3]int) {
	(*arr)[0] = 99
}

type Profile struct {
	name   string
	age    int
	gender string
	mother *Profile // 指针
	father *Profile // 指针
}

func (person *Profile) increaseAge() (bool, error) {
	person.age++
	return true, nil
}

// 定义一个与 Profile 的绑定的方法
func (person Profile) FmtProfile() {
	fmt.Printf("名字：%s\n", person.name)
	fmt.Printf("年龄：%d\n", person.age)
	fmt.Printf("性别：%s\n", person.gender)
}

type Company struct {
	companyName string
	companyAddress string
}

type Staff struct {
	name string
	age int
	gender string
	Company
}

func (staff *Staff) selfDescription() bool {
	fmt.Printf("Name is %v; age is %v; gender is %v; company is %v; address is %v.\n",
		staff.name,
		staff.age,
		staff.gender,
		staff.companyName,
		staff.companyAddress,
	)
	fmt.Printf("Name is %v; age is %v; gender is %v; company is %v; address is %v.\n",
		staff.name,
		staff.age,
		staff.gender,
		staff.Company.companyName,
		staff.Company.companyAddress,
	)
	return true
}

type Goods interface {
	showInfo()
	addStock(num int)
}

type Phone struct {
	price int
	name string
	stock int
}

func (phone Phone) showInfo() {
	fmt.Printf("名称：%v; 价格：%v; 库存：%v", phone.name, phone.price, phone.stock)
}

func (phone *Phone) addStock(num int) {
	phone.stock += num
}

type Computer struct {
	price int
	name string
	stock int
}

func (computer Computer) showInfo() {
	fmt.Printf("名称：%v; 价格：%v; 库存：%v", computer.name, computer.price, computer.stock)
}

func (computer *Computer) addStock(num int) {
	computer.stock += num
}
