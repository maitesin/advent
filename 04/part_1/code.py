import sys

def main(lines):
    c = 0
    for line in lines:
        if process(line.strip()):
            c = c + 1
    print(c)

def process(line):
    words = line.split(' ')
    orig = len(words)
    no_dups = len(set(words))
    return orig == no_dups

if __name__ == "__main__":
    main(open(sys.argv[1], 'r').readlines())
