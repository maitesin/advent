import sys

def main(lines):
    def parse(line):
        depth, range = line.strip().split(': ')
        return (int(depth), int(range))
    levels = [parse(line) for line in lines]
    index = 0
    while True:
        caught = []
        for level in levels:
            if (level[0] + index) % (2 * (level[1] - 1)) == 0:
                caught.append(level[0] * level[1])
        if (len(caught) == 0):
            print(index)
            break
        index += 1

if __name__ == "__main__":
    main(open(sys.argv[1], 'r').readlines())
