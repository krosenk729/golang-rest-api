
## Golang 🐱‍🚀
#### Snippets

###### Variables

A var declaration can include initializers, one per variable. If an initializer is present, the type can be omitted; the variable will take the type of the initializer.

Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type. Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.

```
var i, j int = 1, 2

// is the same as...

var i int = 1
var j int = 2

// or...

i, j := 1, 2

```


Variables declared without an explicit initial value are given their zero value.

The zero value is:
    0 for numeric types,
    false for the boolean type, and
    "" (the empty string) for strings.

Arrays
The type [n]T is an array of n values of type T.

The expression `var a [10]int` declares a variable a as an array of ten integers.

###### Defer
A defer statement defers the execution of a function until the surrounding function returns.

The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

```
package main

import "fmt"

func one() {
	defer fmt.Println("one")

	fmt.Println("hello one")
}

func two() {
	defer fmt.Println("two")
	one()

	fmt.Println("hello tow")
}

func main() {
	defer fmt.Println("defer main")
	two()

	fmt.Println("hello main")
}

```
###### Pointers
Go has pointers. A pointer holds the memory address of a value.

The type *T is a pointer to a T value. Its zero value is nil.


The & operator generates a pointer to its operand.

The * operator denotes the pointer's underlying value. This is known as "dereferencing" or "indirecting".
```
var p *int

i := 42
p = &i

fmt.Println(*p) // read i through the pointer p
*p = 21         // set i through the pointer p
```

```
package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
```

