# gopher

### Welcome to gopher

# Hash Algorithm

It is a mathematical function that takes in a variable-sized inputs such as a file or message , and produces a fixed-size output , known as a hash or digest.

The Out put is unique representation of the input data and even a small chang eto the input data will produce a complately d/t hash value

Hash algorithms are designed to be one-way funcitons , means extremely difficult(if not impossible) to reconstruct the orginal inut data form the hash values.

- MD5
- SHA-1
- SHA-2
- SHA-3 are some examples of hashing algorithms

# Encription

## Synchronous/ Symetric encription

- Single key

## Asynchronous encryption

- public / private key

# Controll Flow

- sequence
- conditional
- loop
  Controll flow refers to the order in which the statements instructions and operations are executed in a computer program . It determines the sequence in which the instructions are executed

The Go Runtime system is a component of a GO programming language that manages and schedule the executrion of GO programs. It includes a number of features that makes it easy to write concurent and parrelal programs in go

The Init function : Each source file can define it's own niladic function to set up what ever state is required. It calls befoe main func

Program Executaion : A complate program is created by linking a single , unimpreted package called the main Package with all the packages it imports. , The main packa must have package name main and declared a funciton mian that takes no arguments and returns no value.

Program execution begins by initializaing the main package and then involing the main . when that function invocatyion returns , the program exits . it does not waait for other (non mnain) goroutines to complete.

# THe Stack & The Heap

In Go , the stack and heap are two regins of memory used for storign variables and data during program execution

## THe Stack

Is a regin of memory used for storing variables that are local to a function or a goroutine. When a function is called or a goroutine is created a new stack frame s created on the stack to store the function arguments , local variables and other data.

The Stack is a LIFO(last-in-first-out) data structure, which means that most recentlly added data is first to be removed.

Stack is faster than heap b/c it is managed nautomatically by the go runtime system, and memory allocation and deallocation is relatively fast.

## The Heap

Is a region of memory used for storing variabls that have a longer lifetime that those stored on the stack. When a variables allocated on the heap , it remains there until it is explicitly deallocated by the program or until the program exists.

The heap is more flexible data structure that the stack but it is generally slower because it requires more overhead for memory allocation and deallocation .

In go , the heap is managed automatically by garbage collector which frees up memory that is no longer being used by the program

# Select Statmenets - for channel adn concurency

## Concurency:

Refers to a code that is written in a concurent desing pattern , this means that the code has the potential ability to execute multiple tasks simultineously , where each task may make progress independently of the other .

    - In go  concurency is achieved using goroutines , a lightweight threads of execution that are managed by the go runtime.
    - A go program can create many goroutines that run concurently, each performing a d/t task.
    - The communication adn synchronization of these goroutines is typically done using channels , which provide a waay for goroutines to exchange data and coordinate their execution.

## Parallelism :

Refers the ability of a program to execute multiple task simultaneously bu utilizing multiple cpus or cores. It helps execution speed up by allowing multiple part of the program to run in pallerlal on d/t processors.

Go makes it easy to wto write concurent code using go rotines and channels.

# nuladic :

Of an operator or function in a program , having no argument.

# Go lang Data types

## Array :

- A numbered sequence of elements of the same type
- Does not change in size
- Used for ggo internals;

## Slice

- Built on top of an array
- holds values of the same type
- changes in size
- has a length and a capacity

## Map

- Key / value storage
- an unordered groups of elements of one type , called the element type , indexed by a set of unique key of another type called the key type

## Struct

- A data structure
- A composite type
- Allows us to collect vallues of d/t type together

# Arrays

### func make

the main built -in functino allocates and initializes an object of type slice , map or cahn(only).

### Multi-dimentional slices

is an array of arrays , where each elements can be accessed using multiples indices.

## Composite / Aggregagte data structure

### Map

A map is a key-value store . This means that you store some value you access that value by a key
Maps are unordered , if you print out all of the keys and values on a map they will printout inrandom order

A map is a perfect data structure when you need to look up data fast.

