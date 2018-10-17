# Go

## CLI : 
 * Go build --> Compiles a bunch of go source code files
 * Go run --> Compiles and execute one or two  files
 * Go fmt --> Formats all the code in each file in the current directory
 * Go install --> Compiles and installs a package
 * Go get --> Downloads the raw source code of someone else's package
 * Go test --> Runs any tests associated with the current project

## Package:
Package == Project == Workspace
A package is a collection of a common source code files

Types of packages: 
 * Executable: Generates a file that we can run.
 * Reusable: Code used as 'helpers'. Good place to put reusable logic

  Example:
  
  Executable:  Defines a package that can be compiled and then *executed*. Must have a func called 'main'. 
  * Package main --> go build --> main.exe

  Reusable: Defines a package that can be used as a dependency (helper code)
  * Package X    --> go build --> nothing!

## Standard libraries for Go
  * golang.org/pkg

  math  io  crypto  encoding  fmt debug

  * The Go Playground at https://play.golang.org/ to quickly run a snippet of code.
  * First, dynamically-typed languages perform type checking at runtime, while statically typed languages perform type checking at compile time.If a script written in a statically-typed language contains errors, it will fail to compile until the errors have been fixed.
  Second, statically-typed languages require you to declare the data types of your variables before you use them, while dynamically-typed languages do not. Consider the two following:
  * Static languages: Go , C++, Java
  * Dynamic languages: JS, Python, Ruby 

Array: Fixed length list of things(more primitive)
Slice: An array that can grow or shrink(more fancy). Every element in a slice must be of same type

func WriteFile(filename string, data []byte, perm os.FileMode) error
