#include <iostream>
#include <string>
#include <sstream>
#include <fstream>
#include <algorithm>
#include <unordered_map>
#include <vector>
#include <stdexcept>
#include <cmath>

struct Image {
  Image(std::string_view input) {
    auto slash = std::find(input.begin(), input.end(), '/');
    size_t pos = 0;
    while(slash != input.end()) {
      raw.push_back(std::string(input.substr(pos, std::distance(input.begin(), slash)-pos)));
      pos = std::distance(input.begin(), slash) + 1;
      slash = std::find(slash + 1, input.end(), '/');
    }
    raw.push_back(std::string(input.substr(pos)));
  }

  std::string single_line() {
    std::ostringstream oss;
    auto it = raw.begin();
    for (; it != raw.end() - 1; ++it) {
      oss << *it << "/";
    }
    oss << *it;
    return oss.str();
  }

  void rotate() {
    std::vector<std::string> rotated(raw.size(), "");
    for (auto && line : raw) {
      for (size_t i = line.size() - 1, j = 0; i > 0; --i, ++j) {
        rotated[j] += line[i];
      }
      rotated[line.size() - 1] += line[0];
    }
    raw = rotated;
  }

  void flip() {
    std::reverse(raw.begin(), raw.end());
  }

  size_t count() {
    size_t c = 0;
    for (auto && line : raw) {
      c += std::count(line.begin(), line.end(), '#');
    }
    return c;
  }

  std::vector<Image> split() {
    std::vector<Image> images;
    if (raw.size()%2 == 0) {
      for (size_t i = 0; i < raw.size(); i += 2) {
        for (size_t j = 0; j < raw.size(); j += 2) {
          std::ostringstream oss;
          oss << raw[i][j] << raw[i][j+1] << '/' << raw[i+1][j] << raw[i+1][j+1];
          images.push_back(Image(oss.str()));
        }
      }
    } else {
      for (size_t i = 0; i < raw.size(); i += 3) {
        for (size_t j = 0; j < raw.size(); j += 3) {
          std::ostringstream oss;
          oss << raw[i][j] << raw[i][j+1] << raw[i][j+2] << '/' << raw[i+1][j] << raw[i+1][j+1] << raw[i+1][j+2] << '/' << raw[i+2][j] << raw[i+2][j+1] << raw[i+2][j+2];
          images.push_back(Image(oss.str()));
        }
      }
    }
    return images;
  }

  std::vector<std::string> raw;
};

struct ArtistBook {
  ArtistBook(std::ifstream & file) {
    std::string line;
    while (std::getline(file, line)) {
      auto first_space = std::distance(line.begin(), std::find(line.begin(), line.end(), ' '));
      auto second_space = first_space + 4;
      rules[line.substr(0, first_space)] = line.substr(second_space, line.size()-1);
    }
  }
  std::unordered_map<std::string, std::string> rules;
};

Image find_image_in_artist_book(const ArtistBook & ab, Image & img) {
  for (size_t t = 0; t < 2; ++t) {
    for (size_t i = 0; i < 4; ++i) {
      auto found = ab.rules.find(img.single_line());
      if (found != ab.rules.end()) {
        return Image(found->second);
      }
      img.rotate();
    }
    img.flip();
  }
  throw std::logic_error("Not rotate or flip image has been found in the ArtistBook");
}

Image merge(const std::vector<Image> & images) {
  std::ostringstream oss;
  size_t size = images.size();
  size_t step = std::sqrt(size);
  for (size_t i = 0; i < size; i += step) {
    for (size_t k = 0; k < images[i].raw.size(); ++k) {
      for (size_t j = 0; j < step; ++j) {
        oss << images[i+j].raw[k];
      }
      oss << '/';
    }
  }
  std::string output(oss.str());
  return Image(output.substr(0, output.size()-1));
}

void iterate(const ArtistBook & ab, Image & image, size_t iterations) {
  for (size_t i = 0; i < iterations; ++i) {
    auto images = image.split();
    std::transform(images.begin(), images.end(), images.begin(), [&ab](auto & img){ return find_image_in_artist_book(ab, img); });
    image = merge(images);
  }
  std::cout << image.count() << std::endl;
}

int main(int argc, char *argv[]) {
  size_t iterations = std::atoll(argv[1]);
  std::ifstream input_file(argv[2], std::ifstream::in);
  ArtistBook ab(input_file);
  Image img(".#./..#/###");
  iterate(ab, img, iterations);
  return 0;
}
