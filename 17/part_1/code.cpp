#include <iostream>
#include <fstream>
#include <vector>
#include <algorithm>

int main(int argc, char *argv[]) {
  // Read content of the file
  std::ifstream input_file(argv[1], std::ifstream::in);
  size_t size;
  input_file >> size;

  // Start processing
  std::vector<size_t> vec = {0};
  auto pos = vec.begin();
  for (size_t i = 1; i <= 2017; ++i) {
    pos = vec.begin() + ((std::distance(vec.begin(), pos) + size)%vec.size()) + 1;
    vec.insert(pos, i);
    pos = std::find(vec.begin(), vec.end(), i);
  }

  // Find next value after 2017
  pos = std::find(vec.begin(), vec.end(), 2017);
  std::cout << *(++pos) << std::endl;

  return 0;
}
