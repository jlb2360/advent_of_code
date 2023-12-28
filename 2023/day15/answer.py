def Hash(s):
    v = 0
    for c in s:
        v += ord(c)
        v *= 17
        v %= 256

    return v

def main():
    f = open('input.txt', 'r')
    s = 0
    line = f.readline()
    sl = line.split(',')

    print(sum(map(Hash, line.split(','))))
    s += sum(map(Hash, line.split(',')))

    print(s)

    boxes = [[] for i in range(256)]
    focal_lengths = {}
    f.seek(0)

    for instruction in line.split(','):
        if '-' in instruction:
            label = instruction[:-1]
            index = Hash(label)
            if label in boxes[index]:
                boxes[index].remove(label)

        else:
            label, length = instruction.split('=')
            length = int(length)

            index = Hash(label)

            if label not in boxes[index]:
                boxes[index].append(label)

            focal_lengths[label] = length

        
    total = 0
    for box_num, box in enumerate(boxes, 1):
        for lens_slot, label in enumerate(box, 1):
            total += focal_lengths[label] * box_num * lens_slot

    print(total)



if __name__ == '__main__':
    main()