- Create a map
- access element in a map
- printout the entire map
- seeing the length of the map
  - len
  - access element

### Map - comma ok idiom

If you look up a non-existent key , the zero value will be returned as the value associated with that non-existent key.

### Structs

#### Is go an object oriented language

Yes and no . Although go has types and methods allows an object oriented styple of programming is not type hierarchy.

The concept of interface in go provides a d/t approach that we belives is easy to use and in some ways more general .

There is no ways to embade types in other types to provide somthing analogous -- but no identical -- to subclassing.

Moreover the methods in go are more general than in c++ or java , they can be defined for any sort of data , even built in types such as plain , unboxed integers . They are not restricted to structs (classes)

##

Unlike oop , in go a type can be created directly form types

```go
  xd := struct {
    first string
    last string
    age int
  }{
    first: "Jems",
    last: "Bond"
    age: 32
  }

```

## composition

Composite data typs aka aggregate data typles aka complex data types , refers to a way of structturing and building complex types by combining multiple simple tpes.

Composition is one of the funcamental principles of object oriented programming and allows you to create more flexible and reusable code .

One of the ways composition is achived is by embeing on estruct type with in another the fields and methods of the embeded "inner" struct become accessaible to hte outer struct.

The Inner type is said to be promoted to the outer type . In addition , mthods of the inner type are also promoted to the outer type , Whcih is similar to inheritance in traditional oop langes but go does not have builtin inheritance mechanism.

```go
  type Engine struct {
    //Engine fields
  }

//Engine method
func (e *Engine) Start() {
  fmt.Println("Engine Starteed")
}

type Car struct {
  Engine // Embeding the engine struct
  //Car specific fields
}

func main(){
  car:=Car{}
  car.Start() //Calling the Start Method form the embeded Engine struct
}




```

By using Composition , You can build complex types by combinign simpler ones , promoting code reuse and modular desing . It provides a way to create relation ship b/n d/t types with out the need to traditional inheritance.

- It is all about types
  - Go has oop aspects , but it's all about type. We create types in go ; user-defined types . We can the have values of that type . We can assign values of a user-defined type to variables
  - Go is object oriented
  1. Encapsulation :
     - a. State("Fields")
     - b. behevior("methods")
     - c. exported & unexported; viewable & not viewable
  2. Reusabblity
  - a. Inhertitance("Embedded types")
  3. Polymorphosm
  - a. interfaces
  4. Overriding
  - a. "promotion"

In Go You dont create classes , you create a type
You dont instantiate , you create values of a type

# Functions

- Structured code / procedural programming is one type of modular code and its purpose is
  - Abastract code
  - Code reusablity
  - More understandable

## syntax of function

```go

func (reciever) identifier(paramerters)(returns) {code}

```

- Parametersa and argumets
  - we define out fucn with paramets (if any)
  - we call our func and pass in arguments (in any)
- Everything in go is pass by value

## Variadic Parameter

A Variadic Parameter is a func which takes an unlimited number of argumerts . When you do this , this is known as a veriadic parmeter . When use the lexical element operator "...T" to signify a variadic Parmameter (There "T" represents some types)

The veriadic parameter mut be in the end

Iff is variadic with final parameter p of type ...T , then within f the type of p is equivalent to type []T. IF f is invokked with no actual arguments for p , the value passed to p is nil . Otherwise the value passed is new slice of type []T with a new underlying array whose successive elements are the acutal arguments, which all must be assignable to T. THe length and capacity of the slice is therefore the number of arguments bound to p and may differ for each call site.

```go
  func Greeting(prefix string , who ...string)
  Greeting("nobody")
  Greeting("Hello:" , "Joe","Anna")
```

### Unfuriing a slice

```go
xi:= []int{1,2,3,4,5,6}
//veriadic function
s:= sum(xi...)
```

## Defer Statment

A "defer" statment invokes a function whose execution is deferred to the moment the surrounding funciton returns , either becasue

- The surrounding funciton executed a return statement
- Reached the end of its function body
- Or because the corresponding gorouting is panicking

