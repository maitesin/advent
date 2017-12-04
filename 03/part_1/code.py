import sys
import math

def main(num):
    square = square_ring(num)
    side_len = square - 1
    path_to_one = math.floor(square/2)
    side_offset_to_middle = path_to_one - 1
    prev_square = square - 2
    print(path_to_one + find_side_trip(num, prev_square * prev_square, side_offset_to_middle, side_len))

def square_ring(num):
    square = 1
    while square * square < num:
        square = square + 2
    return square

def find_side_trip(num, base, offset, side):
    first = base + 1
    if num < first + side:
        return int(math.fabs(num - first - offset))
    elif num < first + side * 2:
        return int(math.fabs(num - first - offset - side))
    elif num < first + side * 3:
        return int(math.fabs(num - first - offset - side * 2))
    else:
        return int(math.fabs(num - first - offset - side * 3))

if __name__ == "__main__":
    main(int(open(sys.argv[1], 'r').readline().strip()))
