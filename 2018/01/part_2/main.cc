#include<string>
#include<iostream>
#include<fstream>
#include<vector>
#include<unordered_map>
#include<algorithm>

std::vector<int64_t> process_file(std::string filename) {
    std::vector<int64_t> freqs;

    char sign;
    int64_t value;

    std::ifstream file(filename);

    while(!file.eof()) {
        file >> sign >> value;
        if (sign == '+') {
            freqs.push_back(value);
        } else {
            freqs.push_back(-value);
        }
    }
    file.close();

    return freqs;
}

int main(int argc, char *argv[]) {

    if (argc != 2) {
        std::cerr << "Bad number of parameters" << std::endl;
        std::cerr << "Usage: " << argv[0] << " <filename>" << std::endl;
        std::exit(-1);
    }

    int64_t pos = 0;
    std::unordered_map<int64_t, bool> visited;

    auto freqs = process_file(argv[1]);

    for (;;) {
        for (auto && freq : freqs) {
            pos += freq;

            auto found = visited.find(pos);
            if (found != std::end(visited)) {
                std::cout << pos << std::endl;
                return 0;
            }
            visited[pos] = true;
        }
    }
}
