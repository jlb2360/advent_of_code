package main

import (
	"bufio"
	l "day11/utilities/mat"
	"log"
	"os"
    "testing"
)



func TestMain1(t *testing.T){
    f, err := os.Open("inputTest.txt")
    if err != nil {
        log.Fatal(err)
    }

    scanner := bufio.NewScanner(f)
    data := l.Mat{}

    for scanner.Scan(){
        line := scanner.Text()
        lis := l.List{}
        var wy int64 = 2
        for _, c := range line{
            if string(c) == "#"{
                wy = 1
            }
        }
        for _, c := range line{
            lis.Insert(string(c), wy, 1)
        }
        data.Insert(&lis)
    }

    data.CheckStarCol(2)
    data.PrintMat()
    data.PrintWeights()


    init := l.Corridinent{Row: 5, Col: 1}
    data.SetInit(5,1)
    
    data.AStar(init)
    num := data.GetDis(9, 4)
    

    if num != 9{
        t.Fatalf("Expected 9 got %v\n",num)
    }
    

}


func TestMain2(t *testing.T){
    f, err := os.Open("inputTest.txt")
    if err != nil {
        log.Fatal(err)
    }

    scanner := bufio.NewScanner(f)
    data := l.Mat{}

    for scanner.Scan(){
        line := scanner.Text()
        lis := l.List{}
        var wy int64 = 2
        for _, c := range line{
            if string(c) == "#"{
                wy = 1
            }
        }
        for _, c := range line{
            lis.Insert(string(c), wy, 1)
        }
        data.Insert(&lis)
    }

    data.CheckStarCol(2)


    init := l.Corridinent{Row: 0, Col: 3}
    data.SetInit(0,3)
    
    data.AStar(init)
    num := data.GetDis(8, 7)
    

    if num != 15{
        t.Fatalf("Expected 15 got %v\n",num)
    }
    

}

func TestMain3(t *testing.T){
    f, err := os.Open("inputTest.txt")
    if err != nil {
        log.Fatal(err)
    }

    scanner := bufio.NewScanner(f)
    data := l.Mat{}

    for scanner.Scan(){
        line := scanner.Text()
        lis := l.List{}
        var wy int64 = 2
        for _, c := range line{
            if string(c) == "#"{
                wy = 1
            }
        }
        for _, c := range line{
            lis.Insert(string(c), wy, 1)
        }
        data.Insert(&lis)
    }

    data.CheckStarCol(2)


    init := l.Corridinent{Row: 2, Col: 0}
    data.SetInit(2,0)
    
    data.AStar(init)
    num := data.GetDis(6, 9)
    

    if num != 17{
        t.Fatalf("Expected 17 got %v\n",num)
    }
    

}

func TestMain4(t *testing.T){
    f, err := os.Open("inputTest.txt")
    if err != nil {
        log.Fatal(err)
    }

    scanner := bufio.NewScanner(f)
    data := l.Mat{}

    for scanner.Scan(){
        line := scanner.Text()
        lis := l.List{}
        var wy int64 = 2
        for _, c := range line{
            if string(c) == "#"{
                wy = 1
            }
        }
        for _, c := range line{
            lis.Insert(string(c), wy, 1)
        }
        data.Insert(&lis)
    }

    data.CheckStarCol(2)


    init := l.Corridinent{Row: 9, Col: 0}
    data.SetInit(9,0)
    
    data.AStar(init)
    num := data.GetDis(9, 4)
    

    if num != 5{
        t.Fatalf("Expected 5 got %v\n",num)
    }
    

}



func TestMain5(t *testing.T){
    f, err := os.Open("inputTest.txt")
    if err != nil {
        log.Fatal(err)
    }

    scanner := bufio.NewScanner(f)
    data := l.Mat{}

    for scanner.Scan(){
        line := scanner.Text()
        lis := l.List{}
        var wy int64 = 2
        for _, c := range line{
            if string(c) == "#"{
                wy = 1
            }
        }
        for _, c := range line{
            lis.Insert(string(c), wy, 1)
        }
        data.Insert(&lis)
    }

    data.CheckStarCol(2)


    gal_rows, gal_cols := data.GetGals()
    sum := 0

    for i, r:= range gal_rows{
        data.Reset()
        init := l.Corridinent{Row: r, Col: gal_cols[i]}
        data.SetInit(r, gal_cols[i])
        data.AStar(init)
        for j:=i+1; j<len(gal_cols); j++{
            num := data.GetDis(gal_rows[j], gal_cols[j])
            sum += int(num)
        }
    }


    if sum != 374{
        t.Fatalf("Expected 374 got %v\n",sum)
    }
    

}

func TestMain6(t *testing.T){
    f, err := os.Open("inputTest.txt")
    if err != nil {
        log.Fatal(err)
    }

    scanner := bufio.NewScanner(f)
    data := l.Mat{}

    for scanner.Scan(){
        line := scanner.Text()
        lis := l.List{}
        var wy int64 = 10
        for _, c := range line{
            if string(c) == "#"{
                wy = 1
            }
        }
        for _, c := range line{
            lis.Insert(string(c), wy, 1)
        }
        data.Insert(&lis)
    }

    data.CheckStarCol(10)


    gal_rows, gal_cols := data.GetGals()
    sum := 0

    for i, r:= range gal_rows{
        data.Reset()
        init := l.Corridinent{Row: r, Col: gal_cols[i]}
        data.SetInit(r, gal_cols[i])
        data.AStar(init)
        for j:=i+1; j<len(gal_cols); j++{
            num := data.GetDis(gal_rows[j], gal_cols[j])
            sum += int(num)
        }
    }


    if sum != 1030{
        t.Fatalf("Expected 1030 got %v\n",sum)
    }
    

}

func TestMain7(t *testing.T){
    f, err := os.Open("inputTest.txt")
    if err != nil {
        log.Fatal(err)
    }

    scanner := bufio.NewScanner(f)
    data := l.Mat{}

    for scanner.Scan(){
        line := scanner.Text()
        lis := l.List{}
        var wy int64 = 100
        for _, c := range line{
            if string(c) == "#"{
                wy = 1
            }
        }
        for _, c := range line{
            lis.Insert(string(c), wy, 1)
        }
        data.Insert(&lis)
    }

    data.CheckStarCol(100)


    gal_rows, gal_cols := data.GetGals()
    sum := 0

    for i, r:= range gal_rows{
        data.Reset()
        init := l.Corridinent{Row: r, Col: gal_cols[i]}
        data.SetInit(r, gal_cols[i])
        data.AStar(init)
        for j:=i+1; j<len(gal_cols); j++{
            num := data.GetDis(gal_rows[j], gal_cols[j])
            sum += int(num)
        }
    }


    if sum !=8410 {
        t.Fatalf("Expected 8410 got %v\n",sum)
    }
    

}
