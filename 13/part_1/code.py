import sys

class Level(object):
    def __init__(self, depth, range):
        self.depth = depth
        self.range = range
        self.sentinel = 0
        self.direction = 0

    def __str__(self):
        return "%s: %s (%d)" % (self.depth, self.range, self.sentinel)

    def next(self):
        if self.direction == 0:
            if self.sentinel == self.range - 1:
                self.direction = 1
                self.sentinel -= 1
            else:
                self.sentinel += 1
        else:
            if self.sentinel == 0:
                self.direction = 0
                self.sentinel = 1
            else:
                self.sentinel -= 1


class Firewall(object):
    def __init__(self, levels):
        self.levels = []
        self.packet = None
        self.caught = []

        index = 0
        for level in levels:
            while level[0] != index:
                self.levels.append(None)
                index += 1
            self.levels.append(Level(level[0], level[1]))
            index += 1

    def has_next(self):
        if self.packet == None:
            return True
        return self.packet < self.levels[-1].depth

    def next(self):
        if self.packet == None:
            # First next executed
            self.packet = 0
            self.check_caught()
            for level in self.levels:
                if level != None:
                    level.next()
        else:
            self.packet += 1
            self.check_caught()
            for level in self.levels:
                if level != None:
                    level.next()

    def check_caught(self):
        if self.levels[self.packet] != None and self.levels[self.packet].sentinel == 0:
            self.caught.append(self.levels[self.packet].depth * self.levels[self.packet].range)

def main(lines):
    def parse(line):
        depth, range = line.strip().split(': ')
        return (int(depth), int(range))
    f = Firewall([parse(line) for line in lines])
    while f.has_next():
        f.next()
    print(sum(f.caught))


if __name__ == "__main__":
    main(open(sys.argv[1], 'r').readlines())
