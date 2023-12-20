package conv

import "fmt"

type Range struct{
    range_start int
    range_len int
    conversion_start int
}

type Node struct{
    info Range
    next *Node
}

type HashNode struct{
    key string
    list *Node
    next *HashNode
    length int
}

type ConvDict struct{
    head *HashNode
}


func (cd *ConvDict) InsertKey(name string){
    hn := &HashNode{key: name, list: nil, next: nil, length: 0}
    if cd.head == nil {
        cd.head = hn 
    } else {
        p := cd.head
        for p.next != nil {
            p = p.next
        }
        p.next = hn
    }
}

func (cd *ConvDict) InsertElem(key string, start int, conv_start int, length int){
    r := Range{range_start: start, range_len: length, conversion_start: conv_start}
    l := &Node{info: r, next: nil}
    
    if cd.head.list == nil {
        lis := cd.head
        lis.list = l
        lis.length = 1
    } else {
        lis := cd.head
        for lis.key != key{
            lis = lis.next
        }

        if lis.list == nil {
            lis.list = l
            lis.length = 1
        } else {
            p := lis.list
            for p.next != nil {
                p = p.next
            }
            p.next = l
            lis.length = lis.length + 1
        }
    }
}

func (cd *ConvDict) PrintKey(key string){
    lis := cd.head
    for lis.key != key {
        lis = lis.next
    }


    fmt.Printf("|   idx |   range_start    |    range_len   |    conversion_start    |\n")
    p := lis.list
    i := 0
    for p != nil {
        fmt.Printf("------------------------------------------------------------------\n")
        fmt.Printf("|   %v  |   %v  |   %v  |   %v  |\n",i,p.info.range_start, p.info.range_len, p.info.conversion_start)
        p = p.next
        i = i + 1
    }
}

func (cd *ConvDict) Keys() []string{
    var k []string

    lis := cd.head
    for lis != nil{
        k = append(k, lis.key)
        lis = lis.next
    } 

    return k
}

func (cd *ConvDict) ConvNum(key string, num int) int{
    lis := cd.head
    for lis.key != key {
        lis = lis.next
    }

    r := lis.list
    min_rs := r.info.range_start
    max_rs := r.info.range_start + r.info.range_len
    for i:=0; i<lis.length; i++{
        if min_rs > r.info.range_start{
            min_rs = r.info.range_start
        }
        if max_rs <= r.info.range_start + r.info.range_len{
            max_rs = r.info.range_start + r.info.range_len
        }

       if num >= r.info.range_start && num < r.info.range_start + r.info.range_len{
           diff := num - r.info.range_start 
           n := r.info.conversion_start + diff
           return n
       }
       r = r.next
        
    }
    if min_rs != 0 && num < max_rs{
        return num
    } else if num >= max_rs{
        return num
    }

    return num
}
