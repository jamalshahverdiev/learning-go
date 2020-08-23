#### This is the simple calculator. Just copy `calculator` folder under `$GOPATH\src\` directory. I have prepared `separated_funcs.go`, `addme.go`, `divide.go`, `multiply.go` and `subtract.go` files like as library under same `calculator` package.

```bash
$ echo $GOPATH
C:\Users\Administrator\go
```

#### To test run `starter.go` file:
```bash
$ go run starter.go

Welcome to the calculator!
Please 'Enter' numbers to use into actions of calculator.
Enter first number: -3
 -3 is a negative integer.exit status 1
```

```bash
$ go run starter.go

Welcome to the calculator!
Please 'Enter' numbers to use into actions of calculator.
Enter first number: 7
7  is a positive integer.
Enter second number: 5
5  is a positive integer.
1) add
2) divide
3) multiply
4) subtract
Input number of the action to use entered numbers and press 'Enter' button: 1
Selected action: 'Add'
The result of 'Addme' method: 12
```
