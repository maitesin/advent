import sys
from enum import Enum

class Group(object):
    def __init__(self, children):
        self.children = children

    def __str__(self):
        output = '["children": '
        for child in self.children:
            output = output + str(child)
        output = output + ']'
        return output

    def count(self):
        return 1 + sum([x._count(2) for x in self.children])

    def _count(self, value):
        return value + sum([x._count(value+1) for x in self.children])

class State(Enum):
    NO_ESCAPE = 0
    ESCAPE = 1

class Garbage(Enum):
    NO_GARBAGE = 0
    GARBAGE = 1

def main(stream):
    print(stream)
    groups = parse(stream)
    print(groups.count())

def parse(stream):
    pos = 0
    children = []
    end = len(stream)
    state = State.NO_ESCAPE
    garbage = Garbage.NO_GARBAGE
    while pos < end:
        if state == State.ESCAPE:
            state = State.NO_ESCAPE
        elif stream[pos] == '!':
            state = State.ESCAPE
        elif stream[pos] == '>':
            garbage = Garbage.NO_GARBAGE
        elif stream[pos] == '<':
            garbage = Garbage.GARBAGE
        elif garbage == Garbage.GARBAGE:
            pass
        elif stream[pos] == '}':
            return Group(children)
        elif stream[pos] == '{':
            group, plus = parse_rec(stream[pos+1:])
            pos = pos + plus + 1
            children.append(group)
        pos = pos + 1

def parse_rec(stream):
    pos = 0
    children = []
    end = len(stream)
    state = State.NO_ESCAPE
    garbage = Garbage.NO_GARBAGE
    while pos < end:
        if state == State.ESCAPE:
            state = State.NO_ESCAPE
        elif stream[pos] == '!':
            state = State.ESCAPE
        elif stream[pos] == '>':
            garbage = Garbage.NO_GARBAGE
        elif stream[pos] == '<':
            garbage = Garbage.GARBAGE
        elif garbage == Garbage.GARBAGE:
            pass
        elif stream[pos] == '}':
            return Group(children), pos
        elif stream[pos] == '{':
            group, plus = parse_rec(stream[pos+1:])
            pos = pos + plus + 1
            children.append(group)
        pos = pos + 1

if __name__ == "__main__":
    main(open(sys.argv[1], 'r').readline().strip())
