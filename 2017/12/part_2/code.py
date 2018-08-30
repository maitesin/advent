import sys

class Groups(object):
    def __init__(self):
        self.groups = [['0']]

    def _is_any_in_list(self, list1, list2):
        for item in list1:
            if item in list2:
                return True
        return False

    def add(self, src, dests):
        group = [src] + dests
        yes_list = []
        no_list = []
        for index, elem in enumerate(self.groups):
            if self._is_any_in_list(group, elem):
                yes_list.append(index)
            else:
                no_list.append(index)
        for pos in yes_list:
            group += self.groups[pos]
        group = list(set(group))
        new_groups = []
        for pos in no_list:
            new_groups.append(self.groups[pos])
        new_groups.append(group)
        self.groups = new_groups

    def signature(self):
        return "%s" % self.groups

def main(lines):
    groups = Groups()
    prev = ""
    curr = groups.signature()
    while prev != curr:
        for line in lines:
            src, dests = parse(line.strip())
            groups.add(src, dests)
        prev = curr
        curr = groups.signature()
    print(len(groups.groups))

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
