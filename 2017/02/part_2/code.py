import sys

def main(content):
    c = 0
    for line in content:
        l = list_of_ints(line.strip())
        v = even_div(l)
        c = c + v
    print(c)

def list_of_ints(line):
    return [int(x) for x in line.split('\t')]

def even_div(l):
    for i, x in enumerate(l):
        for j, y in enumerate(l[i + 1:]):
            if x % y == 0:
                return int(x / y)
            if y % x == 0:
                return int(y / x)
    print("This should never be shown")

if __name__ == "__main__":
    main(open(sys.argv[1], 'r').readlines())
