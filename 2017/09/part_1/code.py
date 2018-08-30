import sys
from enum import Enum

class Group(object):
    def __init__(self, lexer):
        self.children = []
        while lexer.has_next():
            t = lexer.next()
            if t == Token.BEGIN_GROUP:
                self.children.append(Group(lexer))
            elif t == Token.END_GROUP:
                return
            if t == Token.BEGIN_GARBAGE:
                while t != Token.END_GARBAGE:
                    t = lexer.next()
    def count(self):
        return 1 + sum([x._count(2) for x in self.children])

    def _count(self, value):
        return value + sum([x._count(value+1) for x in self.children])

class Token(Enum):
    BEGIN_GROUP = 0
    END_GROUP = 1
    ESCAPE = 2
    BEGIN_GARBAGE = 3
    END_GARBAGE = 4
    COMA = 5
    EVERYTHING_ELSE = 6

class Lexer(object):
    def __init__(self, stream):
        self.stream = stream[1:-1] #Ignore the external group
        self.pos = 0
        self.last = len(self.stream)

    def _get_token(self, value):
        if value == '!':
            return Token.ESCAPE
        elif value == '>':
            return Token.END_GARBAGE
        elif value == '<':
            return Token.BEGIN_GARBAGE
        elif value == '}':
            return Token.END_GROUP
        elif value == '{':
            return Token.BEGIN_GROUP
        elif value == ',':
            return Token.COMA
        else:
            return Token.EVERYTHING_ELSE

    def has_next(self):
        return self.pos != self.last

    def next(self):
        n = self.stream[self.pos]
        self.pos = self.pos + 1
        t = self._get_token(n)
        while t == Token.ESCAPE:
            n = self.stream[self.pos + 1]
            self.pos = self.pos + 2
            t = self._get_token(n)
        return t


def main(stream):
    lexer = Lexer(stream)
    group = Group(lexer)
    print(group.count())

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
            group, plus = parse(stream[pos+1:])
            pos = pos + plus
            children.append(group)
        pos = pos + 1

if __name__ == "__main__":
    main(open(sys.argv[1], 'r').readline().strip())
