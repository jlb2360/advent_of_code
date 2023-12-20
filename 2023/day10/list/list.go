package list

import "fmt"

type Pos struct{
    Pipe string
    Direction string
    row int
    col int
}

type Node struct{
    info string
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
}

type LoL struct{
    head *ListNode
    tail *ListNode
    Length int
}


func (l* List) Insert_back(d string){
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

func (l *List) Insert_front(d string){
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

func (l *List) GetIdx(idx int) string{
    p := l.head
    for i:=0; i<idx; i++{
        p = p.next
    }

    return p.info
}


func (l *LoL) Insert(lis List){
    lisNode := &ListNode{info: lis, next: nil}
    if l.head == nil{
        l.head = lisNode
        l.tail = lisNode
        l.Length = l.Length + 1
    } else {
        p := l.tail
        p.next = lisNode
        l.tail = lisNode
        l.Length = l.Length + 1
    }
}


func (l *LoL) GetIdx(row int, column int) string{
    lis := l.head
    for j:=0; j<row; j++{
        lis = lis.next
    }
    num := lis.info.GetIdx(column)

    return num
}


func (l *LoL) FindS() (int, int){
    lis := l.head
    row := 0
    column := 0
    for lis != nil {
        p := lis.info.head
        column = 0
        for p != nil{
            if p.info == "S"{
                return row, column
            }
            column = column + 1
            p = p.next
        }

        lis = lis.next

        row = row + 1
    }

    return -1, -1
}


func (l *LoL) NextPipe(cPipe string, row int, column int, d string) (int, int, string){
    switch cPipe {
    case "L":{
        if d == "left"{
            return row-1, column, "up"
        }else if d == "down"{
            return row, column+1, "right"
        }
    }
    case "J":{
        if d == "down"{
            return row, column-1, "left"
        } else if d == "right"{
            return row-1, column, "up"
        }
    }
    case "7":{
        if d == "right"{
            return row+1, column, "down"
        } else if d == "up"{
            return row, column-1, "left"
        }
    }
    case "F":{
        if d == "up"{
            return row, column+1, "right"
        } else if d == "left"{
            return row+1, column, "down"
        }
    }
    case "|":{
        if d == "down"{
            return row+1, column, "down"
        } else if d == "up"{
            return row - 1, column, "up"
        }
    }
    case "-":{
        if d == "right"{
            return row, column + 1, "right"
        } else if d == "left"{
            return row, column - 1, "left"
        }
    }
}


    return -1, -1, "None"
}


func (l *LoL) PrintList(){
    list := l.head
    for list != nil {
        p := list.info.head
        for p != nil{
            fmt.Printf("%s",p.info)
            p = p.next
        }
        fmt.Printf("\n")
        list = list.next
    }
}

func (l *LoL) CheckIn(s_r int, s_c int){
    check_lis = List{}
    check_lis.Insert_front("JF")
    for i:=-1; i<=1; i++{
        for j :=-1; j<=1; j++{
            pipe_check := l.GetIdx(s_r+i, s_c+j)
            if i == -1 && j == -1{
                side_check_1 == l.GetIdx(s_r, s_c+j)
                
            }
        }
    }
}
