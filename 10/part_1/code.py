import sys

LENGHT = 256
#LENGHT = 5

def main(line):
    lenghts = [int(x) for x in line.split(',')]
    original = [x for x in range(LENGHT)]
    pos = 0
    skip = 0
    for lenght in lenghts:
        if lenght != 1 and lenght != 0:
            original = process(original, pos, lenght)
        pos = (pos + lenght + skip)%LENGHT
        skip = skip + 1
    print(original[0]*original[1])

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

if __name__ == "__main__":
    main(open(sys.argv[1], 'r').readline().strip())
