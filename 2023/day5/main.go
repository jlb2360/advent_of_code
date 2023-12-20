package main

import (
	"bufio"
	cls "day5/List/conv"
	ls "day5/List/seeds"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)


func insertEl(line *string, cd *cls.ConvDict, key string){

        sl := strings.Split(*line, " ")
        cs, err := strconv.Atoi(sl[0])
        if err != nil {
            log.Fatal(err)
        }
        rs, err := strconv.Atoi(sl[1])
        if err != nil {
            log.Fatal(err)
        }
        length, err := strconv.Atoi(sl[2])
        if err != nil {
            log.Fatal(err)
        }

        cd.InsertElem(key, rs, cs, length)


}


func main(){

    var n1 int

    f, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    scanner := bufio.NewScanner(f)
    seeds := ls.List{}

    var conversion_lines []string

    for scanner.Scan(){
        line := scanner.Text()
        sl := strings.Split(line, " ")
        if sl[0] == "seeds:"{
            for i:=1; i<len(sl); i++{
                num, err := strconv.Atoi(sl[i])
                le, err1 := strconv.Atoi(sl[i+1])
                if err == nil && err1 == nil{
                    seeds.Insert(num, le)
                }
                i = i + 1
            }

        }else{
            conversion_lines = append(conversion_lines, line)
        }
    }

    cd := cls.ConvDict{}
    cd.InsertKey("s2s")
    cd.InsertKey("s2f")
    cd.InsertKey("f2w")
    cd.InsertKey("w2l")
    cd.InsertKey("l2t")
    cd.InsertKey("t2h")
    cd.InsertKey("h2l")
    

    for i:=0; i<len(conversion_lines); i++{
        if conversion_lines[i] == "seed-to-soil map:"{
            i = i + 1
            for conversion_lines[i] != ""{
                insertEl(&conversion_lines[i], &cd, "s2s")

                i = i + 1
            } 
        } else if conversion_lines[i] == "soil-to-fertilizer map:"{
            i = i + 1
            for conversion_lines[i] != ""{
                insertEl(&conversion_lines[i], &cd, "s2f")

                i = i + 1
                
            }
        }else if conversion_lines[i] == "fertilizer-to-water map:"{
            i = i + 1
            for conversion_lines[i] != ""{
                insertEl(&conversion_lines[i], &cd, "f2w")

                i = i + 1
                
            }
        }else if conversion_lines[i] == "water-to-light map:"{
            i = i + 1
            for conversion_lines[i] != ""{
                insertEl(&conversion_lines[i], &cd, "w2l")

                i = i + 1
                
            }
        }else if conversion_lines[i] == "light-to-temperature map:"{
            i = i + 1
            for conversion_lines[i] != ""{
                insertEl(&conversion_lines[i], &cd, "l2t")

                i = i + 1
                
            }
        }else if conversion_lines[i] == "temperature-to-humidity map:"{
            i = i + 1
            for conversion_lines[i] != ""{
                insertEl(&conversion_lines[i], &cd, "t2h")

                i = i + 1
                
            }
        }else if conversion_lines[i] == "humidity-to-location map:"{
            i = i + 1
            for conversion_lines[i] != ""{
                insertEl(&conversion_lines[i], &cd, "h2l")

                i = i + 1

                if i >= len(conversion_lines){
                    break
                }

                
            }
        }


    }

    m := 999999999999999
    for i:=0; i<seeds.Length; i++{
        var prog float64 = float64(i)/float64(seeds.Length)
        fmt.Printf("Progress is %v\n",prog)
        n, le := seeds.GetIdx(i)
        fmt.Printf("Start is %v and length is %v", n, le)
        for j:=n; j<n+le; j++ {
            n1 = j
            for _, k := range cd.Keys(){
                n1 = cd.ConvNum(k, n1)
            }
            if m > n1 {
                m = n1
            }
        }
    }

    
    fmt.Print(m)
}
