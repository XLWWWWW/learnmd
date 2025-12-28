G0

## 常量和变量

### 1、常量

```
/*
1、常量使用关键字 const 定义，用于存储不会改变的数据。
2、存储在常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型。
3、常量的定义格式：const identifier [type] = value，其中常量的定义中 type能够省略 未定义类型的常量会在必要时刻根据上下文来获得相关类型。
4、数字型的常量是没有大小和符号的，并且可以使用任何精度而不会导致溢出：
5、 赋值一个常量时，之后没赋值的常量都会应用上一行的赋值表达式
6、itoa 是 golang中的常量计数器
7、可见性原则 每个新单词的首字母大写，例如：numShips 和 startDate。表示该全局变量能够用于本文件中使用

    但如果你的全局变量希望能够被外部包所使用，则需要将首个单词的首字母也大写
*/
```



- 常量使用关键字 const 定义，用于存储不会改变的数据。

- 存储在常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型。

- 常量的定义格式：const identifier [type] = value，其中常量的定义中 type能够省略 未定义类型的常量会在必要时刻根据上下文来获得相关类型。

- 数字型的常量是没有大小和符号的，并且可以使用任何精度而不会导致溢出：

- 赋值一个常量时，之后没赋值的常量都会应用上一行的赋值表达式

- itoa 是 golang中的常量计数器

- 可见性原则 每个新单词的首字母大写，例如：numShips 和 startDate。表示该全局变量能够用于本文件中使用

  但如果你的全局变量希望能够被外部包所使用，则需要将首个单词的首字母也大写

  

```go
const c string = "C"
const num1 = 122132132131231231231233123211212312412321412312312451231

// 当不明确指明赋值参数时，这个参数会与上一个使用的参数相同 即 a2 = 1 a5 = 33
const (
	aa = 1
	a2
	a3 = 33
	a5
)

```



### 2、变量

- 形式 var identifier type。

- 初始化 当一个变量被声明之后，系统自动赋予它该类型的零值 int 为 0，float32(64) 为 0.0，bool 为 false，string 为空字符串，指针为 nil

- 驼峰法

- 作用域 如果一个变量在函数体外声明，则被认为是全局变量 可以在整个包甚至外部包（被导出后）使用，不管你声明在哪个源文件里或在哪个源文件里调用该变量。

  全局变量的声明语法 var identifier [type] = value 注意使用 = 赋值

- 在函数体内声明的变量称之为局部变量，它们的作用域只在函数体内，参数和返回值变量也是局部变量。

  当在函数体内声明局部变量时，应使用简短声明语法 :=，   使用操作符 := 可以高效地创建一个新的变量，称之为初始化声明。 a := 50 或 b := false

  注意 这只能用于

- Go 编译器的智商能够自动推断其类型 ，并且在编译时就已经完成推断过程。



### 3、数据类型

  - `float32` 精确到小数点后 7 位，`float64` 精确到小数点后 15 位。

  - `int8`（-128 -> 127）

    `int16`（-32768 -> 32767）

    `int32`（-2,147,483,648 -> 2,147,483,647）

    `int64`（-9,223,372,036,854,775,808 -> 9,223,372,036,854,775,807）

  - `uint8`（0 -> 255）

    `uint16`（0 -> 65,535）

    `uint32`（0 -> 4,294,967,295）

    `uint64`（0 -> 18,446,744,073,709,551,615）

  - 复数使用 `re+imI` 来表示，其中 re 代表实数部分，im 代表虚数部分，I 代表根号负 1。

    `complex64` (32 位实数和虚数)

    `complex128` (64 位实数和虚数)

    - 例如   `var c1 complex64 = 5 + 10i`

    - 如果 re 和 im 的类型均为 float32，那么类型为 complex64 的复数 c 可以通过以下方式来获得：

      c = complex(re, im)

    - 函数 real(c) 和 imag(c) 可以分别获得相应的实数和虚数部分。

  - 
    位运算  按位与 &   按位或 |   按位异或 ^   位清除 &^

    位运算只能用于整数类型的变量，且需当它们拥有等长位模式时。即如果 `int32`与`int64`做位运算 则会失效

  - 按位补足 ^  即 m^x，对于无符号 x 使用 “全部位设置为 1” 的规则，对于有符号 x 时使用 m=-1 

    ^10 = -01 ^ 10 = -11

    ```go
    fm.Printf("%08b\n", -1) 	//-0000001
    fm.Printf("%08b\n", 10)		//00001010
    fm.Printf("%08b\n", 100)	//01100100
    fm.Printf("%08b\n", -11)	//-0001011
    fm.Printf("%08b\n", -101)	//-1100101
    
    fm.Println(^10, ^100)		//-11 -101
    ```

    其实做的工作就是相与

  - 位左移 <<：位右移 >>： 

    **和其他语言的移位相同**

  - 在运算时 溢出 不会产生错误，**Go 会简单地将超出位数抛弃**。

    **如果你需要范围无限大的整数或者有理数（意味着只被限制于计算机内存），你可以使用标准库中的 big 包，该包提供了类似 big.Int 和 big.Rat 这样的类型**

  - 字符类型

    严格来说，字符并不是Go语言中的一个类型， 因为只占用 一个字节只是整数的特殊用例。ASCII码吗

    byte 类型是 `uint8` 的别名，对于只占用 1 个字节的传统 ASCII 编码的字符来说，完全没有问题

    因此， `var ch byte = 65` 或 `var ch byte = '\x41'` 或 `var ch byte = 'A'`；

    在 fmt.Printf中使用%c 控制输出 就是字符串

    在 Go 语言中，如果你想进行单个字符串的大小写转换可以使用内置的`unicode`包，

    **具体来说，是通过`unicode.ToLower`将字符由大写转为小写，通过 `unicode.ToUpper` 将字符转为大写**。

    

  - Go 同样支持 Unicode（UTF-8），因此字符同样称为 Unicode 代码点或者 runes， 并在内存中使用 int 来表示 在书写 Unicode 字符时，需要在 16 进制数之前加上前缀 \u 或者 \U。
    因为 Unicode 至少占用 2 个字节，所以我们使用 int16 或者 int 类型来表示。**前缀 \u 则总是紧跟着长度为 4 的 16 进制数**

    如果需要使用到 4 字节，则会加上 \U 前缀；**前缀 \U 紧跟着长度为 8 的 16 进制数**

    unicode 包含了一些针对测试字符的非常有用的函数（其中 ch 代表字符）：

    - 判断是否为字母：`unicode.IsLetter(ch)`
    - 判断是否为数字：`unicode.IsDigit(ch)`
    - 判断是否为空白符号：`unicode.IsSpace(ch)`

  >在Go语言中，`fmt.Printf` 函数提供了许多格式控制符来格式化输出。以下是常用的格式控制符：
  >
  >**通用格式控制符**
  >
  >- `%v`：按值的默认格式输出。
  >- `%+v`：类似 `%v`，但结构体的字段名也会输出。
  >- `%#v`：输出 Go 语法表示的值。
  >- `%T`：输出值的类型。
  >- `%%`：输出百分号本身。
  >
  >**布尔类型**
  >
  >- `%t`：输出 `true` 或 `false`。
  >
  >**整数类型**
  >
  >- `%b`：表示为二进制。
  >- `%c`：表示为对应的 Unicode 字符。
  >- `%d`：表示为十进制。
  >- `%o`：表示为八进制。
  >- `%O`：表示为带前缀 `0o` 的八进制（Go 1.21+）。
  >- `%x`：表示为十六进制，小写字母。
  >- `%X`：表示为十六进制，大写字母。
  >- `%U`：表示为 Unicode 格式 `U+1234`，对应字符的表示形式。
  >
  >**浮点数和复数类型**
  >
  >- `%b`：无小数部分，指数为二进制。
  >- `%e`：指数为 `e` 的科学计数法（小写）。
  >- `%E`：指数为 `E` 的科学计数法（大写）。
  >- `%f`：小数部分但不含指数。
  >- `%F`：与 `%f` 相同。
  >- `%g`：根据实际情况选择 `%e` 或 `%f` 格式。
  >- `%G`：根据实际情况选择 `%E` 或 `%F` 格式。
  >
  >##### 字符串和字节切片
  >
  >- `%s`：输出字符串或字节切片。
  >- `%q`：输出带引号的字符串。
  >- `%x`：每个字节用两个字符的十六进制表示（小写字母）。
  >- `%X`：每个字节用两个字符的十六进制表示（大写字母）。
  >
  >##### 指针
  >
  >- `%p`：表示为十六进制的内存地址。
  >
  >##### 宽度和精度控制
  >
  >- `%[flags][width][.precision]verb`
  >
  >- flags
  >
  >   可以包括：
  >
  >  - `-`：左对齐。
  >  - `+`：总是输出数值的符号。
  >  - `#`：使用 `0`，`0x` 或 `0X` 作为前缀。
  >  - `0`：用零填充。
  >
  >- `width`：最小宽度。
  >
  >- `precision`：精度，通常用于浮点数。
  >
  >
  >
  >### fmt的输入处理
  >
  >Go 语言标准库中用于读取用户输入的函数除了有 `fmt.Scanf` ，还有`fmt.Scan`, `fmt.Scanln`,它们之间有什么细微的差别呢。
  >
  >1. `Scan(&变量名)`: 不需要传递格式化符，只需要传递变量地址即可，不同变量之间用**空格 (不一定是一个空格）或者换行**分开，直到接收完所有指定的变量数，Scan 函数才会返回
  >
  >```go
  >var name string
  >var age int
  >_, err = fmt.Scan(&name, &age)
  >// 可以这样输入 zhangsan 18
  >// 也可以这样输入
  >// zhangsan
  >// 18
  >```
  >
  >1. `Scanln(&变量名)`：它和 Scan 用法类似，但是不同变量之间用空格隔开（不一定是一个空格），换行会让函数提前返回，此时未接收过值的变量为空。
  >
  >2. `Scanf(格式化符，&变量)`：需要指定输入的格式，如果输入不符合指定的格式，从不符合处开始，其后的变量值都为空。
  >
  >  ```go
  >  var a,b int
  >  _,err := fmt.Scanf("%d %d",&a,&b)
  >  ```
  >
  >### ++和--
  >
  >刚才我们使用了i++表示计数值 i 的数值 +1，这里的++ 是后缀递增操作符，它会先返回变量的当前值，然后再将变量增加 1。
  >
  >Go 语言中并没有提供前自增运算符，并且`i++`不能直接用作一个表达式，只能先赋值，再自增



### 4、字符串

字符串是 UTF-8 字符的一个序列（当字符为 ASCII 码时则占用 1 个字节，其它字符根据需要占用 2-4 个字节）。UTF-8 是被广泛使用的编码格式，是文本文件的标准编码，其它包括 XML 和 JSON 在内，也都使用该编码。由于该编码对占用字节长度的不定性，

- Go 中的字符串里面的字符也可能根据需要占用 1 至 4 个字节，这与其它语言如 C++、Java 或者 Python 不同（Java 始终使用 2 个字节）。Go 这样做的好处是不仅减少了内存和硬盘空间占用，同时也不用像其它语言那样需要对使用 UTF-8 字符集的文本进行编码和解码。



- Go 支持以下 2 种形式的字面值：

  - 解释字符串：

    该类字符串使用双引号括起来，其中的相关的转义字符将被替换，这些转义字符包括：

    `\n`：换行符

    `\r`：回车符

    `\t`：tab 键

    `\u` 或 `\U`：Unicode 字符

    `\\`：反斜杠自身

  - 非解释字符串：
  
    该类字符串使用反引号括起来，支持换行，例如：
  
    `This is a raw string \n` 中的 `\n\` 会被原样输出。

- 和 C/C++不一样，Go 中的字符串是根据长度限定，而非特殊字符 `\0`。

- `string` 类型的零值为长度为零的字符串，即空字符串 `""`。

- 一般的比较运算符（`==`、`!=`、`<`、`<=`、`>=`、`>`）通过在内存中按字节比较来实现字符串的对比。你可以通过函数 `len()` 来获取字符串所占的字节长度，例如：`len(str)`。

- 字符串的内容（纯字节）可以通过标准索引法来获取，在中括号 `[]` 内写入索引，索引从 0 开始计数：
  **需要注意的是，这种转换方案只对纯 ASCII 码的字符串有效。**

  **获取字符串中某个字节的地址的行为是非法的，例如：`&str[i]`。**

- **字符串拼接符 `+`**

  两个字符串 `s1` 和 `s2` 可以通过 `s := s1 + s2` 拼接在一起。

  `s2` 追加在 `s1` 尾部并生成一个新的字符串 `s`。由于编译器行尾自动补全分号的缘故，加号 `+` 必须放在第一行。

- 字符串长度

  - `len()` 函数的返回值的类型为 int，表示字符串的 ASCII 字符个数或字节长度。
  - 如果希望按习惯上的字符个数来计算，就需要使用 Go 语言中 UTF-8 包提供的 `RuneCountInString()` 函数，**统计 Uncode 字符数量。**

  ```go
  // count number of characters:
  str1 := "asSASA ddd dsjkdsjs dk"
  fmt.Printf("The number of bytes in string str1 is %d\n", len(str1))  \\22
  fmt.Printf("The number of characters in string str1 is %d\n", utf8.RuneCountInString(str1)) \\22
  str2 := "asSASA ddd dsjkdsjsこん dk"  
  fmt.Printf("The number of bytes in string str2 is %d\n", len(str2))   \\28
  fmt.Printf("The number of characters in string str2 is %d", utf8.RuneCountInString(str2))  \\24
  ```

- Go 中使用 `strings` 包来完成对字符串的主要操作。

  - `HasPrefix()` 判断字符串 `s` 是否以 `prefix` 开头：

    ```go
    strings.HasPrefix(s, prefix string) bool
    ```

  - `HasSuffix()` 判断字符串 `s` 是否以 `suffix` 结尾： `strings.HasSuffix(s, suffix string) bool`

    ```go
    strings.HasSuffix(s, suffix string) bool
    ```

  - `Contains()` 判断字符串 `s` 是否包含 `substr`： 

    ```go
    strings.Contains(s, substr string) bool
    ```

  - `Index()` 返回字符串 `str` 在字符串 `s` 中的索引（`str` 的第一个字符的索引），`-1` 表示字符串 `s` 不包含字符串 `str`：

    ```go
    strings.Index(s, str string) int
    ```

  - `LastIndex()` 返回字符串 `str` 在字符串 `s` 中最后出现位置的索引（`str` 的第一个字符的索引），`-1` 表示字符串 `s` 不包含字符串 `str`：

    ```go
    strings.LastIndex(s, str string) int
    ```

  - 查询非 ASCII 编码的字符在父字符串中的位置，建议使用以下函数来对字符进行定位：

    ```go
    strings.IndexRune(s string, r rune) int
    ```

  - `Replace()` 用于将字符串 `str` 中的前 `n` 个字符串 `old` 替换为字符串 `new`，**并返回一个新的字符串**，如果 `n = -1` 则替换所有字符串 `old` 为字符串 `new`：

    ```go
    strings.Replace(str, old, new string, n int) string
    ```

  - `Count()` 用于计算字符串 `str` 在字符串 `s` 中出现的**非重叠次数**

    ```go
    strings.Count(s, str string) int
    
    var manyG = "gggggggggg"
    fmt.Printf("Number of double g's in %s is: ", manyG)
    fmt.Printf("%d\n", strings.Count(manyG, "gg"))  // 所有的都不能重复 这里最后得到的结果是5
    ```

  - `Repeat()` 用于重复 `count` 次字符串 `s` 并返回一个新的字符串：

    ```go
    strings.Repeat(s, count int) string
    ```

  - `ToLower()` 将字符串中的 Unicode 字符全部转换为相应的小写字符：

    `ToUpper()` 将字符串中的 Unicode 字符全部转换为相应的大写字符：

    ```go
    strings.ToLower(s) string
    strings.ToUpper(s) string
    ```

  - `strings.TrimSpace(s)` 来剔除字符串开头和结尾的空白符号

    `strings.Trim(s,string substr)` 来将开头和结尾的 `substr` 去除掉

    如果你只想剔除开头或者结尾的字符串，则可以使用 `TrimLeft()` 或者 `TrimRight()` 来实现。

  - `strings.Fields(s)` 将会利用 1 个或多个空白符号来作为动态长度的分隔符将字符串分割成若干小块，并返回一个 slice，如果字符串只包含空白符号，则返回一个长度为 0 的 slice。

  - `strings.Split(s, sep)` 用于自定义分割符号来对指定字符串进行分割，同样返回 slice。

  - `Join()` 用于将元素类型为 string 的 slice 使用分割符号来拼接组成一个字符串：

    ```go
    strings.Join(sl []string, sep string) string
    
    str := "The quick brown fox jumps over the lazy dog"
    sl := strings.Fields(str)
    fmt.Printf("Splitted in slice: %v\n", sl)  	//Splitted in slice: [The quick brown fox jumps over the lazy dog]
    for _, val := range sl {
    	fmt.Printf("%s - ", val)				//The - quick - brown - fox - jumps - over - the - lazy - dog -
    }
    fmt.Println()
    str2 := "GO1|The ABC of Go|25"
    sl2 := strings.Split(str2, "|")
    fmt.Printf("Splitted in slice: %v\n", sl2)	//Splitted in slice: [GO1 The ABC of Go 25]
    for _, val := range sl2 {
    	fmt.Printf("%s - ", val)				//GO1 - The ABC of Go - 25 -
    }
    fmt.Println()
    str3 := strings.Join(sl2, ";")
    fmt.Printf("sl2 joined by ;: %s\n", str3) 	//sl2 joined by ;: GO1;The ABC of Go;25
    ```

  - 函数 `strings.NewReader(str)` 用于生成一个 `Reader` 并读取字符串中的内容，然后返回指向该 `Reader` 的指针，

    从其它类型读取内容的函数还有：

    - `Read()` 从 `[]byte` 中读取内容。
    - `ReadByte()` 和 `ReadRune()` 从字符串中读取下一个 `byte` 或者 `rune`。

  - 与字符串相关的类型转换都是通过 `strconv` 包实现的。

    - `strconv.Itoa(i int) string` 返回数字 `i` 所表示的字符串类型的十进制数。  to_string()

    - `strconv.FormatFloat(f float64, fmt byte, prec int, bitSize int) string` 将 64 位浮点型的数字转换为字符串

      其中 `fmt` 表示格式（其值可以是 `'b'`、`'e'`、`'f'` 或 `'g'`），`prec` 表示精度，`bitSize` 则使用 32 表示 `float32`，用 64 表示 `float64`。

    - 将字符串转换为其它类型 **`tp`** 并不总是可能的，可能会在运行时抛出错误 `parsing "…": invalid argument`。

    - `strconv.Atoi(s string) (i int, err error)` 将字符串转换为 `int` 型。

    - `strconv.ParseFloat(s string, bitSize int) (f float64, err error)` 将字符串转换为 `float64` 型。

    >利用多返回值的特性，这些函数会返回 2 个值，第 1 个是转换后的结果（如果转换成功），第 2 个是可能出现的错误，因此，我们一般使用以下形式来进行从字符串到其它类型的转换：
    >
    >```
    >val, err = strconv.Atoi(s)
    >```
    >
    >eg
    >
    >```go
    >package main
    >
    >import (
    >	"fmt"
    >	"strconv"
    >)
    >
    >func main() {
    >	var orig string = "666"
    >	var an int
    >	var newS string
    >
    >	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)	  
    >
    >	an, _ = strconv.Atoi(orig)
    >	fmt.Printf("The integer is: %d\n", an) 
    >	an = an + 5
    >	newS = strconv.Itoa(an)
    >	fmt.Printf("The new string is: %s\n", newS)
    >}
    >```



### 5、日期和时间

`time` 包为我们提供了一个数据类型 `time.Time`（作为值使用）以及显示和测量时间和日期的功能函数。

- 当前时间可以使用 `time.Now()` 获取，或者使用 `t.Day()`、`t.Minute()` 等等来获取时间的一部分；

- 你甚至可以自定义时间格式化字符串，例如： `fmt.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year())` 将会输出 `21.07.2011`。

- `Duration` 类型表示两个连续时刻所相差的纳秒数，类型为 `int64`。`Location` 类型映射某个时区的时间，UTC 表示通用协调世界时间。

- 包中的一个预定义函数 `func (t Time) Format(layout string) string` **可以根据一个格式化字符串来将一个时间 `t` 转换为相应格式的字符串**，你可以使用一些预定义的格式，如：`time.ANSIC` 或 `time.RFC822`。

  一般的格式化设计是通过对于一个标准时间的格式化描述来展现的，这听起来很奇怪（`02 Jan 2006 15:04` 是 Go 语言的诞生时间且自定义格式化时必须以此时间为基准），但看下面这个例子你就会一目了然：

```go
fmt.Println(t.Format("02 Jan 2006 15:04"))   	//21 Jul 2011 10:31


