package main

import (
    "testing"
    "os"
    "log"
    "bufio"
)


func TestMain1(t *testing.T){
    
    f, err := os.Open("inputTest.txt")
    if err != nil {
       log.Fatal(err) 
    }

    defer f.Close()
    scanner := bufio.NewScanner(f)
    sum := 0

    for scanner.Scan(){
        line := scanner.Text()
        num, _ := check_winner(&line)
        sum = sum + num
    }


    if sum != 13{
        t.Fatalf("Expected 13 but got %v", sum)
    }



}

func TestMain2(t *testing.T){
    f, err := os.Open("inputTest.txt")
    if err != nil {
       log.Fatal(err) 
    }

    defer f.Close()
    scanner := bufio.NewScanner(f)
    sum := 0

    var future_copies [26]int
    for i:= 0; i<len(future_copies); i++{
        future_copies[i] = 1
    }
    

    k := 1
    for scanner.Scan(){
        line := scanner.Text()
        _, match:= check_winner(&line)
        sum = sum + future_copies[0]
        copies := future_copies[0]
        for i:= 0; i<len(future_copies)-1; i++{
            future_copies[i] = future_copies[i+1]
        }

        future_copies[len(future_copies)-1] = 1
        



        for j:=0; j<copies; j++{
            for i:=0; i<match; i++{
                future_copies[i] = future_copies[i] + 1 
            }
        }
        k+=1
    }


    if sum != 30{
        t.Fatalf("Expected 30 but got %v", sum)
    }
    
}
