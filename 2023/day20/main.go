package main

import (
	"bufio"
	c "day20/connector"
	q "day20/queue"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
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

		cycle = i

		if found_cycle {
			break
		}
	}

	cycle += 1
	mult := (1000 / cycle) + (1000 % cycle)

	tot := (lo * mult) * (hi * mult)
	fmt.Printf("Part 1: %d\n", tot)
}

func part2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
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

	feed := "lv"
	cycle_lengths := map[string]int{}
	seen := map[string]int{}

	for key, con := range connectors {
		if _, ok := con.(*c.Flip_Flop); ok {
			for i := 0; i < len(con.(*c.Flip_Flop).Destinations); i++ {
				if con.(*c.Flip_Flop).Destinations[i] == feed {
					cycle_lengths[key] = 0
					seen[key] = 0
				}
			}
		} else if _, ok := con.(*c.Conjunction); ok {
			for i := 0; i < len(con.(*c.Conjunction).Destinations); i++ {
				if con.(*c.Conjunction).Destinations[i] == feed {
					cycle_lengths[key] = 0
					seen[key] = 0
				}
			}
		}
	}

	fmt.Printf("Inputs: %v\n", cycle_lengths)

	presses := 0
	for true {
		presses += 1
		queue := q.Queue{}

		queue.Enqueue("broadcaster", "lo", "None")

		for queue.Length > 0 {
			destination, signal, sender := queue.Dequeue()
			if _, ok := connectors[destination]; ok {
				con := connectors[destination]
				signal, destinations, err := con.GetSignal(signal, sender)
				if err != nil {
					continue
				}
				for _, d := range destinations {
					if (d == feed) && (signal == "hi") {
						seen[sender] += 1
						if _, ok := cycle_lengths[sender]; !ok {
							cycle_lengths[sender] = presses
						} else {
							if presses != seen[sender]*cycle_lengths[sender] {
								fmt.Printf("Presses: %d\n", presses)
								fmt.Printf("Cycle Length: %d\n", cycle_lengths[sender])
								return
							}
						}
					}

					checking := true
					for _, val := range cycle_lengths {
						if val == 0 {
							checking = false
						}
					}
					if checking {
						fmt.Printf("Cycle Length: %d\n", cycle_lengths[sender])
					}

					queue.Enqueue(d, signal, destination)
				}

			}
		}

		fmt.Printf("Presses: %d\n", presses)
		fmt.Printf("Cycle Length: %v\n", cycle_lengths)
		fmt.Printf("Seen: %v\n", seen)

	}

	fmt.Printf("Not Found\n")
}
