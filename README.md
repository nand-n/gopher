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
