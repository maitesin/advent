import sys
import math

def main(line):
    steps = [x for x in line.split(',')]
    x = 0
    y = 0
    for step in steps:
        if len(step) == 2:
            if 'n' in step:
                y = y + 1
            if 's' in step:
                y = y - 1
            if 'e' in step:
                x = x + 1
            if 'w' in step:
                x = x - 1
        else:
            if 'n' == step:
                y = y + 2
            else:
                y = y - 2
    ax = abs(x)
    ay = abs(y)
    if ax < ay:
        dif = ay - ax
        print(ax + math.floor(dif/2))
    else:
        print(ax)

if __name__ == "__main__":
    main(open(sys.argv[1], 'r').readline().strip())
