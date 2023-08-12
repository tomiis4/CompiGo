# CompiGo Documentation

## TODO
- [ ] add: types
    - [ ] structs/objects
    - [ ] maps
    - [ ] generics
- [ ] array things
    - [ ] insert
    - [ ] remove
    - [ ] length (lua like `#`)

## File

### run file
```sh
cg run file.cg
```


## Hello world
```cg
// create main function
main => {
    // print "Hello, world!"
    println("Hello, world!")
}
```


## Comments
```c
// one line comment

/*
    multi-line
    comment
*/
```


## Operators
```cg
// change
+  -  /  *  %

// compare
==  !=  <  >  <=  >=

// assign
+=  -=  *=  /=  ++  --
```


## Import package
```cg
// import one package
use "math" from "std.math"

// import multiple packages
use (
    "math" from "std.math"
    "time" from "std.time" as "t"
)

// rename package import
use "time" from "std.time" as "t"
```


## Variables

### Declaration
```cg
// set automatic type
variable := 610

// set type with value
var variable: u32 = 610

// set type without value
var variable: u32

// set mutable variable
mut variable := 610
mut var variable: u32 := 610

// create array
variable := <string>["Hello", "world"]
var variable: <string> = ["Hello", "world"]
```

### Types
#### Integers
| Size     | Signed | Unsigned |
|----------|--------|----------|
| 8-bit    | i8     | u8       |
| 16-bit   | i16    | u16      |
| 32-bit   | i32    | u32      |
| 64-bit   | i64    | u64      |

#### Other Primitive Types
| Type       | Description                                |
|------------|--------------------------------------------|
| string     | Represents a sequence of characters.       |
| bool       | Represents a boolean value: true or false. |
| byte       | Represents a byte-sized value.             |

#### Arrays
Represents a collection of elements of the same data type.

Type definition: 
```<type, length?>```


## Functions
```cg
// function without params/return
name => { /* ... */ }

// function with params
name => (param: string) { /* ... */ }
name => (param1: string, param2: bool) { /* ... */ }

// function with return
name => i64 { return 610 }

// function with params/return
name => (param: string) i64 { return 610 }

// method function
type.name => { /* ... */ }
```


## Logic Statements

### if/else
```cg
if true {
    // ...
} elseif true {
    // ...
} else {
    // ...
}
```

### ternary operator
```cg
// high level explanation: if true then return "y" else return "y"
true ? "x" : "y"
```


## Loops

### for-i
```cg
// create variable I, if is i < 5, add 1 to I
for mut i := 0; i < 5; i++ {
    // break/contiune
}
```

### for-in
```cg
for key, value in <u8>[7, 10, 6] {
    // ...
}
```


## Pointers
```cg
// create empty pointer and allocate memory
var ptr: *i8

// change value of pointer
*ptr = 10

// read value of pointer
*ptr

// get address from variable
&variable
```
