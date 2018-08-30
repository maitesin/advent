#include <iostream>
#include <fstream>
#include <sstream>
#include <string>
#include <algorithm>
#include <vector>

struct Point {
  Point(int64_t x, int64_t y, int64_t z) : x(x), y(y), z(z) {}

  void increment(Point p) {
    x += p.x;
    y += p.y;
    z += p.z;
  }

  int64_t x, y, z;
};

struct Particle {
  Particle (Point p, Point v, Point a, size_t id) : p(std::move(p)), v(std::move(v)), a(std::move(a)), id(id) {}
  Point p, v, a;
  size_t id;
};

Point get_point(std::string_view::iterator begin, std::string_view::iterator end) {
  auto first_coma = std::find(begin, end, ',');
  auto second_coma = std::find(first_coma + 1, end, ',');
  auto x = std::atoll(std::string(begin+1, first_coma).c_str());
  auto y = std::atoll(std::string(first_coma+1, second_coma).c_str());
  auto z = std::atoll(std::string(second_coma+1, end).c_str());
  return Point(x, y, z);
}

Particle get_particle_from_input(std::string_view input, size_t id) {
  auto get_iters = [](auto begin, auto end) {
    auto tmp = std::find(begin, end, '<');
    return std::make_tuple(tmp, std::find(tmp, end, '>'));
  };
  auto [pos_begin, pos_end] = get_iters(input.begin(), input.end());
  auto [vel_begin, vel_end] = get_iters(pos_end, input.end());
  auto [acc_begin, acc_end] = get_iters(vel_end, input.end());
  return Particle(get_point(pos_begin, pos_end), get_point(vel_begin, vel_end), get_point(acc_begin, acc_end), id);
}

void process(std::vector<Particle> & particles) {
  std::vector<size_t> latest;
  auto manhattan = [](const Particle & particle) -> int64_t { return std::llabs(particle.p.x) + std::llabs(particle.p.y) + std::llabs(particle.p.z); };
  auto cmp = [&manhattan](const Particle & p1, const Particle & p2) { return manhattan(p1) < manhattan(p2); };
  size_t threshold = 1000;
  while (true) {
    for (auto && particle : particles) {
      particle.v.increment(particle.a);
      particle.p.increment(particle.v);
    }
    latest.push_back(std::min_element(particles.begin(), particles.end(), cmp)->id);
    if (latest.size() > threshold) {
      latest.erase(latest.begin());
    }
    if (std::count(latest.begin(), latest.end(), *latest.begin()) == threshold) {
      std::cout << *latest.begin() << std::endl;
      break;
    }
  }
}

int main(int /*argc*/, char *argv[]) {
  std::ifstream input_file(argv[1], std::ifstream::in);
  std::string line;
  std::vector<Particle> particles;
  size_t id = 0;

  while(std::getline(input_file, line)) {
    particles.push_back(get_particle_from_input(line, id));
    ++id;
  }

  process(particles);

  return 0;
}
