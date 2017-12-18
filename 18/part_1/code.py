import sys

class Registers(object):
    def __init__(self):
        self.regs = {}
        self.last = None

    def play(self, v):
        self.last = v

    def read(self, k):
        return self.regs.get(k, 0)

    def write(self, k, v):
        self.regs[k] = v

class Instruction(object):
    def __init__(self, input):
        self.type = input[:3]
        if self.type == 'snd' or self.type == 'rcv':
            self.value = input[4:]
        else:
            self.x, self.y = input[4:].split(' ')

    def execute(self, regs):
        if self.type == 'set':
            value = None
            if len(self.y) == 1 and self.y.isalpha():
                value = regs.read(self.y)
            else:
                value = int(self.y)
            regs.write(self.x, value)
        elif self.type == 'add':
            value = None
            if len(self.y) == 1 and self.y.isalpha():
                value = regs.read(self.y)
            else:
                value = int(self.y)
            regs.write(self.x, value + regs.read(self.x))
        elif self.type == 'mul':
            value = None
            if len(self.y) == 1 and self.y.isalpha():
                value = regs.read(self.y)
            else:
                value = int(self.y)
            regs.write(self.x, value * regs.read(self.x))
        elif self.type == 'snd':
            if len(self.value) == 1 and self.value.isalpha():
                regs.play(regs.read(self.value))
            else:
                regs.play(int(self.value))
        elif self.type == 'mod':
            value = None
            if len(self.y) == 1 and self.y.isalpha():
                value = regs.read(self.y)
            else:
                value = int(self.y)
            regs.write(self.x, regs.read(self.x) % value)
        elif self.type == 'rcv':
            v = None
            if len(self.value) == 1 and self.value.isalpha():
                v = regs.read(self.value)
            else:
                v = int(self.value)
            if v != 0:
                return regs.last
        elif self.type == 'jgz':
            value = None
            if len(self.x) == 1 and self.x.isalpha():
                value = regs.read(self.x)
            else:
                value = int(self.x)
            if value > 0:
                return int(self.y)
        else:
            print("This should never be shown")
        return None

def main(lines):
    instructions = []
    for line in lines:
        instructions.append(Instruction(line.strip()))

    regs = Registers()
    pc = 0
    while True:
        result = instructions[pc].execute(regs)
        if result != None:
            if instructions[pc].type == 'rcv':
                print(result)
                break
            elif instructions[pc].type == 'jgz':
                pc += result
                continue
        pc += 1

if __name__ == "__main__":
    main(open(sys.argv[1], 'r').readlines())