Go's defer statement schedules a funciton call (the deferred function) to be run immediately before the funciton executing the surrouning function . It is effective way to deal with situations such as resouces that must be released regardless of which path the function takes to retur.

## Methods

A method is nothing more than a FUNC attached to a TYPE. When you attach a func to a typ it is a mehtod of that type. You attach a func to a type with a reciever

## Interfaces and Polimorphism

An Interface in go defines a set of method signatures.

Polimorphism is the ability of a VALUE of a certain type to also be of another type

---> in Go values can be of more than one type

##### Polimorphism refers to the ability of an object to be of an additional typ e and take on d/t beheviours. In the context of programming it allows values of d/t types to be treated as instance of a common type.

## type Stringer

Stringer is implemented by any value that has a String method , which defines the 'native' format for that value. The string method is used to print values passed as an operand to any fomrat that accepts a string or to an unformatted printer such as print

## Log

### Wrapper function

In go is a function that provides an additional layer of abstraction or functionality around an existing function or method .

It act as an intermediary b/n the caller and teh wrapped function , allowing you to modify inoputs , outputs , or beheviour with out directly modifying the orginal function

rapper functions can be used for various purposes such as

1. Logging : A wrapper function can add logging statements before and after invoking the wrapped functions. This helps in capturing information about the function call , inputs , parameters , return values and any errors that may occur.

# Strings vs byte

In go a string adn a []byte are two d/t types , but they are closely related and can often be converted b/n each other.

A String in go represents a sequence of characters. ITis an immutable type , which means you can not modify individual charactes within a string.

String values are always interprated as UTF-8 encoded Unicode text.

On the other hand a []bye is a slice of bytes where each elements a single byte , it is a mutable type , so you can modify individual bytes within a byte slice . it can be used to represent binary data or text in various encodings

To convert

```go
//string to slice of byte
str:= "Hello"
bytes := []byte(str)

// slice of byte to string
bytes := []byte{72,101,108,111}
str := string(bytes)
```

## byte buffer

A byte buffer is a regin of memory used to temporarily store a sequence of bytes. it provides a data structure for efficient manipulation of byte data. A bye bffer allows you to read and write bytes to and form the bfferm, making it usefull for tasks like data serialization , network communication , file I/O and efficient string manipulation

A byte buffer is a data structure tha provides a convenient interface to manipulate sequences of bytes efficiently. it serves as a temporary storage for byte and enable operations such as reading , writing , appending and resizing byte sequence

## Annonimus Func

Annonimus function are a self-executing func

- no parameter
- 1 parameter

## func expression

function ingo are first class citizens

An expression is a combination of values , variables , operator and functions calls that are evaluated to produce a single value. It can be as simple as literal value or a variable or it can involve complex operations and functions invocations

Expressions are the building blocks of go programs and they are perfomme computations , manipulte values and make decisions based on conditions. They assigned to variables , passed as arguments to functions or used in controll flow statment like if statments and loops.

The Term first calss citizen refers to the status of certain entities such as values , types and functions , that are treated equally and have the same capabilities as other entities in the language. It means that these entities can be assigned to variables m passed as argumets to functions and returned as values from functions , just like any other data type in the language.

In go functions are considered first calss citizen . they can be assiggned to variables , passed as argumetn s to toher functions and retuned as as values form functions. This allows go to support higher order fumnction whre function can operate on toher functions. For example you can define that takes another functions as an argument and then call that function within the body of the functio.

## Closure

Is a form of Anonymous funciton that refer to variables specified outside of the function itself. it is equivalent to accessign global variables that exited before the function's declaration. It works similar to regural fucnion and perform all the tings perfom by any function.

## Recursion funciton

Recursion is the technique of solving a problem by breaking it down in to smaller subproblems of the smae type. In other words , a recursive function that calls itself during its executio.

When a recursive funciton is called it solves a smaller instance of the same problem and them combines the result of the smaller instance with the current instances to obtain the final reunst. The Function continures to call itself on smaller subproblems until it reaches a base case , which is a simple case that can be solved directly wilth out further recursion.

