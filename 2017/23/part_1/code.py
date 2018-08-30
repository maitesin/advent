import sys

class Registers(object):
    def __init__(self):
        self.regs = {}

    def read(self, k):
        return self.regs.get(k, 0)

    def write(self, k, v):
        self.regs[k] = v

class Instruction(object):
    def __init__(self, input):
        self.type = input[:3]
        self.x, self.y = input[4:].split(' ')

    def execute(self, regs):
        if self.type == 'set':
            value = None
            if len(self.y) == 1 and self.y.isalpha():
                value = regs.read(self.y)
            else:
                value = int(self.y)
            regs.write(self.x, value)
        elif self.type == 'sub':
            value = None
            if len(self.y) == 1 and self.y.isalpha():
                value = regs.read(self.y)
            else:
                value = int(self.y)
            regs.write(self.x, regs.read(self.x) - value)
        elif self.type == 'mul':
            value = None
            if len(self.y) == 1 and self.y.isalpha():
                value = regs.read(self.y)
            else:
                value = int(self.y)
            regs.write(self.x, value * regs.read(self.x))
        elif self.type == 'jnz':
            value = None
            if len(self.x) == 1 and self.x.isalpha():
                value = regs.read(self.x)
            else:
                value = int(self.x)
            if value != 0:
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
    muls = 0
    while pc < len(instructions):
        if instructions[pc].type == 'mul':
            muls += 1
        result = instructions[pc].execute(regs)
        if result != None and instructions[pc].type == 'jnz':
            pc += result
        else:
            pc += 1
    print(muls)

if __name__ == "__main__":
    main(open(sys.argv[1], 'r').readlines())
