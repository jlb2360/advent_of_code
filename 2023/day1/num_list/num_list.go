package num_list

import (
	"log"
	"strconv"
)


type NumString struct{
    Num string
    Index int
}

type Node struct{
    info NumString
    next *Node
}

type List struct{
   head *Node
   first *Node
   last *Node
}




func (l *List) Insert(d NumString){
    list := &Node{info: d, next: nil}
    if l.head == nil {
        l.head = list
        l.first = list
        l.last = list
    } else {
        p := l.head
        for p.next != nil {
            p = p.next
        }
        p.next = list
    }

    if l.first != nil {
        first := l.first
        if first.info.Index > d.Index{
            l.first = list
        }
    }

    if l.last != nil {
        last := l.last
        if last.info.Index < d.Index{
            l.last = list
        }
    }
}

func (l *List) GrabNum() int{
    f := l.first
    la := l.last

    n, err := strconv.Atoi(f.info.Num + la.info.Num)
    if err != nil{
        log.Fatal(err)
    }

    return n
}

