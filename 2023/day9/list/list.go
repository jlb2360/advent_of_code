package list

import "fmt"

type Node struct{
    info int
    next *Node
}

type List struct{
    head *Node
    tail *Node
    Length int
}

type ListNode struct{
    info List
    next *ListNode
    prev *ListNode
}

type LoL struct{
    head *ListNode
    tail *ListNode
    Length int
}


func (l *LoL) Insert(lis List){
    lisN := &ListNode{info: lis, next: nil, prev: nil}
    if l.head == nil{
        l.head = lisN
        l.tail = lisN
        l.Length = 1
    } else {
        p := l.head
        for p.next != nil{
            p = p.next
        }
        lisN.prev = p
        p.next = lisN
        l.tail = lisN
        l.Length = l.Length + 1
    }

}

func (l *List) Insert(d int){
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

func (l *List) Insert_front(d int){
    list := &Node{info: d, next: nil}
    if l.head == nil{
        l.head = list
        l.tail = list
        l.Length = 1
    } else {
        p := l.head
        list.next = p
        l.head = list
        l.Length = l.Length + 1
    }
}

func (l *List) CalcDif() List{
    dif_list := List{}
    p := l.head
    for p.next != nil {
        dif := p.next.info - p.info
        p = p.next
        dif_list.Insert(dif)
    }

    return dif_list
}

func (l *List) CheckZeros() bool{
    p := l.head
    for p != nil{
        if p.info != 0{
            return false
        }
        p = p.next
    }

    return true
}

func (l *List) GetLast() int{
    return l.tail.info
}

func (l *List) GetFirst() int{
    return l.head.info
}

func (l *LoL) Cald_dif() int{
    list := l.tail
    for list.prev != nil{
        new_n := list.info.GetLast() + list.prev.info.GetLast()
        list = list.prev
        list.info.Insert(new_n)
    }

    return list.info.GetLast()
}

func (l *LoL) Cald_back_dif() int{
    list := l.tail
    for list.prev != nil{
        new_n := list.prev.info.GetFirst() - list.info.GetFirst()
        list = list.prev
        list.info.Insert_front(new_n)
    }

    return list.info.GetFirst()
}

func (l *List) PrintList(){
    p := l.head
    for p != nil{
        fmt.Printf("%v ", p.info)
        p = p.next
    }

    fmt.Printf("\n")
}
