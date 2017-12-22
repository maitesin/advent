#include <iostream>
#include <string>
#include <fstream>
#include <algorithm>
#include <vector>
#include <map>
#include <cmath>

enum direction {UP, LEFT, DOWN, RIGHT};

direction turn_left(direction current) {
  switch (current) {
    case UP:
      return LEFT;
    case LEFT:
      return DOWN;
    case DOWN:
      return RIGHT;
    case RIGHT:
      return UP;
    default:
      std::cout << "This should never be seen" << std::endl;
  }
}

direction turn_right(direction current) {
  switch (current) {
    case UP:
      return RIGHT;
    case LEFT:
      return UP;
    case DOWN:
      return LEFT;
    case RIGHT:
      return DOWN;
    default:
      std::cout << "This should never be seen" << std::endl;
  }
}

std::map<std::pair<int64_t, int64_t>, bool> get_grid(std::ifstream & input_file) {
  std::map<std::pair<int64_t, int64_t>, bool> grid;
  std::vector<std::string> input;
  std::string line;

  while(std::getline(input_file, line)) {
    input.push_back(line);
  }
  int64_t mid_way = std::floor(input.size()/2.0);
  int64_t size = input.size();
  for (int64_t i = 0; i < size; ++i) {
    for (int64_t j = 0; j < size; ++j) {
      if (input[i][j] == '#') {
        grid[std::make_pair(i - mid_way, j - mid_way)] = true;
      }
    }
  }
  return grid;
}

int main(int /*argc*/, char *argv[]) {
  size_t iters = std::atoll(argv[1]);
  std::ifstream input_file(argv[2], std::ifstream::in);
  auto grid = get_grid(input_file);

  size_t count = 0;
  int64_t i = 0, j = 0;
  direction dir = direction::UP;
  for (size_t t = 0; t < iters; ++t) {
    auto pair = std::make_pair(i, j);
    auto found = grid.find(pair);
    if (found != grid.end()) {
      dir = turn_right(dir);
      grid.erase(pair);
    } else {
      dir = turn_left(dir);
      grid[pair] = true;
      ++count;
    }
    switch (dir) {
      case UP:
        --i;
        break;
      case LEFT:
        --j;
        break;
      case DOWN:
        ++i;
        break;
      case RIGHT:
        ++j;
        break;
      default:
        std::cout << "NEVER" << std::endl;
    }
  }

  std::cout << count << std::endl;
  return 0;
}
