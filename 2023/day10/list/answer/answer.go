package answer

import "fmt"

type Answer struct {
        Distance int
        Pipe string
        Col int
        Row int
}


type Node struct {
    info Answer
    next *Node
}


type List struct {
    head *Node
    tail *Node
    Length int
    Maxim *Node
}


func (l *List) Insert_back(d Answer){
    list := &Node{info: d, next: nil}
    if l.head == nil{
        l.head = list
        l.tail = list
        l.Length = 1
        l.Maxim = list
    } else {
        p := l.tail
        p.next = list
        l.tail = list
        l.Length = l.Length + 1
        if list.info.Distance > l.Maxim.info.Distance{
            l.Maxim = list
        }
    }
}

func (l *List) Insert_front(d Answer){
    list := &Node{info: d, next: nil}
    if l.head == nil{
        l.head = list
        l.tail = list
        l.Length = 1
        l.Maxim = list
    } else {
        p:=l.head
        list.next = p
        l.head = list
        l.Length = l.Length + 1
        if list.info.Distance > l.Maxim.info.Distance{
            l.Maxim = list
        }
    }
}

func (l *List) GetMax() Answer{
    return l.Maxim.info
}


func (l *List) PrintList(){
    p := l.head
    fmt.Printf("    distance    |    pipe    |    column    |    row    |\n")
    for p != nil {
        fmt.Printf("    %v    |    %s    |    %v    |    %v    |\n", p.info.Distance, p.info.Pipe, p.info.Col, p.info.Row)
        p = p.next
    }
}

func (l *List) MinList(l1 *List, l2 *List) List{
    min_lis := List{}
    if l1.Length == l2.Length{
        p1 := l1.head
        p2 := l2.head
        for p1 != nil {
        if p1.info.Distance < p2.info.Distance{
            min_lis.Insert_back(p1.info)
        } else if p2.info.Distance < p1.info.Distance{
            min_lis.Insert_back(p2.info)
        } else {
            min_lis.Insert_back(p1.info)
        }
        p1 = p1.next
        p2 = p2.next
    }
    }
    
    return min_lis
    
}

