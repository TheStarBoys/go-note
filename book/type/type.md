# 数据类型

## 基本数据类型

| 名称         | 宽度（字节） | 零值  | 说明                                                   |
| ------------ | ------------ | ----- | ------------------------------------------------------ |
| bool         | 1            | false | 布尔类型。true, false                                  |
| byte         | 1            | 0     | 字节类型。可以看作是uint8。                            |
| rune         | 4            | 0     | rune类型。用于存储Unicode码点，可以看作是int32。       |
| int/uint     | -            | 0     | 其宽度与平台有关。例如，32位机器，int宽度为4。         |
| int8/uint8   | 1            | 0     | 8位二进制数表示的整型                                  |
| int16/uint16 | 2            | 0     | 16位二进制数表示的整型                                 |
| int32/uint32 | 4            | 0     | 32位二进制数表示的整型                                 |
| int64/uint64 | 8            | 0     | 64位二进制数表示的整型                                 |
| float32      | 4            | 0.0   | 32位二进制数表示的浮点数类型                           |
| float64      | 8            | 0.0   | 64位二进制数表示的浮点数类型                           |
| complex64    | 8            | 0.0   | 64位二进制数表示的复数类型。实部和虚部由float32表示。  |
| complex128   | 16           | 0.0   | 128位二进制数表示的复数类型。实部和虚部由float64表示。 |
| string       | -            | ""    | 字符串类型。本质上是[]byte，其值不可变。               |

### 整数

**表示方式有3种：**

1. 十进制表示法：`123`
2. 八进制表示法：`0116`
3. 十六进制表示法：`0x4e`

### 浮点数

**表示方式：**

1. 十进制：`56.78`
2. 指数：
   - `12e+2 // 表示12 * (10 ^ 2)`
   - `12e-3 // 表示12 / (10 ^ 3)`

### 复数

`12e+2 + 43.4e-3i`

### 布尔型

一个布尔类型的值只有两种：`true` 和 `false`。`if` 和 `for` 语句的条件部分都是布尔类型的值，用 `==` 和 `<` 等比较操作也会产生布尔类型的值。

### rune

**表示方式：**

1. 该rune字面量所对应的字符，例如：
   - `'a'`、 `'中'`
2. 使用`"\x"`为前导并后跟2位16进制数。这种方式可以表示宽度为一个字节的值，即一个ASCII编码值。
3. 使用`"\"`为前导并后跟3位8进制数，这种方式可以表示宽度为一个字节的值，即一个ASCII编码值。
4. 使用`"\u"`为前导并后跟4位16进制数。它只能用于表示两个字节宽度的值。这种方式即为Unicode编码规范中的UCS-2表示法。**不过，UCS-2在不久之后就会别废止**。
5. 使用`"\U"`为前导并后跟8位16进制数。这种方式即为Unicode编码规范中的UCS-4表示法。UCS-4表示法已经成为Unicode编码规范和相关国际标准中的规范编码格式。

**转义字符：**

| 转义字符 | Unicode码点 | 说明                         |
| -------- | ----------- | ---------------------------- |
| \a       | U+0007      | 告警铃声或蜂鸣声             |
| \b       | U+0008      | 退格符                       |
| \f       | U+000C      | 换页符                       |
| \n       | U+000A      | 换行符                       |
| \r       | U+000D      | 回车符                       |
| \t       | U+0009      | 水平制表符                   |
| \v       | U+000b      | 垂直制表符                   |
| \\\      | U+005c      | 反斜杠                       |
| \\'      | U+0027      | 单引号。仅在rune字面量中有效 |
| \\"      | U+0022      | 双引号。仅在rune字面量中有效 |

**注意**，在rune字面量中，除了在上面表格中出现的转义字符之外的以 `"\"` 为前导的字符序列都是不合法的。当然，在上表中的转义符 `“\"”`也不能在rune字面量中出现。

### 字符串

在底层，一个字符串值即是一个字节的序列。

#### 字符串长度

字符串的长度即是底层字节序列中字节的个数。

长度为0的序列与一个空字符串对应。

一个字符串的长度在编译器就可以确定。

#### 字符串字面量

它代表了一个连续的字符序列。其中，每一个字符都会被隐含地以Unicode编码规范的UTF-8编码格式编码为若干字节。

它有两种表示格式：

- 原生字符串字面量
- 解释型字符串字面量



