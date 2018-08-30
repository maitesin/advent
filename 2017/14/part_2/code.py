import sys

LENGHT = 256

def main(input):
    matrix = []
    for num in range(128):
        word = input + "-" + str(num)
        matrix.append(binary_array(knot_hash(word)))
    find_contiguous(matrix)

def find_contiguous(matrix):
    count = 1
    for i in range(128):
        for j in range(128):
            if matrix[i][j] == '#':
                seek_and_destroy(matrix, i, j, count)
                count += 1
    print(count-1)

def seek_and_destroy(matrix, i, j, count):
    matrix[i][j] = count
    if i > 0:
        if matrix[i-1][j] == '#':
            seek_and_destroy(matrix, i-1, j, count)
    if j > 0:
        if matrix[i][j-1] == '#':
            seek_and_destroy(matrix, i, j-1, count)
    if i < 127:
        if matrix[i+1][j] == '#':
            seek_and_destroy(matrix, i+1, j, count)
    if j < 127:
        if matrix[i][j+1] == '#':
            seek_and_destroy(matrix, i, j+1, count)

def binary_array(hash):
    array = []
    for char in hash:
        value = int(char, 16)
        small = []
        index = 4
        while value != 0:
            small.append('#' if value % 2 == 1 else '.')
            value = value >> 1
            index -= 1
        while index != 0:
            small.append('.')
            index -= 1
        small.reverse()
        array += small
    return array

def knot_hash(line):
    lenghts = [ord(c) for c in line] + [17, 31, 73, 47, 23]
    original = [x for x in range(LENGHT)]
    pos = 0
    skip = 0
    for times in range(64):
        for lenght in lenghts:
            if lenght != 1 and lenght != 0:
                original = process(original, pos, lenght)
            pos = (pos + lenght + skip)%LENGHT
            skip = skip + 1
    return hash(original)

def process(orig, pos, lenght):
    end = (pos + lenght)%LENGHT
    if pos > end:
        # Circular case
        rev = orig[pos:] + orig[:end]
        rev.reverse()
        orig = rev[lenght-end:] + orig[end:pos] + rev[:lenght-end]
    else:
        # Normal case
        rev = orig[pos:end]
        rev.reverse()
        orig = orig[:pos] + rev + orig[end:]
    return orig

def hash(numbers):
    h = ""
    for iter in range(16):
        tmp = numbers[0 + iter * 16] ^ numbers[1 + iter * 16] ^ numbers[2 + iter* 16] ^ numbers[3 + iter * 16] ^ numbers[4 + iter * 16] ^ numbers[5 + iter * 16] ^ numbers[6 + iter * 16] ^ numbers[7 + iter * 16]        ^ numbers[8 + iter * 16] ^ numbers[9 + iter * 16] ^ numbers[10 + iter * 16] ^ numbers[11 + iter * 16] ^ numbers[12 + iter * 16] ^ numbers[13 + iter * 16] ^ numbers[14 + iter * 16] ^ numbers[15 + iter * 16]
        tmp = hex(tmp)
        if len(tmp) == 3:
            h = h + '0'
        h = h + tmp[2:]
    return h

if __name__ == "__main__":
    main(open(sys.argv[1], 'r').readline().strip())
