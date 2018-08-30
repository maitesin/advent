import sys

def main(lines):
    map_list = build(lines)
    root = find_root_(map_list)
    find_unbalanced_(root, map_list)

def build(lines):
    elems = {}
    for line in lines:
        if '->' in line:
            key, weight, deps = line_deps(line)
            elems[key] = [deps, weight]
        else:
            key, weight = line_nodeps(line)
            elems[key] = [None, weight]
    return elems

def line_nodeps(line):
    key, w = line.split(' ')
    return key, int(w.strip()[1:-1])

def line_deps(line):
    key, w = line_nodeps(line.split('->')[0].strip())
    l = line.split('->')[1].strip()
    deps = l.split(', ')
    return key, w, deps

def find_root_(elems):
    for key, deps in elems.items():
        if deps[0] != None:
            return find_root(elems, key)

def find_root(elems, key):
    for k, deps in elems.items():
        if deps[0] != None and key in deps[0]:
            return find_root(elems, k)
    return key

def find_unbalanced_(root, elems):
    ws = []
    deps, weight = elems[root]
    if deps != None:
        for dep in deps:
            ws.append(find_unbalanced_(dep, elems))
        print (ws, '+', weight, '=', sum(ws) + weight)
        if len(set(ws)) != 1:
            print("Not only one value", ws)
        return sum(ws) + weight
    else:
        return weight

if __name__ == "__main__":
    main(open(sys.argv[1], 'r').readlines())
