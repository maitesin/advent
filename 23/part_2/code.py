import math

def main(lower, upper, step):
    not_primes = []
    for num in range(lower, upper + 1, step):
        if not all(num%i != 0 for i in range(2, int(math.sqrt(num))+1)):
            not_primes.append(num)
    print(len(not_primes))

if __name__ == "__main__":
    main(106700, 123700, 17)
