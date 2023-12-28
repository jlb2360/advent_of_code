package main

import (
	"bufio"
	c "day20/connector"
	q "day20/queue"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainPart1(t *testing.T) {

	expected := 32000000

	f, err := os.Open("inputTest1.txt")
	if err != nil {
		t.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	connectors := make(map[string]c.Connector)

	for scanner.Scan() {
		line := scanner.Text()
		c_name, outputs := strings.Split(line, " -> ")[0], strings.Split(line, " -> ")[1]

		if c_name == "broadcaster" {
			outputs := strings.Split(outputs, ", ")
			connectors["broadcaster"] = &c.Broadcaster{Destinations: outputs}
		} else if c_name == "output" {
			connectors["output"] = &c.Output{}
		} else {
			if string(c_name[0]) == "%" {
				c_name = strings.TrimPrefix(c_name, "%")
				outputs := strings.Split(outputs, ", ")
				connectors[c_name] = &c.Flip_Flop{On: false, Destinations: outputs}
			} else {
				c_name = strings.TrimPrefix(c_name, "&")
				outputs := strings.Split(outputs, ", ")
				connectors[c_name] = &c.Conjunction{Brain: make(map[string]string), Destinations: outputs}
			}
		}
	}

	for key, con := range connectors {
		if _, ok := con.(*c.Flip_Flop); ok {
			for i := 0; i < len(con.(*c.Flip_Flop).Destinations); i++ {
				d := con.(*c.Flip_Flop).Destinations[i]
				if _, ok := connectors[d]; ok {
					conjuct := connectors[d]
					if _, ok := conjuct.(*c.Conjunction); ok {
						conjuct.(*c.Conjunction).Brain[key] = "lo"
					}
				}
			}
		} else if _, ok := con.(*c.Conjunction); ok {
			for i := 0; i < len(con.(*c.Conjunction).Destinations); i++ {
				d := con.(*c.Conjunction).Destinations[i]
				if _, ok := connectors[d]; ok {
					conjuct := connectors[d]
					if _, ok := conjuct.(*c.Conjunction); ok {
						conjuct.(*c.Conjunction).Brain[key] = "lo"
					}
				}
			}
		}
	}

	lo := 0
	hi := 0
	var cycle int
	for i := 0; i < 1000; i++ {
		queue := q.Queue{}

		queue.Enqueue("broadcaster", "lo", "None")
		lo += 1

		for queue.Length > 0 {
			destination, signal, sender := queue.Dequeue()
			if _, ok := connectors[destination]; ok {
				con := connectors[destination]
				signal, destinations, err := con.GetSignal(signal, sender)
				if err != nil {
					continue
				}
				for _, d := range destinations {
					if signal == "lo" {
						lo++
					} else {
						hi++
					}
					queue.Enqueue(d, signal, destination)
				}

			}
		}

		found_cycle := true
		//checking if cycle completed and if so, break
		for _, con := range connectors {
			if _, ok := con.(*c.Flip_Flop); ok {
				if con.(*c.Flip_Flop).On {
					found_cycle = false
					break
				}
			}
			if _, ok := con.(*c.Conjunction); ok {
				for _, value := range con.(*c.Conjunction).Brain {
					if value == "hi" {
						found_cycle = false
						break
					}
				}
			}
		}

		if found_cycle {
			cycle = i
			break
		}
	}

	cycle += 1
	mult := (1000 / cycle) + (1000 % cycle)

	tot := (lo * mult) * (hi * mult)

	assert.Equal(t, expected, tot)
}

func TestMainPart1_2(t *testing.T) {

	expected := 11687500

	f, err := os.Open("inputTest2.txt")
	if err != nil {
		t.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	connectors := make(map[string]c.Connector)

	for scanner.Scan() {
		line := scanner.Text()
		c_name, outputs := strings.Split(line, " -> ")[0], strings.Split(line, " -> ")[1]

		if c_name == "broadcaster" {
			outputs := strings.Split(outputs, ", ")
			connectors["broadcaster"] = &c.Broadcaster{Destinations: outputs}
		} else if c_name == "output" {
			connectors["output"] = &c.Output{}
		} else {
			if string(c_name[0]) == "%" {
				c_name = strings.TrimPrefix(c_name, "%")
				outputs := strings.Split(outputs, ", ")
				connectors[c_name] = &c.Flip_Flop{On: false, Destinations: outputs}
			} else {
				c_name = strings.TrimPrefix(c_name, "&")
				outputs := strings.Split(outputs, ", ")
				connectors[c_name] = &c.Conjunction{Brain: make(map[string]string), Destinations: outputs}
			}
		}
	}

	for key, con := range connectors {
		if _, ok := con.(*c.Flip_Flop); ok {
			for i := 0; i < len(con.(*c.Flip_Flop).Destinations); i++ {
				d := con.(*c.Flip_Flop).Destinations[i]
				if _, ok := connectors[d]; ok {
					conjuct := connectors[d]
					if _, ok := conjuct.(*c.Conjunction); ok {
						conjuct.(*c.Conjunction).Brain[key] = "lo"
					}
				}
			}
		} else if _, ok := con.(*c.Conjunction); ok {
			for i := 0; i < len(con.(*c.Conjunction).Destinations); i++ {
				d := con.(*c.Conjunction).Destinations[i]
				if _, ok := connectors[d]; ok {
					conjuct := connectors[d]
					if _, ok := conjuct.(*c.Conjunction); ok {
						conjuct.(*c.Conjunction).Brain[key] = "lo"
					}
				}
			}
		}
	}

	lo := 0
	hi := 0
	var cycle int
	for i := 0; i < 1000; i++ {
		queue := q.Queue{}

		queue.Enqueue("broadcaster", "lo", "None")
		lo += 1

		for queue.Length > 0 {
			destination, signal, sender := queue.Dequeue()
			if _, ok := connectors[destination]; ok {
				con := connectors[destination]
				signal, destinations, err := con.GetSignal(signal, sender)
				if err != nil {
					continue
				}
				for _, d := range destinations {
					if signal == "lo" {
						lo++
					} else {
						hi++
					}
					queue.Enqueue(d, signal, destination)
				}

			}
		}

		found_cycle := true
		//checking if cycle completed and if so, break
		for _, con := range connectors {
			if _, ok := con.(*c.Flip_Flop); ok {
				if con.(*c.Flip_Flop).On {
					found_cycle = false
					break
				}
			}
			if _, ok := con.(*c.Conjunction); ok {
				for _, value := range con.(*c.Conjunction).Brain {
					if value == "hi" {
						found_cycle = false
						break
					}
				}
			}
		}

		if found_cycle {
			cycle = i
			break
		}
	}

	cycle += 1
	mult := (1000 / cycle) + (1000 % cycle)

	tot := (lo * mult) * (hi * mult)

	assert.Equal(t, expected, tot)
}
