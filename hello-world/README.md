# Write some code
Get started with Hello, World.

1. Open a command prompt and cd to your home directory.
	On Linux or Mac:
	```bash
	cd
	```
	On Windows:
	```bash
	cd %HOMEPATH%
	```

2. Create a hello directory for your first Go source code.

	For example, use the following commands:
	```bash
	mkdir hello
	cd hello
	```

3. Enable dependency tracking for your code.

	When your code imports packages contained in other modules, you manage those dependencies through your code's own module. That module is defined by a go.mod file that tracks the modules that provide those packages. That go.mod file stays with your code, including in your source code repository.

	To enable dependency tracking for your code by creating a go.mod file, run the go mod init command, giving it the name of the module your code will be in. The name is the module's module path.

	In actual development, the module path will typically be the repository location where your source code will be kept. For example, the module path might be github.com/mymodule. If you plan to publish your module for others to use, the module path must be a location from which Go tools can download your module. For more about naming a module with a module path, see Managing dependencies.

	For the purposes of this tutorial, just use example/hello.
	```bash
	$ go mod init example/hello
	go: creating new go.mod: module example/hello
	```

4. In your text editor, create a file hello.go in which to write your code.

5. Paste the following code into your hello.go file and save the file.
	```go
	package main

	import "fmt"

	func main() {
		fmt.Println("Hello, World!")
	}
	```

	This is your Go code. In this code, you:

	- Declare a main package (a package is a way to group functions, and it's made up of all the files in the same directory).
	- Import the popular fmt package, which contains functions for formatting text, including printing to the console. This package is one of the standard library packages you got when you installed Go.
	- Implement a main function to print a message to the console. A main function executes by default when you run the main package.

6. Run your code to see the greeting.
	```bash
	$ go run .
	Hello, World!
	```

	The go run command is one of many go commands you'll use to get things done with Go. Use the following command to get a list of the others:
	```bash
	$ go help
	```
