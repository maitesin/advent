#include <iostream>
#include <fstream>
#include <algorithm>

int main(int argc, char *argv[]) {
  // Read content of the file
  std::ifstream input_file(argv[1], std::ifstream::in);
  size_t input;
  input_file >> input;

  // Start processing
  size_t next = 1;
  for (size_t size = 1, pos = 1, i = 1; i <= 50000000; ++i, ++size) {
    pos = (pos + input) % size + 1;
    if (pos == 1) {
      next = i;
    }
  }

  // Find next value after 0
  std::cout << next << std::endl;

  return 0;
}
