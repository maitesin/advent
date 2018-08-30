import sys

class Generator(object):
    def __init__(self, value, factor, mul):
        self.value = value
        self.factor = factor
        self.div = 2147483647
        self.mul = mul

    def next(self):
        new = int((self.value * self.factor) % self.div)
        self.value = new
        while new % self.mul != 0:
            new = int((self.value * self.factor) % self.div)
            self.value = new
        return new

def main(lines):
    init_a = get_input_value(lines[0])
    init_b = get_input_value(lines[1])
    gen_a = Generator(init_a, 16807, 4)
    gen_b = Generator(init_b, 48271, 8)
    count = 0
    for i in range(5000000):
        bits_a = get_bits(gen_a.next())
        bits_b = get_bits(gen_b.next())
        if bits_a == bits_b:
            count += 1
    print(count)

def get_input_value(line):
    return int(line.split(' ')[4].strip())

def get_bits(value):
    bits = ""
    index = 0
    while index < 16:
        bits = str(value % 2) + bits
        value = value >> 1
        index += 1
        if value == 0:
            break
    while index < 16:
        bits = '0' + bits
        index += 1
    return bits

if __name__ == '__main__':
    main(open(sys.argv[1], 'r').readlines())
