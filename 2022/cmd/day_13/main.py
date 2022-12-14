#!/usr/bin/env python3

import sys
from functools import cmp_to_key


def main(content):
    pairs = content.strip().split("\n\n")
    accum = 0

    for i, pair in enumerate(pairs):
        p1, p2 = pair.strip().split("\n")
        if check_pair(eval(p1), eval(p2)) > 0:
            accum += 1 + i

    print(f'Part 1: {accum}')

    packets = [
        [[6]],
        [[2]]
    ]
    for pair in pairs:
        p1, p2 = pair.strip().split("\n")
        packets.append(eval(p1))
        packets.append(eval(p2))

    packets = sorted(packets, key=cmp_to_key(check_pair), reverse=True)
    i6 = 0
    i2 = 0
    for i, packet in enumerate(packets):
        if packet == [[6]]:
            i6 = i + 1
        elif packet == [[2]]:
            i2 = i + 1
    print(f'Part 2: {i2 * i6}')


def check_pair(p1, p2):
    if not type(p1) is list:
        p1 = [p1]
    if not type(p2) is list:
        p2 = [p2]
    for (l, r) in zip(p1, p2):
        if type(l) is list or type(r) is list:
            res = check_pair(l, r)
        else:
            res = r - l
        if res != 0:
            return res
    return len(p2) - len(p1)


if __name__ == '__main__':
    with open(sys.argv[1], 'r') as f:
        main(f.read())

