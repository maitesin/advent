import sys

class Point:
    def __init__(self, x, y):
        self._x = x
        self._y = y
    
    def __add__(self, step):
        if step.direction == 'U':
            return Point(x=self._x, y=self._y + step.distance)
        elif step.direction == 'D':
            return Point(x=self._x, y=self._y - step.distance)
        elif step.direction == 'L':
            return Point(x=self._x - step.distance, y=self._y)
        elif step.direction == 'R':
            return Point(x=self._x + step.distance, y=self._y)
        else:
            raise ValueError(f'Invalid direction found in step {step}')

    @property
    def x(self):
        return self._x

    @property
    def y(self):
        return self._y

    def __repr__(self):
        return f'Point(x={self._x}, y={self._y})'


class Step:
    def __init__(self, info):
        self._direction = info[0]
        self._distance = int(info[1:])
    
    @property
    def direction(self):
        return self._direction
    
    @property
    def distance(self):
        return self._distance
    
    def __repr__(self):
        return f'Step(direction={self._direction}, distance=={self._distance})'

class Cable:
    def __init__(self, steps):
        self._steps = [Step(info) for info in steps]
    
    def offsets(self):
        current = Point(0, 0)
        # U, D, L, R offsets
        u = 0
        d = 0
        l = 0
        r = 0

        for step in self._steps:
            current += step
            if current.y > u:
                u = current.y
            elif current.y < d:
                d = current.y
            elif current.x > r:
                r = current.x
            elif current.x < l:
                l = current.x

        return u, abs(d), abs(l), r
    
    def find_cross(self, cable):
        pass

    def __repr__(self):
        return f'Cable(steps={self._steps})'


def main(filename):
    with open(filename, 'r') as f:
        cable1 = Cable(f.readline().split(','))
        cable2 = Cable(f.readline().split(','))
        print('Cable1={!r}\n Cable2={!r}'.format(cable1, cable2))
        print(cable1.offsets())
        print(cable2.offsets())


if __name__ == "__main__":
    main(sys.argv[1])