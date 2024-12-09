import sys


def get_start():
    for r, row in enumerate(grid):
        for c, val in enumerate(row):
            if val == "^":
                return (r, c)


with open(sys.argv[1], "r") as f:
    grid = list(map(list, map(str.strip, f.readlines())))

num_rows = len(grid)
num_cols = len(grid[0])

r, c = get_start()
print(r, c)
dr, dc = -1, 0
visited = set()

steps = 0
while True:
    steps += 1
    visited.add((r, c))
    if not (0 <= r + dr < num_rows and 0 <= c + dc < num_cols):
        break
    if grid[r + dr][c + dc] == "#":
        dc, dr = -dr, dc
    
    r += dr
    c += dc
        
    print(r, c)
    if steps == 4000:
        break

print(f"Part 1: {len(visited)}")


start_r, start_c = get_start()


def check_for_loop():
    r, c = start_r, start_c
    dr, dc = -1, 0
    visited = set()

    while True:
        if (r, c, dr, dc) in visited:
            return True
        visited.add((r, c, dr, dc))
        if not (0 <= r + dr < num_rows and 0 <= c + dc < num_cols):
            return False
        if grid[r + dr][c + dc] == "#":
            dc, dr = -dr, dc
        else:
            r += dr
            c += dc


part2 = 0
for ro in range(num_rows):
    for co in range(num_cols):
        if grid[ro][co] != ".":
            continue
        grid[ro][co] = "#"
        if check_for_loop():
            part2 += 1
        grid[ro][co] = "."

print(f"Part 2: {part2}")
