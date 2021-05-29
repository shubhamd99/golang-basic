### Go

![golang](https://i.imgur.com/9y787WC.png)

* Golang - Go is a statically typed, compiled programming language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson. Go is syntactically similar to C, but with memory safety, garbage collection, structural typing, and CSP-style concurrency. 

* Concurrency - In computer programming, concurrency is ability of a computer to deal with multiple things at once. A general purpose PC might have just one CPU core which does all the processing and computation.

* Docs - https://golang.org/doc/

* Dowload Go - https://golang.org/doc/install

* `go version`

* Microservice architecture, or simply microservices, is a distinctive method of developing software systems that tries to focus on building single-function modules with well-defined interfaces and operations. The trend has grown popular in recent years as Enterprises look to become more Agile and move towards a DevOps and continuous testing.

###### Advantages

1. Developed at Google
2. Statically Typed language
3. Community Support
4. Simplicity
5. Faster Compile Time
6. Garbage Collection
7. Builtin Concurrency
8. Standard Libraries - In the end you will get standalone binary file to run the code

###### Tips

1. We cannot declare unused variables, go will throw compliation error

###### Plugins used for VSCode

1. Go - https://marketplace.visualstudio.com/items?itemName=golang.Go
2. Docker - https://marketplace.visualstudio.com/items?itemName=ms-azuretools.vscode-docker
3. Live Share - https://marketplace.visualstudio.com/items?itemName=MS-vsliveshare.vsliveshare

###### Useful Commands

1. `open $GO_PATH` (Open GO Path Directory in Finder)
2. `go run main.go` (Run go file)
3. `go env` (To see all the Environment variables used by Go module)
4. `go run --work main.go` (Show the temp location where the binary is stored)
5. `go build` OR `go build -o shubham` (Take the program compiles it and generates the binary exe)
6. `go install` (The go install command behaves almost identically to go build , but instead of leaving the executable in the current directory, or a directory specified by the -o flag, it places the executable into the `$GOPATH/bin` directory)

7. `go get github.com/user/repo` To Install the package
8. `go mod init github.com/shubhamd99/golang-basic` (Create Go MOD file - Package Manager)
9.  `go mod tidy` (Remove `// indirect` comments in go mod file and remove the un-necessary packages as well)

10. `go get ./...` (Install all the necessary packages for the project after cloning fresh project)
11. `go clean -cache -modcache -i -r` (Clear all local downloaded cache)
12. `go build` (This also download all the packages for the project from go.mod file)

###### Packages Used
```
go get go.uber.org/zap
go get github.com/julienschmidt/httprouter
go get rsc.io/quote
```

###### Go Get Types
```
1. go get github.com/julienschmidt/httprouter (Its same as go get github.com/julienschmidt/httprouter@latest)
2. go get github.com/julienschmidt/httprouter@v1.2 (with version specified)
3. 
```

![releases](https://i.imgur.com/2frdq08.png)

###### Directory Structure
1. `open $GO_PATH`
2. src folder is the place where we will interact the most
3. Inside `src` folder we will have different hosts like github.com. Every package out there is on the some host. Go library uses this host to pull out the data.
4. Inside hosts folder we will have our github username folder
5. Inside github username folder we will have our projects
6. `bin` is the directory where golang stores all its binary
7. `pkg` is the directory where golang stores and pulls the dependencies and mods

![directory_go](https://i.imgur.com/lpW8reG.png)

###### GO Directory Image View
![d_go](https://i.imgur.com/mR7wavZ.png)

###### File Types
![file_types](https://i.imgur.com/rSJYoqz.png)

###### Packages
In go project we have to declare main package which can have multiple packages link to it together. We have to declare it on the top of the file.

```go
package main --> package a --> package b
```

###### [Data Types](https://www.tutorialspoint.com/go/go_data_types.htm)

In the Go programming language, data types refer to an extensive system used for declaring variables or functions of different types.

1. Boolean types - (a) true (b) false
2. Numeric types - a) integer types or b) floating point values
3. String types - Strings are immutable types that is once created, it is not possible to change the contents of a string.
4. Derived types
   
   (a) Pointer types
   (b) Array types
   (c) Structure types
   (d) Union types and
   (e) Function types
   f) Slice types
   g) Interface types
   h) Map types
   i) Channel Types

###### [Variables](https://www.tutorialspoint.com/go/go_variables.htm)

A variable is nothing but a name given to a storage area that the programs can manipulate. Each variable in Go has a specific type, which determines the size and layout of the variable's memory, the range of values that can be stored within that memory, and the set of operations that can be applied to the variable.

```
var  i, j, k int;
var  c, ch byte;
var  f, salary float32;

c := 200 // Shorthand variable assignment

const dev = "Hello" // Constants
```

###### Integers Types