t := time.Now()
fmt.Println(t) // e.g. Wed Dec 21 09:52:14 +0100 RST 2011
fmt.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year())
// 21.12.2011
t = time.Now().UTC()
fmt.Println(t)          // Wed Dec 21 08:52:14 +0000 UTC 2011
fmt.Println(time.Now()) // Wed Dec 21 09:52:14 +0100 RST 2011
// calculating times:
week = 60 * 60 * 24 * 7 * 1e9 // must be in nanosec
week_from_now := t.Add(time.Duration(week))
fmt.Println(week_from_now) // Wed Dec 28 08:52:14 +0000 UTC 2011
// formatting times:
fmt.Println(t.Format(time.RFC822)) // 21 Dec 11 0852 UTC
fmt.Println(t.Format(time.ANSIC))  // Wed Dec 21 08:56:34 2011
// The time must be 2006-01-02 15:04:05
fmt.Println(t.Format("02 Jan 2006 15:04")) // 21 Dec 2011 08:52
s := t.Format("20060102")
fmt.Println(t, "=>", s)
// Wed Dec 21 08:52:14 +0000 UTC 2011 => 20111221
```

### 6、指针


- Go 语言的取地址符是 `&`，放到一个变量前使用就会返回相应变量的内存地址。

  ```go
  var i1 = 5
  fmt.Printf("An integer: %d, it's location in memory: %p\n", i1, &i1)
  ```

  这个地址可以存储在一个叫做指针的特殊数据类型中，在本例中这是一个指向 int 的指针，即 `i1`：此处使用 `*int` 表示。如果我们想调用指针 `intP`，我们可以这样声明它：

  ```
  var intP *int
  intP = &i1
  ```

- **一个指针变量可以指向任何一个值的内存地址** 它指向那个值的内存地址，在 32 位机器上占用 4 个字节，在 64 位机器上占用 8 个字节，并且与它所指向的值的大小无关。当然，可以声明指针指向任何类型的值来表明它的原始性或结构性；
- **解引用**，你可以在指针类型前面加上 `*` 号（前缀）来获取指针所指向的内容，这里的 `*` 号是一个类型更改器。使用一个指针引用一个值被称为间接引用。
- 当一个指针被定义后没有分配到任何变量时，它的值为 `nil`。
- 一个指针变量通常缩写为 `ptr`。

- 你不能获取字面量或常量的地址，例如：

  ```go
  const i = 5
  ptr := &i //error: cannot take the address of i
  ptr2 := &10 //error: cannot take the address of 10
  ```

- 对于经常导致 C 语言内存泄漏继而程序崩溃的指针运算（所谓的指针算法，如：`pointer+2`，移动指针指向字符串的字节数或数组的某个位置）是不被允许的。

  因此 `p++` 在 Go 语言的代码中是不合法的。

- 指针的一个高级应用是你可以传递一个变量的引用（如函数的参数），这样不会传递变量的拷贝。指针传递是很廉价的，只占用 4 个或 8 个字节。当程序在工作中需要占用大量的内存，或很多变量，或者两者都有，使用指针会减少内存占用和提高效率。被指向的变量也保存在内存中，直到没有任何指针指向它们，所以从它们被创建开始就具有相互独立的生命周期。

- 指针也可以指向另一个指针，并且可以进行任意深度的嵌套，导致你可以有多级的间接引用，但在大多数情况这会使你的代码结构不清晰。

- 对一个空指针的反向引用是不合法的，并且会使程序崩溃：

### 7、控制结构

- if else

  ```go
  if condition1 {
  	// do something	
  } else if condition2 {
  	// do something else	
  } else {
  	// catch-all or default
  }
  ```

  go中允许在if中先执行一个语句 接着 在进行判断

  ```go
  // eg1
  if err := file.Chmod(0664); err != nil {
  	fmt.Println(err)
  	return err
  }
  ```

- switch case

  ```go
  switch var1 {
  	case val1:
  		...
  	case val2:
  		...
  	default:
  		...
  }
  
  
  
  switch {
  	case condition1:
  		...
  	case condition2:
  		...
  	default:
  		...
  }
  
  
  // 同样 switch 也能够包含一个初始化语句
  switch initialization {
  	case val1:
  		...
  	case val2:
  		...
  	default:
  		...
  }
  switch a, b := x[i], y[j]; {
  	case a < b: t = -1
  	case a == b: t = 0
  	case a > b: t = 1
  }
  ```

  - 每一个 `case` 分支都是唯一的，从上至下逐一测试，直到匹配为止。（ **Go 语言使用快速的查找算法**来测试 `switch` 条件与 `case` 分支的匹配情况，直到算法匹配到某个 `case` 或者进入 `default` 条件为止。）

  - 一旦成功地匹配到某个分支，在执行完相应代码后就会退出整个 `switch` 代码块，也就是说您不需要特别使用 `break` 语句来表示结束。

  - 可选的 `default` 分支可以出现在任何顺序，但最好将它放在最后。它的作用类似与 if-else 语句中的 `else`，表示不符合任何已给出条件时，执行相关语句。

  - case中的情况能够哼

    ```go
    var num1 int = 100
    
    	switch num1 {
    	case 98, 99:
    		fmt.Println("It's equal to 98")
    	case 100: 
    		fmt.Println("It's equal to 100")
    	default:
    		fmt.Println("It's not equal to 98 or 100")
    	}
    ```



- for 循环

  ```go
  for i:=0; i<5; i++ {
  	for j:=0; j<10; j++ {
  		println(j)
  	}
  }
  
  
  for i >= 0 {
  	i = i - 1
  	fmt.Printf("The variable i is now: %d\n", i)
  }
  
  // 没有条件时  就构成了无限循环
  
  ```

  - **for-range 结构**

    用途 和python意向，用于迭代任何一个集合

    `for ix, val := range coll { }`

    - ix 索引 val 实际的值
    - **`val` 始终为集合中对应索引的值拷贝，因此它一般只具有只读性质，对它所做的任何修改都不会影响到集合中原有的值**

- goto与标签

  `for`、`switch` 或 `select` 语句都可以配合标签 (label) 形式的标识符使用，即某一行第一个以冒号 (`:`) 结尾的单词

  标签的名称是大小写敏感的，为了提升可读性，一般建议使用全部大写字母

  ```go
  package main
  
  import "fmt"
  
  func main() {
  
  LABEL1:
  	for i := 0; i <= 5; i++ {
  		for j := 0; j <= 5; j++ {
  			if j == 4 {
  				continue LABEL1
  			}
  			fmt.Printf("i is: %d, and j is: %d\n", i, j)
  		}
  	}
  }
  
  //本例中，continue 语句指向 LABEL1，当执行到该语句的时候，就会跳转到 LABEL1 标签的位置。
  // 当 j==4 和 j==5 的时候，没有任何输出
  
  ```

  

  ```go
  package main
  
  func main() {
  	i:=0
  	HERE:
  		print(i)
  		i++
  		if i==5 {
  			return
  		}
  		goto HERE
  }
  
  // 上面的代码会输出 01234。
  
  ```

  **特别注意** 使用标签和 `goto` 语句是不被鼓励的：它们会很快导致非常糟糕的程序设计，而且总有更加可读的替代方案来实现相同的需求

  如果您必须使用 `goto`，应当只使用正序的标签（标签位于 `goto` 语句之后），但注意标签和 `goto` 语句之间不能出现定义新变量的语句，否则会导致编译失败。

  ```go
  // compile error goto2.go:8: goto TARGET jumps over declaration of b at goto2.go:8
  package main
  
  import "fmt"
  
  func main() {
  		a := 1
  		goto TARGET // compile error
  		b := 9
  	TARGET:  
  		b += a
  		fmt.Printf("a is %v *** b is %v", a, b)
  }
  ```

  

## 函数

- Go 是**编译型语言**，所以**函数编写的顺序是无关紧要的**；鉴于可读性的需求，最好把 `main()` 函数写在文件的前面，其他函数按照一定逻辑顺序进行编写（例如函数被调用的顺序）、
- `init()` 函数 
  - 变量除了可以在全局声明中初始化，也可以在 init() 函数中初始化。这是一类非常特殊的函数，它不能够被人为调用，而是在每个包完成初始化后自动执行，并且执行优先级比 main() 函数高。
  - 每个源文件可以包含多个 init() 函数，同一个源文件中的 init() 函数会按照从上到下的顺序执行，如果一个包有多个源文件包含 init() 函数的话，则官方鼓励但不保证以文件名的顺序调用。初始化总是以单线程并且按照包的依赖关系顺序执行。
- main 函数 与 init 函数 不能拥有参数和返回值
- Go 里面有三种类型的函数：
  - 普通的带有名字的函数
  - 匿名函数或者 lambda 函数
  - 方法（Methods）

- 在 Go 里面函数重载是不被允许的。这将导致一个编译错误：

- 函数值（functions value）之间可以相互比较：如果它们引用相同的函数或者都是 nil 的话，则认为它们是相同的函数。
- 函数不能在其它函数里面声明（不能嵌套），不过我们可以通过使用匿名函数（参考 [第 6.8 节](https://learnku.com/docs/the-way-to-go/introduce/06.8.md)）来破除这个限制。
- 目前 Go 没有泛型（generic）的概念，也就是说它不支持那种支持多种类型的函数。不过在大部分情况下可以通过接口（interface），特别是空接口与类型选择（type switch，参考 第 11.12 节）与 / 或者通过使用反射（reflection，参考 第 6.8 节）来实现相似的功能。



### 参数与返回值

- go函数也可以返回零个或多个值，这是他的一大特性，能够使用return 返回一组值

- 函数定义时，它的形参一般是有名字的，不过我们也可以定义没有形参名的函数，只有相应的形参类型，就像这样：`func f(int, int, float64)`。

- 按值传递和按引用传递

  - **Go 默认使用按值传递来传递参数**，也就是传递参数的副本。函数接收参数副本之后，**在使用变量的过程中可能对副本的值进行更改，但不会影响到原来的变量**，比如 `Function(arg1)`。

  - 如果你希望函数可以直接修改参数的值，而不是对参数的副本进行操作，你需要将参数的地址（变量名前面添加 & 符号，比如 &variable）传递给函数，这就是按引用传递，比如 `Function(&arg1)`，此时传递给函数的是一个指针。如果传递给函数的是一个指针(32位或64位)，指针的值（一个地址）会被复制，但指针的值所指向的地址上的值不会被复制；我们可以通过这个指针的值来修改这个值所指向的地址上的值。

    传递指针给函数不但可以节省内存（因为没有复制变量的值），而且赋予了函数直接修改外部变量的能力，所以被修改的变量不再需要使用 `return` 返回。

  - 在函数调用时，像切片（slice）、字典（map）、接口（interface）、通道（channel）这样的引用类型都是**默认使用引用传递**

  - 函数定义形式

    参数为 a,b,c 都为int类型 返回值为int类型

    ```go
    func MultiPly3Nums(a int, b int, c int) int {
        // var product int = a * b * c
        // return product
        return a * b * c
    }
    
    // 当需要返回多个非命名返回值时，需要使用 () 把它们括起来，比如 (int, int)
    func getX2AndX3(input int) (int, int) {
        return 2 * input, 3 * input
    }
    
    // 命名返回值 能够省略 直接使用return 命令即可
    // return 或 return var 都是可以的。
    // 即使函数使用了命名返回值，你依旧可以无视它而返回明确的值。
    // 尽量使用命名返回值：会使代码更清晰、更简短，同时更加容易读懂。
    func getX2AndX3_2(input int) (x2 int, x3 int) {
        x2 = 2 * input
        x3 = 3 * input
        // return x2, x3
        return
    }
    ```

    

### 变长参数

- 如果函数的最后一个参数是采用 `...type` 的形式，那么这个函数就可以处理一个变长的参数，这个长度可以为 0，这样的函数称为变长函数。

  ```go
  func myFunc(a, b, arg ...int) {}
  
  func Greeting(prefix string, who ...string)
  Greeting("hello:", "Joe", "Anna", "Eileen")
  // 在 Greeting 函数中，变量 who 的值为 []string{"Joe", "Anna", "Eileen"}。
  ```

  

- 如果参数被存储在一个 slice 类型的变量 slice 中，则可以通过 slice... 的形式来传递参数调用变参函数。

  ```go
  package main
  
  import "fmt"
  
  func main() {
      x := min(1, 3, 2, 0)
      fmt.Printf("The minimum is: %d\n", x)
      slice := []int{7,9,3,5,1}
      x = min(slice...)
      fmt.Printf("The minimum in the slice is: %d", x)
  }
  
  func min(s ...int) int {
      if len(s)==0 {
          return 0
      }
      min := s[0]
      for _, v := range s {
          if v < min {
              min = v
          }
      }
      return min
  }
  ```

- 一个接受变长参数的函数可以将这个参数作为其它函数的参数进行传递

  

  ```go
  func F1(s ...string) {
  	F2(s...)
  	F3(s)
  }
  
  func F2(s ...string) { }
  // 变长参数可以作为对应类型的 slice 进行二次传递。
  func F3(s []string) { }
  ```

- 变长参数的类型并不是都相同,使用多个参数进行传递并不是明治的选择

  - 定义一个结构类型，假设它叫 `Options`，用以存储所有可能的参数：

    ```go
    type Options struct {
    	par1 type1,
    	par2 type2,
    	...
    }
    ```

    

  - 使用空接口：

    如果一个变长参数的类型没有被指定，则可以使用默认的空接口 `interface{}`，这样就可以接受任何类型的参数（详见[第 11.9 节](https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/11.9.md) ）。该方案不仅可以用于长度未知的参数，还可以用于任何不确定类型的参数。一般而言我们会使用一个 for-range 循环以及 `switch` 结构对每个参数的类型进行判断：

    ```go
    func typecheck(..,..,values … interface{}) {
    	for _, value := range values {
    		switch v := value.(type) {
    			case int: …
    			case float: …
    			case string: …
    			case bool: …
    			default: …
    		}
    	}
    }
    ```

    

### defer和追踪

- 关键字 `defer` 允许我们**推迟到函数返回之前（或任意位置执行 `return` 语句之后）一刻才执行某个语句或函数**（为什么要在返回之后才执行这些语句？因为 `return` 语句同样可以包含一些操作，而不是单纯地返回某个值）。

  ```go
  func function1() {
  	fmt.Printf("In function1 at the top\n")
  	defer function2()
  	fmt.Printf("In function1 at the bottom!\n")
  }
  
  func function2() {
  	fmt.Printf("Function2: Deferred until the end of the calling function!")
  }
  ```

  

- 使用 `defer` 的语句同样可以接受参数

  ```go
  func a() {
  	i := 0
  	defer fmt.Println(i)
  	i++
  	return
  }
  ```

  

  

- 当有多个 defer 行为被注册时，它们会以逆序执行（类似栈，即后进先出）

  ```go
  func f() {
  	for i := 0; i < 5; i++ {
  		defer fmt.Printf("%d ", i)
  	}
  }
  ```

  

-   因此，defer的特性就是在最后删除，它一般用于释放某些已分配的资源

  - 关闭文件流

  - 解锁一个加锁的资源

  - 打印最终报告

  - 关闭数据库链接

  - 一个基础但十分实用的实现代码执行追踪的方案就是在进入和离开某个函数打印相关的消息

    > 例子
    >
    > ```go
    > package main
    > 
    > import "fmt"
    > 
    > func trace(s string)   { fmt.Println("entering:", s) }
    > func untrace(s string) { fmt.Println("leaving:", s) }
    > 
    > func a() {
    > 	trace("a")
    > 	defer untrace("a")
    > 	fmt.Println("in a")
    > }
    > 
    > func b() {
    > 	trace("b")
    > 	defer untrace("b")
    > 	fmt.Println("in b")
    > 	a()
    > }
    > 
    > func main() {
    > 	b()
    > }
    > 
    > 
    > // 更精简一点 这个有点帅吧
    > package main
    > 
    > import "fmt"
    > 
    > func trace(s string) string {
    > 	fmt.Println("entering:", s)
    > 	return s
    > }
    > 
    > func un(s string) {
    > 	fmt.Println("leaving:", s)
    > }
    > 
    > func a() {
    > 	defer un(trace("a"))
    > 	fmt.Println("in a")
    > }
    > 
    > func b() {
    > 	defer un(trace("b"))
    > 	fmt.Println("in b")
    > 	a()
    > }
    > 
    > func main() {
    > 	b()
    > }
    > ```
    >
    > 

### 内置函数

Go 语言拥有一些不需要进行导入操作就可以使用的内置函数。它们有时可以针对不同的类型进行操作

| 名称                             | 说明                                                         |
| -------------------------------- | ------------------------------------------------------------ |
| `close()`                        | 用于管道通信                                                 |
| `len()`、`cap()`                 | `len()` 用于返回某个类型的长度或数量（字符串、数组、切片、`map` 和管道）；<br />`cap()` 是容量的意思，用于返回某个类型的最大容量（只能用于数组、切片和管道，不能用于 `map`） |
| `new()`、`make()`                | `new()` 和 `make()` 均是用于分配内存：`new()` 用于值类型和用户定义的类型，如自定义结构，`make` 用于内置引用类型（切片、`map` 和管道）。<br />它们的用法就像是函数，但是将类型作为参数：`new(type)`、`make(type)`。<br />**`new(T)` 分配类型 `T` 的零值并返回其地址**，**也就是指向类型 `T` 的指针**。它也可以被用于基本类型：`v := new(int)`。<br />**`make(T)` 返回类型 `T` 的初始化之后的值，因此它比 `new()` 进行更多的工作**。**`new()` 是一个函数，不要忘记它的括号**。 |
| `copy()`、`append()`             | 用于复制和连接切片                                           |
| `panic()`、`recover()`           | 两者均用于错误处理机制                                       |
| `print()`、`println()`           | 底层打印函数，在部署环境中建议使用 `fmt` 包                  |
| `complex()`、`real ()`、`imag()` | 用于创建和操作复数                                           |



### 递归函数

```go
func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}
```

在使用递归函数时经常会遇到的一个重要问题就是栈溢出：一般出现在大量的递归调用导致的程序栈内存分配耗尽。这个问题可以通过一个名为 [懒惰求值](https://zh.wikipedia.org/wiki/惰性求值) 的技术解决，在 Go 语言中，我们可以使用管道 (channel) 和 goroutine（详见[第 14.8 节](https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/14.8.md)）来实现。

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%d is even: is %t\n", 16, even(16)) // 16 is even: is true
	fmt.Printf("%d is odd: is %t\n", 17, odd(17))
	// 17 is odd: is true
	fmt.Printf("%d is odd: is %t\n", 18, odd(18))
	// 18 is odd: is false
}

func even(nr int) bool {
	if nr == 0 {
		return true
	}
	return odd(RevSign(nr) - 1)
}

func odd(nr int) bool {
	if nr == 0 {
		return false
	}
	return even(RevSign(nr) - 1)
}

func RevSign(nr int) int {
	if nr < 0 {
		return -nr
	}
	return nr
}
```



