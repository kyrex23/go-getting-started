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

Oops, I forgot it... 😅

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


## Contributing

We welcome contributions to this repository! If you have a new feature, a bug fix or any improvements to existing code,
please follow our [contribution guidelines](./CONTRIBUTING.md)

We appreciate your help in making this resource more valuable for others! 🙂

## License

This repository is licensed under the [MIT License](./LICENSE).

Feel free to use, modify and distribute the content as long as you adhere to the terms of the license.
