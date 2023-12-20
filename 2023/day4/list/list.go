package list

import (
    "errors"
    "fmt"
)


type Node struct{
    info int
    next *Node
}

type List struct{
   head *Node
   max  *Node
   Length int
}




func (l *List) Insert(d int){
    list := &Node{info: d, next: nil}
    if l.head == nil {
        l.head = list
        l.max = list
        l.Length += 1
    } else {
        p := l.head
        for p.next != nil {
            p = p.next
        }
        p.next = list
        l.Length += 1
    }

    if l.max.info < d{
        l.max = list
    }
}

func (l *List) GetMax() int{
   return l.max.info 
}

func (l *List) CheckIn(d int) bool{
    p := l.head
    for p != nil {
        if p.info == d {
            return true
        }
        p = p.next
    }
    return false
}

func (l *List) GrabIdx(idx int) (int, error){
    if idx <= 0{
        return -1, errors.New("Index out of range")
    }
    p := l.head
   for i := 1; i<idx; i++ {
        p = p.next
   }
   return p.info, nil

}

func (l *List) PrintList(){

    p := l.head
    for p != nil {
        fmt.Printf(" %v", p.info)
        p = p.next
    }
    fmt.Printf("\n")
}
