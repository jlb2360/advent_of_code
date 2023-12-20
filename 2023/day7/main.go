package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
    l "day7/list"
)

func get_power(cards [5]int) int{
    
    pairs := [5]int{1,1,1,1,1}
    for i, c1 := range cards{
        for j, c2 := range cards{
            if i != j {
                if c1 == c2{
                    pairs[i] = pairs[i] + 1
                }
            }
        }
    }

    p := 0
    for i:=0; i<5; i++{
        if p < pairs[i]{
            p = pairs[i]
        }
    }


    if p == 1{
        return 1
    } else if p == 2 {
        n2 := 0
        for i := 0; i<5; i++{
            if pairs[i] == 2{
                n2 = n2 + 1
            }
        }
        if n2 == 4{
            return 3
        }
        return 2
    } else if p == 3 {
        for i := 0; i<5; i++{
            if pairs[i] == 2{
                return 5
            }
        }
        return 4
    } else if p==4{
        return 6
    } else if p==5{
        return 7
    }

    


    return 0

}

func main(){
    f, err := os.Open("inputTest.txt")
    if err != nil{
        log.Fatal(err)
    }

    scanner := bufio.NewScanner(f)
    var h [5]int

    hands := l.List{}

    for scanner.Scan(){
        line := scanner.Text()

        ls := strings.Split(line, " ")

        for j, c := range ls[0]{
            n, err := strconv.Atoi(string(c))
            if err == nil{
                h[j] = n
            } else{
                switch string(c) {
                case "T":
                    h[j]=10
                case "J":
                    h[j]=11
                case "Q":
                    h[j]=12
                case "K":
                    h[j]=13
                case "A":
                    h[j]=14
                default:
                    fmt.Printf("Got a weird string %s",string(c))
                    
                }
            }
            
        }


        b, err := strconv.Atoi(ls[1])
        if err != nil{
            log.Fatal(err)
        }

        p := get_power(h)

        ha := l.Hand{Cards: h, Power: p, Bid: b}

        hands.Insert(ha)
    }

    hands.PrintList()
    sum := hands.SumRank()

    fmt.Printf("The sum is: %v", sum)
}
