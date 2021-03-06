# GoLang Improtant Differences

[Golang tour](https://go.dev/tour/)
[Introduction Docs](https://go.dev/doc/effective_go#introduction)
## IF Statements
Since if and switch accept an initialization statement, it's common to see one used to set up a local variable.
```go
if err := file.Chmod(0664); err != nil {
  log.Print(err)
  return err
}
```
In the Go libraries, you'll find that when an if statement doesn't flow into the next statement—that is, the body ends in break, continue, goto, or return—the unnecessary else is omitted.
```go
f, err := os.Open(name)
if err != nil {
  return err
}
codeUsing(f)
```
This is an example of a common situation where code must guard against a sequence of error conditions. The code reads well if the successful flow of control runs down the page, eliminating error cases as they arise. Since error cases tend to end in return statements, the resulting code needs no else statements.
```go
f, err := os.Open(name)
if err != nil {
  return err
}
d, err := f.Stat()
if err != nil {
  f.Close()
  return err
}
codeUsing(f, d)
```

## Redeclaration and reassignment

An aside: The last example in the previous section demonstrates a detail of how the := short declaration form works. The declaration that calls os.Open reads,
```go
f, err := os.Open(name)
```
This statement declares two variables, f and err. A few lines later, the call to f.Stat reads,
```go
d, err := f.Stat()
```
which looks as if it declares d and err. Notice, though, that err appears in both statements. This duplication is legal: err is declared by the first statement, but only re-assigned in the second. This means that the call to f.Stat uses the existing err variable declared above, and just gives it a new value.

In a := declaration a variable v may appear even if it has already been declared, provided:

this declaration is in the same scope as the existing declaration of v (if v is already declared in an outer scope, the declaration will create a new variable §),
the corresponding value in the initialization is assignable to v, and
there is at least one other variable that is created by the declaration.

## For
The Go for loop is similar to—but not the same as—C's. It unifies for and while and there is no do-while. There are three forms, only one of which has semicolons.
```go
// Like a C for
for init; condition; post { }

// Like a C while
for condition { }

// Like a C for(;;)
for { }
```
Short declarations make it easy to declare the index variable right in the loop.
```go
sum := 0
for i := 0; i < 10; i++ {
sum += i
}
```
If you're looping over an array, slice, string, or map, or reading from a channel, a range clause can manage the loop.
```go
for key, value := range oldMap {
newMap[key] = value
}
```
If you only need the first item in the range (the key or index), drop the second:
```go
for key := range m {
if key.expired() {
delete(m, key)
}
}
```
If you only need the second item in the range (the value), use the blank identifier, an underscore, to discard the first:
```go
sum := 0
for _, value := range array {
sum += value
}
```
The blank identifier has many uses, as described in a later section.

For strings, the range does more work for you, breaking out individual Unicode code points by parsing the UTF-8. Erroneous encodings consume one byte and produce the replacement rune U+FFFD. (The name (with associated builtin type) rune is Go terminology for a single Unicode code point. See the language specification for details.) The loop
```go
for pos, char := range "日本\x80語" { // \x80 is an illegal UTF-8 encoding
fmt.Printf("character %#U starts at byte position %d\n", char, pos)
}
```
prints

character U+65E5 '日' starts at byte position 0
character U+672C '本' starts at byte position 3
character U+FFFD '�' starts at byte position 6
character U+8A9E '語' starts at byte position 7
Finally, Go has no comma operator and ++ and -- are statements not expressions. Thus if you want to run multiple variables in a for you should use parallel assignment (although that precludes ++ and --).
```go
// Reverse a
for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
a[i], a[j] = a[j], a[i]
}
```

## Type switch
A switch can also be used to discover the dynamic type of interface variable. Such a type switch uses the syntax of a type assertion with the keyword type inside the parentheses. If the switch declares a variable in the expression, the variable will have the corresponding type in each clause. It's also idiomatic to reuse the name in such cases, in effect declaring a new variable with the same name but a different type in each case.
```go
var t interface{}
t = functionOfSomeType()
switch t := t.(type) {
default:
fmt.Printf("unexpected type %T\n", t)     // %T prints whatever type t has
case bool:
fmt.Printf("boolean %t\n", t)             // t has type bool
case int:
fmt.Printf("integer %d\n", t)             // t has type int
case *bool:
fmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
case *int:
fmt.Printf("pointer to integer %d\n", *t) // t has type *int
}
```

## Defer
The deferred call's arguments are evaluated immediately, but the function 
call is not executed until the surrounding function returns
Deferred function calls are pushed onto a stack. When a function returns, 
its deferred calls are executed in last-in-first-out order.

## Pointer
Unlike C, Go has no pointer arithmetic .
To access the field X of a struct when we have the struct pointer p we could write (*p).X. 
However, that notation is cumbersome, so the language 
permits us instead to write just p.X, without the explicit dereference.

## Slices
An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array. In practice, slices are much more common than arrays.

The type []T is a slice with elements of type T.

A slice is formed by specifying two indices, a low and high bound, separated by a colon:
```go
a[low : high]
```
This selects a half-open range which includes the first element, but excludes the last one.

The following expression creates a slice which includes elements 1 through 3 of a:
```go
a[1:4]
```

A slice does not store any data, it just describes a section of an underlying array.

Changing the elements of a slice modifies the corresponding elements of its underlying array.

Other slices that share the same underlying array will see those changes.

these slice expressions are equivalent:
```go
a[0:10]
a[:10]
a[0:]
a[:]
```
The length of a slice is the number of elements it contains.

The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.

Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.

The make function allocates a zeroed array and returns a slice that refers to that array:
```go
a := make([]int, 5)  // len(a)=5
```
To specify a capacity, pass a third argument to make:
```go
b := make([]int, 0, 5) // len(b)=0, cap(b)=5

b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:]      // len(b)=4, cap(b)=4
```

The resulting value of append is a slice containing all the elements of the original slice plus the provided values.

If the backing array of s is too small to fit all the given values a bigger array will be allocated. The returned slice will point to the newly allocated array.

## Function closures
Go functions may be closures. A closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.

For example, the adder function returns a closure. Each closure is bound to its own sum variable.

## Methods

Go does not have classes. However, you can define methods on types.

A method is a function with a special receiver argument.

You can only declare a method with a receiver whose type is defined 
in the same package as the method. You cannot declare a method with a 
receiver whose type is defined in another package (which includes the built-in types such as int).

## pointer receivers in methods

This means the receiver type has the literal syntax *T for some type T. (Also, T cannot itself be a pointer such as *int.)

For example, the Scale method here is defined on *Vertex.

Methods with pointer receivers can modify the value to which the receiver points (as Scale does here). Since methods often need to modify their receiver, pointer receivers are more common than value receivers.

Try removing the * from the declaration of the Scale function on line 16 and observe how the program's behavior changes.

With a value receiver, the Scale method operates on a copy of the original Vertex value. (This is the same behavior as for any other function argument.) The Scale method must have a pointer receiver to change the Vertex value declared in the main function.

## Methods and pointer indirection
Comparing the previous two programs, you might notice that functions with a pointer argument must take a pointer:

var v Vertex
ScaleFunc(v, 5)  // Compile error!
ScaleFunc(&v, 5) // OK
while methods with pointer receivers take either a value or a pointer as the receiver when they are called:

var v Vertex
v.Scale(5)  // OK
p := &v
p.Scale(10) // OK
For the statement v.Scale(5), even though v is a value and not a pointer, the method with the pointer receiver is called automatically. That is, as a convenience, Go interprets the statement v.Scale(5) as (&v).Scale(5) since the Scale method has a pointer receiver.


## 	Interfaces
An interface type is defined as a set of method signatures.

A value of interface type can hold any value that implements those methods.

## Interfaces are implemented implicitly
A type implements an interface by implementing its methods. There is no explicit declaration of intent, no "implements" keyword.

Implicit interfaces decouple the definition of an interface from its implementation, which could then appear in any package without prearrangement.

```go
type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}
```

Under the hood, interface values can be thought of as a tuple of a value and a concrete type:

(value, type)
An interface value holds a value of a specific underlying concrete type.

## Interface values with nil underlying values
If the concrete value inside the interface itself is nil, the method will be called with a nil receiver.

In some languages this would trigger a null pointer exception, but in Go it is common to write methods that gracefully handle being called with a nil receiver (as with the method M in this example.)

## The empty interface 🎉
The interface type that specifies zero methods is known as the empty interface:

interface{}
An empty interface may hold values of any type. (Every type implements at least zero methods.)

Empty interfaces are used by code that handles values of unknown type. For example, fmt.Print takes any number of arguments of type interface{}.


## Stringers
One of the most ubiquitous interfaces is Stringer defined by the fmt package.

type Stringer interface {
String() string
}
A Stringer is a type that can describe itself as a string. The fmt package (and many others) look for this interface to print values.

## Errors
Go programs express error state with error values.

The error type is a built-in interface similar to fmt.Stringer:

```go
type error interface {
Error() string
}

```
(As with fmt.Stringer, the fmt package looks for the error interface when printing values.)

Functions often return an error value, and calling code should handle errors by testing whether the error equals nil.
```go

i, err := strconv.Atoi("42")
if err != nil {
fmt.Printf("couldn't convert number: %v\n", err)
return
}
fmt.Println("Converted integer:", i)
```
A nil error denotes success; a non-nil error denotes failure.

## Readers
The io package specifies the io.Reader interface, which represents the read end of a stream of data.

The Go standard library contains many implementations of this interface, including files, network connections, compressors, ciphers, and others.

The io.Reader interface has a Read method:

func (T) Read(b []byte) (n int, err error)
Read populates the given byte slice with data and returns the number of bytes populated and an error value. It returns an io.EOF error when the stream ends.

The example code creates a strings.Reader and consumes its output 8 bytes at a time.

## Images
Package image defines the Image interface:

```go
package image

type Image interface {
ColorModel() color.Model
Bounds() Rectangle
At(x, y int) color.Color
}

```
Note: the Rectangle return value of the Bounds method is actually an image.Rectangle, as the declaration is inside package image.

(See the documentation for all the details.)

The color.Color and color.Model types are also interfaces, but we'll ignore that by using the predefined implementations color.
RGBA and color.RGBAModel. These interfaces and types are specified by the image/color package

## Channels
Channels are a typed conduit through which you can send and receive values with the channel operator, <-.
```go

ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and
// assign value to v.
```
(The data flows in the direction of the arrow.)

Like maps and slices, channels must be created before use:

```go
ch := make(chan int)

```
By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables.

The example code sums the numbers in a slice, distributing the work between two goroutines. Once both goroutines have completed their computation, it calculates the final result.

## Goroutines
A goroutine is a lightweight thread managed by the Go runtime.

```go
go f(x, y, z)
//starts a new goroutine running

f(x, y, z)
```
The evaluation of f, x, y, and z happens in the current goroutine and the execution of f happens in the new goroutine.

Goroutines run in the same address space, so access to shared memory must be synchronized. The sync package provides useful primitives, although you won't need them much in Go as there are other primitives. (See the next slide.)

## Buffered Channels
![](Learning/MoreTypes/buffered%20Channel%20creating%20deadlock.png)

## Range and Close
A sender can close a channel to indicate that no more values will be sent. Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression: after

v, ok := <-ch
ok is false if there are no more values to receive and the channel is closed.

The loop for i := range c receives values from the channel repeatedly until it is closed.

Note: Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.

Another note: Channels aren't like files; you don't usually need to close them. Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a range loop.

## Select
The select statement lets a goroutine wait on multiple communication operations.

A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.

## Default Selection
The default case in a select is run if no other case is ready.

Use a default case to try a send or receive without blocking:
```go

select {
case i := <-c:
// use i
default:
// receiving from c would block
}

```

## Exercise: Equivalent Binary Trees
There can be many different binary trees with the same sequence of values stored in it. For example, here are two binary trees storing the sequence 1, 1, 2, 3, 5, 8, 13.

![](https://go.dev/tour/static/img/tree.png)
A function to check whether two binary trees store the same sequence is quite complex in most languages. We'll use Go's concurrency and channels to write a simple solution.

This example uses the tree package, which defines the type:
```go

type Tree struct {
Left  *Tree
Value int
Right *Tree
}
Continue description on next page.


```

## sync.Mutex
We've seen how channels are great for communication among goroutines.

But what if we don't need communication? What if we just want to make sure only one goroutine can access a variable at a time to avoid conflicts?

This concept is called mutual exclusion, and the conventional name for the data structure that provides it is mutex.

Go's standard library provides mutual exclusion with sync.Mutex and its two methods:

*Lock*
*Unlock*
We can define a block of code to be executed in mutual exclusion by surrounding it with a call to Lock and Unlock as shown on the Inc method.

We can also use defer to ensure the mutex will be unlocked as in the Value method.


# Where to Go from here...
You can get started by installing Go.
[GO documentations](https://go.dev/doc/)
[how to write go code](https://go.dev/doc/code)
[Package reference](https://go.dev/pkg/)
[language spec](https://go.dev/ref/spec)

Once you have Go installed, the Go Documentation is a great place to continue. It contains references, tutorials, videos, and more.

To learn how to organize and work with Go code, read How to Write Go Code.

If you need help with the standard library, see the package reference. For help with the language itself, you might be surprised to find the Language Spec is quite readable.

To further explore Go's concurrency model, watch [Go Concurrency Patterns (slides)](https://www.youtube.com/watch?v=f6kdp27TYZs) and [Advanced Go Concurrency Patterns (slides)](https://www.youtube.com/watch?v=QDDwwePbDtw) and read the [Share Memory by Communicating codewalk](https://go.dev/doc/codewalk/sharemem/).

To get started writing web applications, watch [A simple programming environment (slides)](https://vimeo.com/53221558) and read the [Writing Web Applications tutorial](https://go.dev/doc/articles/wiki/)

The [First Class Functions in Go codewalk](https://go.dev/doc/codewalk/functions/) gives an interesting perspective on [Go's function types](https://go.dev/blog/)

The Go Blog has a large archive of informative Go articles.

Visit the Go home page for more.