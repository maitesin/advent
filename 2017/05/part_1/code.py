import sys

def main(lines):
    instructions = prepare_instructions(lines)
    sp = SaveProcess(instructions)
    while not sp.is_done():
        sp.process_next()
    print(sp.gc)

def prepare_instructions(lines):
    return [int(x.strip()) for x in lines]

class SaveProcess(object):
    def __init__(self, instructions):
        self.inst = instructions
        self.gc = 0
        self.pc = 0

    def is_done(self):
        return self.pc >= len(self.inst)

    def process_next(self):
        next_pc = self.inst[self.pc]
        self.inst[self.pc] = next_pc + 1
        self.pc = self.pc + next_pc
        self.gc = self.gc + 1

if __name__ == "__main__":
    main(open(sys.argv[1], 'r').readlines())
