import sys

class State(object):
    def __init__(self, z, o):
        self.z_value = z[0]
        self.z_move = z[1]
        self.z_state = z[2]
        self.o_value = o[0]
        self.o_move = o[1]
        self.o_state = o[2]

    def zero(self, tape, cursor):
        tape[cursor] = self.z_value
        return self.z_state, cursor + self.z_move

    def one(self, tape, cursor):
        tape[cursor] = self.o_value
        return self.o_state, cursor + self.o_move

class TuringMachine(object):
    def __init__(self, states, start_state):
        self.tape = [0]
        self.cursor = 0
        self.states = states
        self.current = start_state

    def next_step(self):
        if self.tape[self.cursor] == 0:
            self.current, self.cursor = self.states[self.current].zero(self.tape, self.cursor)
        else:
            self.current, self.cursor = self.states[self.current].one(self.tape, self.cursor)
        if self.cursor < 0:
            self.tape = [0] + self.tape
            self.cursor = 0
        elif self.cursor >= len(self.tape):
            self.tape += [0]
            self.cursor = len(self.tape) - 1

    def diagnostic_checksum(self):
        return sum(self.tape)

def main(start_state, steps):
    states = {'A': State([1, 1, 'B'],[0, -1, 'D']), 'B':State([1, 1, 'C'],[0, 1, 
'F']),'C':State([1, -1, 'C'],[1, -1, 'A']),'D':State([0, 
-1, 'E'],[1, 1, 'A']),'E':State([1, -1, 'A'],[0, 1, 'B']),'F':State([0, 1, 'C'],[0, 1, 'E'])}
    tm = TuringMachine(states, start_state)
    for iter in range(steps):
        tm.next_step()
    print(tm.diagnostic_checksum())

if __name__ == "__main__":
    main('A', 12317297)
