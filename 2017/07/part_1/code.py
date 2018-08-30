import sys

def main(lines):
    map_list = build(lines)
    print(find_root_(map_list))

def build(lines):
    elems = {}
    for line in lines:
        if '->' in line:
            key, deps = line_deps(line)
            elems[key] = deps
        else:
            key = line_nodeps(line)
            elems[key] = None
    return elems

def line_nodeps(line):
    return line.split(' ')[0]

def line_deps(line):
    key = line_nodeps(line)
    l = line.split('->')[1].strip()
    deps = l.split(', ')
    return key, deps

def find_root_(elems):
    for key, deps in elems.items():
        if deps != None:
            return find_root(elems, key)

def find_root(elems, key):
    for k, deps in elems.items():
        if deps != None and key in deps:
            return find_root(elems, k)
    return key


if __name__ == "__main__":
    main(open(sys.argv[1], 'r').readlines())
