<div align="center">
    <h1>
        Capital Gains
    </h1>
</div>

This project is about a Command Line Application that calulates how much tax
you should pay based on the profit or losses of a stock market investment.

> [!NOTE]
> No external libraries were used in this project

# Installation

First of all, make sure you have Go or Docker installed to run the application,
you can also modify the input.txt file and put there the stocks operations.

Once that's installed there are two options to run the program:

1. Using Go
Enter the folder and run:

``` bash
go get; make build; make run
```

2. Using Docker
To build the container use the following command:

```bash
docker build -t capitalgains
```

> [!NOTE]
> If the build completes successfully it means that there was no
> errors in the tests

To run the project just run:

```bash
docker run capitalgains
```

> [!NOTE]
> In case you can to modify the input.txt to add your test cases
> you will need to re build the image

# Tests

To run tests use the command:

```bash
make test;
```

# Documentation

## Why Go?

In evaluating JavaScript, Python, and Go for my project, JavaScript was quickly
dismissed due to its browser-centric design, which feels less natural in console
applications. While Python excels in data management, its reliance on implicit
behaviors and syntax "magic" raised concerns about error-proneness.

Ultimately, Go emerged as the preferred choice. Its clean syntax, ease of use,
compilation, strong typing, and robust error handling make it a standout
language for this job. Go's simplicity and performance, coupled with its
reliability, align perfectly with the project's requirements.

## Folder Structure

As the application is small the best option best option is a flat structure.

``` bash
root/
    ├── cmd/
       ├── main.go # Entry point for the application
       ├── constants.go
       ├── stock.go
       ├── operation.go
       ├── utils.go
       ├── tests/
    ├── .gitignore
    ├── go.mod
    ├── go.sum
    └── README.md
```

It's a basic structure but self explanatory, the stock and operation files
are like classes, each file has a struct and each struct has its own methods.
