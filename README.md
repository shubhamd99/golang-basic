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

###### Run

1. RUN - `go run main.go`
2. Check Connection - `curl -v -d 'Shubham' localhost:9090`
3. Create binary exe - `go build`


##### Workspace

![workspace](https://i.imgur.com/Rt1sBu3.png)

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
d =  42;
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