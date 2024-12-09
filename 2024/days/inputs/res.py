import re


def get_lines(test=False):
    with open(("test" if test else "input_test") + ".txt") as file:
        for ln in file:
            yield ln.strip()


def sum_muls(muls):
    """a mul is a tuple of two integer strings"""
    return sum([int(mul[0]) * int(mul[1]) for mul in muls])


def solution_1(memory_dumps):
    p = re.compile(r"mul\((\d+),(\d+)\)")
    return sum([sum_muls(p.findall(dump)) for dump in memory_dumps])


def solution_2(memory_dumps):
    return solution_1(
        [
            re.split(r"don\'t\(\)", do_section)[0]
            for do_section in re.split(r"do\(\)", "".join(memory_dumps))
        ]
    )


print("1:", solution_1(get_lines()))
print("2:", solution_2(get_lines()))