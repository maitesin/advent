#include <iostream>
#include <fstream>
#include <vector>
#include <algorithm>

struct log {
    std::uint8_t year, month, day, hour, minute;
    std::string info;

    void print() {
        std::cout << year << "-" << month << "-" << day
        << " " << hour << ":" << minute << "->" << info << std::endl;
    }
};

std::vector<log> read_file(std::string filename) {
    std::ifstream file(filename);
    std::vector<log> entries;
    std::string tmp;

    while (!file.eof()) {   
        std::getline(file, tmp);
        std::cout << tmp << std::endl;
    }

    return entries;
}

int main(int argc, char *argv[]) {

    if (argc != 2) {
        std::cerr << "Bad number of parameters" << std::endl;
        std::cerr << "Usage: " << argv[0] << " <filename>" << std::endl;
        return EXIT_FAILURE;
    }

    auto entries = read_file(argv[1]);

    for (auto && entry : entries) {
        entry.print();
    }

    return EXIT_SUCCESS;
}