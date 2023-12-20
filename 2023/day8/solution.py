from math import gcd
from functools import reduce

def follow_instructions(instructions, current_node):
    for instruction in instructions:
        if instruction == "L":
            current_node = network[current_node][0]
        elif instruction == "R":
            current_node = network[current_node][1]
    return current_node


with open("inputs/8.txt", "r") as file:
    input = file.read().splitlines()

instructions = input[0]
inst_len = len(instructions)
network = {}
steps = 0

for line in input[2:]:
    node, connections = line.split(" = ")
    connections = connections[1:-1].split(", ")
    network[node] = connections

current_node = "AAA"
while current_node != "ZZZ":
    current_node = follow_instructions(instructions, current_node)
    steps += inst_len

print("Solution 1: ", steps)
