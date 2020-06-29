package main

import "fmt"

func main()  {
	//三、变量声明
	//1、声明格式var 变量名类型，变量声明了必须使用
	//2、声明整型变量，该变量默认值为0
	//3、同一个{}里，声明的变量名是唯一的
	var a int = 10
	fmt.Printf("a = ",a)
	//4、可以同时声明多个变量，中间用逗号隔开
	//var a,b,c int


	//七、匿名变量:丢弃数据不进行处理，_匿名变量配合函数返回值才有价值
	//_,i,_,b := 1,2,3,4

	//八、数据置换
	//使用多重赋值的方式实现(交换两个变量的值)
	b,c := 10,20
	b,c = c,b
    fmt.Println(b,c)

	//九、输出格式
	//Printf()函数不换行（"\n"能换行），Println()换行
	fmt.Printf("b=%d,c=%d\n",b,c)
	//十、接收输入:用于接收键盘输入
	var age int
	fmt.Println("请输入你的年龄")
	fmt.Scan(&age)
	fmt.Printf("a=%d",age)
	//十一、变量命名规范
	//关键字：
	//break       default        func         interface        select
	//case        defer         go           map              struct
	//chan        else           goto           package          switch
	//const        fallthrough    if             range            type
	//continue    for             import        return            var

	//预定义名字
	//true false iota nil
	//    	int int8 int16 int32 int64
	//    	uint uint8 uint16 uint32 uint64 uintptr
	//    	float32 float64 complex128 complex64
	//    	bool byte rune string error
	//
	//    	make len cap new append copy close delete
	//    	complexrealimag
	//    	panic recover
    //驼峰命名法
}
