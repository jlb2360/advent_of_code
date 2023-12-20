package list

import "fmt"


type Node struct{
    info int
    next *Node
}

type List struct{
    head *Node
}

func (l *List) Insert(d int){
    list := &Node{info: d, next: nil}
    if l.head == nil {
        l.head = list
    } else {
        p := l.head
        for p.next != nil {
            p = p.next
        }
        p.next = list
    }
}


func (l *List) GetIdx(idx int) int{
    p := l.head
    for i:=0; i<idx; i++{
        p = p.next
    }

    return p.info
}

func (l *List) PrintList(){
    p := l.head
    for p != nil{
        fmt.Printf("%v\n",p.info)
        p = p.next
    }
}
