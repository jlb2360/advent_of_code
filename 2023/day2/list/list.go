package list


type Node struct{
    info int
    next *Node
}

type List struct{
   head *Node
   max  *Node
}




func (l *List) Insert(d int){
    list := &Node{info: d, next: nil}
    if l.head == nil {
        l.head = list
        l.max = list
    } else {
        p := l.head
        for p.next != nil {
            p = p.next
        }
        p.next = list
    }

    if l.max.info < d{
        l.max = list
    }
}

func (l *List) GetMax() int{
   return l.max.info 
}