原生字符串字面量：是在两个反引号 \`  之间的字符序列。在两个反引号之间的，除了反引号以外的其他字符都是合法的，不存在任何转义字符。

解释性字符串字面量：是被两个引号 `"` 包含的字符序列。 解释性字符串中的转义字符都会被成功转义。举个例子：

```go
func main() {
	str := "\101"
	fmt.Println(str) // 'A'
	str = "\U00004E00"
	fmt.Println(str) // 中文字符 '一'
}
```



## 复合数据类型

### 数组

一个数组就是一个由若干想同类型的元素组成的序列，在Go语言中被称为Array。

#### 类型表示法

在声明数组的时候，需要指明它的长度和类型，例如：

```go
[n]T
```

其中，n代表数组的长度，T代表数组中的元素类型。

**注意**：只要类型声明中的数组长度不同，**即使两个数组的元素类型相同，它们也是不同的类型**。例如： `[2]string` 和 `[3]string` 就是两个不同的类型。数组的长度一旦确定，就无法在任何时候改变它。也即是说，**数组的长度是固定的**。

**长度**：

- 可以由非负整数字面量代表
- 也可以由一个表达式代表。

例如：

```go
[2]byte
[2*3*4]byte
```

**数组的类型**：

- 可以是预定义类型：string, byte, int 等
- 也可以是自定义类型，例如 `Person` 结构体。

#### 值表示法

```go
func main() {
	arr1 := [5]int{1, 2, 3, 4, 5}
	arr2 := [5]int{1, 2, 3, 4}
	arr3 := [...]int{1, 2, 3}
	arr4 := [5]int{1: 1, 2: 3, 4: 6, 0: 5}

	// arr: [1 2 3 4 5], len = 5, cap = 5
	fmt.Printf("arr: %v, len = %d, cap = %d\n", arr1, len(arr1), cap(arr1))
	
	// arr: [1 2 3 4 0], len = 5, cap = 5
	fmt.Printf("arr: %v, len = %d, cap = %d\n", arr2, len(arr2), cap(arr2))
	
	// arr: [1 2 3], len = 3, cap = 3
	fmt.Printf("arr: %v, len = %d, cap = %d\n", arr3, len(arr3), cap(arr3))
	
	// arr: [5 1 3 0 6], len = 5, cap = 5
	fmt.Printf("arr: %v, len = %d, cap = %d\n", arr4, len(arr4), cap(arr4))
}
```



### 切片

切片底层就是引用了数组，使用方式和数组很类似，但其长度是可以动态扩容的。

#### 类型表示法

```go
[]T
```

其中 `T` 是切片中元素的类型。

#### 值表示法

和数组类似。

#### 操作

**len**

切片的零值是 `nil` ，在 `nil` 切片上应用内建函数 `len()` ，将会得到 0 。

**cap**

切片的容量就是其底层数组的长度。

在 `nil` 切片上应用内建函数 `cap()` ，将会得到 0 。

**切片操作**

基本操作：

```go
func main() {
	array := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("array: %T\n", array) // array: [5]int
	// 对数组进行切片操作，将获得引用这个底层数组的切片
	slice := array[:]
	fmt.Printf("slice: %T\n", slice) // slice: []int
	// 对该切片进行更改，将更改底层数组的值
	slice[1] = 3
	fmt.Println("slice:", slice) // slice: [1 3 3 4 5]
	fmt.Println("array:", array) // array: [1 3 3 4 5]

	// 还可以将数组按照下标切成不同的区间
	slice = array[1:]
	fmt.Println("slice:", slice) // slice: [3 3 4 5]
}
```

利用切片操作来模拟**队列**：

```go
func main() {
	queue := []int{} // []
	// 入队
	queue = append(queue, 1) // [1]
	queue = append(queue, 2) // [1 2]
	// 出队
	queue = queue[1:] // [2]
	queue = queue[1:] // []
}
```

利用切片操作模拟**栈**：

```go
func main() {
	stack := []int{} // []
	// 入栈
	stack = append(stack, 1) // [1]
	stack = append(stack, 2) // [1 2]
	// 出栈
	stack = stack[:len(stack)-1] // [1]
	stack = stack[:len(stack)-1] // []
}
```

**append**

利用 `append` 在切片后追加一个元素：

```go
slice := []int{1, 2, 3}
slice = append(slice, 4) // [1 2 3 4]
```

利用 `append` 追加多个元素：

```go
slice := []int{1, 2, 3}
slice = append(slice, 4, 5, 6) // [1 2 3 4 5 6]
```

