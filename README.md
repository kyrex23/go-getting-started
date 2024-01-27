<img align="left" width="60" height="60" alt="Go icon"
src="https://cdn.icon-icons.com/icons2/2699/PNG/512/golang_logo_icon_171073.png">

# Getting Started with Go

Welcome to the "Getting Started with Go" repository! 👋

This repository is a collection of beginner-friendly tutorials aimed at helping you learn and get started with the Go
programming language.

### Table of contents

- [Introduction](#introduction)
- [Contributing](#contributing)
- [License](#license)


## Introduction

Go, also known as Golang, is a statically typed, compiled language designed for simplicity and efficiency.

It has gained popularity for its clean syntax, strong standard library, and excellent performance.

Whether you're a complete beginner or an experienced programmer looking to learn Go, this repository is a great place
to start.

### Usage of workspaces to manage a multi-module repository

This repository contains multiple Go projects, so workspaces will be used to handle them more easily.

You can create a workspace by running the following command at the root of the repository:
```bash
go work init .
```

This will create a `go.work` file in which to include existing modules, as shown below:
```go
go 1.21.2

use (
	./hello-world
)
```

You can then execute a module with the following command:
```bash
go run ./hello-world
```


## Official Tutorials

> Note: this tutorials are obtained from the [official website](https://go.dev/doc/)

Here's a list of tutorials available in this repository to help you kickstart your journey with Go:

1. [Hello, World!](./hello-world/README.md) -- A simple tutorial to get your first Go program up and running.

2. [Import external packages](./import-external-packages/README.md) -- A tutorial to import public shared packages to use them into your code.

3. [Modules](./create-module/README.md) -- A tutorial introducing functions, error handling, arrays, maps, unit testing and compiling.


Feel free to explore these tutorials in any order that suits your learning style.

Each tutorial comes with step-by-step instructions and sample code to help you understand Go concepts better.


## Contributing

We welcome contributions to this repository! If you have a new feature, a bug fix or any improvements to existing code,
please follow our [contribution guidelines](./CONTRIBUTING.md)

We appreciate your help in making this resource more valuable for others! 🙂

## License

This repository is licensed under the [MIT License](./LICENSE).

Feel free to use, modify and distribute the content as long as you adhere to the terms of the license.
