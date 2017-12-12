import sys

def main(lines):
    connected = ['0']
    prev = 0
    curr = len(connected)
    while prev != curr:
        for line in lines:
            src, dests = parse(line.strip())
            if src in connected:
                add_all(connected, src, dests)
                continue
            for dest in dests:
                if dest in connected:
                    add_all(connected, src, dests)
                    break
        prev = curr
        curr = len(connected)
    print(len(connected))

def parse(line):
    src, dests = line.split(' <-> ')
    if ',' in dests:
        return src, dests.split(', ')
    else:
        return src, [dests]

def add_all(connected, src, dests):
    if not src in connected:
        connected.append(src)
    for dest in dests:
        if not dest in connected:
            connected.append(dest)

if __name__ == "__main__":
    main(open(sys.argv[1], 'r').readlines())
