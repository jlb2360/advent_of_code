package tree

import (
	"errors"
	"fmt"
)

type Decision struct{
    Spot string
    Left string
    Right string
}

type Node struct{
    info Decision
    next *Node
}

type List struct{
    head *Node
    tail *Node
    Length int
}

func (l *List) Insert(d Decision){
    list := &Node{info: d, next: nil}
    if l.head == nil{
        l.head = list
        l.tail = list
        l.Length = 1
    } else {
        p := l.tail
        p.next = list
        l.tail = list
        l.Length = l.Length + 1
    }
}

func (l *List) GetIDX(idx int) Decision{
    p := l.head
    for i:=0; i<idx; i++{
        p = p.next
    }
    return p.info
}

func (l *List) PrintList(){
    p := l.head
    for p != nil{
        fmt.Printf("%s (%s, %s)\n", p.info.Spot, p.info.Left, p.info.Right)
        p = p.next
    }
}


func (l *List) FindSpot(d string) (Decision, error){
    p := l.head
    for p != nil {
        if p.info.Spot == d{
            return p.info, nil
        }
        p = p.next
    }

    return Decision{Spot: "", Left: "", Right: ""}, errors.New("No Decision")
}

func (l *List) Exchange(idx int, d Decision){
    list := &Node{info: d, next: nil}
    p:=l.head
    if idx == 0{
        list.next = p.next
        l.head = list
    } else{
    for i:=0; i<idx-1; i++{
        p = p.next
    }
    temp := p.next
    list.next = temp.next
    p.next = list
}
}
