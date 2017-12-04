#include <iostream>
#include <array>

size_t calculate_position(const std::array<std::array<size_t, 50>, 50> & matrix, size_t i, size_t j) {
  return matrix[i - 1][j - 1] + matrix[i - 1][j] + matrix[i][j - 1] + matrix[i][j + 1] + matrix[i + 1][j] + matrix[i + 1][j + 1] + matrix[i - 1][j + 1] + matrix[i + 1][j - 1];
}

size_t go_up(std::array<std::array<size_t, 50>, 50> & matrix, size_t i, size_t j, size_t threshold);

size_t go_right(std::array<std::array<size_t, 50>, 50> & matrix, size_t i, size_t j, size_t threshold) {
  while(matrix[i][j - 1] == 0) {
    matrix[i][j] = calculate_position(matrix, i, j);
    if (threshold < matrix[i][j]) {
      return matrix[i][j];
    }
    ++i;
  }
  matrix[i][j] = calculate_position(matrix, i, j);
  return go_up(matrix, i, j - 1, threshold);
}

size_t go_down(std::array<std::array<size_t, 50>, 50> & matrix, size_t i, size_t j, size_t threshold) {
  while(matrix[i + 1][j] == 0) {
    matrix[i][j] = calculate_position(matrix, i, j);
    if (threshold < matrix[i][j]) {
      return matrix[i][j];
    }
    ++j;
  }
  matrix[i][j] = calculate_position(matrix, i, j);
  return go_right(matrix, i + 1, j, threshold);
}

size_t go_left(std::array<std::array<size_t, 50>, 50> & matrix, size_t i, size_t j, size_t threshold) {
  while(matrix[i][j + 1] == 0) {
    matrix[i][j] = calculate_position(matrix, i, j);
    if (threshold < matrix[i][j]) {
      return matrix[i][j];
    }
    --i;
  }
  matrix[i][j] = calculate_position(matrix, i, j);
  return go_down(matrix, i, j + 1, threshold);
}

size_t go_up(std::array<std::array<size_t, 50>, 50> & matrix, size_t i, size_t j, size_t threshold) {
  while(matrix[i - 1][j] == 0) {
    matrix[i][j] = calculate_position(matrix, i, j);
    if (threshold < matrix[i][j]) {
      return matrix[i][j];
    }
    --j;
  }
  matrix[i][j] = calculate_position(matrix, i, j);
  return go_left(matrix, i - 1, j, threshold);
}

size_t find_next_biggest_number(size_t threshold, std::array<std::array<size_t, 50>, 50> & matrix) {
  size_t const mid = 24;
  matrix[mid][mid] = 1;
  return go_up(matrix, 25, 24, threshold);
}

int main(void) {
  std::array<std::array<size_t, 50>, 50> matrix;
  for (size_t i = 0; i < 50; ++i) {
    matrix[i].fill(0);
  }

  size_t threshold = 265149;

  std::cout << find_next_biggest_number(threshold, matrix) << std::endl;

  return 0;
}