### 回调函数

函数可以作为其它函数的参数进行传递，然后在其它函数内调用执行，一般称之为回调

```go
func main() {
	callback(1, Add)
}

func Add(a, b int) {
	fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
}

func callback(y int, f func(int, int)) {
	f(y, 2) // this becomes Add(1, 2)
}
```



### 匿名函数和闭包

- 当我们不希望给函数起名字的时候，可以使用匿名函数，例如：`func(x, y int) int { return x + y }`。

- 这样的一个函数不能够独立存在（编译器会返回错误：`non-declaration statement outside function body`），

  但可以被赋值于某个变量，即保存函数的地址到变量中：`fplus := func(x, y int) int { return x + y }`，然后通过变量名对函数进行调用：`fplus(3,4)`。

- 数列表的第一对括号必须紧挨着关键字 `func`，因为匿名函数没有名称。花括号 `{}` 涵盖着函数体，最后的一对括号表示对该匿名函数的调用。

  ```go
  func() {
  	sum := 0
  	for i := 1; i <= 1e6; i++ {
  		sum += i
  	}
  }()
  ```

- 匿名函数同样被称之为闭包（函数式语言的术语）：它们被允许调用定义在其它环境下的变量。闭包可使得某个函数捕捉到一些外部状态，例如：函数被创建时的状态。另一种表示方式为：**一个闭包继承了函数所声明时的作用域。**这种状态（作用域内的变量）都被共享到闭包的环境中，因此这些变量可以在闭包中被操作，直到被销毁，


- 关键字 `defer` 经常配合匿名函数使用，它可以用于改变函数的命名返回值。

  ```go
  package main
  
  import "fmt"
  
  func f() (ret int) {
  	defer func() {
  		ret++
  	}()
  	return 1
  }
  func main() {
  	fmt.Println(f())
  }
  
  // f函数的返回值为2 因为 f()的匿名函数被调用 defer 关键字使得该在返回之后被调用
  ```

  

- 将函数作为返回值

  ```go
  func Add2() func(b int) int {
  	return func(b int) int {
  		return b + 2
  	}
  }
  
  func Adder(a int) func(b int) int {
  	return func(b int) int {
  		return a + b
  	}
  }
  func main() {
  	// 无参数时 第二次传入的参数是用于 内部的匿名函数的
  	p2 := Add2()
  	fmt.Printf("Call Add2 for 3 gives: %v\n", p2(3))
  	// 第一个参数赋值给a 第二个参数用于内部的匿名参数
  	TwoAdder := Adder(2)
  	fmt.Printf("The result is: %v\n", TwoAdder(3))
  }
  
  ```

- **闭包函数保存并积累其中的变量的值，不管外部函数退出与否，它都能够继续操作外部函数中的局部变量**

  ```go
  func Adder() func(int) int {
  	var x int
  	return func(delta int) int {
  		x += delta
  		return x
  	}
  }
  
  func main() {
  	var f = Adder()
  	fmt.Print(f(1), " - ")
  	fmt.Print(f(20), " - ")
  	fmt.Print(f(300))
  }
  // 三次调用函数 f 的过程中函数 Adder() 中变量 delta 的值分别为：1、20 和 300。
  // 在多次调用中，变量 x 的值是被保留的，即 0 + 1 = 1，然后 1 + 20 = 21，最后 21 + 300 = 321：
  // 闭包函数保存并积累其中的变量的值，不管外部函数退出与否，它都能够继续操作外部函数中的局部变量
  ```

  


  > - 斐波那契数列 使用闭包处理 其中a，b的值 在每次调用后 实际上都会保留到下一次使用. 因此 每一次返回的都是增加后的数
  >
  >
  > ```go
  > package main
  > 
  > // fib returns a function that returns
  > // successive Fibonacci numbers.
  > func fib() func() int {
  > 	a, b := 1, 1
  > 	return func() int {
  > 		a, b = b, a+b
  > 		return b
  > 	}
  > }
  > 
  > func main() {
  > 	f := fib()
  > 	// Function calls are evaluated left-to-right.
  > 	// println(f(), f(), f(), f(), f())
  > 	for i := 0; i <= 9; i++ {
  > 		println(i+2, f())
  > 	}
  > }
  > ```
  >
  > - 工厂函数
  >
  >   一个返回值为另一个函数的函数可以被称之为工厂函数，这在您需要创建一系列相似的函数的时候非常有用
  >
  >   书写一个工厂函数而不是针对每种情况都书写一个函数
  >
  >   ```go
  >   func MakeAddSuffix(suffix string) func(string) string {
  >   	return func(name string) string {
  >   		if !strings.HasSuffix(name, suffix) {
  >   			return name + suffix
  >   		}
  >   		return name
  >   	}
  >   }
  >                                     
  >   addBmp := MakeAddSuffix(".bmp")
  >   addJpeg := MakeAddSuffix(".jpeg")
  >   addBmp("file") // returns: file.bmp
  >   addJpeg("file") // returns: file.jpeg
  >   ```

- 使用闭包调试

  当您在分析和调试复杂的程序时，无数个函数在不同的代码文件中相互调用，如果这时候能够准确地知道哪个文件中的具体哪个函数正在执行，对于调试是十分有帮助的。

  您可以使用 `runtime` 或 `log` 包中的特殊函数来实现这样的功能。包 `runtime` 中的函数 `Caller()` 提供了相应的信息，因此可以在需要的时候实现一个 `where()` 闭包函数来打印函数执行的位置：

  ```go
  where := func() {
  	_, file, line, _ := runtime.Caller(1)
  	log.Printf("%s:%d", file, line)
  }
  where()
  // some code
  where()
  // some more code
  where()
  
  
  // 您也可以设置 log 包中的 flag 参数来实现：
  
  log.SetFlags(log.Llongfile)
  log.Print("")
  
  // 或使用一个更加简短版本的 where() 函数：
  var where = log.Print
  func func1() {
  where()
  ... some code
  where()
  ... some code
  where()
  }
  ```

  

### 统计函数执行时间

```go
start := time.Now()
longCalculation()
end := time.Now()
delta := end.Sub(start)
fmt.Printf("longCalculation took this amount of time: %s\n", delta)
```

### 通过内存缓存来提升性能

当在进行大量的计算时，提升性能最直接有效的一种方式就是避免重复计算。

通过在内存中缓存和重复利用相同计算的结果，称之为内存缓存。最明显的例子就是生成斐波那契数列的程序

要计算数列中第 n 个数字，需要先得到之前两个数的值，但很明显绝大多数情况下前两个数的值都是已经计算过的。

即每个更后面的数都是基于之前计算结果的重复计算,而我们要做就是将第 n 个数的值存在数组中索引为 n 的位置,然后在数组中查找是否已经计算过，如果没有找到，则再进行计算。





## 集合

go 明显受到 Python 的影响,拥有**集合**的概念，它是可以包含大量条目 (item) 的数据结构，例如数组、切片和 `map`

### 数组

- 数组是具有相同 **唯一类型** 的一组已编号且**长度固定**的数据项序列（这是一种同构的数据结构）；

  这种类型可以是任意的原始类型例如整型、字符串或者自定义类型。

  数组长度必须是一个**常量表达式**，并且必须是一个非负整数。**数组长度也是数组类型的一部分**，所以 `[5]int` 和 `[10]int` 是属于不同类型的

- 数组元素可以通过 **索引**（位置）来读取（或者修改），索引从 `0` 开始，第一个元素索引为 `0`

- 元素的数目（也称为长度或者数组大小）**必须是固定的并且在声明该数组时就给出**（编译时需要知道数组长度以便分配内存）；**数组长度最大为 2GB。**

- Go 语言中的数组是一种 **值类型**（不像 C/C++ 中是指向首元素的指针），所以可以通过 `new()` 来创建： `var arr1 = new([5]int)`。

  那么这种方式和 `var arr2 [5]int` 的区别是什么呢？`arr1` 的类型是 `*[5]int`，而 `arr2` 的类型是 `[5]int`

  这样的结果就是当把一个数组赋值给另一个时，需要再做一次数组内存的拷贝操作。例如：

  ```
  arr2 := *arr1
  arr2[2] = 100
  ```

  这样两个数组就有了不同的值，在赋值后修改 `arr2` 不会对 `arr1` 生效。

- 几种常用的赋值方法

  - `var arrAge = [5]int{18, 20, 15, 22, 16}`

    5个值都被正确赋值

  - `[10]int {1, 2, 3}`

    这是一个有 10 个元素的数组，除了前三个元素外其他元素都为 `0`

  - `var arrLazy = [...]int{5, 6, 7, 8, 22}`

    其中 ... 能够被省略

  - `var arrKeyValue = [5]string{3: "Chris", 4: "Ron"}`

    key value 语法

    

- 多维数组

  数组通常是一维的，但是可以用来组装成多维数组，例如：`[3][5]int`，`[2][2][2]float64`。



- 将数组传递给函数

  把一个大数组传递给函数会消耗很多内存。有两种方法可以避免这种情况：

  - 传递数组的指针

    ```go
    package main
    import "fmt"
    
    func main() {
    	array := [3]float64{7.0, 8.5, 9.1}
    	x := Sum(&array) // Note the explicit address-of operator
    	// to pass a pointer to the array
    	fmt.Printf("The sum of the array is: %f", x)
    }
    
    func Sum(a *[3]float64) (sum float64) {
    	for _, v := range a { // derefencing *a to get back to the array is not necessary!
    		sum += v
    	}
    	return
    }
    ```

  - 使用数组的切片

​    

### 数组切片

- 数组的切片和python一样 是左闭右开的

- 切片是可索引的，并且可以由 `len()` 函数获取长度。

- 给定项的切片索引可能比相关数组的相同元素的索引小。和数组不同的是，切片的长度可以在运行时修改，最小为 0， 最大为相关数组的长度：切片是一个 **长度可变的数组**。

- 切片提供了计算容量的函数 `cap()` 可以测量切片最长可以达到多少：**它等于切片的长度 + 数组除切片之外的长度**。

  如果 `s` 是一个切片，`cap(s)` 就是从 `s[0]` 到数组末尾的数组长度。

  切片的长度永远不会超过它的容量，所以对于切片 `s` 来说该不等式永远成立：`0 <= len(s) <= cap(s)`。

- **多个切片如果表示同一个数组的片段，它们可以共享数据；因此一个切片和相关数组的其他切片是共享存储的**，相反，不同的数组总是代表不同的存储。数组实际上是切片的构建块。

- **优点** 因为切片是引用，所以它们不需要使用额外的内存并且比使用数组更有效率，所以在 Go 代码中切片比数组更常用。

