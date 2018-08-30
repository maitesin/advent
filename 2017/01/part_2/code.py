import sys

def main(content):
    l = len(content)
    h = int(l/2)
    c = 0
    for i, _ in enumerate(content):
        if content[i] == content[(i + h)%l]:
            c = c + int(content[i])
    print(c)

if __name__ == "__main__":
    main(open(sys.argv[1], 'r').readline().strip())

