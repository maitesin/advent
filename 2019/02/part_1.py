import sys

def main(filename):
    with open(filename, 'r') as f:
        code = [int(x) for x in f.readline().split(',')]
        code[1] = 12
        code[2] = 2
        execute(code, 0)
        print(code[0])
        

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