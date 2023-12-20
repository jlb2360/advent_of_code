package list

import (
	"fmt"
)




type Hand struct{
    Cards [5]int
    Power int
    Bid   int
}

type Node struct{
    info Hand
    next *Node
    prev *Node
}

type List struct{
    head *Node
}


func (l *List) Insert(d Hand){
    list := &Node{info: d, next: nil, prev: nil}
    if l.head == nil{
        l.head = list
    } else {
        p := l.head
        if p.next == nil{
            if p.info.Power < d.Power{
                list.prev = p
                p.next = list
                return
            } else if p.info.Power > d.Power{
                
                p.prev = list
                list.next = p
                l.head = list
                return
            } else if p.info.Power == d.Power{
                check := check_tie(p, &d)
                if check == true{
                    p.prev = list
                    list.next = p
                    l.head = list
                    return
                } else if check == false{
                    list.prev = p
                    p.next = list
                    return
                }
            }
        }
        if p.info.Power < d.Power{
            list.prev = p
            p.next = list
            return
        } else if p.info.Power > d.Power{
                
            prev := p.prev

            p.prev = list
            list.next = p
            list.prev = prev
            return
        } else if p.info.Power == d.Power{
            check := check_tie(p, &d)
            if check == true{
                prev := p.prev

                p.prev = list
                list.next = p
                list.prev = prev
                return
            } else if check == false{
                list.prev = p
                p.next = list
                return
            }
        }
        if p.next.info.Power < d.Power{
            for p.next.info.Power < d.Power{
                p = p.next
                if p.next == nil{
                    if p.info.Power < d.Power{
                        list.prev = p
                        p.next = list
                        return
                    } else if p.info.Power > d.Power{
                    
                        prev := p.prev

                        p.prev = list
                        list.next = p
                        list.prev = prev
                        return

                    } else if p.info.Power == d.Power{
                        check := check_tie(p, &d)
                        if check == true{
                            prev := p.prev

                            p.prev = list
                            list.next = p
                            list.prev = prev
                            return
                        } else if check == false{
                            list.prev = p
                            p.next = list
                            return
                        }
                    }
                }
            }
        }

        if p.next.info.Power > d.Power{
            temp := p.next
            temp.prev = list
            list.prev = p
            list.next = temp
            p.next = list
            return
        }

        if p.next.info.Power == d.Power{
            for check_tie(p.next, &d){
                p = p.next
                if p.next != nil{
                    if p.next.info.Power > d.Power{
                        temp := p.next
                        temp.prev = list
                        list.prev = p
                        list.next = temp
                        p.next = list
                        return
                    }
                } else {
                    if p.info.Power < d.Power{
                        list.prev = p
                        p.next = list
                        return
                    } else if p.info.Power > d.Power{
                    
                        prev := p.prev

                        p.prev = list
                        list.next = p
                        list.prev = prev
                        return
                    } else if p.info.Power == d.Power{
                        check := check_tie(p, &d)
                        if check == true{
                            prev := p.prev

                            p.prev = list
                            list.next = p
                            list.prev = prev
                            return
                        } else if check == false{
                            list.prev = p
                            p.next = list
                            return
                        }
                    }
                }
            }

            temp := p.next
            temp.prev = list
            list.prev = p
            list.next = temp
            p.next = list
            return
        }

    }
}

func check_tie(p *Node, d *Hand) bool{
    for i:=0; i<5; i++{
        if p.info.Cards[i] < d.Cards[i]{
            return true
        } else if p.info.Cards[i] > d.Cards[i]{
            return false
        }
    }

    return false
}

func (l *List) PrintList(){
    p := l.head
    i := 0
    for p != nil{
        fmt.Printf("Hand %v: %v with power: %v with bid: %v\n", i, p.info.Cards, p.info.Power, p.info.Bid)
        p = p.next
        i = i + 1
    }
}

func (l *List) SumRank() int{
    i := 1
    sum := 0
    p := l.head
    for p != nil {
        sum = sum + (i*p.info.Bid)
        p = p.next
        i = i + 1
    }

    return sum
}
