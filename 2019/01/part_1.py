import sys

def main(filename):
    with open(filename, 'r') as f:
        total_value = 0
        for line in f.readlines():
            total_value += (int(line) // 3) - 2
        print(total_value)

if __name__ == "__main__":
    main(sys.argv[1])