![int types](https://i.imgur.com/3GzT2vT.png)


###### [Arrays](https://www.tutorialspoint.com/go/go_arrays.htm)

Go programming language provides a data structure called the array, which can store a fixed-size sequential collection of elements of the same type. 

```
var variable_name [SIZE] variable_type

var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
```

###### [Slices](https://www.tutorialspoint.com/go/go_slice.htm)

Go Slice is an abstraction over Go Array. Go Array allows you to define variables that can hold several data items of the same kind but it does not provide any inbuilt method to increase its size dynamically or get a sub-array of its own. Slices overcome this limitation.

```
var numbers []int /* a slice of unspecified size */
/* numbers == []int{0,0,0,0,0}*/
numbers = make([]int,5,5) /* a slice of length 5 and capacity 5*/
```

###### [Maps](https://www.tutorialspoint.com/go/go_maps.htm)

Go provides another important data type named map which maps unique keys to values. A key is an object that you use to retrieve a value at a later date. 

```
var map_variable map[key_data_type]value_data_type

var countryCapitalMap map[string]string
/* create a map*/
countryCapitalMap = make(map[string]string)

/* insert key-value pairs in the map*/
countryCapitalMap["France"] = "Paris"
countryCapitalMap["Italy"] = "Rome"
countryCapitalMap["Japan"] = "Tokyo"
countryCapitalMap["India"] = "New Delhi"

/* print map using keys*/
for country := range countryCapitalMap {
    fmt.Println("Capital of",country,"is",countryCapitalMap[country])
}

/* test if entry is present in the map or not*/
capital, ok := countryCapitalMap["United States"]

/* if ok is true, entry is present otherwise entry is absent*/
if(ok){
    fmt.Println("Capital of United States is", capital)  
} else {
    fmt.Println("Capital of United States is not present") 
}
```

###### [Struct](https://www.tutorialspoint.com/go/go_structures.htm)

Collection of data types. Go arrays allow you to define variables that can hold several data items of the same kind. Structure is another user-defined data type available in Go programming, which allows you to combine data items of different kinds.

```
type Books struct {
   title string
   author string
   subject string
   book_id int
}
```

###### Loops

```
arr := []int{1, 2, 3, 4, 5, 6, 7}

for key, value := range arr {
	fmt.Println(key, value)
}
```

###### If and Switch

```
if i >= 0 || j >= 0 {
    fmt.Println("i, j is greater than 0")
}
	

switch o := 2 + 3; o {
case 1, 3, 5, 7, 9:
   fmt.Println("Odd")
case 2, 4, 8:
   fmt.Println("Even")
default:
   fmt.Println("default")
}

```

###### Defer, Panic and Recover

A defer statement defers the execution of a function until the surrounding function returns. The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

Panic is a built-in function that stops the ordinary flow of control and begins panicking. 

Recover is a built-in function that regains control of a panicking goroutine. Recover is only useful inside deferred functions. During normal execution, a call to recover will return nil and have no other effect. 

###### [Pointers](https://www.tutorialspoint.com/go/go_pointers.htm)

As you know, every variable is a memory location and every memory location has its address defined which can be accessed using ampersand (&) operator, which denotes an address in memory. 

A pointer is a variable whose value is the address of another variable, i.e., direct address of the memory location. 

```
var var_name *var-type

var ip *int        /* pointer to an integer */
var fp *float32    /* pointer to a float */

var  ptr *int

fmt.Printf("The value of ptr is : %x\n", ptr  )
```

###### [Functions](https://www.tutorialspoint.com/go/go_functions.htm)

A function is a group of statements that together perform a task. 

```go
/* function returning the max between two numbers */

func max(num1, num2 int) int {
   /* local variable declaration */
   result int

   if (num1 > num2) {
      result = num1
   } else {
      result = num2
   }
   return result 
}
```
The init function runs before the main function, generally used for initialization
```go
func init() {
	// init
}
```

The main function is the entry point of the Go module
```go
func main() {
	// execute
}
```

###### [Interfaces](https://www.tutorialspoint.com/go/go_functions.htm)

Collection of behaviours. Go programming provides another data type called interfaces which represents a set of method signatures.

###### [Goroutines](https://gobyexample.com/goroutines)

A goroutine is a lightweight thread of execution.

Suppose we have a function call f(s). Here’s how we’d call that in the usual way, running it synchronously.

```f("direct")```

To invoke this function in a goroutine, use `go f(s)`. This new goroutine will execute concurrently with the calling one.

To Identify Race Conditions run ```go run -race .\Main.go```

###### [Channels](https://tour.golang.org/concurrency/2)

Channels are a typed conduit through which you can send and receive values with the channel operator, <-.

```
ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and assign value to v.
```

Like maps and slices, channels must be created before use:

```
ch := make(chan int)
```

By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables.