Recurstion can be a powerful technique for solving problmes that exibit a recursive structire , such as tree traversal , graph traversal , ana many mathematical calculations. However it is important to design recursive functions carefully , ensuring that they have well-defined base case and property handle the termination condition to avoid infinite recursion.

## Wrapper function

A wrapper function is a function that provides an additional layer of abstraction or functionality around an existing funciton or methor. It acts as intermediary b/n the caller and wrapped fucntion allowing you to modify ilnputs , outputs or behevious , without directly modifying the orginal funciton. A wrapper function wraps or modifies another function's beheviour.

# Testing in go

Go privides a built int tesitng freamwork that simplifies the process of testing go cod. Here is an overview of the file structe , nameing conventionss , and how testing works in go

- Test files : live iln the same packages as the code they test. They are named with the following conversion 'filename_test.go' where filename is the name of the file thea contains the code you want to test.

- Test functions : munst start with the word 'Test' Followed by a word that starts with a capital letter . The signature of a test function 'fucn TestXxx(\*testing.T)' whre Xxxx does not start with lower case later

  ```go
    package main

    fucn Add(a int , b int) int {
      return a + b
    }
  ```

## unit test

Unit test is a type of softwate testing where individual componnets or units of a softwaree are tested. The purpouse is to validate that each unit of the software performs as designed,
A unit test is the smallerst testable part of any software. It usually has one or a few inputs and usuallly a single output.

In go programming a unit test usually tests a single funcion , method or struct , the goal is to confirm that the beheviour is correct.

Unit tets in ggo are typically written using the built - in testing package , 'testing' .

A Unit test is a subset of test . Tests in software can take many forms such as unit tests , integration test , functionsl tests , system test ...

An Integration test woumld test b.n multiple components of a system to ensure they work together correctyl.

## Documents

Godocs includes comments - and produes documentation as html or plain text . The end result is the documentation tightyly coupled with the code it documents. for example godocs's web interface you can navigate form a function's documetna tot its impllementation with one click.

## interfaces & mock tesing

Interface in go are a powerfull tool for abstraction and can be especially usefull when you want to write unit tests for function that interact with a db . By creating an interface that describes the beheviour of a database interactions your code needs , you can swap out the real database for a mock one in your tests.

### Interface remainder

Interfaces in go are a pwerfull tool for abstraction and can be especially usefull when you want to write unit tests for functions that interact with a database.

By createin an interface that describes the beheviour of the database interactions your code needs . , you can swap out the real database for a mock one in your test

### Annonimus Func

Annonimus func , aka funcion literal or lamada funciton is a way to define a function without giving its name , instead you can directyly declare and define a function inline when ever it is needed .

Annonemous funcs are commonly used when you want to pass a function as an argument to another function or when you need to define a short lived funciton for a specific task.

### Callback functions

Callback function are very common in may programming languages including go. Callbacks are essentially fucntions that are passed as argumetns to other functions and are intended to be called (or "executed") later in your program

# Pointers

Pointers provide a way to pass references to data instead of copying the entire data , which can be especially beneficial for large data structure .

## \* (Asterisk)

- Is used to declare a pointer variable.
  ex: \*int gives the memory addres of the type
- Is used to dereference a pointer, which means accessign the value stored at the memory address pointed by the pointer variable.

## & (Ampersand)

- Is used to get the memory addrtess of a variable  
   ex: &num gives the memory address of the variable num.
- Is commonly used when working with pointers , passing variables memory address to functions , or initializing pointers with existing variables

The Pointers is a memory adress

In go , a pointer refers to a varibale that holds the memory adress. Pointers are a powerfull feature that allows you to directly manipulate memory and build complex data structure like linkedlist , trees and more .

### Dereferencing pointers

Is geting the value out of the memory adddress or simply getting value of that memeory address , that allows to directly manipulate the vlaue.

### Pass by Value

Will pass the value of the variable into the method , the orginal variable 'copy' the value into another memory location and pass the newly created one into the method , so any thing that happens to the varianle incide the method will not affect teh orginal variable value.