- 声明切片的格式是： `var identifier []type`（不需要说明长度）。

  一个切片在未初始化之前默认为 `nil`，长度为 0。

  切片的初始化格式是：`var slice1 []type = arr1[start:end]`。

  - `var slice1 []type = arr1[:]`  == `var slice1 []type = arr1[0:len(arr1)]`

  - `arr1[2:]` 和 `arr1[2:len(arr1)]` 相同，都包含了数组从第三个到最后的所有元素。

    `arr1[:3]` 和 `arr1[0:3]` 相同，包含了从第一个到第三个元素（不包括第四个）。

  - 如果你想去掉 `slice1` 的最后一个元素，只要 `slice1 = slice1[:len(slice1)-1]`。

  - 一个由数字 1、2、3 组成的切片可以这么生成：`s := [3]int{1,2,3}[:]`（注：应先用 `s := [3]int{1, 2, 3}` 生成数组, 再使用 `s[:]` 转成切片）甚至更简单的 **`s := []int{1,2,3}`**。

  - ![img](https://github.com/unknwon/the-way-to-go_ZH_CN/raw/master/eBook/images/7.2_fig7.2.png?raw=true)

  - 切片在定义后 支持增长 能够增长到之前提到的cap的长度的位置 实际上cap的长度是你切片的起始位置到原数组末尾的位置

    
    
    ```go
    package main
    import "fmt"
    
    func main() {
    	var arr1 [6]int
    	var slice1 []int = arr1[2:5] // item at index 5 not included!
    
    	// load the array with integers: 0,1,2,3,4,5
    	for i := 0; i < len(arr1); i++ {
    		arr1[i] = i
    	}
    
    	// print the slice
    	for i := 0; i < len(slice1); i++ {
    		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
    	}
    
    	fmt.Printf("The length of arr1 is %d\n", len(arr1))
    	fmt.Printf("The length of slice1 is %d\n", len(slice1))
    	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))
    
    	// grow the slice 这里在原有切片的长度基础上 又增加了1
    	slice1 = slice1[0:4]
    	for i := 0; i < len(slice1); i++ {
    		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
    	}
    	fmt.Printf("The length of slice1 is %d\n", len(slice1))
    	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))
    
    	// grow the slice beyond capacity
    	//slice1 = slice1[0:7 ] // panic: runtime error: slice bound out of range
        
        for i := 0; i < len(slice1); i++ {
    		slice1[i] = 33 + i
    	}
    	fmt.Println(slice1)
    	fmt.Println(arr1)
    
    	for i := 0; i < len(arr1); i++ {
    		arr1[i] = 777 + i
    	}
    	fmt.Println(slice1)
    	fmt.Println(arr1)
        
        
    }
    ```
    
  
- **切片的改变会影响原数组，原数组的改变会影响切片**

  因此，在使用时，如果你有一个函数需要对数组做操作，你可能总是需要把参数声明为切片。当你调用该函数时，把数组分片，创建为一个切片引用并传递给该函数。

  相应的，函数中的入参也需要设置为 切片

  ```go
  func sum(a []int) int {
  	s := 0
  	for i := 0; i < len(a); i++ {
  		s += a[i]
  	}
  	return s
  }
  
  func main() {
  	var arr = [5]int{0, 1, 2, 3, 4}
  	sum(arr[:])
  }
  ```

- 使用make创建一个切片

  - 当相关数组还没有定义时，我们可以使用 `make()` 函数来创建一个切片，同时创建好相关数组：`var slice1 []type = make([]type, len)`。

  - 也可以简写为 `slice1 := make([]type, len)`，这里 `len` 是数组的长度并且也是 `slice` 的初始长度。即 `s2 := make([]int, 10)`，那么 `cap(s2) == len(s2) == 10`。

  - 如果你想创建一个 `slice1`，它不占用整个数组，而只是占用以 `len` 为个数个项，那么只要：`slice1 := make([]type, len, cap)`。

    `make()` 的使用方式是：`func make([]T, len, cap)`，其中 `cap` 是可选参数。

    所以下面两种方法可以生成相同的切片:

    ```go
    make([]int, 50, 100)
    new([100]int)[0:50]
    ```

    > - `new(T)` 为每个新的类型 `T` 分配一片内存，初始化为 `0` 并且返回类型为 `*T` 的内存地址：这种方法 **返回一个指向类型为 `T`，值为 `0` 的地址的指针**，它适用于值类型如数组和结构体（参见[第 10 章](https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/10.0.md)）；它相当于 `&T{}`。
    >
    > - `make(T)` **返回一个类型为 T 的初始值**，它只适用于 3 种内建的引用类型：切片、`map` 和 `channel`（参见[第 8 章](https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/08.0.md)和[第 13 章](https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/13.0.md)）。
    >
    > - <img src="https://github.com/unknwon/the-way-to-go_ZH_CN/raw/master/eBook/images/7.2_fig7.3.png?raw=true" alt="img" style="zoom:50%;" />
    >
    >   > *new 获取的是存储指定变量内存地址的一个变量，对于变量内部结构并不会执行相应的初始化操作， 所以 slice、map、channel 需要 make 进行初始化并获取对应的内存地址，而非 new 简单的获取内存地址*

- 如果想要往切片添加元素，可以使用`append`函数在切片末尾添加元素，切片会自动扩展。

- 多维切片

  和数组一样，切片通常也是一维的，但是也可以由一维组合成高维。通过分片的分片（或者切片的数组），长度可以任意动态变化，所以 Go 语言的多维切片可以任意切分。而且，内层的切片必须单独分配（通过 `make()` 函数）。

- bytes包

  类型 `[]byte` 的切片十分常见，Go 语言有一个 `bytes` 包专门用来提供这种类型的操作方法。

  `bytes` 包和字符串包十分类似，而且它还包含一个十分有用的类型 `Buffer`:

  ```go
  import "bytes"
  
  type Buffer struct {
  	...
  }
  ```

  - 这是一个长度可变的 `bytes` 的 buffer，提供 `Read()` 和 `Write()` 方法，因为读写长度未知的 `bytes` 最好使用 `buffer`。

  - `Buffer` 可以这样定义：`var buffer bytes.Buffer`。

    或者使用 `new()` 获得一个指针：`var r *bytes.Buffer = new(bytes.Buffer)`。

    或者通过函数：`func NewBuffer(buf []byte) *Buffer`，创建一个 `Buffer` 对象并且用 `buf` 初始化好；`NewBuffer` 最好用在从 `buf` 读取的时候使用。

  - **通过 buffer 串联字符串**

    我们创建一个 `buffer`，通过 `buffer.WriteString(s)` 方法将字符串 `s` 追加到后面，最后再通过 `buffer.String()` 方法转换为 `string`：

    ```go
    var buffer bytes.Buffer
    for {
    	if s, ok := getNextString(); ok { //method getNextString() not shown here
    		buffer.WriteString(s)
    	} else {
    		break
    	}
    }
    fmt.Print(buffer.String(), "\n")
    ```

    

注意 ，使用for_range遍历数组或者是切片时，第一个返回值 `ix` 是数组或者切片的索引，第二个是在该索引位置的值；他们都是仅在 `for` 循环内部可见的局部变量。**`value` 只是 `slice1` 某个索引位置的值的一个拷贝，不能用来修改 `slice1` 该索引位置的值。**



#### 切片的复制和追加

- `func append(s[]T, x ...T) []T`

  - 其中 `append()` 方法将 0 个或多个具有相同类型 `s` 的元素追加到切片后面并且返回新的切片；

  - 追加的元素必须和原切片的元素是同类型。

  - 如果 `s` 的容量不足以存储新增元素，`append()` 会分配新的切片来保证已有切片元素和新增元素的存储。因此，返回的切片可能已经指向一个不同的相关数组了。`append()` 方法总是返回成功，除非系统内存耗尽了。
  - 如果你想将切片 `y` 追加到切片 `x` 后面，只要将第二个参数扩展成一个列表即可：`x = append(x, y...)`。

​    

- `func copy(dst, src []T) int` 

  将类型为 `T` 的切片从源地址 `src` 拷贝到目标地址 `dst`，覆盖 `dst` 的相关元素，并且返回拷贝的元素个数。    

  ```go
  func AppendByte(slice []byte, data ...byte) []byte {
  	m := len(slice)
  	n := m + len(data)
  	if n > cap(slice) { // if necessary, reallocate
  		// allocate double what's needed, for future growth.
  		newSlice := make([]byte, (n+1)*2)
  		copy(newSlice, slice)
  		slice = newSlice
  	}
  	slice = slice[0:n]
  	copy(slice[m:n], data)
  	return slice
  }
  ```

  

#### 字符串、数组和切片的应用

- **从字符串生成字节切片**

  假设 `s` 是一个字符串（本质上是一个字节数组），那么就可以直接通过 `c := []byte(s)` 来获取一个字节的切片 `c` 。另外，您还可以通过 `copy()` 函数来达到相同的目的：`copy(dst []byte, src string)`。

  > 我们知道，Unicode 字符会占用 2 个字节，有些甚至需要 3 个或者 4 个字节来进行表示。如果发现错误的 UTF8 字符，则该字符会被设置为 `U+FFFD` 并且索引向前移动一个字节。和字符串转换一样，您同样可以使用 `c := []int32(s)` 语法，这样切片中的每个 `int` 都会包含对应的 Unicode 代码，因为字符串中的每次字符都会对应一个整数。类似的，您也可以将字符串转换为元素类型为 `rune` 的切片：`r := []rune(s)`。

- 获取字符串的某一部分

  字符串也支持切片

  使用 `substr := str[start:end]` 可以从字符串 str 获取到从索引 `start` 开始到 `end-1` 位置的子字符串。同样的，`str[start:]` 则表示获取从 `start` 开始到 `len(str)-1` 位置的子字符串。而 `str[:end]` 表示获取从 0 开始到 `end-1` 的子字符串。

- **字符串和切片的内存结构**

  在内存中，一个字符串实际上是一个双字结构，即一个指向实际数据的指针和记录字符串长度的整数

  切片的内存结构也是一样的

  <img src="https://github.com/unknwon/the-way-to-go_ZH_CN/raw/master/eBook/images/7.6_fig7.4.png" alt="img" style="zoom:50%;" />

- 修改字符串中的某个字符

  - Go 语言中的字符串是不可变的 和python是一样的

  - 因此，如果你想对字符串的某一个位置进行修改，就需要 先将字符串转换为字节数组 让后在进行修改 最后转换为字符串

    ```go
    s := "hello"
    c := []byte(s)
    c[0] = 'c'
    s2 := string(c) // s2 == "cello"
    ```



排序及搜索切片和数组

- 标准库提供了 `sort` 包来实现常见的搜索和排序操作。

  例如，您可以使用 `sort` 包中的函数 `func Ints(a []int)` 来实现对 `int` 类型的切片排序

  可以通过函数 `IntsAreSorted(a []int) bool` 来检查，如果返回 `true` 则表示已经被排序。

- 想要在数组或切片中搜索一个元素，该数组或切片必须先被排序（因为标准库的搜索算法使用的是二分法，然后，您就可以使用函数 `func SearchInts(a []int, n int) int` 进行搜索，并返回对应结果的索引值。

  同样，这里的函数还有，

  `func SearchFloat64s(a []float64, x float64) int`
  `func SearchStrings(a []string, x string) int`



**切片的底层指向一个数组，该数组的实际容量可能要大于切片所定义的容量。只有在没有任何切片指向的时候，底层的数组内存才会被释放，这种特性有时会导致程序占用多余的内存。**



### Map

`map` 是一种特殊的数据结构：**一种元素对 (pair) 的无序集合**，pair 的一个元素是 key，对应的另一个元素是 value，所以这个结构也称为关联数组或字典。这是一种快速寻找值的理想结构：给定 key，对应的 value 可以迅速定位。

`map` 是引用类型，可以使用如下声明：

- 未初始化的 `map` 的值是 `nil`。

- 在声明的时候不需要知道 `map` 的长度，`map` 是可以动态增长的。

- key 可以是任意可以用 `==` 或者 `!=` 操作符比较的类型，比如 `string`、`int`、`float32(64)`。所以数组、切片和结构体不能作为 key 

- value,我们可以存储任意值，但是使用这种类型作为值时需要先做一次类型断言

- **`map` 传递给函数的代价很小：在 32 位机器上占 4 个字节，64 位机器上占 8 个字节，无论实际上存储了多少数据。**

  通过 key 在 `map` 中寻找值是很快的，比线性查找快得多，但是仍然比从数组和切片的索引中直接读取要慢 100 倍；所以如果你很在乎性能的话还是建议用切片来解决问题。

- 常用的 `len(map1)` 方法可以获得 `map` 中的 pair 数目，这个数目是可以伸缩的，因为 map-pairs 在运行时可以动态添加和删除。

```go
var map1 map[keytype]valuetype
// 例如
var map1 map[string]int
```

#### 初始化方法

-  `map` 可以用 `{key1: val1, key2: val2}` 的描述方法来初始化，就像数组和结构体一样。

- `map` 是 **引用类型** 的： 内存用 `make()` 方法来分配。

  `map` 的初始化：`var map1 = make(map[keytype]valuetype)`。

  或者简写为：`map1 := make(map[keytype]valuetype)`。

- **不要使用 `new()`，永远用 `make()` 来构造 `map`**

  **注意** 如果你错误地使用 `new()` 分配了一个引用对象，你会获得一个空引用的指针，相当于声明了一个未初始化的变量并且取了它的地址：

- 和数组不同，`map` 可以根据新增的 key-value 对动态的伸缩，因此它不存在固定长度或者最大限制。

  但是你也可以选择标明 `map` 的初始容量 `capacity`，就像这样：`make(map[keytype]valuetype, cap)`。

- map的键值有明确的规定，但是，对于value没有明确的规定，

  因此，能够使用切片作为map的值

  ```go
  mp1 := make(map[int][]int)  // value 为 int切片
  mp2 := make(map[int]*[]int)	// value 为 int切片指针
  ```

#### 测试键值对是否存在

- 在map中，对于不存在的值。如果 `map` 中不存在 `key1`，`val1` 就是一个值类型的空值。

  `val1, isPresent = map1[key1]`

  `isPresent` 返回一个 `bool` 值：如果 `key1` 存在于 `map1`，`val1` 就是 `key1` 对应的 `value` 值，并且 `isPresent` 为 `true`；

  如果 `key1` 不存在，`val1` 就是一个空值，并且 `isPresent` 会返回 `false`。

  

- 删除键值

  从 `map1` 中删除 `key1`：直接 `delete(map1, key1)` 就可以。如果 `key1` 不存在，该操作不会产生错误。



#### map的遍历

- 可以使用 `for` 循环读取 `map`：

  同样，在遍历时，value也还是原map值的一个拷贝

  ```go
  for key, value := range map1 {
  	...
  }
  
  for _, value := range map1 {
  	...
  }
  
  for key := range map1 {
  	fmt.Printf("key is: %d\n", key)
  }
  ```

- **注意 `map` 不是按照 key 的顺序排列的，也不是按照 value 的序排列的。**



#### map类型的切片

- 假设我们想获取一个 `map` 类型的切片，我们必须使用两次 `make()` 函数，第一次分配切片，第二次分配切片中每个 `map` 元素。

  ```go
  package main
  import "fmt"
  
  func main() {
  	// Version A:
      // 与初始化 不同 初始化是 make(map[int]int,5)
      // 这里时创建 切片，切片的每一个部分都是map
  	items := make([]map[int]int, 5)
  	for i:= range items {
  		items[i] = make(map[int]int, 1)
  		items[i][1] = 2
  	}
  	fmt.Printf("Version A: Value of items: %v\n", items)
      // Version A: Value of items: [map[1:2] map[1:2] map[1:2] map[1:2] map[1:2]]
  
  	// Version B: NOT GOOD!
  	items2 := make([]map[int]int, 5)
  	for _, item := range items2 {
  		item = make(map[int]int, 1) // 同样，在遍历时，value也还是原map值的一个拷贝
  		item[1] = 2 // This 'item' will be lost on the next iteration.
  	}
  	fmt.Printf("Version B: Value of items: %v\n", items2)
      // Version B: Value of items: [map[] map[] map[] map[] map[]]
      
  }
  ```



#### map的排序

`map` 默认是无序的，不管是按照 key 还是按照 value 默认都不排序（详见第 8.3 节）。

如果你想为 `map` 排序，需要将 key（或者 value）拷贝到一个切片，再对切片排序

```go
// the telephone alphabet: 

package main
import (
	"fmt"
	"sort"
)

var (
	barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
							"delta": 87, "echo": 56, "foxtrot": 12,
							"golf": 34, "hotel": 16, "indio": 87,
							"juliet": 65, "kili": 43, "lima": 98}
)

func main() {
	fmt.Println("unsorted:")
	for k, v := range barVal {
		fmt.Printf("Key: %v, Value: %v / ", k, v)
	}
    
    // 逻辑就是 将所有的key 拿出来排序，按照排序后的key值输出 map[key]
	keys := make([]string, len(barVal))
	i := 0
	for k, _ := range barVal {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	fmt.Println()
	fmt.Println("sorted:")
	for _, k := range keys {
		fmt.Printf("Key: %v, Value: %v / ", k, barVal[k])
	}
}
```





#### map键值对调

- map的特性能够保证键值都是唯一的，如果希望进行 key和value的对调，首先要确定 value是唯一的，因为若是 value不唯一，对调后的 `map` 很可能没有包含原 `map` 的所有键值对！



## package

这里列举一些常用的包，和其中包含的不同方法



### regexp包

正则表达式包

- 匹配

  变量 `ok` 将返回 `true` 或者 `false`  以表示 是否匹配正确

  ```go
  ok, _ := regexp.Match(pat, []byte(searchIn))
  
  ok, _ := regexp.MatchString(pat, searchIn)
  
  ```

- compile

  `Compile()` 函数也可能返回一个错误，我们在使用时忽略对错误的判断是因为我们确信自己正则表达式是有效的。当用户输入或从数据中获取正则表达式的时候，我们有必要去检验它的正确性。

  ```go
  re, _ := regexp.Compile(pat)
  ```

### 锁和 sync 包

在一些复杂的程序中，通常通过不同线程执行不同应用来实现程序的并发。

`map` 类型是不存在锁的机制来实现这种效果（出于对性能的考虑），所以 map 类型是非线程安全的。当并行访问一个共享的 `map` 类型的数据，`map` 数据将会出错。

`sync.Mutex` 是一个互斥锁，它的作用是守护在临界区入口来确保同一时间只能有一个线程进入临界区。

假设 `info` 是一个需要上锁的放在共享内存中的变量。通过包含 `Mutex` 来实现的一个典型例子如下：

```go
import  "sync"

type Info struct {
	mu sync.Mutex
	// ... other fields, e.g.: Str string
}
```



如果一个函数想要改变这个变量可以这样写:

```go
func Update(info *Info) {
	info.mu.Lock()
    // critical section:
    info.Str = // new value
    // end critical section
    info.mu.Unlock()
}
```



还有一个很有用的例子是通过 `Mutex` 来实现一个可以上锁的共享缓冲器:

```go
gtype SyncedBuffer struct {
	lock 	sync.Mutex
	buffer  bytes.Buffer
}
```

### 精密计算和 big 包

对于整数的高精度计算 Go 语言中提供了 `big` 包，被包含在 `math` 包下：**有用来表示大整数的 `big.Int`** 和**表示大有理数的 `big.Rat` 类型**（可以表示为 2/5 或 3.1416 这样的分数，而不是无理数或 π）。**这些类型可以实现任意位类型的数字，只要内存足够大。**缺点是更大的内存和处理开销使它们使用起来要比内置的数字类型慢很多。



- 大的整型数字是通过 `big.NewInt(n)` 来构造的，其中 `n` 为 `int64` 类型整数。
- 而大有理数是通过 `big.NewRat(n, d)` 方法构造。`n`（分子）和 `d`（分母）都是 `int64` 型整数。
- 因为 Go 语言不支持运算符重载，所以所有大数字类型都有像是 `Add()` 和 `Mul()` 这样的方法。它们作用于作为 receiver 的整数和有理数，**大多数情况下它们修改 receiver 并以 receiver 作为返回结果**。因为没有必要创建 `big.Int` 类型的临时变量来存放中间结果，所以运算可以被链式地调用，并节省内存。

### 自定义包和可见性

```go
package pack1
var Pack1Int int = 42	// 变量大写开头 表示在引用时能够被其他文件读取
var pack1Float = 3.14	// 小写开头 这是私有变量

func ReturnStr() string {
	return "Hello main!"
}
```

```go
package main
import (
	"fmt"
	"./pack1"
)

func main() {
	var test1 string
	test1 = pack1.ReturnStr()
	fmt.Printf("ReturnStr from package1: %s\n", test1)
	fmt.Printf("Integer from package1: %d\n", pack1.Pack1Int)
	// fmt.Printf("Float from package1: %f\n", pack1.pack1Float)
}
```

- 如果包 pack1 和我们的程序在同一路径下，我们可以通过 `"import ./pack1"` 这样的方式来引入

- 为了区分，不同包存放在不同的目录下，每个包（所有属于这个包中的 go 文件）都存放在和包名相同的子目录下：

- 当使用 `.` 作为包的别名时，你可以不通过包名来使用其中的项目。例如：`test := ReturnStr()`。

  ```go
  import . "./pack1"
  ```

- Import with `_` :

  ```
  import _ "./pack1/pack1"
  ```

  `pack1` 包只导入其副作用，也就是说，只执行它的 `init()` 函数并初始化其中的全局变量

- 包中的`init()` 函数是不能被调用的。

>1、一个包能分成多个源文件么？
>
>这是可以的，但是需要
>
>- 所有文件必须在文件开头声明相同的包名。
>- 将多个源文件放在同一个目录中。
>
>- 当你使用 `go build` 或 `go install` 命令构建程序时，Go编译器会自动将同一包中的所有源文件一起编译。
>- 同一包中的源文件可以互相引用其中定义的变量、函数、结构体、接口等。而无需导入其他文件。



- godoc 工具在显示自定义包中的注释也有很好的效果：注释必须以 `//` 开始并无空行放在声明（包，类型，函数）前。godoc 会为每个文件生成一系列的网页。

  >- 命令行下进入目录下并输入命令：
  >
  >  `godoc -http=:6060 -goroot="."`
  >
  >（`.` 是指当前目录，`-goroot` 参数可以是 `/path/to/my/package1` 这样的形式指出 `package1` 在你源码中的位置或接受用冒号形式分隔的路径，无根目录的路径为相对于当前目录的相对路径）
  >
  >- 在浏览器打开地址：[http://localhost:6060](http://localhost:6060/)
  >
  >然后你会看到本地的 godoc 页面从左到右一次显示出目录中的包：

- 自定义包的目录结构

  ```go
  /home/user/goprograms
  	ucmain.go	(uc 包主程序)
  	Makefile (ucmain 的 makefile)
  	ucmain
  	src/uc	 (包含 uc 包的 go 源码)
  		uc.go
  	 	uc_test.go
  	 	Makefile (包的 makefile)
  	 	uc.a
  	 	_obj
  			uc.a
  		_test
  			uc.a
  	bin		(包含最终的执行文件)
  		ucmain
  	pkg/linux_amd64
  		uc.a	(包的目标文件)
  ```

  

### go install

go install 是 Go 中自动包安装工具：如需要将包安装到本地它会从远端仓库下载包：检出、编译和安装一气呵成。

- 使用命令安装：`go install tideland-cgl.googlecode.com/hg`。

- 可执行文件 `hg.a` 将被放到 `$GOROOT/pkg/linux_amd64/tideland-cgl.googlecode.com` 目录下，源码文件被放置在 `$GOROOT/src/tideland-cgl.googlecode.com/hg` 目录下，同样有个 `hg.a` 放置在 `_obj` 的子目录下。

- 现在就可以在 go 代码中使用这个包中的功能了，例如使用包名 cgl 导入：

```
import cgl "tideland-cgl.googlecode.com/hg"
```





## 结构和方法

### 结构体 struct

- 结构体是复合类型 (composite types)，当需要定义一个类型，它由一系列属性组成，每个属性都有自己的类型和值的时候，就应该使用结构体，它把数据聚集在一起。然后可以访问这些数据，就好像它是一个独立实体的一部分。结构体也是值类型，因此可以通过 **new** 函数来创建。
- 组成结构体类型的那些数据称为 **字段 (fields)**。每个字段都有一个类型和一个名字；在一个结构体中，字段名字必须是唯一的。

```go
type identifier struct {
    field1 type1
    field2 type2
    ...
}
// 1、field1 field2 都是

// 对于一些简单的结构体
type T struct {a, b int}
```

- 结构体的字段可以是任何类型，甚至是结构体本身，也可以是函数或是接口

- 结构体中的使用

  ```go
  var s T
  s.a = 5
  s.b = 8
  ```

- 使用new创建

  - 使用 `new()` 函数给一个新的结构体变量分配内存，**它返回指向已分配内存的指针**：`var t *T = new(T) `**这里 t的类型是 *T 表示指针**

  ```go
  var t *T
  t = new(T)
  
  t := new(T)
  ```

  - 同样，也能够 使用声明的方式，`var t T`也会给t分配内存，并零值化内存，**但是这个时候t是类型T**

  - 在这两种方式中，`t` 通常被称做类型 T 的一个实例 (instance) 或对象 (object)。调用字段的方法都是 . (这里两种调用的方式都是. 没有像c中做进一步的->)

    使用点号符可以获取结构体字段的值：`structname.fieldname`。

    在 Go 语言中这叫 **选择器 (selector)**。无论变量是一个结构体类型还是一个结构体类型指针，都使用同样的 **选择器符 (selector-notation)** 来引用结构体的字段：

  ```go
  package main
  
  import "fmt"
  
  type struct1 struct {
  	i1  int
  	f1  float32
  	str string
  }
  
  func main() {
  	ms := new(struct1)
  	ms.i1 = 10
  	ms.f1 = 15.5
  	ms.str = "Chris"
      
  	fmt.Printf("The int is: %d\n", ms.i1)
  	fmt.Printf("The float is: %f\n", ms.f1)
  	fmt.Printf("The string is: %s\n", ms.str)
      fmt.Printf("The string is: %s\n", (*ms).str)  // 这样也是合法的
  	fmt.Println(ms)
  
  	var mb struct1
  	mb.i1 = 1231
  	mb.f1 = 13.33
  	mb.str = "fuck you so hard"
  
  	fmt.Printf("The int is: %d\n", mb.i1)
  	fmt.Printf("The float is: %f\n", mb.f1)
  	fmt.Printf("The string is: %s\n", mb.str)
  	fmt.Println(ms)
  }
  ```

- 混合字面量语法初始化语法

  ```go
  type Interval struct {
      start int
      end   int
  }
  intr := Interval{0, 3}            (A)
  intr := Interval{end:5, start:1}  (B)
  intr := Interval{end:5}           (C)
  
  ```

- 类型 `struct1` 在定义它的包 `pack1` 中必须是唯一的，它的完全类型名是：`pack1.struct1`。

- Go 语言中，结构体和它所包含的数据在内存中是以连续块的形式存在的，即使结构体中嵌套有其他的结构体，这在性能上带来了很大的优势。

  ![img](https://github.com/unknwon/the-way-to-go_ZH_CN/raw/master/eBook/images/10.1_fig10.2.jpg?raw=true)

- **递归结构体**

  结构体类型可以通过引用自身来定义。这在定义链表或二叉树的元素（通常叫节点）时特别有用，此时节点包含指向临近节点的链接（地址）。如下所示，链表中的 `su`，树中的 `ri` 和 `le` 分别是指向别的节点的指针。

  ![img](https://github.com/unknwon/the-way-to-go_ZH_CN/raw/master/eBook/images/10.1_fig10.3.jpg?raw=true)

  ```go
  type Node struct {
      data    float64
      su      *Node
  }
  
  type Node struct {
      pr      *Node
      data    float64
      su      *Node
  }
  
  ```

  

- **结构体转换**

  Go 中的类型转换遵循严格的规则。当为结构体定义了一个 `alias` 类型时，此结构体类型和它的 `alias` 类型都有相同的底层类型，它们可以如示例 10.3 那样互相转换，同时需要注意其中非法赋值或转换引起的编译错误。

  ```go
  package main
  import "fmt"
  
  type number struct {
      f float32
  }
  
  type nr number   // alias type
  
  func main() {
      a := number{5.0}
      b := nr{5.0}
      // var i float32 = b   // compile-error: cannot use b (type nr) as type float32 in assignment
      // var i = float32(b)  // compile-error: cannot convert b (type nr) to type float32
      // var c number = b    // compile-error: cannot use b (type nr) as type number in assignment
      // needs a conversion:
      var c = number(b)
      fmt.Println(a, b, c)
  }
  
  ```

  

- **工厂模式初始化**

  Go 语言不支持面向对象编程语言中那样的构造子方法，但是可以很容易的在 Go 中实现 “构造子工厂”方法。为了方便通常会为类型定义一个工厂

  ```go
  type File struct {
      fd      int     // 文件描述符
      name    string  // 文件名
  }
  
  // 构造一个工厂，返回指针就行
  func NewFile(fd int, name string) *File {
      if fd < 0 {
          return nil
      }
  
      return &File{fd, name}
  }
  
  ```

  > - 通过之前提到的 将结构体的名称设置为以小写字母开头 那么别的文件无法使用，再通过定义大写字母开头的工厂，整个问题就能得到解决
  >
  > ```go
  > type matrix struct {
  >     ...
  > }
  > 
  > func NewMatrix(params) *matrix {
  >     m := new(matrix) // 初始化 m
  >     return m
  > }
  > ```
  >
  > - 再提一下 make和new的问题
  >   - 在 数组、切片、map时尽量使用make
  >   - 在对于自定义结构体是 使用 new
  >   - 在自定义结构体时，不能使用make，会在编译时报错，
  >   - 在map初始化时， new返回的是一个指向 `nil` 的指针，它尚未被分配内存，因此后续的操作会在运行时发生错误。

- 带标签的结构体

  结构体中的字段除了有名字和类型外，还可以有一个可选的标签 (tag)：它是一个附属于字段的字符串，可以是文档或其他的重要标记。标签的内容不可以在一般的编程中使用，只有包 `reflect` 能获取它。

  ```go
  package main
  
  import (
  	"fmt"
  	"reflect"
  )
  
  type TagType struct { // tags 指的是后面的这些
  	field1 bool   "An important answer"
  	field2 string "The name of the thing"
  	field3 int    "How much there are"
  }
  
  func main() {
  	tt := TagType{true, "Barak Obama", 1}
  	for i := 0; i < 3; i++ {
  		refTag(tt, i)
  	}
  }
  
  func refTag(tt TagType, ix int) {
  	ttType := reflect.TypeOf(tt)
  	ixField := ttType.Field(ix)
  	fmt.Printf("%v\n", ixField.Tag)
  }
  ```

  

- **匿名字段和内嵌结构体**

  结构体可以包含一个或多个 **匿名（或内嵌）字段**，即这些字段没有显式的名字，只有字段的类型是必须的，此时类型就是字段的名字。匿名字段本身可以是一个结构体类型，即 **结构体可以包含内嵌结构体**。

  **在一个结构体中对于每一种数据类型只能有一个匿名字段。**

  ```go
  package main
  
  import "fmt"
  
  type innerS struct {
  	in1 int
  	in2 int
  }
  
  type outerS struct {
  	b    int
  	c    float32
  	int  // anonymous field
  	innerS //anonymous field
  }
  
  func main() {
  	outer := new(outerS)
  	outer.b = 6
  	outer.c = 7.5
  	outer.int = 60
  	outer.in1 = 5
  	outer.in2 = 10
  
  	fmt.Printf("outer.b is: %d\n", outer.b)
  	fmt.Printf("outer.c is: %f\n", outer.c)
  	fmt.Printf("outer.int is: %d\n", outer.int)
  	fmt.Printf("outer.in1 is: %d\n", outer.in1)
  	fmt.Printf("outer.in2 is: %d\n", outer.in2)
  
  	// 使用结构体字面量
  	outer2 := outerS{6, 7.5, 60, innerS{5, 10}}
  	fmt.Println("outer2 is:", outer2)
  }
  ```

  

- 命名冲突

  当两个字段拥有相同的名字（可能是继承来的名字）时该怎么办呢？

  1. 外层名字会覆盖内层名字（但是两者的内存空间都保留），这提供了一种重载字段或方法的方式；
  2. 如果相同的名字在同一级别出现了两次，如果这个名字被程序使用了，将会引发一个错误（不使用没关系）。没有办法来解决这种问题引起的二义性，必须由程序员自己修正。

  ```go  
  // 例子：
  
  type A struct {a int}
  type B struct {a, b int}
  
  type C struct {A; B}
  var c C
  // 规则 2：使用 c.a 是错误的，到底是 c.A.a 还是 c.B.a 呢？会导致编译器错误：ambiguous DOT reference c.a disambiguate with either c.A.a or c.B.a。
  
  // 指明具体的类别下才能够使用
  ccc := C{A{1, 2}, B{3}}
  fmt.Println(ccc.A.a, ccc.B.a)
  ```



### 方法 method

在 Go 语言中，结构体就像是类的一种简化形式，**Go 方法是作用在接收者 (receiver) 上的一个函数，接收者是某种类型的变量**。因此方法是一种特殊类型的函数。

- 接收者类型可以是（几乎）任何类型，不仅仅是结构体类型：任何类型都可以有方法，甚至可以是函数类型，可以是 `int`、`bool`、`string` 或数组的别名类型。但是接收者不能是一个接口类型。

- 最后接收者不能是一个指针类型，但是它可以是任何其他允许类型的指针。

- 在go中，一个类型加上它的方法等价于面向对象中的一个类。一个重要的区别是：在 Go 中，类型的代码和绑定在它上面的方法的代码可以不放置在一起，它们可以存在在不同的源文件，唯一的要求是：它们必须是同一个包的。

- 对于类型 `T`（或 `*T`）上的所有方法的集合叫做类型 `T`（或 `*T`）的方法集 (method set)。

- **因为方法是函数，所以同样的，不允许方法重载**，即对于一个类型只能有一个给定名称的方法。

  **但是如果基于接收者类型，是有重载的：具有同样名字的方法可以在 2 个或多个不同的接收者类型上存在**，比如在同一个包里这么做是允许的：

  ```go
  func (a *denseMatrix) Add(b Matrix) Matrix
  func (a *sparseMatrix) Add(b Matrix) Matrix
  ```

- 定义方法的一般格式如下：

  ```go
  func (recv receiver_type) methodName(parameter_list) (return_value_list) { ... }
  ```

  > 其中，普通函数的定义方法为 `func funcname(parameter_list) (return_value_list) { ... }`
  >
  > 匿名函数函数的定义方法为 `func (parameter_list) (return_value_list) { ... }`

  - 在方法名之前，`func` 关键字之后的括号中指定 receiver。

  - 如果 `recv` 是 receiver 的实例，`Method1` 是它的方法名，那么方法调用遵循传统的 `object.name` 选择器符号：`recv.Method1()`。

  - 如果 `recv` 是一个指针，Go 会自动解引用。

  - 如果方法不需要使用 `recv` 的值，可以用 **_** 替换它，比如：

    ```go
    func (_ receiver_type) methodName(parameter_list) (return_value_list) { ... }
    ```

  - `recv` 就像是面向对象语言中的 `this` 或 `self`，但是 Go 中并没有这两个关键字。随个人喜好，你可以使用 `this` 或 `self` 作为 receiver 的名字。下面是一个结构体上的简单方法的例子：

    ```go
    package main
    
    import "fmt"
    
    type TwoInts struct {
    	a int
    	b int
    }
    func (tn *TwoInts) AddThem() int {
    	return tn.a + tn.b
    }
    
    func (tn *TwoInts) AddToParam(param int) int {
    	return tn.a + tn.b + param
    }
    
    func main() {
    	two1 := new(TwoInts)
    	two1.a = 12
    	two1.b = 10
    
    	fmt.Printf("The sum is: %d\n", two1.AddThem())
    	fmt.Printf("Add them to the param: %d\n", two1.AddToParam(20))
    
    	two2 := TwoInts{3, 4}
    	fmt.Printf("The sum is: %d\n", two2.AddThem())
    }
    
    // 对于非自定义类型
    package main
    
    import "fmt"
    
    type IntVector []int
    
    func (v IntVector) Sum() (s int) {
    	for _, x := range v {
    		s += x
    	}
    	return
    }
    
    func main() {
    	fmt.Println(IntVector{1, 2, 3}.Sum()) // 输出是6
    }
    
    ```

  - **类型和作用在它上面定义的方法必须在同一个包里定义**，这就是为什么不能在 `int`、`float32(64)` 或类似这些的类型上定义方法。试图在 `int` 类型上定义方法会得到一个编译错误：

    >但是有一个间接的方式：可以先定义该类型（比如：`int` 或 `float32(64)`）的别名类型，然后再为别名类型定义方法。或者像下面这样将它作为匿名类型嵌入在一个新的结构体中。当然方法只在这个别名类型上有效。
    >
    >```go
    >package main
    >
    >import (
    >	"fmt"
    >	"time"
    >)
    >
    >type myTime struct {
    >	time.Time //anonymous field
    >}
    >
    >func (t myTime) first3Chars() string {
    >	return t.Time.String()[0:3]
    >}
    >func main() {
    >	m := myTime{time.Now()}
    >	// 调用匿名 Time 上的 String 方法
    >	fmt.Println("Full time now:", m.String())
    >	// 调用 myTime.first3Chars
    >	fmt.Println("First 3 chars:", m.first3Chars())
    >}
    >
    >/* Output:
    >Full time now: Mon Oct 24 15:34:54 Romance Daylight Time 2011
    >First 3 chars: Mon
    >*/
    >```

    >## 函数和方法的区别
    >
    >- 函数将变量作为参数：`Function1(recv)`
    >
    >  方法在变量上被调用：`recv.Method1()`
    >
    >- 在接收者是指针时，方法可以改变接收者的值（或状态），这点函数也可以做到（当参数作为指针传递，即通过引用调用时，函数也可以改变参数的状态）。
    >
    >- **方法没有和数据定义（结构体）混在一起：它们是正交的类型；表示（数据）和行为（方法）是独立的。**
    >
    >- **接收者必须有一个显式的名字，这个名字必须在方法中被使用。**
    >
    >  `receiver_type` 叫做 **（接收者）基本类型**，这个类型必须在和方法同样的包中被声明。
    >
    >- 鉴于性能的原因，**`recv` 最常见的是一个指向 `receiver_type` 的指针**（因为我们不想要一个实例的拷贝，如果按值调用的话就会是这样），特别是在 receiver 类型是结构体时，就更是如此了。
    >
    >  对于某一个方法的调用，这里的指针和原始值好像没有什么更细致的区分，影响的只是是否能够修改原数据结构中的值
    >
    >  其中相互的取址或者解引用的转化是自动进行的
    >
    >  但是 使用指针进行传递时，由于指针的大小只为4或者8个字节，因此，性能更好。
    >
    >  ```go
    >  package main
    >  
    >  import (
    >  	"fmt"
    >  )
    >  
    >  type B struct {
    >  	thing int
    >  }
    >  
    >  func (b *B) change() { b.thing = 1 }
    >  
    >  func (b B) write() string { return fmt.Sprint(b) }
    >  
    >  func main() {
    >  	var b1 B // b1 是值
    >  	b1.change()
    >  	fmt.Println(b1.write())
    >  
    >  	b2 := new(B) // b2 是指针
    >  	b2.change()
    >  	fmt.Println(b2.write())
    >  }
    >  
    >  /* 输出：
    >  {1}
    >  {1}
    >  */
    >  ```
    >
    >  

- **函数、方法、结构体、结构体中的字段**

  只有这些都由名字控制，只有大写字母开头才能够在调入包时用于其他文件包，即导出

  因此，这里能够控制方法 使其作为外部调用时访问或者是设置的方法

  ```go
  package person
  
  type Person struct {
  	firstName string
  	lastName  string
  }
  
  func (p *Person) FirstName() string {
  	return p.firstName
  }
  
  func (p *Person) SetFirstName(newName string) {
  	p.firstName = newName
  }
  
  
  // 另一个包中导入使用
  package main
  
  import (
  	"./person"
  	"fmt"
  )
  
  func main() {
  	p := new(person.Person)
  	// p.firstName undefined
  	// (cannot refer to unexported field or method firstName)
  	// p.firstName = "Eric" 
  	p.SetFirstName("Eric")
  	fmt.Println(p.FirstName()) // Output: Eric
  }
  ```

- **并发访问对象**

  对象的字段（属性）**不应该由 2 个或 2 个以上的不同线程在同一时间去改变。**如果在程序发生这种情况，为了安全并发访问，可以使用包 `sync`（参考[第 9.3 节](https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/09.3.md)中的方法。在[第 14.17 节](https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/14.17)中我们会通过 goroutines 和 channels 探索另一种方式。



- **内嵌类型的方法和继承**

  **当一个匿名类型被内嵌在结构体中时，匿名类型的可见方法也同样被内嵌**，这在效果上等同于外层类型 **继承** 了这些方法：**将父类型放在子类型中来实现亚型**。

  这个机制提供了一种简单的方式来模拟经典面向对象语言中的子类和继承相关的效果，也类似 Ruby 中的混入 (mixin)。

  ```go
  type Engine interface {
  	Start()
  	Stop()
  }
  
  type Car struct {
  	Engine
  }
  
  
  func (c *Car) GoToWorkIn() {
  	// get in car
  	c.Start()
  	// drive to work
  	c.Stop()
  	// get out of car
  }
  
  //  在如一个例子，如
  package main
  
  import (
  	"fmt"
  	"math"
  )
  
  type Point struct {
  	x, y float64
  }
  
  func (p *Point) Abs() float64 {
  	return math.Sqrt(p.x*p.x + p.y*p.y)
  }
  
  type NamedPoint struct {
  	Point
  	name string
  }
  
  func main() {
  	n := &NamedPoint{Point{3, 4}, "Pythagoras"}
  	fmt.Println(n.Abs()) // 打印 5
  }
  
  // 可以覆写方法（像字段一样）：和内嵌类型方法具有同样名字的外层类型的方法会覆写内嵌类型对应的方法。
  // 在示例 10.18 method4.go 中添加：
  
  func (n *NamedPoint) Abs() float64 {
  	return n.Point.Abs() * 100.
  }
  ```

  **因为一个结构体可以嵌入多个匿名类型，所以实际上我们可以有一个简单版本的多重继承，就像：`type Child struct { Father; Mother}`**



### 在类型中嵌入功能  

在Go中，**当一个结构体被嵌入到另一个结构体中时，嵌入的结构体的字段和方法会被“提升”到外部结构体的命名空间中**。

这意味着你可以直接通过外部结构体访问嵌入结构体的字段和方法，就像它们是外部结构体本身的一部分一样。

主要有两种方法来实现在类型中嵌入功能：

A：聚合（或组合）：包含一个所需功能类型的具名字段。

```go
package main

import (
	"fmt"
)

type Log struct {
	msg string
}

type Customer struct {
	Name string
	log  *Log
}

func main() {
	c := new(Customer)
	c.Name = "Barak Obama"
	c.log = new(Log)
	c.log.msg = "1 - Yes we can!"
	// shorter
	c = &Customer{"Barak Obama", &Log{"1 - Yes we can!"}}
	c.Log().Add("2 - After me the world will be a better place!")
    
	fmt.Println(c.log)
	fmt.Println(c.Log())  

}

func (l *Log) Add(s string) {
	l.msg += "\n" + s
}

func (l *Log) String() string {
	return l.msg
}

func (c *Customer) Log() *Log {
	return c.log
}

```



B：内嵌：内嵌（匿名地）所需功能类型，像前一节 10.6.5 所演示的那样。

```go
package main

import (
	"fmt"
)

type Log struct {
	msg string
}

type Customer struct {
	Name string
	Log
}

func main() {
	c := &Customer{"Barak Obama", Log{"1 - Yes we can!"}}
	c.Add("2 - After me the world will be a better place!")
	fmt.Println(c)
	fmt.Println(c.String())

}

func (l *Log) Add(s string) {
	l.msg += "\n" + s
}

func (l *Log) String() string {
	return l.msg
}

func (c *Customer) String() string {
	return c.Name + "\nLog:" + fmt.Sprintln(c.Log.String())
}

```



### 多重继承问题

多重继承指的是类型获得多个父类型行为的能力，它在传统的面向对象语言中通常是不被实现的（C++ 和 Python 例外）。因为在类继承层次中，多重继承会给编译器引入额外的复杂度。但是在 Go 语言中，通过在类型中嵌入所有必要的父类型，可以很简单的实现多重继承。

```go
package main

import (
	"fmt"
)

type Camera struct{}

func (c *Camera) TakeAPicture() string {
	return "Click"
}

type Phone struct{}

func (p *Phone) Call() string {
	return "Ring Ring"
}

type CameraPhone struct {
	Camera
	Phone
}

func main() {
	cp := new(CameraPhone)
	fmt.Println("Our new CameraPhone exhibits multiple behaviors...")
	fmt.Println("It exhibits behavior of a Camera: ", cp.TakeAPicture())
	fmt.Println("It works like a Phone too: ", cp.Call())
}
```

### 通用方法和方法命名

在编程中一些基本操作会一遍又一遍的出现，比如打开 (Open)、关闭 (Close)、读 (Read)、写 (Write)、排序(Sort) 等等，并且它们都有一个大致的意思：打开 (Open)可以作用于一个文件、一个网络连接、一个数据库连接等等。具体的实现可能千差万别，但是基本的概念是一致的。

在 Go 语言中，通过使用接口，标准库广泛的应用了这些规则，在标准库中这些通用方法都有一致的名字，比如 `Open()`、`Read()`、`Write()`等。想写规范的 Go 程序，就应该遵守这些约定，给方法合适的名字和签名，就像那些通用方法那样。这样做会使 Go 开发的软件更加具有一致性和可读性。



### 类型的 String() 方法和格式化描述符

如果类型定义了 `String()` 方法，它会被用在 `fmt.Printf()` 中生成**默认的输出(%v参数)**：等同于使用格式化描述符 `%v` 产生的输出。

还有 `fmt.Print()` 和 `fmt.Println()` 也会自动使用 `String()` 方法。

```go
package main

import (
	"fmt"
	"strconv"
)

type TwoInts struct {
	a int
	b int
}

func main() {
	two1 := new(TwoInts)
	two1.a = 12
	two1.b = 10
	fmt.Printf("two1 is: %v\n", two1)		//two1 is: (12/10)
	fmt.Println("two1 is:", two1)			//two1 is: (12/10)
	fmt.Printf("two1 is: %T\n", two1)		//two1 is: *main.TwoInts  格式化描述符 %T 会给出类型的完全规格
	fmt.Printf("two1 is: %#v\n", two1)		//two1 is: &main.TwoInts{a:12, b:10} %#v 会给出实例的完整输出，包括它的字段（在程序自动生成 Go 代码时也很有用）。
}

func (tn *TwoInts) String() string {
	return "(" + strconv.Itoa(tn.a) + "/" + strconv.Itoa(tn.b) + ")"
}
```

不要在 `String()` 方法里面调用涉及 `String()` 方法的方法，它会导致意料之外的错误，比如下面的例子，它导致了一个无限递归调用（`TT.String()` 调用 `fmt.Sprintf`，而 `fmt.Sprintf` 又会反过来调用 `TT.String()`），很快就会导致内存溢出：

```go
type TT float64

func (t TT) String() string {
    return fmt.Sprintf("%v", t)
}
t.String()
```

### 垃圾回收和 SetFinalizer

Go 开发者不需要写代码来释放程序中不再使用的变量和结构占用的内存，在 Go 运行时中有一个独立的进程，即垃圾收集器 (GC)，会处理这些事情，它搜索不再使用的变量然后释放它们的内存。可以通过 `runtime` 包访问 GC 进程。

通过调用 `runtime.GC()` 函数可以显式的触发 GC，但这只在某些罕见的场景下才有用，比如当内存资源不足时调用 `runtime.GC()`，它会在此函数执行的点上立即释放一大片内存，此时程序可能会有短时的性能下降（因为 `GC` 进程在执行）。

**如果需要在一个对象 `obj` 被从内存移除前执行一些特殊操作，比如写到日志文件中，可以通过如下方式调用函数来实现：**

```go
runtime.SetFinalizer(obj, func(obj *typeObj))
```

`func(obj *typeObj)` 需要一个 `typeObj` 类型的指针参数 `obj`，特殊操作会在它上面执行。`func` 也可以是一个匿名函数。

在对象被 GC 进程选中并从内存中移除以前，`SetFinalizer` 都不会执行，即使程序正常结束或者发生错误。



## 接口

Go 语言不是一种 *“传统”* 的面向对象编程语言：它里面没有类和继承的概念。

但是 Go 语言里有非常灵活的 **接口** 概念，通过它可以实现很多面向对象的特性。接口提供了一种方式来 **说明** 对象的行为：如果谁能搞定这件事，它就可以用在这儿。

接口定义了一组方法（方法集），但是这些方法不包含（实现）代码：它们没有被实现（它们是抽象的）。接口里也不能包含变量。

### 定义与使用

```go
//通过如下格式定义接口：
type Namer interface {
    Method1(param_list) return_type
    Method2(param_list) return_type
    ...
}
```

> - （按照约定，只包含一个方法的）接口的名字由方法名加 `er` 后缀组成，例如 `Printer`、`Reader`、`Writer`、`Logger`、`Converter` 等等。还有一些不常用的方式（当后缀 `er` 不合适时），比如 `Recoverable`，此时接口名以 `able` 结尾，或者以 `I` 开头（像 `.NET` 或 `Java` 中那样）。
> - Go 语言中的接口都很简短，通常它们会包含 0 个、最多 3 个方法。
> - 在 Go 语言中接口可以有值，一个接口类型的变量或一个 **接口值** ：`var ai Namer`，`ai` 是一个多字（multiword）数据结构，它的值是 `nil`。它本质上是一个指针，虽然不完全是一回事。**指向接口值的指针是非法的，它们不仅一点用也没有，还会导致代码错误。**
> - 类型（比如结构体）可以实现某个接口的方法集；这个实现可以描述为，该类型的变量上的每一个具体方法所组成的集合，包含了该接口的方法集。实现了 `Namer` 接口的类型的变量可以赋值给 `ai`（即 `receiver` 的值），方法表指针（method table ptr）就指向了当前的方法实现。当另一个实现了 `Namer` 接口的类型的变量被赋给 `ai`，`receiver` 的值和方法表指针也会相应改变。

- **类型不需要显式声明它实现了某个接口：接口被隐式地实现。多个类型可以实现同一个接口**。

- **实现某个接口的类型（除了实现接口方法外）可以有其他的方法**。

- **一个类型可以实现多个接口**。

- **接口类型可以包含一个实例的引用， 该实例的类型实现了此接口（接口是动态类型）**。

- 即使接口在类型之后才定义，二者处于不同的包中，被单独编译：只要类型实现了接口中的方法，它就实现了此接口。

```go
package main

import "fmt"

// 结构体 Square 实现了接口 Shaper 。
type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func main() {
	sq1 := new(Square)
	sq1.side = 5

	var areaIntf Shaper
    // 将一个 Square 类型的变量赋值给一个接口类型的变量：areaIntf = sq1 。
    // 现在接口变量包含一个指向 Square 变量的引用，通过它可以调用 Square 上的方法 Area()
    // 接口变量里包含了接收者实例的值和指向对应方法表的指针。
    // 这就是go的多态实现。
	areaIntf = sq1
	// shorter,without separate declaration:
	// areaIntf := Shaper(sq1)
	// or even:
	// areaIntf := sq1
	fmt.Printf("The square has area: %f\n", areaIntf.Area())
}
```

> - 如果 `Shaper` 有另外一个方法 `Perimeter()`，但是 `Square` 没有实现它，即使没有人在 `Square` 实例上调用这个方法，编译器也会给出错误。
>
> ```go
> package main
> 
> import "fmt"
> 
> type Shaper interface {
> 	Area() float32
> }
> 
> type Square struct {
> 	side float32
> }
> 
> func (sq *Square) Area() float32 {
> 	return sq.side * sq.side
> }
> 
> type Rectangle struct {
> 	length, width float32
> }
> 
> func (r Rectangle) Area() float32 {
> 	return r.length * r.width
> }
> 
> func main() {
> 
> 	r := Rectangle{5, 3} // Area() of Rectangle needs a value
> 	q := &Square{5}      // Area() of Square needs a pointer
> 	// shapes := []Shaper{Shaper(r), Shaper(q)}
> 	// or shorter
> 	shapes := []Shaper{r, q}
> 	fmt.Println("Looping through shapes for area ...")
> 	for n, _ := range shapes {
> 		fmt.Println("Shape details: ", shapes[n])
> 		fmt.Println("Area of this shape is: ", shapes[n].Area())
> 	}
> }
> ```
>
> ```go
> package main
> 
> import "fmt"
> 
> type stockPosition struct {
> 	ticker     string
> 	sharePrice float32
> 	count      float32
> }
> 
> /* method to determine the value of a stock position */
> func (s stockPosition) getValue() float32 {
> 	return s.sharePrice * s.count
> }
> 
> type car struct {
> 	make  string
> 	model string
> 	price float32
> }
> 
> /* method to determine the value of a car */
> func (c car) getValue() float32 {
> 	return c.price
> }
> 
> /* contract that defines different things that have value */
> type valuable interface {
> 	getValue() float32
> }
> 
> func showValue(asset valuable) {
> 	fmt.Printf("Value of the asset is %f\n", asset.getValue())
> }
> 
> func main() {
> 	var o valuable = stockPosition{"GOOG", 577.20, 4}
> 	showValue(o)
> 	o = car{"BMW", "M3", 66500}
> 	showValue(o)
> }
> ```
>
> 

###  接口嵌套接口

一个接口可以包含一个或多个其他的接口，这相当于直接将这些内嵌接口的方法列举在外层接口中一样。

比如接口 `File` 包含了 `ReadWrite` 和 `Lock` 的所有方法，它还额外有一个 `Close()` 方法。

```go
type ReadWrite interface {
    Read(b Buffer) bool
    Write(b Buffer) bool
}

type Lock interface {
    Lock()
    Unlock()
}

type File interface {
    ReadWrite
    Lock
    Close()
}
```

### 类型断言

接口变量只有被赋值后 才会达到类似多态的运行效果，因此 需要监测对象的赋值情况。

- 测试 `areaIntf` 里是否有一个包含 `*Square` 类型的变量，结果是确定的；然后我们测试它是否包含一个 `*Circle` 类型的变量，结果是否定的。
- 如果忽略 `areaIntf.(*Square)` 中的 `*` 号，会导致编译错误

- `t,ok = varI.(*type)` 当ok为1 时表示含有变量信息

```go
package main

import (
	"fmt"
	"math"
)
type Square struct {
	side float32
}
type Circle struct {
	radius float32
}
type Shaper interface {
	Area() float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (ci *Circle) Area() float32 {
	return ci.radius * ci.radius * math.Pi
}

func main() {
	var areaIntf Shaper
	sq1 := new(Square)
	sq1.side = 5

	areaIntf = sq1
	// Is Square the type of areaIntf?
	if t, ok := areaIntf.(*Square); ok {
		fmt.Printf("The type of areaIntf is: %T\n", t)
	}
	if u, ok := areaIntf.(*Circle); ok {
		fmt.Printf("The type of areaIntf is: %T\n", u)
	} else {
		fmt.Println("areaIntf does not contain a variable of type Circle")
	}
}

```

- 接口变量的类型也可以使用一种特殊形式的 `switch` 来检测：**type-switch** 

```go
switch t := varI.(type) {
case *Square:
	fmt.Printf("Type Square %T with value %v\n", t, t)
case *Circle:
	fmt.Printf("Type Circle %T with value %v\n", t, t)
case nil:
	fmt.Printf("nil value: nothing to check?\n")
default:
	fmt.Printf("Unexpected type %T\n", t)
}
```

### 方法集和接口

假定 `v` 是一个值，然后我们想测试它是否实现了 `Stringer` 接口，可以这样做：

```go
type Stringer interface {
    String() string
}

if sv, ok := v.(Stringer); ok {
    fmt.Printf("v implements String(): %s\n", sv.String()) // note: sv, not v
}
```

接口是一种契约，实现类型必须满足它，它描述了类型的行为，规定类型可以做什么。接口彻底将类型能做什么，以及如何做分离开来，使得相同接口的变量在不同的时刻表现出不同的行为，这就是多态的本质。



在方法的那一节上，我们能够知道，作用于变量上的方法实际上并没有明确区分 指针和实例

但是当碰到接口类型值时，这会变得有点复杂，原因是接口变量中存储的具体值是不可寻址的，幸运的是，如果使用不当编译器会给出错误。考虑下面的程序：

```go
package main

import (
	"fmt"
)

type List []int

func (l List) Len() int {
	return len(l)
}

func (l *List) Append(val int) {
	*l = append(*l, val)
}

type Appender interface {
	Append(int)
}

func CountInto(a Appender, start, end int) {
	for i := start; i <= end; i++ {
		a.Append(i)
	}
}

type Lener interface {
	Len() int
}

func LongEnough(l Lener) bool {
	return l.Len()*10 > 42
}

func main() {
	// A bare value
	var lst List
	// 在 lst 上调用 CountInto 时会导致一个编译器错误，因为 CountInto 需要一个 Appender，而它的方法 Append 只定义在指针上。
	// CountInto(lst, 1, 10)
    
    // 在 lst 上调用 LongEnough 是可以的，因为 Len 定义在值上。
	if LongEnough(lst) { // VALID:Identical receiver type
		fmt.Printf("- lst is long enough\n")
	}

	// 在 plst 上调用 CountInto 是可以的，因为 CountInto 需要一个 Appender，并且它的方法 Append 定义在指针上。 在 plst 上调用 LongEnough 也是可以的，因为指针会被自动解引用。
	plst := new(List)
	CountInto(plst, 1, 10) //VALID:Identical receiver type
	if LongEnough(plst) {
		// VALID: a *List can be dereferenced for the receiver
		fmt.Printf("- plst is long enough\n")
	}
}
```

这里有点绕，首先，不看接口的实现，单单看方法的实现，

`func (l *List) Append(val int) {}`、`func (l List) Len() int {}` 如果没有接口的实现，对于传入的是 *List 还是 List 都没有其他的区别，因为 go中会自动进行取址或这是解引用的手段以匹配函数

但是如果将接口和方法结合在一起，

- 首先，`CountInto` 需要一个 `Appender` 类型的参数，因此你传递给 `CountInto` 的值必须满足 `Appender` 接口。

- `Appender` 接口要求实现 `Append(int)` 方法。

  由于 `List` 的 `Append` 方法的接收者是 `*List`，所以只有 `*List` 类型满足 `Appender` 接口，而 `List` 类型不满足。

- 而`Lener` 接口要求实现 `Len()` 方法。

  `List` 类型的 `Len()` 方法的接收者是值类型，因此 `List` 和 `*List` 都满足 `Lener` 接口。这就是为什么你可以通过 `lst` 和 `plst` 都调用 `Len()` 方法，并且都能够匹配 `Lener` 接口的原因。



> 这里继续对上面`func CountInto(a Appender, start, end int){}`、 `func LongEnough(l Lener) bool {}` 这两个函数进行坍探讨
>
> - 首先，他们的定义格式 表明了他是一个函数，并且 其中的某一个参数类型为 接口
>
> - 还记得上一节的例子吗，对于一个接口变量，能够将其赋值为实现了其所有接口的某一个类型的实例，接着根据这样的方法实现多态
>
> - **接口作为函数参数**
>
>   **当一个函数接受一个接口类型作为参数时，该函数可以接受任何实现了这个接口的类型的实例。并且正确调用该实例实现的方法，这也是多态的实现**
>
>   这是因为接口定义的是一组方法的集合，只要某个类型实现了这些方法，它就可以被赋值给该接口类型的变量。
>   
> - 另外，正如上面说的，如果这些以接口为入参的函数中，使用了不同接收类型的 方法
>
>   那么，如果函数中调用的方法为T接收类型的方法，那么该函数能够接收 T 或者是 T* 的 类型的实例
>
>   如果函数中调用的方法为T* 接收类型的方法，那么该函数只能够接收 T* 的 类型的实例
>
> - 那么，这样 通过 接口 能够实现类似 多态的行为
>
>   通过 接口+类型方法实现+入参为接口的函数 就能够实现类似 函数模板的功能。
>





> 几个接口的例子
>
> ### 1、使用sorter接口排序
>
> 一个很好的例子是来自标准库的 `sort` 包，要对一组数字或字符串排序，只需要实现三个方法：反映元素个数的 `Len()`方法、比较第 `i` 和 `j` 个元素的 `Less(i, j)` 方法以及交换第 `i` 和 `j` 个元素的 `Swap(i, j)` 方法。
>
> ```go
> type Sorter interface {
>  Len() int
>  Less(i, j int) bool
>  Swap(i, j int)
> }
> 
> func Sort(data Sorter) {
>  for pass := 1; pass < data.Len(); pass++ {
>      for i := 0;i < data.Len() - pass; i++ {
>          if data.Less(i+1, i) {
>              data.Swap(i, i + 1)
>          }
>      }
>  }
> }
> ```
>
> 现在如果我们想对一个 `int` 数组进行排序，所有必须做的事情就是：为数组定一个类型并在它上面实现 `Sorter` 接口的方法：
>
> ```go
> type IntArray []int
> func (p IntArray) Len() int           { return len(p) }
> func (p IntArray) Less(i, j int) bool { return p[i] < p[j] }
> func (p IntArray) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
> 
> 
> // 调用
> data := []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
> a := sort.IntArray(data) //conversion to type IntArray from package sort
> sort.Sort(a)
> ```
>
> 那么，将其封装为包，就能够跨文件调用
>
> ```go
> package sort
> 
> type Sorter interface {
> 	Len() int
> 	Less(i, j int) bool
> 	Swap(i, j int)
> }
> 
> func Sort(data Sorter) {
> 	for pass := 1; pass < data.Len(); pass++ {
> 		for i := 0; i < data.Len()-pass; i++ {
> 			if data.Less(i+1, i) {
> 				data.Swap(i, i+1)
> 			}
> 		}
> 	}
> }
> 
> func IsSorted(data Sorter) bool {
> 	n := data.Len()
> 	for i := n - 1; i > 0; i-- {
> 		if data.Less(i, i-1) {
> 			return false
> 		}
> 	}
> 	return true
> }
> 
> // Convenience types for common cases
> type IntArray []int
> 
> func (p IntArray) Len() int           { return len(p) }
> func (p IntArray) Less(i, j int) bool { return p[i] < p[j] }
> func (p IntArray) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
> 
> type StringArray []string
> 
> func (p StringArray) Len() int           { return len(p) }
> func (p StringArray) Less(i, j int) bool { return p[i] < p[j] }
> func (p StringArray) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
> 
> // Convenience wrappers for common cases
> func SortInts(a []int)       { Sort(IntArray(a)) }
> func SortStrings(a []string) { Sort(StringArray(a)) }
> 
> func IntsAreSorted(a []int) bool       { return IsSorted(IntArray(a)) }
> func StringsAreSorted(a []string) bool { return IsSorted(StringArray(a)) }
> ```
>
> ```go
> package main
> 
> import (
> 	// 这里的导入路径是相对于你的 `GOPATH` 或者 Go module 的根目录
> 	sort "awesomeProject/sort"
> 	"fmt"
> )
> 
> func ints() {
> 	data := []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
> 	a := sort.IntArray(data) //conversion to type IntArray
> 	sort.Sort(a)
> 	if !sort.IsSorted(a) {
> 		panic("fails")
> 	}
> 	fmt.Printf("The sorted array is: %v\n", a)
> }
> 
> func strings() {
> 	data := []string{"monday", "friday", "tuesday", "wednesday", "sunday", "thursday", "", "saturday"}
> 	a := sort.StringArray(data)
> 	sort.Sort(a)
> 	if !sort.IsSorted(a) {
> 		panic("fail")
> 	}
> 	fmt.Printf("The sorted array is: %v\n", a)
> }
> 
> type day struct {
> 	num       int
> 	shortName string
> 	longName  string
> }
> 
> type dayArray struct {
> 	data []*day
> }
> 
> func (p *dayArray) Len() int           { return len(p.data) }
> func (p *dayArray) Less(i, j int) bool { return p.data[i].num < p.data[j].num }
> func (p *dayArray) Swap(i, j int)      { p.data[i], p.data[j] = p.data[j], p.data[i] }
> 
> func days() {
> 	Sunday := day{0, "SUN", "Sunday"}
> 	Monday := day{1, "MON", "Monday"}
> 	Tuesday := day{2, "TUE", "Tuesday"}
> 	Wednesday := day{3, "WED", "Wednesday"}
> 	Thursday := day{4, "THU", "Thursday"}
> 	Friday := day{5, "FRI", "Friday"}
> 	Saturday := day{6, "SAT", "Saturday"}
> 	data := []*day{&Tuesday, &Thursday, &Wednesday, &Sunday, &Monday, &Friday, &Saturday}
> 	a := dayArray{data}
> 	// 因为 dayArray 的方法的参数 都为 *dayArray 所以这里只支持 指针类型的参数传入
> 	sort.Sort(&a)
> 	if !sort.IsSorted(&a) {
> 		panic("fail")
> 	}
> 	for _, d := range data {
> 		fmt.Printf("%s ", d.longName)
> 	}
> 	fmt.Printf("\n")
> }
> 
> func main() {
> 	ints()
> 	strings()
> 	days()
> }
> ```
>
> 2、读取和写入
>
> `io` 包提供了用于读和写的接口 `io.Reader` 和 `io.Writer`：
>
> ```go
> type Reader interface {
>     Read(p []byte) (n int, err error)
> }
> 
> type Writer interface {
>     Write(p []byte) (n int, err error)
> }
> ```
>
> 只要类型实现了读写接口，提供 `Read` 和 `Write` 方法，就可以从它读取数据，或向它写入数据。
>
> 一个对象要是可读的，它必须实现 `io.Reader` 接口，这个接口只有一个签名是 `Read(p []byte) (n int, err error)` 的方法，它从调用它的对象上读取数据，并把读到的数据放入参数中的字节切片中，然后返回读取的字节数和一个 `error` 对象，如果没有错误发生返回 `nil`，如果已经到达输入的尾端，会返回 `io.EOF("EOF")`，如果读取的过程中发生了错误，就会返回具体的错误信息。
>
> 类似地，一个对象要是可写的，它必须实现 `io.Writer` 接口，这个接口也只有一个签名是 `Write(p []byte) (n int, err error)` 的方法，它将指定字节切片中的数据写入调用它的对象里，然后返回实际写入的字节数和一个 `error` 对象（如果没有错误发生就是 `nil`）。





### 空接口

#### 概念

- **空接口或者最小接口** 不包含任何方法，它对实现不做任何要求：

```go
type Any interface {}
```

- **任何其他类型都实现了空接口**，`any` 或 `Any` 是空接口一个很好的别名或缩写。

  因此，对于 一个接口变量，能够赋予其任何类型的值

  ```go
  package main
  import "fmt"
  
  var i = 5
  var str = "ABC"
  
  type Person struct {
  	name string
  	age  int
  }
  
  type Any interface{}
  
  func main() {
  	var val Any
  	val = 5
  	fmt.Printf("val has the value: %v\n", val)
  	val = str
  	fmt.Printf("val has the value: %v\n", val)
  	pers1 := new(Person)
  	pers1.name = "Rob Pike"
  	pers1.age = 55
  	val = pers1
  	fmt.Printf("val has the value: %v\n", val)
  	switch t := val.(type) {
  	case int:
  		fmt.Printf("Type int %T\n", t)
  	case string:
  		fmt.Printf("Type string %T\n", t)
  	case bool:
  		fmt.Printf("Type boolean %T\n", t)
  	case *Person:
  		fmt.Printf("Type pointer to Person %T\n", t)
  	default:
  		fmt.Printf("Unexpected type %T", t)
  	}
  }
  ```

  > 每个 `interface {}` 变量在内存中占据两个字长：一个用来存储它包含的类型，另一个用来存储它包含的数据或者指向数据的指针。

#### 构建通用类型或包含不同类型变量的数组

就是通过使用空接口。让我们给空接口定一个别名类型 `Element`：`type Element interface{}`

然后定义一个容器类型的结构体 `Vector`，它包含一个 `Element` 类型元素的切片：

```go
type Element interface{}
type Vector struct {
	a []Element
}

// 定义一个 `At()` 方法用于返回第 `i` 个元素：
func (p *Vector) At(i int) Element {
	return p.a[i]
}

// 再定一个 Set() 方法用于设置第 i 个元素的值：
func (p *Vector) Set(i int, e Element) {
	p.a[i] = e
}
```

- 因为任何类型都是西安了空接口，因此，`Vector` 里能放任何类型的变量

- `Vector` 中存储的所有元素都是 `Element` 类型，要得到它们的原始类型（unboxing：拆箱）需要用到类型断言。

- **复制数据切片至空接口切片**

  假设你有一个 `myType` 类型的数据切片，你想将切片中的数据复制到一个空接口切片中，因为它们俩在内存中的布局是不一样的，所以这里不能够直接赋值

  这里需要使用 通过遍历的方式 一个一个的显式赋值>

  ```go
  var dataSlice []myType = FuncReturnSlice()
  var interfaceSlice []interface{} = make([]interface{}, len(dataSlice))
  
  // 错误 不能够直接赋值
  // var interfaceSlice []interface{} = dataSlice
  
  
  for i, d := range dataSlice {
      interfaceSlice[i] = d
  }
  ```



### 接口到接口

**Go语言中的接口到接口的转换实际上是接口类型的赋值，而不仅仅是转换。这种赋值是基于接口的实现情况进行的。**

#### 基本概念

1. **接口赋值**：一个接口类型可以赋值为另一个接口类型，只要被赋值的类型实现了目标接口中的所有方法。
2. **空接口**：`interface{}` 是Go语言中的空接口，它可以持有任何类型的值，因为它没有定义任何方法。空接口常用于需要接受任意类型的情况。

> 接口之间的转换并不是直接的赋值，而是通过类型断言和接口实现来完成的。
>
> 通过类型断言，你可以将一个接口类型的变量转换为另一个接口类型
>
> 实际上接口的转换依赖于 `a, ok := b.(A);` 这样的类型检查功能
>
> **类型断言**：
>
> - **成功**：如果 `empty` 实际上持有的值实现了 `A` 接口，那么转换会成功。此时，`ok` 的值为 `true`，并且 `ai` 将持有 `empty` 转换后的 `A` 类型的值。
>
>   **失败**：如果 `empty` 实际上持有的值没有实现 `A` 接口，那么转换会失败。此时，`ok` 的值为 `false`，并且 `ai` 的值将是接口的零值，即 `nil`。
>
> 另一种方法 **type switch**
>
> > 只会提取出类型于case中的各个类型进行比较，
> >
> > 在这个例子中，**`v`** 是接口值的具体类型，在每个 `case` 分支中，`v` 的类型与该 `case` 条件中指定的类型一致。
> >
> > ```go
> > switch v := i.(type) {
> > case int:
> >     fmt.Println("int:", v)
> > case string:
> >     fmt.Println("string:", v)
> > case bool:
> >     fmt.Println("bool:", v)
> > default:
> >     fmt.Println("unknown type")
> > }
> > ```

```go
package main

import "fmt"

// 定义接口A
type A interface {
    MethodA() string
}

// 定义接口B
type B interface {
    MethodB() string
}

// 实现接口A的类型
type MyType struct{}

func (m MyType) MethodA() string {
    return "MethodA"
}

func (m MyType) MethodB() string {
    return "MethodB"
}

func main() {
    var a A
    var b B
    myType := MyType{}

    // 将 MyType 赋值给接口A
    a = myType

    // 由于 myType 同时实现了 A 和 B 接口，可以将 a 转换为 B
    if b, ok := a.(B); ok {
        fmt.Println(b.MethodB()) // Output: MethodB
    } else {
        fmt.Println("a does not implement B")
    }

    // 将 MyType 赋值给接口B
    b = myType

    // 由于 myType 同时实现了 A 和 B 接口，可以将 b 转换为 A
    if a, ok := b.(A); ok {
        fmt.Println(a.MethodA()) // Output: MethodA
    } else {
        fmt.Println("b does not implement A")
    }
    // 空接口变量 在经过赋值后
    var empty interface{}
    empty = myType
 
    // 应为 Mytype{}同时实现了A接口和B接口 因此 能够通过类型断言
    
    ai, ok := empty.(A)
    if ok {
        fmt.Println(ai.MethodA()) // Output: MethodA
    } else {
        fmt.Println("empty does not implement A")
    }

    // 类型断言从空接口中提取接口B
    bi, ok := empty.(B)
    if ok {
        fmt.Println(bi.MethodB()) // Output: MethodB
    } else {
        fmt.Println("empty does not implement B")
    }
}
```



### 接口与动态类型

- Go 没有类：数据（结构体或更一般的类型）和方法是一种松耦合的正交关系。

- Go中，接收一个（或多个）接口类型作为参数的函数，其**实参**可以是任何实现了该接口的类型的变量。 *实现了某个接口的类型可以被传给任何以此接口为参数的函数*。

  这意味着对象可以根据提供的方法被处理（例如，作为参数传递给函数），而忽略它们的实际类型：它们能做什么比它们是什么更重要。

  ```go
  package main
  
  import "fmt"
  
  type IDuck interface {
  	Quack()
  	Walk()
  }
  
  func DuckDance(duck IDuck) {
  	for i := 1; i <= 3; i++ {
  		duck.Quack()
  		duck.Walk()
  	}
  }
  
  type Bird struct {
  	// ...
  }
  
  func (b *Bird) Quack() {
  	fmt.Println("I am quacking!")
  }
  
  func (b *Bird) Walk()  {
  	fmt.Println("I am walking!")
  }
  
  func main() {
  	b := new(Bird)
  	DuckDance(b)
  }
  ```

- **Go 的动态实现，通常需要编译器静态检查的支持**：**当变量被赋值给一个接口类型的变量时，编译器会检查其是否实现了该接口的所有函数**。

  如果方法调用作用于像 `interface{}` 这样的“泛型”上，你可以通过类型断言来检查变量是否实现了相应接口。

  因此 Go 提供了动态语言的优点，却没有其他动态语言在运行时可能发生错误的缺点。

  ```go
  // 例如，我们希望使用不同的类型表示 XML 输出流中的不同实体。然后我们为 XML 定义一个如下的“写”接口（甚至可以把它定义为私有接口）：
  type xmlWriter interface {
  	WriteXML(w io.Writer) error
  }
  // 现在我们可以实现适用于该流类型的任何变量的 StreamXML() 函数，并用类型断言检查传入的变量是否实现了该接口；如果没有，我们就调用内建的 encodeToXML() 来完成相应工作：
  
  // Exported XML streaming function.
  func StreamXML(v interface{}, w io.Writer) error {
  	if xw, ok := v.(xmlWriter); ok {
  		// It’s an  xmlWriter, use method of asserted type.
  		return xw.WriteXML(w)
  	}
  	// No implementation, so we have to use our own function (with perhaps reflection):
  	return encodeToXML(v, w)
  }
  
  // Internal XML encoding function.
  func encodeToXML(v interface{}, w io.Writer) error {
  	// ...
  }
  ```

  

### 接口的提取

Go 接口可以让开发者找出自己写的程序中的类型。**假设有一些拥有共同行为的对象，并且开发者想要抽象出这些行为，这时就可以创建一个接口来使用。**

```go
//multi_interfaces_poly.go
package main
import "fmt"
type Shaper interface {
	Area() float32
}
type TopologicalGenus interface {
	Rank() int
}
type Square struct {
	side float32
}
func (sq *Square) Area() float32 {
	return sq.side * sq.side
}
func (sq *Square) Rank() int {
	return 1
}
type Rectangle struct {
	length, width float32
}
func (r Rectangle) Area() float32 {
	return r.length * r.width
}
func (r Rectangle) Rank() int {
	return 2
}
func main() {
	r := Rectangle{5, 3} // Area() of Rectangle needs a value
	q := &Square{5}      // Area() of Square needs a pointer
	shapes := []Shaper{r, q}
	fmt.Println("Looping through shapes for area ...")
	for n, _ := range shapes {
		fmt.Println("Shape details: ", shapes[n])
		fmt.Println("Area of this shape is: ", shapes[n].Area())
	}
	topgen := []TopologicalGenus{r, q}
	fmt.Println("Looping through topgen for rank ...")
	for n, _ := range topgen {
		fmt.Println("Shape details: ", topgen[n])
		fmt.Println("Topological Genus of this shape is: ", topgen[n].Rank())
	}
}
```



### 显式地指明类型实现了某个接口

如果你希望满足某个接口的类型显式地声明它们实现了这个接口，你可以向接口的方法集中添加一个具有描述性名字的方法。例如：

```go
type Fooer interface {
	Foo()
	ImplementsFooer()
}
```

类型 Bar 必须实现 `ImplementsFooer` 方法来满足 `Fooer` 接口，以清楚地记录这个事实。

```go
type Bar struct{}
func (b Bar) ImplementsFooer() {}
func (b Bar) Foo() {}
```

空接口和函数重载

在 Go 语言中函数重载可以用可变参数 `...T` 作为函数最后一个参数来实现。

如果我们把 `T` 换为空接口，那么可以知道任何类型的变量都是满足 `T` (空接口）类型的，这样就允许我们传递任何数量任何类型的参数给函数，即重载的实际含义。

```go
fmt.Printf(format string, a ...interface{}) (n int, errno error)
```



### 接口的继承

当一个类型包含（内嵌）另一个类型（实现了一个或多个接口）的指针时，这个类型就可以使用（另一个类型）所有的接口方法。

有用的接口可以在开发的过程中被归纳出来。添加新接口非常容易，因为已有的类型不用变动（仅仅需要实现新接口的方法）。
已有的函数可以扩展为使用接口类型的约束性参数：通常只有函数签名需要改变。对比基于类的 OO 类型的语言在这种情况下则需要适应整个类层次结构的变化。

## 反射

- 反射是**用程序检查其所拥有的结构**，尤其是类型的一种能力；这是元编程的一种形式。

- 反射可以**在运行时检查类型和变量**，例如：它的大小、它的方法以及它能“动态地”调用这些方法。这对于没有源代码的包尤其有用。这是一个强大的工具，除非真得有必要，否则应当避免使用或小心使用。

- 变量的最基本信息就是**类型**和**值**：反射包的 `Type` 用来表示一个 Go 类型，反射包的 `Value` 为 Go 值提供了反射接口。

- 两个简单的函数，`reflect.TypeOf` 和 `reflect.ValueOf`，返回被检查对象的类型和值。

  ```go
  func TypeOf(i interface{}) Type
  func ValueOf(i interface{}) Value
  ```

  另一个是 `Type` 和 `Value` 都有 `Kind()` 方法返回一个常量来表示类型：`Uint`、`Float64`、`Slice` 等等。同样 `Value` 有叫做 `Int()` 和 `Float()` 的方法可以获取存储在内部的值（跟 `int64` 和 `float64` 一样）

  `reflect.ValueOf(x).Float()` 返回这个 `float64` 类型的实际值；同样的适用于 `Int(), Bool(), Complex(), String()`

  ```go
  var m myInt
  m = 3
  
  fmt.Println(reflect.TypeOf(m))		// main.myInt
  fmt.Println(reflect.ValueOf(m))		// 3
  
  v := reflect.TypeOf(m)
  	
  fmt.Println(v.Kind())				// int
  
  cc := float32(1.332)
  fmt.Println(reflect.TypeOf(cc))		// float32
  fmt.Println(reflect.ValueOf(cc))	// 1.332
  ```

### 通过反射修改值

- 如果要修改值，必须传递变量的指针。
- 使用 `Elem()` 方法获取指针指向的值，这样才能修改实际值。
- 使用相应的 `Set` 方法来设置新值，比如 `SetInt`、`SetString` 等。、
- **`CanSet()` 检查**：使用 `CanSet()` 检查 `reflect.Value` 是否可以修改。如果 `CanSet()` 返回 `false`，意味着当前值不可修改。
- **类型安全**：反射操作绕过了Go的静态类型检查，使用时要格外小心，确保类型匹配，以避免运行时错误

  ```go
  v := reflect.ValueOf(x)
  
  // setting a value:
  // v.SetFloat(3.1415) // Error: will panic: reflect.Value.SetFloat using unaddressable value
  
  fmt.Println("settability of v:", v.CanSet())
  v = reflect.ValueOf(&x) // Note: take the address of x.
  fmt.Println("type of v:", v.Type())
  fmt.Println("settability of v:", v.CanSet())
  v = v.Elem()
  fmt.Println("The Elem of v is: ", v)
  fmt.Println("settability of v:", v.CanSet())
  v.SetFloat(3.1415) // this works!
  fmt.Println(v.Interface())
  fmt.Println(v)
  ```



### 反射结构类型

有些时候需要反射一个结构类型。`NumField()` 方法返回结构内的字段数量；通过一个 `for` 循环用索引取得每个字段的值 `Field(i)`。

```go  
package main

import (
	"fmt"
	"reflect"
)

type NotknownType struct {
	s1, s2, s3 string
}

func (n NotknownType) String() string {
	return n.s1 + " - " + n.s2 + " - " + n.s3
}

type T struct {
	A int
	B string
}

// variable to investigate:
var secret interface{} = NotknownType{"Ada", "Go", "Oberon"}

func main() {
	value := reflect.ValueOf(secret) // <main.NotknownType Value>
	typ := reflect.TypeOf(secret)    // main.NotknownType
	// alternative:
	// typ := value.Type()  // main.NotknownType
	fmt.Println(typ)
	knd := value.Kind() // struct
	fmt.Println(knd)

	// iterate through the fields of the struct:
	for i := 0; i < value.NumField(); i++ {
		fmt.Printf("Field %d: %v\n", i, value.Field(i))
		// 结构中只有被导出字段（首字母大写）才是可设置的；
		// value.Field(i).SetString("C#")
	}

	// call the first method, which is String():
	results := value.Method(0).Call(nil)
	fmt.Println(results) // [Ada - Go - Oberon]

	t := T{23, "skidoo"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
    /*
    	0: A int = 23
		1: B string = skidoo
    */
	s.Field(0).SetInt(77)
	s.Field(1).SetString("Sunset Strip")
	fmt.Println("t is now", t)	// t is now {77 Sunset Strip}
}

```



  



### printf和反射

`fmt` 包中的 `Printf()`（以及其他格式化输出函数）都会使用反射来分析它的 `...` 参数。

`Printf()` 的函数声明为：

```go
func Printf(format string, args ... interface{}) (n int, err error)
```

- `Printf()` 中的 `...` 参数为空接口类型。`Printf()` 使用反射包来解析这个参数列表。所以，`Printf()` 能够知道它每个参数的类型。

  因此格式化字符串中只有 `%d` 而没有 `%u` 和 `%ld`，因为它知道这个参数是 unsigned 还是 long。这也是为什么 `Print()` 和 `Println()` 在没有格式字符串的情况下还能如此漂亮地输出。

```go
package main

import (
	"os"
	"strconv"
)

type Stringer interface {
	String() string
}

type Celsius float64

func (c Celsius) String() string {
	return strconv.FormatFloat(float64(c),'f', 1, 64) + " °C"
}

type Day int

var dayName = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

func (day Day) String() string {
	return dayName[day]
}

func print(args ...interface{}) {
	for i, arg := range args {
		if i > 0 {os.Stdout.WriteString(" ")}
		switch a := arg.(type) { // type switch
			case Stringer:	os.Stdout.WriteString(a.String())
			case int:		os.Stdout.WriteString(strconv.Itoa(a))
			case string:	os.Stdout.WriteString(a)
			// more types
			default:		os.Stdout.WriteString("???")
		}
	}
}

func main() {
	print(Day(1), "was", Celsius(18.36))  // Tuesday was 18.4 °C
}
```





### 反射的核心功能

Go 的反射机制提供了在运行时对类型进行操作的能力，允许你在不事先知道类型具体结构的情况下：

- **获取类型信息**：包括字段、方法、类型名称等。
- **获取和修改字段的值**：通过指针访问和操作内部字段。
- **动态调用方法**：根据方法名在运行时调用方法。



1. **获取类型信息**：

   - 你可以通过 `reflect.TypeOf` 获取对象的具体类型。`Type` 类型提供了对象的类型信息，如名称、种类（`Kind`）、字段、方法等。

   - 对于自定义类型来说，

     - `t := reflect.TypeOf(p)`

     - `t.NumField()` 获取字段数量。

       `t.Field(i)` 获取第 `i` 个字段的信息，包括字段名(`t.Field(i).Name`) 和类型(`t.Field(i).Type`) 和 值 （`field.Interface()`）。

     - `t.NumMethod()` 获取方法数量。

       `t.Method(i)` 获取第 `i` 个方法的信息，包括方法名(`t.Method(i).Name`)和方法类型(`t.Method(i).Type`)。

2. **获取和修改值**：

   - 通过 `reflect.ValueOf` 获取对象的值。

     `Value` 类型提供了获取和修改对象值的方法。

     - 要修改值，通常需要使用指向该值的指针，`t = reflect.Valueof(&aaa)`

     - 并调用 `t.Elem()` 方法来获取可修改的值。
     - 对于不同的类型 使用 各自的Set方法修改值

   - 对于 自定义类型，还是

     - `t = reflect.Valueof(&aaa).Elem()`
     - 接着通过 `Field(i).SetXXX()` 来进行赋值

3. **动态调用方法**：

   **获取方法**：

   - 使用 `reflect.ValueOf` 获取对象的 `reflect.Value`。
   - 使用 `reflect.Value.MethodByName` 或 `reflect.Type.Method` 获取指定方法的 `reflect.Value`。

   **调用方法**：

   - 使用 `reflect.Value.Call` 方法调用方法。
   - `Call` 方法接收一个 `[]reflect.Value` 类型的参数，用于传递方法调用所需的参数。
   - 如果方法有返回值，你可以从 `Call` 方法的结果中获取返回值

   ```go
   package main
   
   import (
   	"fmt"
   	"reflect"
   )
   
   // 定义一个自定义结构体类型
   type Person struct {
   	Name string
   	Age  int
   }
   
   // 定义一个方法
   func (p Person) Greet(greeting string) {
   	fmt.Printf("%s, my name is %s and I am %d years old.\n", greeting, p.Name, p.Age)
   }
   
   func main() {
   	// 创建一个 Person 实例
   	p := Person{Name: "Alice", Age: 30}
   
   	// 获取 Person 的 reflect.Value
   	v := reflect.ValueOf(p)
   
   	// 获取方法的 reflect.Value
   	method := v.MethodByName("Greet")
   	if method.IsValid() {
   		// 准备参数
   		params := []reflect.Value{reflect.ValueOf("Hello")}
   
   		// 调用方法
   		method.Call(params)
   	} else {
   		fmt.Println("Method not found")
   	}
   }
   
   ```

   

## 总结：Go 中的面向对象



Go 没有类，而是松耦合的类型、方法对接口的实现。

OO 语言最重要的三个方面分别是：**封装、继承和多态**

- 封装（数据隐藏）：和别的 OO 语言有 4 个或更多的访问层次相比，Go 把它简化为了 2 层

  **1）包范围内的：通过标识符首字母小写，*对象*只在它所在的包内可见**

  **2）可导出的：通过标识符首字母大写，*对象*对所在包以外也可见**

类型只拥有自己所在包中定义的方法。

- 继承：用组合实现：**内嵌一个（或多个）包含想要的行为（字段和方法）的类型**；该类型能够直接使用内层的类型的方法和属性，多重继承可以通过内嵌多个类型实现

- 多态：用接口实现：

  **某个类型的实例可以赋给它所实现的任意接口类型的变量。** 通过赋值后的接口变量 便能够动态的选择相对应的方法进行执行

  **某个传入参数为接口类型的函数 允许传入任意实现了该接口的类型** 

  - 在不涉及接口时，方法对于的接收类型对于 时 T 还是 *T不做进一步的要求，go会自动执行相应的操作以匹配类型
  - 但是涉及到接口，并且 传入参数为接口类型的函数中调用了相关类型的方法时，*T 需要 保证函数的接口处的传入类型为 *T 而T不做进一步的要求

  

  类型和接口是松耦合的，并且多重继承可以通过实现多个接口实现。



## 文件操作

- 使用 `bufio.NewScanner()` 函数创建一个 `bufio.Scanner` 对象，需要传入 `io.Reader` 接口，通常是 `os.Stdin`，表示标准输入输出流（键盘输入
- 使用 `scanner.Scan()` 方法可以逐行扫描文本，它会在成功扫描一行后返回 `true`，如果到达文件末尾或发生错误，则返回 `false`。
- 使用 `scanner.Text()` 方法可以获取当前行的文本内容作为字符串。
- `bufio.NewScanner` 能够读取完整的一行是因为它默认使用换行符 `\n` 作为分隔符，只有遇到换行才会停止当前行的输入，但你可以使用 `scanner.Split` 方法来自定义分隔符。
