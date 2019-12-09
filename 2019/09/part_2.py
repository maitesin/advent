import sys

class IntComputer:
    def __init__(self, code, io_initial_value = 0):
        self._io = io_initial_value
        self._code = code
        self._relative_base = 0
        self._pos = 0
        self._halted = False
    

    def get_operand(self, pos, mode):
        if mode == 0: # position mode
            return self._code[self._code[pos]]
        elif mode == 1: # immediate mode
            return self._code[pos]
        elif mode == 2: # relative mode
            return self._code[self._relative_base + self._code[pos]]
        else:
            raise ValueError("Invalid mode used: {}".format(mode))

    def store_value(self, pos, mode, value):
        if mode == 0: # position mode
            self._code[self._code[pos]] = value
        elif mode == 1: # immediate mode
            self._code[pos] = value
        elif mode == 2: # relative mode
            self._code[self._relative_base + self._code[pos]] = value
        else:
            raise ValueError("Invalid mode used: {}".format(mode))
    
    def _execute(self):
        a, b, c, optcode = IntComputer._parse_instruction(str(self._code[self._pos]))
        
        if optcode == 1:
            # Add
            oper1 = self.get_operand(self._pos + 1, c)
            oper2 = self.get_operand(self._pos + 2, b)
            self.store_value(self._pos + 3, a, oper1 + oper2)

            self._pos += 4
        elif optcode == 2:
            # Mult
            oper1 = self.get_operand(self._pos + 1, c)
            oper2 = self.get_operand(self._pos + 2, b)
            self.store_value(self._pos + 3, a, oper1 * oper2)

            self._pos += 4
        elif optcode == 3:
            # Save input
            self.store_value(self._pos + 1, c, self._io)
            
            self._pos += 2
        elif optcode == 4:
            # Load input
            self._io = self.get_operand(self._pos + 1, c)

            self._pos += 2
        elif optcode == 5:
            oper1 = self.get_operand(self._pos + 1, c)
            oper2 = self.get_operand(self._pos + 2, b)
            if oper1 != 0:
                self._pos = oper2
            else:
                self._pos += 3
        elif optcode == 6:
            oper1 = self.get_operand(self._pos + 1, c)
            oper2 = self.get_operand(self._pos + 2, b)
            if oper1 == 0:
                self._pos = oper2
            else:
                self._pos += 3
        elif optcode == 7:
            oper1 = self.get_operand(self._pos + 1, c)
            oper2 = self.get_operand(self._pos + 2, b)
            self.store_value(self._pos + 3, a, 1 if oper1 < oper2 else 0)

            self._pos += 4
        elif optcode == 8:
            oper1 = self.get_operand(self._pos + 1, c)
            oper2 = self.get_operand(self._pos + 2, b)
            self.store_value(self._pos + 3, a, 1 if oper1 == oper2 else 0)

            self._pos += 4
        elif optcode == 9:
            oper = self.get_operand(self._pos + 1, c)
            self._relative_base += oper

            self._pos += 2
        elif optcode == 99:
            # Exit
            self._halted = True
            return
        else:
            raise ValueError('Invalid optcode found: {}'.format(optcode))

    def run(self):
        while not self._halted:
            self._execute()
    
    def get_io_value(self):
        return self._io
    
    @staticmethod
    def _parse_instruction(instruction):
        a = 0
        b = 0
        c = 0
        de = 0

        inst = instruction

        if len(inst) == 5:
            a = int(inst[0])
            inst = inst[1:]
        
        if len(inst) == 4:
            b, c, de = int(inst[0]), int(inst[1]), int(inst[2:])
        elif len(inst) == 3:
            c, de = int(inst[0]), int(inst[1:])
        elif len(inst) == 2 or len(inst) == 1:
            de = int(inst)
        else:
            raise ValueError("Wrong instruction: {}".format(instruction))

        return a, b, c, de

def main(filename):
    with open(filename, 'r') as f:
        code = [int(x) for x in f.readline().split(',')] + [0] * 1000
        computer = IntComputer(code, 2)
        computer.run()
        print(computer.get_io_value())


if __name__ == "__main__":
    main(sys.argv[1])