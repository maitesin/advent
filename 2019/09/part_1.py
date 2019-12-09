import sys

input_value = 1
relative_base = 0

def main(filename):
    with open(filename, 'r') as f:
        code = [int(x) for x in f.readline().split(',')] + [0] * 1000
        execute(code, 0)
        print(input_value)

def parse_instruction(instruction):
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

def get_operand(code, pos, mode):
    if mode == 0: # position mode
        return code[code[pos]]
    elif mode == 1: # immediate mode
        return code[pos]
    elif mode == 2: # relative mode
        return code[relative_base + code[pos]]
    else:
        raise ValueError("Invalid mode used: {}".format(mode))

def store_value(code, pos, mode, value):
    if mode == 0: # position mode
        code[code[pos]] = value
    elif mode == 1: # immediate mode
        code[pos] = value
    elif mode == 2: # relative mode
        code[relative_base + code[pos]] = value
    else:
        raise ValueError("Invalid mode used: {}".format(mode))

def execute(code, pos):
    global input_value
    global relative_base
    a, b, c, optcode = parse_instruction(str(code[pos]))
    
    if optcode == 1:
        # Add
        oper1 = get_operand(code, pos + 1, c)
        oper2 = get_operand(code, pos + 2, b)
        store_value(code, pos + 3, a, oper1 + oper2)

        execute(code, pos + 4)
    elif optcode == 2:
        # Mult
        oper1 = get_operand(code, pos + 1, c)
        oper2 = get_operand(code, pos + 2, b)
        store_value(code, pos + 3, a, oper1 * oper2)

        execute(code, pos + 4)
    elif optcode == 3:
        # Save input
        store_value(code, pos + 1, c, input_value)
        
        execute(code, pos + 2)
    elif optcode == 4:
        # Load input
        input_value = get_operand(code, pos + 1, c)

        execute(code, pos + 2)
    elif optcode == 5:
        oper1 = get_operand(code, pos + 1, c)
        oper2 = get_operand(code, pos + 2, b)
        if oper1 != 0:
            execute(code, oper2)
        else:
            execute(code, pos + 3)
    elif optcode == 6:
        oper1 = get_operand(code, pos + 1, c)
        oper2 = get_operand(code, pos + 2, b)
        if oper1 == 0:
            execute(code, oper2)
        else:
            execute(code, pos + 3)
    elif optcode == 7:
        oper1 = get_operand(code, pos + 1, c)
        oper2 = get_operand(code, pos + 2, b)
        store_value(code, pos + 3, a, 1 if oper1 < oper2 else 0)

        execute(code, pos + 4)
    elif optcode == 8:
        oper1 = get_operand(code, pos + 1, c)
        oper2 = get_operand(code, pos + 2, b)
        store_value(code, pos + 3, a, 1 if oper1 == oper2 else 0)

        execute(code, pos + 4)
    elif optcode == 9:
        oper = get_operand(code, pos + 1, c)
        relative_base += oper

        execute(code, pos + 2)
    elif optcode == 99:
        # Exit
        return
    else:
        raise ValueError('Invalid optcode found: {}'.format(optcode))

if __name__ == "__main__":
    main(sys.argv[1])