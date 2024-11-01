package main

import (
	"fmt"
	"time"
)

func main() {

    fmt.Println("Hello! It's calculator by LEDev =D")
    fmt.Println("Write simple math question (example: 2+2, 4-1.5, 4*3.5, 5.5/1.1) or q to quit")

    for {
        fmt.Print(" > ")

        var question string
        fmt.Scanln(&question)

        if question == "q" {
            break
        }

        time1 := time.Now()
        
        var num1 float64
        var num2 float64

        var action int8 = -1
        // 0+, 1-, 2*, 3/

        var temp int64 = -1
        var temp2 int64 = -1

        ok := true

        for i := 0; i < len(question); i++ {

            var simv int32 = int32(question[i])

            if simv >= 48 && simv <= 57 {
                // It's num
                var num int8 = int8(simv - 48)
                // fmt.Print("num: ")
                // fmt.Println(num)
                if action == -1 {
                    if temp == -1 {
                        num1*=10
                        num1+=float64(num)
                    } else {
                        num1+=float64(num)/float64(temp)
                        temp*=10
                    }
                } else {
                    if temp2 == -1 {
                        num2*=10
                        num2+=float64(num)
                    } else {
                        num2+=float64(num)/float64(temp2)
                        temp2*=10
                    }
                }

            } else if simv == 44 || simv == 46 {
                // It's , or .
                // fmt.Println("point")
                if action == -1 {
                    if temp == -1 {
                        temp = 10
                    } else {
                        ok = false
                        break
                    }
                } else {
                    if temp2 == -1 {
                        temp2 = 10
                    } else {
                        ok = false
                        break
                    }
                }

            } else if simv >= 42 && simv <= 47 {
                // It's action

                if action != -1 {
                    ok = false
                    break
                }

                switch simv {
                    case 43:
                        action = 0
                    case 45:
                        action = 1
                    case 42:
                        action = 2
                    case 47:
                        action = 3
                    default:
                        ok = false
                }
                // if ok {
                //     fmt.Print("action: ")
                //     fmt.Println(action)
                // }

            } else {
                ok = false
                break
            }
        }

        if action == -1 || !ok {
            fmt.Println("Incorrect input!")
            continue
        }

        if action == 3 && num2 == 0 {
            fmt.Println("You can't divide by zero!")
            continue
        }

        // fmt.Println(num1)
        // fmt.Println(num2)

        var result float64
        switch action {
            case 0:
                result = num1+num2
            case 1:
                result = num1-num2
            case 2:
                result = num1*num2
            case 3:
                result = num1/num2
        }

        fmt.Print("= ")
        fmt.Println(result)

        time2 := time.Now()
        elapsed := float64(time2.Sub(time1).Nanoseconds()) / 1000000.0
        fmt.Print("Elapsed time: ")
        fmt.Print(elapsed)
        fmt.Println(" ms")
    }

    fmt.Print("Goodbye!")
    time.Sleep(time.Second)
    fmt.Print(".")
    time.Sleep(time.Second)
    fmt.Print(".")
    time.Sleep(time.Second)
    fmt.Println()
}
