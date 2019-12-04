import sys

def has_double_digits(value):
    if len(set(value)) != 5:
        return False
    found = False
    for pos in range(len(value) - 1):
        if value[pos] == value[pos + 1]:
            found = True
    return found

def has_six_digits(value):
    return len(value) == 6

def has_increasing_values(value):
    for pos in range(len(value) - 1):
        if value[pos] > value[pos + 1]:
            return False
    return True

def valid_passwords_generator(lower, upper):
    for value in range(lower, upper + 1):
        value_str = str(value)
        if has_six_digits(value_str) and has_double_digits(value_str) and has_increasing_values(value_str):
            yield value

def count_number_of_passwords(lower, upper):
    number_of_passwords = 0
    for _ in valid_passwords_generator(lower, upper):
        number_of_passwords += 1
    return number_of_passwords

def main(filename):
    with open(filename, 'r') as f:
        lower, upper = map(int, f.readline().split('-'))
        return count_number_of_passwords(lower, upper)

if __name__ == '__main__':
    print(main(sys.argv[1]))