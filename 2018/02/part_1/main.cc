#include <iostream>
#include <fstream>
#include <algorithm>

int main(int argc, char *argv[]) {

    if (argc != 2) {
        std::cerr << "Bad number of parameters" << std::endl;
        std::cerr << "Usage: " << argv[0] << " <filename>" << std::endl;
        return EXIT_FAILURE;
    }

    std::ifstream file(argv[1]);
    std::string id;
    std::uint64_t two_count = 0, three_count = 0;
    bool two_found, three_found;


    while (!file.eof()) {
        // clear the found flag states
        two_found = three_found = false;

        file >> id;
        std::sort(begin(id), end(id));

        auto it = std::begin(id);
        auto end = std::end(id);

        while (it != end) {
            char c = *it;
            auto end_of_same = std::find_if_not(it, end, [&c](auto current){ return current == c;});
            auto distance = std::distance(it, end_of_same);
            it = end_of_same++;
            two_found = distance == 2 ? true : two_found;
            three_found = distance == 3 ? true : three_found;
        }

        if (two_found) two_count++;
        if (three_found) three_count++;
    }

    std::cout << two_count * three_count << std::endl;

    return EXIT_SUCCESS;
}