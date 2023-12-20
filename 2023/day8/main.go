package main

import (
	"bufio"
	t "day8/tree"
	"fmt"
	"log"
	"os"
	"strings"
)

func find_steps(places t.List, dir string) int{
    var spot t.Decision
    var err error
    spot, err = places.FindSpot("AAA")
    if err != nil {
        log.Fatal(err)
    }
    iter:= 0
    steps := 0
    for spot.Spot != "ZZZ"{
        iter = iter + 1
    for _, d := range dir{
        
        
        if string(d) == "L"{
            spot, err = places.FindSpot(spot.Left)
            if err != nil{
                log.Fatal(err)
            }
            steps = steps + 1
        } else if string(d) == "R"{
            spot, err = places.FindSpot(spot.Right)
            if err != nil{
                log.Fatal(err)
            }
            steps = steps + 1
        }

        if spot.Spot == "ZZZ"{
            return steps
        }
        }
    }

    return steps
}

func checks_spot(starts t.List) bool{

    spot := starts.GetIDX(6)
    if string(spot.Spot[2]) == "Z"{
        return false
    }

    for i:=0; i<starts.Length; i++{
        spot:=starts.GetIDX(i)
        if string(spot.Spot[2]) != "Z"{
            return true
        }
    }

    return false
}


func find_all(places t.List, starts t.List, dir string) int{
    steps := 0
    iter := 0
    looping := true
    for looping{
    for _, d := range dir{
        //for i:=0; i<starts.Length; i++{
            spot := starts.GetIDX(6)
            if string(d) == "L"{
                new_s, err := places.FindSpot(spot.Left)
                if err != nil {
                    log.Fatal(err)
                }
                starts.Exchange(6, new_s)
            } else if string(d) == "R"{
                new_s, err := places.FindSpot(spot.Right)
                if err != nil{
                    log.Fatal(err)
                }
                starts.Exchange(6, new_s)
            }
        //}

        steps = steps + 1

        looping = checks_spot(starts)
        if looping == false{
            return steps
        }

    }

    looping = checks_spot(starts)
    iter = iter + 1
    }


    return steps

}


func main(){
    f, err := os.Open("input.txt")
    if err != nil{
        log.Fatal(err)
    }

    scanner := bufio.NewScanner(f)

    places := t.List{}
    starts := t.List{}
    var directions string

    i := 0
    for scanner.Scan(){
        line := scanner.Text()
        if i == 0{
            directions = line
        } else if i > 1 {
            ls := strings.Split(line, "=")
            pos := ls[0][0:3]
            lss := strings.Split(ls[1], ",")
            le := lss[0][2:len(ls[0])+1]
            ri := lss[1][1:4]
            d := t.Decision{Spot: pos, Left: le, Right: ri}
            places.Insert(d)


            if string(pos[2]) == "A"{
                starts.Insert(d)
            }
        
        }

        i = i + 1
        
    }

    //sum := find_steps(places, directions)

    //fmt.Printf("Number of Steps %v\n", sum)

    sum := find_all(places, starts, directions)

    fmt.Printf("Number of Steps %v\n", sum)

}
