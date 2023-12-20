package mat

import (
	"fmt"
	"math"
)



type Corridinent struct{
    Row int
    Col int
}

type Point struct{
    object string
    visited  bool
    distance int64
    weight_y int64
    weight_x int64
}

type Node struct{
    info Point
    next *Node
    prev *Node
}

type List struct{
    head *Node
    tail *Node
    Length int
}

type ListNode struct{
    info *List
    next *ListNode
}

type Mat struct{
    head *ListNode
    tail *ListNode
    Length int
}


func (l *List) Insert(d string, wy int64, wx int64){
    po := Point{object: d, visited: false, distance: 20000, weight_y: wy, weight_x: wx}
    lis := &Node{info: po, next: nil, prev: nil}
    if l.head == nil{
        l.head = lis
        l.tail = lis
        l.Length = 1
    } else {
        p := l.tail
        lis.prev = p
        p.next = lis
        l.tail = lis
        l.Length += 1
    }
}

func (l *List) PrintList(){
    p := l.head
    for p != nil {
        fmt.Printf("%s",p.info.object)
        p = p.next
    }
}

func (l *List) PrintWeights(){
    p := l.head
    for p != nil {
        fmt.Printf("(%v,%v)",p.info.weight_y, p.info.weight_x)
        p = p.next
    }
}


func (l *List) PrintDistance(){
    p := l.head
    for p != nil {
        fmt.Printf("%v ",p.info.distance)
        p = p.next
    }
}


func (l *List) PrintVisited(){
    p := l.head
    for p != nil {
        fmt.Printf("%v ",p.info.visited)
        p = p.next
    }
}

func (l *List) GetIdx(col int) *Node{
    p := l.head
    for i:=0; i<col;i++{
        p = p.next
    }

    return p
}


func (l *List) GetDis(col int) int64{
    p := l.head
    for i:=0; i<col;i++{
        p = p.next
    }

    return p.info.distance
}

func (l *List) CheckStar() bool{
    p := l.head
    for p != nil{
        if p.info.object == "#"{
            return false
        }
        p = p.next
    }
    return true
}

func (l *List) ChangeWeight(col int, wx int64){
    p := l.head
    for i:=0;i<col;i++{
        p=p.next
    }

    p.info.weight_x = wx
}

func (l *List) CheckVisited() bool{
    p := l.head
    for p != nil{
        if p.info.visited == false{
            return false
        }
        p = p.next
    }

    return true
}



//--------------------------------------------------
//       Matrix Implementation
//--------------------------------------------------

func (l *Mat) Insert(d *List){
    lis := &ListNode{info: d, next: nil}
    if l.head == nil{
        l.head = lis
        l.tail = lis
        l.Length = 1

    } else {
        p := l.tail
        p.next = lis
        l.tail = lis
        l.Length += 1

    }
}

func (l *Mat) PrintMat(){
    lis := l.head
    for lis != nil{
        lis.info.PrintList()
        fmt.Printf("\n")
        lis = lis.next
    }
}

func (l *Mat) PrintWeights(){
    lis := l.head
    for lis != nil{
        lis.info.PrintWeights()
        fmt.Printf("\n")
        lis = lis.next
    }
}
func (l *Mat) PrintDistance(){
    lis := l.head
    for lis != nil{
        lis.info.PrintDistance()
        fmt.Printf("\n")
        lis = lis.next
    }
}


func (l *Mat) PrintVisited(){
    lis := l.head
    for lis != nil{
        lis.info.PrintVisited()
        fmt.Printf("\n")
        lis = lis.next
    }
}

func (l *Mat) GetIdx(row int, col int) *Node{
    lis := l.head
    for i:=0;i<row;i++{
        lis = lis.next
    }
    st := lis.info.GetIdx(col)

    return st
}


func (l *Mat) GetDis(row int, col int) int64{
    lis := l.head
    for i:=0;i<row;i++{
        lis = lis.next
    }
    st := lis.info.GetDis(col)

    return st
}

func (l *Mat) Shape() (int, int){
    lis := l.head
    return l.Length, lis.info.Length
}


func (l *Mat) ChangeWeight(row int, col int, wx int64){
    lis := l.head
    for i:=0;i<row;i++{
        lis = lis.next
    }

    lis.info.ChangeWeight(col, wx)
}

