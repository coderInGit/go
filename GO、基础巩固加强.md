1、指针
（1）GO语言中的指针
GO语言指针和C语言指针不同：
 - 默认值 nil
 - 操作符“&”取变量地址，“*”通过指针访问目标对象
 - 不支持指针运算，不支持 “->”运算符，直接用“.”访问目标成员

func main(){
    a := 100
    fmt.Printf("&a=%p\n",&a)

    var p *int = nil
    p = &a
    fmt.Printf("p = %p\n",p)
    fmt.Printf("a = %d,p = %d\n",a,*p)

    *p = 324
    fmt.Printf("a = %d,p = %d\n",a,*p)
}

2、函数 New()
表达式new(T)将创建一个T类型的匿名变量，所做的是为T类型的新值分配并清零一块内存空间，然后将这块内存空间的地址作为结果返回，而这个结果就是指向这个新的T类型值的指针值，返回的指针类型为*T。
new创建的内存空间位于heap上，空间的默认值为数据类型默认值。如：new(int) 则 *p为0，new(bool) 则 *p为false

3、指针做函数参数
地址传递（传引用）

4、slice
(1)、切片论述
切片并不是数组或数组指针，它通过内部指针和相关属性引⽤数组⽚段，以实现变⻓⽅案。

（2）创建切片
- 自动推导类型创建slice
s1 := [] int {1, 2, 3, 4}  创建 有 4 个元素的切片，分别为：1234
- 借助make创建 slice，格式：make(切片类型，长度，容量)
s2 := make([]int, 5, 10)	len(s2) = 5, cap(s2) = 10
- make时，没有指定容量，那么 长度==容量
s3 := make([]int, 5)  len(s3) = 5, cap(s3) = 5

（3）、切片截取
操作	含义
s[n]	切片s中索引位置为n的项
s[:]	从切片s的索引位置0到len(s)-1处所获得的切片
s[low:]	从切片s的索引位置low到len(s)-1处所获得的切片
s[:high]	从切片s的索引位置0到high处所获得的切片，len=high
s[low:high]	从切片s的索引位置low到high处所获得的切片，len=high-low
s[low : high : max]	从切片s的索引位置low到high处所获得的切片，len=high-low，cap=max-low
len(s)	切片s的长度，总是<=cap(s)
cap(s)	切片s的容量，总是>=len(s)
利用数组创建切片。切片在操作过程中，是直接操作原数组。切片是数组的引用！因此，在go语言中，我们常常使用切片代替数组。

（4）、切片做函数参数
切片作为函数参数时，传引用。


（5）、常用操作参数
append() 函数可以向 slice 尾部添加数据，可以自动为切片扩容。常常会返回新的 slice 对象

copy函数

函数 copy 在两个 slice 间复制数据，复制⻓度以 len 小的为准，两个 slice 指向同⼀底层数组。直接对应位置覆盖。

5、map
map值可以是任意类型，没有限制。map里所有键的数据类型必须是相同的，值也必须如此，但键和值的数据类型可以不相同。

对于map而言，可以使用len()函数，但不能使用cap()函数。

（1）删除 
使用delete()函数，指定key值可以方便的删除map中的k-v映射。

（2）map做函数参数
与slice 相似，在函数间传递映射并不会制造出该映射的一个副本，不是值传递，而是引用传递

6、结构体
（1）结构体初始化
- 普通变量：
type Student struct{
    id int 
    name string
    sex byte
    age int 
    addr string
}

func main(){
    //1、顺序初始化，必须每个成员都初始化
    var s1 Student = Student{1,"luffy",'m',18,"EastSea"}
    s2 := Student(2,"sanji",'f',20,"EastSea")

    //2、指定初始化某个成员，没有初始化的成员为零值
    s4 := Student{id:2,name:"Zoro"}
}

- 指针变量
type Student struct{
    id int 
    name string
    sex byte
    age int 
    addr string
}

func main(){
    var s5 *Student = &Student(3,"Nami",'m',16,"EastSea")
    s6 := &Student(4,"ro",'m',3,"NorthSea")
}

- 使用结构体成员
普通变量

var s1 Student = Student{1,"Luffy,'m',18,"EastSea"}
fmt.Printf("id=%d,name=%s,sex=%c,age=%d,addr=%s\n",s1.id,s1.name,s1.sex,s1.age,s1.addr)

//2、成员变量赋值

var s2 Student
s2.id = 2
s2.name = "Sanji"
s2.sex = 'd'
s2.age = 20
s2.addr = "WestSea"

fmt.Println(s2)

指针变量
//3、现分配空间
s3 := new(Student)
s3.id = 3
s3.name = "Nami"
fmt.Println(s3) //&{3 Nami 0 0}


//4、普通变量和指针变量类型打印
var s4 Student = Student{4,"Sanji",'m',18,"EastSea"}
fmt.Printf(“s4 =%v,&s4 =%v\n”,s4,&s4)//s4 = {4 sanji 109 18 sz},&s4 = &{4 Sanji 109 18 EastSea}

var p *Student = &s4
//p.成员和（*p）成员 操作是等价的
p.id = 5
(*p).name = "ro"
fmt.Println(p,*p,s4)

在Go语言中，普通结构体变量 和 结构体指针变量发访问成员的方法一致。

（2）、结构体比较

如果结构体的全部成员都是可以比较的，那么结构体也是可以比较的，那样的话两个结构体将可以使用 == 或 != 运算符进行比较，但不支持 > 或 < 。

func main(){
    s1 := Student{1,"Luffy,'m',18,"EastSea"}
    s2 := Student{1,"Luffy,'m',18,"EastSea"}
    fmt.Println("s1 == s2",s1 == s2)
    fmt.Println("s1 != s2",s1 != s2)
}

（3）、作函数参数

- 传值

func printValue(stu Student){
    su.id = 250
    fmt.Println("printvalue stu =",stu)
}

func main(){
    var s Student = Student(1,"Luffy",'m',18,"EastSea")

    printvalue(s) //值传递，形参不会影响到实参
}
传参过程中，实参会将自己的值拷贝一份给形参。因此结构体“传值”操作几乎不会在实际开发中被使用到。
近乎100%的使用都采用“传址”的方式，将结构体的引用传递给所需函数。

- 传引用
func printPointer(p *Student){
    su.id = 250
    fmt.Println("printPointer p =",p)
}

func main(){
    var s student = Student{1,"Luffy",'m',18,"EastSea"}
    printPointer(&s) // 传引用（地址），形参修改会影响到实参
    fmt.Println("main s = ",s)
}

7、文件操作
(1)、 字符串处理函数
除Contains、Join、Trim、Replace等我们学过的字符串处理函数之外，以下函数也常常会被用到。

- 字符串分割
func Split(s,sep string)[]string //功能：把s字符按照sep分割，返回slice

- 按空格拆分字符串 Fields

- 判断字符串后缀 HasSuffix

(2)、文件操作常用API