利用 `append` 追加一个切片：

```go
slice := []int{1, 2, 3}
slice = append(slice, []int{4, 5, 6}...) // [1 2 3 4 5 6]
```

利用 `append` 得到一个值一样，但地址不一样的新切片：

```go
slice := []int{1, 2, 3}
slice = append([]int{}, slice...) // [1 2 3]
```

利用 `append` 删除一个元素：

```go
slice := []int{1, 2, 3}
slice = append(slice[:1], slice[2:]...) // [1 3]
```

当 `append` 引起切片扩容时：

```go
func main() {
	// 底层数组
	array := [...]int{1, 2, 3, 4, 5}
	// 当切片通过append发生扩容时
	slice := array[:]
	// slice: [1 2 3 4 5], len = 5, cap = 5
	fmt.Printf("slice: %v, len = %d, cap = %d\n", slice, len(slice), cap(slice))
	slice = append(slice, 6, 7)
	// slice: [1 2 3 4 5 6 7], len = 7, cap = 10
	fmt.Printf("slice: %v, len = %d, cap = %d\n", slice, len(slice), cap(slice))
	// 修改切片值将不再影响array，因为其底层数组已经发生改变
	slice[2] = 10
	fmt.Println("slice:", slice) // slice: [1 2 10 4 5 6 7]
	fmt.Println("array:", array) // array: [1 2 3 4 5]
}
```

**copy**

利用 `copy` 将 srcSlice 中的值复制给 dstSlice，其返回值是被复制元素的个数：

```go
n := copy(dstSlice, srcSlice)
```



#### 遍历

`for range` 遍历

```go
func main() {
	slice := []int{1, 3, 2, 6}
	for i, v := range slice {
		fmt.Printf("index: %d, value: %d\n", i, v)
	}
	// Output:
	// index: 0, value: 1
	// index: 1, value: 3
	// index: 2, value: 2
	// index: 3, value: 6
}
```



### 字典

在Go语言中，字典类型的官方称谓是 `Map`，它是哈希表（Hash Table）的一个实现。`Map` 是无序的。

#### 类型表示法

```go
map[K]T
```

其中 `K` 是键的类型，`T` 是元素类型，也就是键值对中，值的类型。

`Map` 的键必须是可比较的，即键的值可以作为比较操作符 `==` 和 `!=` 的操作数。

比如其键的类型不能是：

- 函数类型
- 字典类型
- 切片类型

如果其键的类型是接口类型，那么就要求在程序运行期间，该类型的字典值中的每一个键值的动态类型都是可比较的，否则会引起 `panic`。

**合法的例子：**

```go
map[int]string
map[string]struct{name, department string}
```



**不合法的例子：**

```go
map[[]int]string
map[map[int]string]string
```

**利用技巧将不合法的键转为可比较的键：**

```go
func main() {
	// 比如我们期望将 []int 作为键，但它不合法，怎么办呢？
	// 首先声明键类型为string
	m := make(map[string]int)
	slice := []int{1, 2, 3}
	fmt.Println(helper(slice)) // ['\x01' '\x02' '\x03']
	// 通过辅助函数就可以将一个切片作为键了
	m[helper(slice)] = 1
}

// 通过辅助函数，将切片转为字符串表示
func helper(slice []int) string {
	return fmt.Sprintf("%q", slice)
}
```

#### 值表示法

```go
map[string]bool{"Vim": true, "Emacs": true, "LiteIDE": false, "Notepad": false}
```

#### 声明和初始化

内置的make函数可以创建一个map：

```go
m := make(map[string]bool)
```

 map字面值的语法创建map，同时还可以指定一些最初的key/value： 

```go
m := map[string]bool{
    "Vim": true,
    "Emacs": true,
    "LiteIDE": false,
    "Notepad": false,
}
```



#### 遍历

`for range` 遍历

可以看出 `Map` 是无序的。

```go
func main() {
	m := make(map[int]int)
	m[1] = 2
	m[2] = 4
	m[3] = 1
	for k, v := range m {
		fmt.Printf("key: %d, value: %d\n", k, v)
	}
	// Output:
	// 第一次结果：
	// key: 3, value: 1
	// key: 1, value: 2
	// key: 2, value: 4

	// 第二次结果：
	// key: 1, value: 2
	// key: 2, value: 4
	// key: 3, value: 1
}
```



### 结构体



### JSON



### 文本和HTML模板



## 类型转换