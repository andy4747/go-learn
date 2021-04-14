package exercises


import (
    "fmt"
)

func Fib(){
    a := FibonacciIterative(30)
    fmt.Println(a)

    b := FibonacciRecursive(30)
    fmt.Println(b)
}

func FibonacciIterative(n int) int {
    a,b := 0,1
    for i:=0;i<n-1;i++{
        c := a+b
        a,b = b,c
    }
    return b
}

func FibonacciRecursive(n int) int {
    if n<=1 {
        return n
    }
    return FibonacciRecursive(n-1)+FibonacciRecursive(n-2)
}