### Pass by reference

Will pass the memory location instead of the value , It passess the 'container' of the variables incide the memory will affect the orginal variable

In Go , Referecnce typ is pointer - a type of data the referes to , or points to , the locatoin in memory of a value. Go have serveral references type including pointers , slices , maps , channels , functions and interfaces

1. Pointers : A pointer holds the memory address of a value . It allows you to directly access and modify the memory location of a value
2. Slices : is a descriptor of an array segment. It includes a ponter to the array , the length of the segment and its's capacity (the maximum length of the segment)
3. Maps : is a powerfull data structure that associates values of one type (the key) with values of another type (the vlue). It's an unordered collectoin of key-value pairs
4. Channels : are used for communications b/n goroutines (go term for threads) . They allow you topass data from one goroutine to another
5. Functions : are first-class citezens , meaning they can be assigned to variables , passed a n arguments to other functions and returned as values form other functiosn .
6. Interfaces : is type represents a set of method signatures. It provides a way to specify the beheviour of an object. If somthing can dothis , then it can be used here.

In go all data is passed by value , which means that when ever you pass data to function go creates a copyof that data and assign the copy of a prameter variable. The function can do whetever it wants to copy with out affecting the orginal data.

A mutable value is a value that can be changed . In go slices , maps , and pointers are mutable data types . Even though they are passed by value , they still behave as if they ware passed by references becaluse the 'Value' that copied and passed is reference to the underlying data , not the actual data

## Pointer and value semantic

Semantic means relating in language or logic

#### Value Semantic

Value semantic refers to when the actual data of a variable is passed to a function or assigned to another variable. This means that the new varibale or function parameter gets a completely independent copy of the data.

## Pointer & Value Semantics heuristics

There are some general guidance you can follow when deciding whether to use pointer or values semantics in go

1. Use Value Semantics when Possible

   - Value semantics are simpler and usually safer , since they don't involve shared state or require you to think about memoryo management.
   - As a rule of thumb , if a function doesn't need to modify its input or the data you're working with is small (like built in type or small structs ) use value semantics

2. Use Pointer Semanitcs for llarger Data

   - Copying large structs or arrays can be inefficient
   - If the data you're working with is large , you might want to use pointer semantics to avoid the cost of copying the data. A rule of thumb: 64 bytes or larger , use pointers

3. Use Pointer Semantics for mutablity

   - If a function or mehtod needs to modify its reciever or an input paramter , you will need to use pointer semanitcs .
   - This is common use coase for methods that need to update the state of a struc.

4. Consitency :

   - It is important to be consitent if some fucnitons on type use pointer semantics and other use value semantics , this can leads to confusion. Typically , onces a type has a mehtod with pointer semantics , all methods on that type should have pointer semanitcs.

5. Pointer Semantics when interfacing with other code

   - If you're interfaciing with other code (like a library or a system call) you might need to use ponter semantics .

- "The Majority of bugs that we gets in software have to do with the mutation of memory " --> Bill Kenedy

# The Heap vs Stack

### Value Semantics

In GO , when a funciton recieves a vlaue (not pointer) , it gets its own copy of that value. This copy is typically placed on the stack , which is fast and doesn't invole any form of garbage collection . Once the funciotn retuns this memory can be instantly reclaimed,

### Ponter Semantics & the heap

When a function recieves a pointer or retuns a ponter to local variables it indicates to the compiler that this value could be shared across goroutine boundaries or could persist after the function returns. TO ensure that the data will remain available , the go compiler must allocate it on the heap , rather than on the stack . Heap allocation is more expensive and requires garbase collaction.

# Method Sets

Is the set of methods attached to a type. This concept is key to Go's interface mechanism and it is associated with both the value types and pointer type

- The Method set type T consists of all methods with reciever type T.
- The Methdo set of type *T consists of all method with reciever *T or T.

The Idea of the method set is integeral to how interfacecs are implemented andused in go .

An Interface isn go defines a method set and any type whose method set is superset of the interface's method set is considered to implement that interface.

