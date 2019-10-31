#include <iostream>
#include <fstream>
#include <string>
#include <vector>
#include <algorithm>

struct diff {
    std::string str;
    std::uint8_t value;
};

diff diffIds(std::string id1, std::string id2) {
    std::uint8_t value = 0;
    std::string str = "";
    for (std::uint8_t i = 0; i < id1.size(); ++i) {
        if (id1[i] == id2[i]) {
            str += id1[i];
        } else {
            value++;
        }
    }

    return diff{str, value};
}

int main(int argc, char *argv[]) {

    if (argc != 2) {
        std::cerr << "Bad number of parameters" << std::endl;
        std::cerr << "Usage: " << argv[0] << " <filename>" << std::endl;
        return EXIT_FAILURE;
    }

    std::ifstream file(argv[1]);
    std::vector<std::string> ids;
    std::string id;

    while (!file.eof()) {
        file >> id;
        ids.push_back(id);
    }

    diff smallest{"", 100};

    for (auto id1 = std::begin(ids); id1 < std::end(ids); id1++) {
        for (auto id2 = id1 + 1; id2 < std::end(ids); id2++) {
            diff local = diffIds(*id1, *id2);
            if (local.value < smallest.value) {
                smallest = local;
            }
        }
    }

    std::cout << "Smallest difference: " << smallest.str << "; " << smallest.value << std::endl;

    return EXIT_SUCCESS;
}