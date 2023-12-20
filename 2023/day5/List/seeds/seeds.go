package seeds

import "fmt"

type Seeds struct{
    start int
    length int
}

type Node struct {
    info Seeds
    next *Node
}


type List struct{
    head *Node
    Length int
}


func (l *List) Insert(begin int, le int){

    s := Seeds{start: begin, length: le}

    list := &Node{info: s, next: nil}
    if l.head == nil {
        l.head = list
        l.Length = 1
    } else {
        p := l.head
        for p.next != nil {
            p = p.next
        }

        p.next = list
        l.Length = l.Length + 1
    }
}

func (l *List) PrintList(){
    if l.head != nil {
        p := l.head
        for p != nil {
            fmt.Printf("%v\n",p.info.start)
            fmt.Printf("%v\n",p.info.length)
            p = p.next
        }
    }
}

func (l *List) GetIdx(idx int) (int, int){
    p := l.head
    for i:=0; i<idx; i++{
        if p.next != nil {
            p = p.next
        }
    }

    return p.info.start, p.info.length
    
}
