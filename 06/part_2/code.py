import sys
sys.setrecursionlimit(10000)

def main(line):
	registers = [int(x) for x in line.split('\t')]
	dh = DebuggerHelp(registers)
	print(dh.redistribute())

class DebuggerHelp(object):
	def __init__(self, registers):
		self.regs = registers
		self.prev = None

	def _max_position_(self):
		return self.regs.index(max(self.regs))

	def _count_(self, regs):
		if self.prev != None:
			if self.prev._compare_one_(regs):
				return 1
			else:
				return self.prev._count_(regs) + 1

	def _compare_one_(self, regs):
		for x, y in zip(self.regs, regs):
			if x != y:
				return False
		return True

	def _compare_(self, regs):
		for x, y in zip(self.regs, regs):
			if x != y:
				if self.prev == None:
					return False
				else:
					return self.prev._compare_(regs)
		return True

	def redistribute(self):
		pos = self._max_position_()
		old_pos = self.regs[pos]
		self.regs[pos] = 0
		l = len(self.regs)
		for i in range(1, old_pos + 1):
			index = (pos + i) % l
			self.regs[index] = self.regs[index] + 1
		if self.prev != None:
			if self.prev._compare_(self.regs):
				return self._count_(self.regs)
			else:
				dh2 = DebuggerHelp(list(self.regs))
				dh2.prev = self
				return dh2.redistribute()
		else:
			dh2 = DebuggerHelp(list(self.regs))
			dh2.prev = self
			return dh2.redistribute()

if __name__ == "__main__":
	main(open(sys.argv[1], 'r').readline().strip())
