package mat

import (
	"errors"
	"fmt"
	"strconv"
)

type Element struct {
    Entry   string
    IsDigit bool
    IsShape bool
}


type Node struct {
    info Element
    next *Node
}

type Row struct {
    head *Node
    nextRow *Row
}

type Mat struct {
    head *Row
}


func (m *Mat) InsertRow(){
    new_r := &Row{head: nil, nextRow: nil}
    if m.head == nil {
       m.head = new_r 
    } else {
        p := m.head
        for p.nextRow != nil {
            p = p.nextRow
        }
        p.nextRow = new_r
    }
}

func (m *Mat) InsertElement(d Element, row_idx int) (error){
    list := &Node{info: d, next: nil}
    if m.head == nil {
        m.InsertRow()
        row := m.head
        row.head = list
    } else {
        row := m.head
        for j := 1; j < row_idx; j++ {
            if row.nextRow != nil {
                row = row.nextRow
            } else {
                return errors.New("Index out of range")
            }
        }
        if row.head == nil {
            row.head = list
        } else {
            p := row.head
            for p.next != nil {
                p = p.next
            }
            p.next = list
        }
    }

    return nil
}


func (m *Mat) PrintMat(){
    row := m.head
    for row != nil {
        p := row.head
        for p != nil {
            fmt.Printf("%s",p.info.Entry)
            p = p.next
        }
        fmt.Printf("\n")
        row = row.nextRow
    }


}

func (m *Mat) GrabIndex(row_idx int, col_idx int) (Element, error){

    if col_idx <= 0 || row_idx <= 0 {
        return Element{Entry: "n", IsDigit: false, IsShape: false}, errors.New("Index out of range")
    }
    row := m.head
    for j:=1; j < row_idx; j++ {
        if row.nextRow != nil {
            row = row.nextRow
        } else {
            return Element{Entry: "n", IsDigit: false, IsShape: false}, errors.New("Index out of range")
        }
    }
    
    p := row.head
    for i:=1; i < col_idx; i++{
        if p.next != nil {
            p = p.next
        } else {
            return Element{Entry: "n", IsDigit: false, IsShape: false}, errors.New("Index out of range")
        }
    }

    return p.info, nil
}


func (m *Mat) GetSumParts(){
    row := m.head
    j := 1
    num := ""
    sum := 0
    for row != nil {
        i := 1
        p := row.head
        for p != nil {
            if p.info.IsDigit{
                scan_row_begin := j-1
                scan_col_begin := i-1
                num = num + p.info.Entry

                for p.next.info.IsDigit{
                    p = p.next
                    num = num + p.info.Entry
                    i = i + 1
                    if p.next == nil {
                        break
                    }
                }

                scan_row_end := j + 1
                scan_col_end := i + 1


                for k := scan_row_begin; k<=scan_row_end; k++{
                    for n := scan_col_begin; n<=scan_col_end; n++{
                        data, err := m.GrabIndex(k,n)
                        if err == nil{
                            if data.IsShape{
                                temp, err1 := strconv.Atoi(num)
                                if err1 == nil {
                                    sum = sum + temp
                                }
                            }
                        }
                    }
                }


            } else {
                num = ""
            }
            p = p.next
            i = i+1

        }
        row = row.nextRow
        j = j+1
    }

    fmt.Printf("Sum is %v\n", sum)
}

func (m *Mat) SumGears(){
    row := m.head
    j := 1
    sum := 0
    for row != nil {
        i := 1
        p := row.head
        for p != nil {
            num_of_digits := 0
            if p.info.IsShape{
                gr := 1
               if p.info.Entry == "*"{
                    for k := j-1; k<=j+1; k++{
                        for n := i-1; n<=i+1; n++{
                            data, err := m.GrabIndex(k,n)
                            if err == nil{
                                if data.IsDigit{
                                    num := data.Entry
                                    num_of_digits = num_of_digits + 1

                                    nd := n-1
                                    data1, err2 := m.GrabIndex(k,nd)
                                    if err2 == nil {
                                        for data1.IsDigit{
                                            num = data1.Entry + num
                                            nd = nd -1
                                            data1, err2 = m.GrabIndex(k, nd)
                                            if err2 != nil {
                                                break
                                            } 
                                        }
                                    } else {
                                        fmt.Printf("%s", err2)
                                    }

                                    n = n+1
                                    data1, err2 = m.GrabIndex(k, n)
                                    if err2 == nil {
                                        for data1.IsDigit{
                                           num = num + data1.Entry
                                           n = n+1
                                           data1, err2 = m.GrabIndex(k,n)
                                           if err2 != nil {
                                                break
                                           } 

                                        }
                                    }


                                    


                                    temp, err1 := strconv.Atoi(num)
                                    if err1 == nil {
                                        gr = gr*temp
                                    }

                                }
                            }
                            
                        }

                    }

                    if num_of_digits == 2 {
                        sum = sum + gr
                    }

               }

            }
            p = p.next


            i = i+1

        }
        row = row.nextRow
        j = j+1
    }

    fmt.Printf("Sum is %v\n", sum)

}