func (l *Mat) CheckStarCol(wx int64){
    var cols []int
    for col:=0; col<l.head.info.Length; col++{
        stars := true
        for row:=0; row<l.Length; row++{
            if l.GetIdx(row, col).info.object == "#"{
                stars = false
            }
        }

        if stars == true{
            cols = append(cols, col)
        }
    }

    for _, col := range cols{
        for row:=0; row<l.Length; row++{
            l.ChangeWeight(row, col, wx)
        }
        }
}


func (m *Mat) GetGals() ([]int, []int){
    var gal_rows []int
    var gal_cols []int
    lis := m.head
    for row:=0; row<m.Length; row++{
        p := lis.info.head
        for col:=0; col < lis.info.Length; col++{
            if p.info.object == "#"{
                gal_rows = append(gal_rows, row)
                gal_cols = append(gal_cols, col)
            }
            p = p.next
        }

        lis = lis.next
    }

    return gal_rows, gal_cols
}

func (m *Mat) CheckVisited() bool{
    lis := m.head
    for lis != nil{
        check := lis.info.CheckVisited()
        if check == false{
            return false
        }
        lis = lis.next
    }
    return true
}



func (m *Mat) AStar(current Corridinent){
    calculating := false
    rs := m.Length
    cs := m.head.info.Length
    

    k := 0
    tot := rs*cs
    for calculating == false{
        //prog := float64(k)/float64(tot)
        //fmt.Printf("Progress Loop: %v\n", prog)
        //check states around calculating
        row := current.Row
        col := current.Col

        next := Corridinent{Row: row, Col: col}

        var point *Node


        point = m.GetIdx(row, col)
        for i:=row-1;i<=row+1;i++{
                if (i >= 0) && (i < rs){
                    new_p := m.GetIdx(i, col)
                    new_d := point.info.distance + new_p.info.weight_y 
                    if new_p.info.visited == false {
                        if new_p.info.distance > new_d{
                            new_p.info.distance = new_d
                        }
                    }
                    
                }
        }

                new_prev := point.prev
                new_next := point.next
                if new_prev != nil{
                new_d := point.info.distance + new_prev.info.weight_x 
                if new_prev.info.visited == false {
                    if new_prev.info.distance > new_d{
                        new_prev.info.distance = new_d
                    }
                }
                }

                if new_next != nil{
                    if new_next.info.visited == false{
                        new_d := point.info.distance + new_next.info.weight_x 
                        if new_next.info.distance > new_d{
                            new_next.info.distance = new_d
                        }
                    }
                }

        point.info.visited = true
        next_d := math.Inf(1)

        lis := m.head
        for i:=0;i<=rs;i++{
            new_p := lis.info.head
            for j:=0; j<=cs;j++{
                if (i>=0) && (i <rs) && (j>=0) && (j<cs){
                    if new_p.info.visited == false{
                        if next_d > float64(new_p.info.distance){
                            next_d = float64(new_p.info.distance)
                            next.Row = i
                            next.Col = j
                        }
                    }
                    new_p = new_p.next
                }
            }

            if lis.next != nil{
                lis = lis.next
            }
        }

        current = next

        k += 1
        if k == tot{
            calculating = true
        }
    }
}

func (m *Mat) SetInit(row int, col int){
    lis := m.head
    for i:=0;i<row;i++{
        lis = lis.next
    }
    lis.info.SetInit(col)
}

func (l *List) SetInit(col int){
    p := l.head
    for i:=0;i<col;i++{
        p = p.next
    }
    p.info.distance=0
}


func (m *Mat) SetDist(row int, col int, d int64){
    lis := m.head
    for i:=0;i<row;i++{
        lis = lis.next
    }
    lis.info.SetDist(col, d)
}

func (l *List) SetDist(col int, d int64){
    p := l.head
    for i:=0;i<col;i++{
        p = p.next
    }
    p.info.distance=d
}


func (m *Mat) SetVisited(row int, col int, d bool){
    lis := m.head
    for i:=0;i<row;i++{
        lis = lis.next
    }
    lis.info.SetVisited(col, d)
}

func (l *List) SetVisited(col int, d bool){
    p := l.head
    for i:=0;i<col;i++{
        p = p.next
    }
    p.info.visited=d
}


func (m *Mat) Reset(){
    lis := m.head
    for lis != nil{
        lis.info.Reset()
        lis = lis.next
    }
}

func (l *List) Reset(){
    p := l.head
    for p!=nil{
        p.info.distance = 9223372036854775807
        p.info.visited = false
        p = p.next
    }
}
