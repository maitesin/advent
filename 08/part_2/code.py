import sys

class Bank(object):
    def __init__(self):
        self.regs = {}
        self.max = -sys.maxsize

    def inc(self, reg, value):
        self.regs[reg] = self.regs.get(reg, 0) + value
        if self.regs[reg] > self.max:
            self.max = self.regs[reg]

    def dec(self, reg, value):
        self.regs[reg] = self.regs.get(reg, 0) - value
        if self.regs[reg] > self.max:
            self.max = self.regs[reg]

    def check(self, reg, op, value):
        reg_value = self.regs.get(reg, 0)
        if op == "==":
            return reg_value == value
        elif op == "!=":
            return reg_value != value
        elif op == ">=":
            return reg_value >= value
        elif op == "<=":
            return reg_value <= value
        elif op == ">":
            return reg_value > value
        elif op == "<":
            return reg_value < value
        else:
            raise ValueError("The operation '%s' is not supported" % op)

def main(lines):
    b = Bank()
    ins = extract_instructions(lines)
    for i in ins:
        cond = i[1]
        if b.check(cond[0], cond[1], cond[2]):
            op = i[0]
            if op[1] == 'inc':
                b.inc(op[0], op[2])
            elif op[1] == 'dec':
                b.dec(op[0], op[2])
            else:
                raise ValueError("The registers are limited to 'inc' and 'dec'. '%s' is not supported" % op[1])
    print(b.max)

def extract_instructions(lines):
    l = []
    for line in lines:
        inc_dec, cond = line.strip().split(' if ')
        regadd, addi, valadd = inc_dec.strip().split(' ')
        regcmp, opcmp, valcmp = cond.strip().split(' ')
        l.append([[regadd, addi, int(valadd)],[regcmp, opcmp, int(valcmp)]])
    return l

if __name__ == "__main__":
    main(open(sys.argv[1], 'r').readlines())
