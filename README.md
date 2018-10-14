# Go

# CLI : 
 Go build --> Compiles a bunch of go source code files
 Go run --> Compiles and execute one or two  files
 Go fmt --> Formats all the code in each file in the current directory
 Go install --> Compiles and installs a package
 Go get --> Downloads the raw source code of someone else's package
 Go test --> Runs any tests associated with the current project

#Package:
Package == Project == Workspace
A package is a collection of a common source code files

Types of packages: 
  Executable: Generates a file that we can run.
  Reusable: Code used as 'helpers'. Good place to put reusable logic

  Example: 
  Executable:  Defines a package that can be compiled and then *executed*. Must have a func called 'main'. 
  Package main --> go build --> main.exe
  Reusable: Defines a package that can be used as a dependency (helper code)
  Package X    --> go build --> nothing!

  # Standerd libraries for Go
  golang.org/pkg
  math  io  crypto  encoding  fmt debug

