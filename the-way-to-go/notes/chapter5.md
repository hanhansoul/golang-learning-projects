# 控制结构

# if-else 结构

```go
if condition1 {
	// do something	
} else if condition2 {
	// do something else	
} else {
	// catch-all or default
}
```

这里举一些有用的例子：

1. 判断一个字符串是否为空：
   - `if str == "" { ... }`
   - `if len(str) == 0 {...}`

2. 判断运行 Go 程序的操作系统类型，这可以通过常量 `runtime.GOOS` 来判断(第 2.2 节)。

```
 if runtime.GOOS == "windows"	 {
 	.	..
 } else { // Unix-like
 	.	..
 }
```

3. 函数 `Abs()` 用于返回一个整型数字的绝对值:

```
 func Abs(x int) int {
 	if x < 0 {
 		return -x
 	}
 	return x	
 }
```

4. `isGreater` 用于比较两个整型数字的大小:

```
 func isGreater(x, y int) bool {
 	if x > y {
 		return true	
 	}
 	return false
 }
```

在第四种情况中，if 可以包含一个初始化语句（如：给一个变量赋值）。这种写法具有固定的格式（在初始化语句后方必须加上分号）：

```
if initialization; condition {
	// do something
}

val := 10
if val > max {
	// do something
}

if val := 10; val > max {
	// do something
}
```

但要注意的是，使用简短方式 `:=` 声明的变量的作用域只存在于 if 结构中（在 if 结构的大括号之间，如果使用 if-else 结构则在 else 代码块中变量也会存在）。



# 测试多返回值函数的错误

Go 语言的函数经常使用两个返回值来表示执行是否成功：返回某个值以及 true 表示成功；返回零值（或 nil）和 false 表示失败。

**习惯用法**

```
value, err := pack1.Function1(param1)
if err != nil {
	fmt.Printf("An error occured in pack1.Function1 with parameter %v", param1)
	return err
}
// 未发生错误，继续执行：
```

**习惯用法**

```
if err := file.Chmod(0664); err != nil {
	fmt.Println(err)
	return err
}
```

**习惯用法**

```
if value, ok := readData(); ok {
…
}
```



# switch 结构

相比较 C 和 Java 等其它语言而言，Go 语言中的 switch 结构使用上更加灵活。它接受任意形式的表达式：

```go
switch var1 {
	case val1:
		...
	case val2:
		...
	default:
		...
}
```

一旦成功地匹配到某个分支，在执行完相应代码后就会退出整个 switch 代码块，也就是说您不需要特别使用 `break`语句来表示结束。

因此，程序也不会自动地去执行下一个分支的代码。如果在执行完每个分支的代码后，还希望继续执行后续分支的代码，可以使用 `fallthrough` 关键字来达到目的。

因此：

```
switch i {
	case 0: // 空分支，只有当 i == 0 时才会进入分支
	case 1:
		f() // 当 i == 0 时函数不会被调用
}
```

并且：

```
switch i {
	case 0: fallthrough
	case 1:
		f() // 当 i == 0 时函数也会被调用
}
```

switch 语句的第二种形式是不提供任何被判断的值（实际上默认为判断是否为 true），然后在每个 case 分支中进行测试不同的条件。

```
switch {
	case condition1:
		...
	case condition2:
		...
	default:
		...
}
```

例如：

```
switch {
	case i < 0:
		f1()
	case i == 0:
		f2()
	case i > 0:
		f3()
}
```

switch 语句的第三种形式是包含一个初始化语句：

```
switch initialization {
	case val1:
		...
	case val2:
		...
	default:
		...
}
```

这种形式可以非常优雅地进行条件判断：

```
switch result := calculate(); {
	case result < 0:
		...
	case result > 0:
		...
	default:
		// 0
}
```



# for 结构

如果想要重复执行某些语句，Go 语言中您只有 for 结构可以使用。

## 基于计数器的迭代

基本形式为：

```go
for 初始化语句; 条件语句; 修饰语句 {}
```

## 基于条件判断的迭代

for 结构的第二种形式是没有头部的条件判断迭代（类似其它语言中的 while 循环），基本形式为：`for 条件语句 {}`。

您也可以认为这是没有初始化语句和修饰语句的 for 结构，因此 `;;` 便是多余的了。

```go
package main

import "fmt"

func main() {
	var i int = 5

	for i >= 0 {
		i = i - 1
		fmt.Printf("The variable i is now: %d\n", i)
	}
}
```

## 无限循环

条件语句是可以被省略的，如 `i:=0; ; i++` 或 `for { }` 或 `for ;; { }`（`;;` 会在使用 gofmt 时被移除）：这些循环的本质就是无限循环。最后一个形式也可以被改写为 `for true { }`，但一般情况下都会直接写 `for { }`。

## for-range 结构

这是 Go 特有的一种的迭代结构，您会发现它在许多情况下都非常有用。它可以迭代任何一个集合（包括数组和 map）。语法上很类似其它语言中 foreach 语句，但您依旧可以获得每次迭代所对应的索引。一般形式为：`for ix, val := range coll { }`。

要注意的是，`val` 始终为集合中对应索引的值拷贝，因此它一般只具有只读性质，对它所做的任何修改都不会影响到集合中原有的值（**译者注：如果 `val` 为指针，则会产生指针的拷贝，依旧可以修改集合中的原值**）。

```
for pos, char := range str {
...
}
```



# Break 与 continue

略。



# 标签与 goto

略。