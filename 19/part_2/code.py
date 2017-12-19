import sys
from enum import Enum

class Direction(Enum):
    UP = 1
    DOWN = 2
    LEFT = 3
    RIGHT = 4

def change(lines, i, j, dir, count):
    if dir == Direction.DOWN or dir == Direction.UP:
        if lines[i][j-1] == '-' or (lines[i][j-1] >= 'A' and lines[i][j-1] <= 'Z'):
            return i, j-1, Direction.LEFT, count
        elif lines[i][j+1] == '-' or (lines[i][j+1] >= 'A' and lines[i][j+1] <= 'Z'):
            return i, j+1, Direction.RIGHT, count
    elif dir == Direction.RIGHT or dir == Direction.LEFT:
        if lines[i-1][j] == '|' or (lines[i-1][j] >= 'A' and lines[i-1][j] <= 'Z'):
            return i-1, j, Direction.UP, count
        elif lines[i+1][j] == '|' or (lines[i+1][j] >= 'A' and lines[i+1][j] <= 'Z'):
            return i+1, j, Direction.DOWN, count
    return -1, -1, None, count

def process(lines, i, j, dir, count):
    if lines[i][j] == ' ':
        return -1, -1, None, count
    count += 1
    if lines[i][j] == '+':
        return change(lines, i, j, dir, count)
    if dir == Direction.DOWN:
        i += 1
    elif dir == Direction.RIGHT:
        j += 1
    elif dir == Direction.UP:
        i -= 1
    elif dir == Direction.LEFT:
        j -= 1
    return i, j, dir, count

def main(lines):
    lines = [line[:-1] for line in lines]
    i = j = 0
    j = lines[0].index('|')
    dir = Direction.DOWN
    count = 0 
    while True:
        if dir != None:
            i, j, dir, count = process(lines, i, j, dir, count)
        else:
            break
    print(count)

if __name__ == "__main__":
    main(open(sys.argv[1], 'r').readlines())
