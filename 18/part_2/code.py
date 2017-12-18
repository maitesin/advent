import sys

class Registers(object):
    def __init__(self, p):
        self.regs = {}
        self.regs['p'] = p
        self.count = 0

    def play(self):
        self.count += 1

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

    def __str__(self):
        if hasattr(self, 'value'):
            return "%s %s" % (self.type, self.value)
        else:
            return "%s %s %s" % (self.type, self.x, self.y)

    def execute(self, regs, read_q, write_q):
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
                regs.play()
                write_q.append(regs.read(self.value))
            else:
                regs.play()
                write_q.append(int(self.value))
        elif self.type == 'mod':
            value = None
            if len(self.y) == 1 and self.y.isalpha():
                value = regs.read(self.y)
            else:
                value = int(self.y)
            regs.write(self.x, regs.read(self.x) % value)
        elif self.type == 'rcv':
            if len(read_q) > 0:
                elem = read_q.pop(0)
                regs.write(self.value, elem)
                return 1
            else:
                return -1
        elif self.type == 'jgz':
            value = None
            if len(self.x) == 1 and self.x.isalpha():
                value = regs.read(self.x)
            else:
                value = int(self.x)
            if value > 0:
                if len(self.y) == 1 and self.y.isalpha():
                    return regs.read(self.y)
                else:
                    return int(self.y)
        else:
            print("This should never be shown")
        return None

def main(lines):
    instructions = []
    for line in lines:
        instructions.append(Instruction(line.strip()))

    l = len(instructions)
    regs_0 = Registers(0)
    regs_1 = Registers(1)
    queue_0 = []
    queue_1 = []
    blocked_0 = False
    blocked_1 = False
    pc_0 = 0
    pc_1 = 0
    while True:
        if pc_0 >= l and pc_1 >= l:
            break
        if blocked_0 and blocked_1:
            print("Deadlock")
            break
        if pc_0 < l:
            result_0 = instructions[pc_0].execute(regs_0, queue_0, queue_1)
            if result_0 != None:
                if instructions[pc_0].type == 'jgz':
                    pc_0 += result_0
                elif instructions[pc_0].type == 'rcv':
                    blocked_0 = result_0  == -1
                    if not blocked_0:
                        pc_0 += 1
                else:
                    print("This is a lie from 0")
            else:
                pc_0 += 1
        if pc_1 < l:
            result_1 = instructions[pc_1].execute(regs_1, queue_1, queue_0)
            if result_1 != None:
                if instructions[pc_1].type == 'jgz':
                    pc_1 += result_1
                elif instructions[pc_1].type == 'rcv':
                    blocked_1 = result_1 == -1
                    if not blocked_1:
                        pc_1 += 1
                else:
                    print("This is a lie from 1")
            else:
                pc_1 += 1
    print(regs_1.count)

if __name__ == "__main__":
    main(open(sys.argv[1], 'r').readlines())
