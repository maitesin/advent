import sys

def main(content):
    c = 0
    for line in content:
        l = list_of_ints(line.strip())
        min, max = min_max(l)
        c = c + (max - min)
    print(c)

def list_of_ints(line):
    return [int(x) for x in line.split('\t')]

def min_max(l):
    min = sys.maxsize
    max = -sys.maxsize
    for x in l:
        if x < min:
            min = x
        if x > max:
            max = x
    return min, max

if __name__ == "__main__":
    main(open(sys.argv[1], 'r').readlines())
