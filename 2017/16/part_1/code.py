import sys

def main(input):
    actions = input.split(',')
    programs = ['a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p']
    for action in actions:
        if action[0] == 's':
            programs = spin(programs, int(action[1:]))
        elif action[0] == 'x':
            pos1, pos2 = action[1:].split('/')
            programs = exchange(programs, int(pos1), int(pos2))
        elif action[0] == 'p':
            a, b = action[1:].split('/')
            programs = partner(programs, a, b)
        else:
            print("This should never be printed")
            return
    print(''.join(programs))

def spin(progs, pos):
    return progs[-pos:] + progs[:-pos]

def exchange(progs, i, j):
    if i == j:
        return progs
    if j < i:
        i, j = j, i
    return progs[:i] + [progs[j]] + progs[i+1:j] + [progs[i]] + progs[j+1:]

def partner(progs, a, b):
    a_pos = progs.index(a)
    b_pos = progs.index(b)
    return exchange(progs, a_pos, b_pos)

if __name__ == "__main__":
    main(open(sys.argv[1], 'r').readline().strip())
