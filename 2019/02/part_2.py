import sys

def main(filename):
    with open(filename, 'r') as f:
        code = [int(x) for x in f.readline().split(',')]
        for noun in range(100):
            for verb in range(100):
                local_code = code.copy()
                local_code[1] = noun
                local_code[2] = verb
                execute(local_code, 0)
                if local_code[0] == 19690720:
                    print(f'noun={noun}; verb={verb} -> {100 * noun + verb}')
                    return

def execute(code, pos):
    optcode = code[pos]
    if optcode == 1:
        # Add
        oper1 = code[code[pos + 1]]
        oper2 = code[code[pos + 2]]
        code[code[pos + 3]] = oper1 + oper2
    elif optcode == 2:
        # Mult
        oper1 = code[code[pos + 1]]
        oper2 = code[code[pos + 2]]
        code[code[pos + 3]] = oper1 * oper2
    elif optcode == 99:
        # Exit
        return
    else:
        raise ValueError(f'Invalid optcode found: {optcode}')
    
    execute(code, pos + 4)

if __name__ == "__main__":
    main(sys.argv[1])