A crucial thing to remember is that in GO, if you define a method with a pointer reciever , the method is only int he method set of pointer type . Tlhis is important int he context of interfaces b/c if an interface requires a method that's defined on the pointer (not the value) then you can use a pointer to that type to satisfy these interface, not a value of the type.

# Gemerics

Will help us achive the dry prenciples , by alowing us to define dynamic type.
With Generics , you can declare and use funciton or types that are written to work with any of a set of types provided by calling code

## Concret type vs interface type

### Concrete type

is a type that you can directly instanciate or create a value from. this means that the type can directly represent a set of values and that you can create an instace of this type with out any additional informatioon.

Example of concreate type : int , bool , float64 arrays , slices , maps , and structs month the others . These type shave specific predetermined storage layout characterstcs.

```go
type Employee struct {
  Name string
  Age int
}
emp := Employee{Name:"John" , Age:2}
```

### Interface type

Set of methods or types but does not define specific data layout ro instace.
Interface types are abstract - they represent beheviour or type but not a specific set of values.

if you declare an interface like 'io.Reader' you can not directly create an instacne of io.Reader , instead you create instances concreate type taht satisfy the ioReader interface

```go

type Reader interface {
  Read(p [byte])(n int , err error)
}
```

In the latter case io.Reader is an interface type and any type that implements the Read method with the correct signature is said to satisfy the io.Reader interface

- In go if the Variable name start wtih capital letter it means it is viewable outside of the package

# Application

## Marshal

Marshal traverse teh value v recursivel. if an encountered value implements the Marshaller interface and is not a nil pointer , Marshal calls its MarshalJson Method to produce JSON . If no marshalJson Method is present but the value implements encoding .

## Unmarshal

Unmaprsahl parses the JSON-encoded data and stores the result in the vlaue pointed to by v.

Unmparshal uses inverse of the encodeing that Marshal uses , allocating maps , slices and pointers as necessary .

To Umparshal json int o a pointer , unmarshal first handles the case of the JSON being the JSON literal null. In that case , unmpardhsla sets the pointer to nil . Other wise unmarshal unmarshals the json into the value pointed at by the pointer if the pointer is nil , unmarshal allocates a new vlaue for it to point to.

## Writer interface

Writer is the interface that wraps the basic writer method
Writer writes len(p) bytes from p to the underlying data stream. it returns teh number of bytes written from p (0<=n<=len(p)>>) and any error encountered that caused the writer to stop early. Writer must reaturn a non-nil error if it reatns n<len(p) Writer must bot modify the slice data even temporaril.

```go
  type Writer interface {
    Writer(p [byte] (n int , err error))
  }
```

### Decode

Decode reads the next JSON-encoe value from its input and stores it inthe value pointed to by v

## Bcrypt

Is a package in go that implements Provos and Mazières’s bcrypt adaptive hashing algorithm to calculate hash

# Concurency

Concurency is the composition of independently executing processes.

## Concurency vs Parallerism

Paralallism is simulatenious execution of (possibly related ) computations,
Concurency is about dealing with lots of things at onece , parallerism is about lots of things at once

## method sets

A type may have a method set associated with it. THe method set of an interface type is its interface. The method set of any other type T consists of all methods declared with reciever type T. The method set of the corresponding pointer type *T is the set of all methods declared with reciever *T or T (that is , it also contains the method set of T). Futrhter rules apply to structs containing anonymous fields as described in the section on struct types any other types has an emty methd set. in a method set , each method must have a unique non-blank mehtod name.

The method sets of a type determines the interfaces that the type implements and the methods that can be called using a reciever of that type.

## Go statements

Go statements starts with the execution of a fnction call as an independent concuerent thread of control or gorounte with the same address space.

THe function value and parameters are evaluated as usual in the calling goroutine , but unlike with a regular call , program execution does not wait for the invoked functions to complate. Instead , the funcion begins executing independently in a new goroutine. When the function terminates , its goroitine also terminates. If the function has any return values , they are discarded withn the funcion complates.
