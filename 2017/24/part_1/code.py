import sys

def combinations(bridge, value, ports):
    nexts = [port for port in ports if value in port]
    if not nexts:
        return sum([p[0] + p[1] for p in bridge])
    counts = []
    for next in nexts:
        index = ports.index(next)
        counts.append(combinations(bridge + [next], next[0] if next[0] != value else next[1], ports[:index] + ports[index+1:]))
    return max(counts)

def main(lines):
    ports = []
    for port in lines:
        x, y = port.strip().split('/')
        ports.append((int(x), int(y)))
    starters = []
    for port in ports:
        if 0 in port:
            starters.append(port)
    counts = []
    for start in starters:
        index = ports.index(start)
        counts.append(combinations([start], start[0] if start[0] != 0 else start[1], ports[:index] + ports[index+1:]))
    print(max(counts))


if __name__ == "__main__":
    main(open(sys.argv[1], 'r').readlines())
