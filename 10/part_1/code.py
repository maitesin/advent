import sys

LENGHT = 255

def main(line):
    lenghts = [int(x) for x in line.split(',')]
    original = [x for x in range(LENGHT)]
    pos = 0
    skip = 0
    for lenght in lenghts:
        if lenght != 1:
            original = process(original, pos, (pos+lenght)%LENGHT)
        pos = (pos + lenght + skip)%LENGHT
        skip = skip + 1
    print(original[0]*original[1])

def process(orig, pos, lenght):
    if pos > lenght:
        # Circular case
        rev = orig[pos:] + orig[:lenght]
        rev.reverse()
        orig = rev[lenght:] + orig[lenght:pos] + rev[:lenght]
    elif pos == lenght:
        # All selected
        rev = orig[pos:] + orig[:lenght]
        rev.reverse()
        orig = rev[LENGHT-pos:] + rev[:LENGHT-pos]
    else:
        # Normal case
        rev = orig[pos:lenght]
        rev.reverse()
        orig = orig[:pos] + rev + orig[lenght:]
    return orig

if __name__ == "__main__":
    main(open(sys.argv[1], 'r').readline().strip())
