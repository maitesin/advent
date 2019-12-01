import sys

def main(filename):
    with open(filename, 'r') as f:
        total_value = 0
        for line in f.readlines():
            value = int(line)
            total_local = 0
            while value > 0:
                value = calculate_fuel(value)
                total_local += value
            total_value += total_local
        print(total_value)

def calculate_fuel(weight):
    fuel_weight = (weight // 3) - 2
    return fuel_weight if fuel_weight > 0 else 0

if __name__ == "__main__":
    main(sys.argv[1])