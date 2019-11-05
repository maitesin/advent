#include <iostream>
#include <fstream>
#include <vector>
#include <algorithm>
#include <unordered_map>

struct entry {
    std::uint32_t id;
    std::uint32_t x, y;
    std::uint32_t x_offset, y_offset;
};

struct tile {
    bool marked, twice;
    std::uint32_t id;
    tile() : marked(0), twice(0) {}
};

std::vector<std::vector<tile>> generate_matrix(std::uint32_t x,std::uint32_t y) {
    std::vector<std::vector<tile>> matrix;

    for (std::uint32_t i = 0; i < x; ++i) {
        matrix.push_back(std::vector<tile>(y));
    }

    return matrix;
}

int main(int argc, char *argv[]) {

    if (argc != 2) {
        std::cerr << "Bad number of parameters" << std::endl;
        std::cerr << "Usage: " << argv[0] << " <filename>" << std::endl;
        return EXIT_FAILURE;
    }

    std::ifstream file(argv[1]);
    std::vector<entry> entries;
    char c;
    entry tmp;

    while (!file.eof()) {
        file >> c;
        file >> tmp.id;
        file >> c;
        file >> tmp.x_offset;
        file >> c;
        file >> tmp.y_offset;
        file >> c;
        file >> tmp.x;
        file >> c;
        file >> tmp.y;

        entries.push_back(tmp);
    }

    entries.pop_back(); // remove the doubled last entry

    std::uint32_t max_x = 0, max_y = 0;

    for (auto && entry : entries) {
        auto x = entry.x + entry.x_offset;
        if (x > max_x)
            max_x = x;

        auto y = entry.y + entry.y_offset;
        if (y > max_y)
            max_y = y;
    }

    auto matrix = generate_matrix(max_x, max_y);
    std::unordered_map<std::uint32_t, bool> no_overlaped;

    for (auto && entry : entries) {
        bool overlap = false;
        for (auto i = entry.x_offset; i < entry.x_offset + entry.x; ++i) {
            for (auto j = entry.y_offset; j < entry.y_offset + entry.y; ++j) {
                if (matrix[i][j].marked) {
                    matrix[i][j].twice = true;
                    overlap = true;
                } else {
                    matrix[i][j].marked = true;
                    matrix[i][j].id = entry.id;
                }
            }
        }
        if (!overlap) {
            no_overlaped[entry.id] = true;
        } else {
            for (auto i = entry.x_offset; i < entry.x_offset + entry.x; ++i) {
                for (auto j = entry.y_offset; j < entry.y_offset + entry.y; ++j) {
                    if (matrix[i][j].twice) {
                        no_overlaped.erase(matrix[i][j].id);
                    }
                }
            }
        }
    }

    for (auto && entry : no_overlaped) {
        std::cout << entry.first << std::endl;
    }

    return EXIT_SUCCESS;
}