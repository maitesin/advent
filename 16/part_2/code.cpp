#include <iostream>
#include <string>
#include <fstream>
#include <sstream>
#include <vector>
#include <algorithm>
#include <iterator>
#include <memory>

struct Action {
  virtual void perform(std::string & vec) const = 0;
};

struct Spin : public Action {
  Spin(std::string_view action) : len(std::atoi(std::string(action.substr(1)).c_str())) {}
  void perform(std::string & vec) const {
    std::rotate(vec.rbegin(), vec.rbegin() + len, vec.rend());
  }
  short len;
};

struct Exchange : public Action {
  Exchange(std::string_view action) : i(std::atoi(std::string(action.substr(1, std::distance(std::begin(action), std::find(std::begin(action), std::end(action), '/')))).c_str())), j(std::atoi(std::string(action.substr(std::distance(std::begin(action), std::find(std::begin(action), std::end(action), '/'))+1)).c_str())) {}
  void perform(std::string & vec) const {
    std::iter_swap(vec.begin() + i, vec.begin() + j);
  }
  short i, j;
};

struct Partner : public Action {
  Partner(std::string_view action) : a(action[1]), b(action[3]) {}
  void perform(std::string & vec) const {
    auto a_iter = std::find(vec.begin(), vec.end(), a);
    auto b_iter = std::find(vec.begin(), vec.end(), b);
    std::iter_swap(a_iter, b_iter);
  }
  char a, b;
};

size_t find_cicle(std::string progs, const std::vector<std::unique_ptr<Action>> & actions) {
  size_t index = 1;
  std::string original = progs;
  while (true) {
    for (auto && action : actions) {
      action->perform(progs);
    }
    if (progs == original) return index;
    ++index;
  }
}

int main(int argc, char *argv[]) {
  std::vector<std::unique_ptr<Action>> actions;
  std::string progs("abcdefghijklmnop");
  std::ifstream input_file(argv[1], std::ifstream::in);
  std::string action;
  size_t mod;

  while(std::getline(input_file, action, ',')) {
    switch(action[0]) {
      case 's':
        actions.push_back(std::make_unique<Spin>(action));
        break;
      case 'x':
        actions.push_back(std::make_unique<Exchange>(action));
        break;
      case 'p':
        actions.push_back(std::make_unique<Partner>(action));
        break;
      default:
        std::cerr << "You should never see this" << std::endl;
    }
  }

  mod = 1000000000 % find_cicle(progs, actions);

  for (size_t t = 0; t < mod; ++t) {
    for (auto &&act : actions) {
      act->perform(progs);
    }
  }

  std::cout << progs << std::endl;

  return 0;
}
