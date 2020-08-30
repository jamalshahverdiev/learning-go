package calculator

import (
    "fmt"
    "os"
)

func MultipleInts() (int, int) {
        fmt.Print(`
Welcome to the calculator!
Please 'Enter' numbers to use into actions of calculator.`)
        var f, s int
        fmt.Print("\nEnter first number: ")
        fmt.Scan(&f)
        positiveIntegerValue(f)
        fmt.Print("\nEnter second number: ")
        fmt.Scan(&s)
        positiveIntegerValue(s)
        return f, s
}

func positiveIntegerValue(piv int) {
        if piv < 0 {
                fmt.Printf(" %d is a negative integer.", piv)
                os.Exit(1)
        } else if piv == 0 {
                fmt.Printf("%d is a positive zero integer value.", piv)
                os.Exit(1)
        }
    
        fmt.Printf("%d  is a positive and on-zere integer value.", piv)
}
