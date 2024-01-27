# Call code in an external package

When you need your code to do something that might have been implemented by someone else, you can look for a package that has functions you can use in your code.

1. Make your printed message a little more interesting with a function from an external module.

	1. Visit pkg.go.dev and search for a "quote" package.
	2. Locate and click the rsc.io/quote package in search results (if you see rsc.io/quote/v3, ignore it for now).
	3. In the Documentation section, under Index, note the list of functions you can call from your code. You'll use the Go function.
	4. At the top of this page, note that package quote is included in the rsc.io/quote module.
	5. You can use the pkg.go.dev site to find published modules whose packages have functions you can use in your own code.
	6. Packages are published in modules -- like rsc.io/quote -- where others can use them. Modules are improved with new versions over time, and you can upgrade your code to use the improved versions.

2. In your Go code, import the rsc.io/quote package and add a call to its Go function.

	After adding the highlighted lines, your code should include the following:

	```go
	package main

	import "fmt"

	import "rsc.io/quote"

	func main() {
		fmt.Println(quote.Go())
	}
	```

3. Add new module requirements and sums.

	Go will add the quote module as a requirement, as well as a go.sum file for use in authenticating the module.

	For more, see [Authenticating modules](https://go.dev/ref/mod#authenticating) in the Go Modules Reference.

	```bash
	$ go mod tidy
	go: finding module for package rsc.io/quote
	go: found rsc.io/quote in rsc.io/quote v1.5.2
	```

4. Run your code to see the message generated by the function you're calling.
	```bash
	$ go run .
	Don't communicate by sharing memory, share memory by communicating.
	```

	Notice that your code calls the Go function, printing a clever message about communication.

	When you ran `go mod tidy`, it located and downloaded the rsc.io/quote module that contains the package you imported. By default, it downloaded the latest version.