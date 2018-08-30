import sys

def combinations(bridge, value, ports):
    nexts = [port for port in ports if value in port]
    if not nexts:
        return [(sum([p[0] + p[1] for p in bridge]), len(bridge))]
    counts = []
    for next in nexts:
        index = ports.index(next)
        counts += combinations(bridge + [next], next[0] if next[0] != value else next[1], ports[:index] + ports[index+1:])
    return counts

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
        counts += combinations([start], start[0] if start[0] != 0 else start[1],
                ports[:index] + ports[index+1:])
    def get_max_lenght(item):
        return item[1]
    max_lenght = sorted(counts, key=get_max_lenght)[-1][1]
    unpacked = [x[0] for x in counts if x[1] == max_lenght]
    print(max(unpacked))


if __name__ == "__main__":
    main(open(sys.argv[1], 'r').readlines())
