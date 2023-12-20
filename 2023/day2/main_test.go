package main


import (
    "testing"
    "os"
    "bufio"
)

func TestMain1(t *testing.T){
    f, err := os.Open("test1.txt") 
    
    if err != nil {
        t.Fatalf("could not open file")
    }

  
    defer f.Close()
    scanner := bufio.NewScanner(f)

    var sum int = 0

    for scanner.Scan() {
        line := scanner.Text()
        occ, num := check_balls(&line)
        if occ == true {
            sum = sum + num
        }
    }

    if sum != 8 {
        t.Fatalf("Did not get 8 from ID sum we got %v instead", sum)
    }




}
