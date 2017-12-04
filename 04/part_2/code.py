import sys

def main(lines):
    c = 0
    for line in lines:
        if process(line.strip()):
            c = c + 1
    print(c)

def process(line):
    words = line.split(' ')
    for i, w in enumerate(words):
        for w2 in words[i+1:]:
            if is_anagram(w, w2):
                return False
    return True

def is_anagram(w1, w2):
    if not len(w1) == len(w2):
        return False
    for c in w1:
        if not c in w2:
            return False
    for c in w2:
        if not c in w1:
            return False
    return True

if __name__ == "__main__":
    main(open(sys.argv[1], 'r').readlines())
