import sys

class Point:
    def __init__(self, x, y):
        self._x = x
        self._y = y

    @property
    def x(self):
        return self._x

    @property
    def y(self):
        return self._y

    def manhattan_distance(self):
        return abs(self.x) + abs(self.y)

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

    def __eq__(self, other):
        return self.x == other.x and self.y == other.y

    def __hash__(self):
        return hash((self.x, self.y))

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

        extended_steps = []
        for step in self._steps:
            extended_steps += Cable._break_into_single_step_steps(step)

        self._points = []
        current_point = Point(0, 0)
        for step in extended_steps:
            current_point += step
            self._points.append(current_point)

    @staticmethod
    def _break_into_single_step_steps(step):
        return [Step(f'{step.direction}1')] * step.distance

    @property
    def steps(self):
        return self._steps

    @property
    def points(self):
        return self._points

    def find_crosses(self, cable):
        return set(self.points) & set(cable.points)

    def __repr__(self):
        return f'Cable(steps={self._steps})'


def main(filename):
    with open(filename, 'r') as f:
        cable1 = Cable(f.readline().split(','))
        cable2 = Cable(f.readline().split(','))

        crosses = cable1.find_crosses(cable2)
        smallest_distance = sys.maxsize
        for cross in crosses:
            distance = cross.manhattan_distance()
            if distance < smallest_distance:
                smallest_distance = distance
        print(smallest_distance)


if __name__ == "__main__":
    main(sys.argv[1])
