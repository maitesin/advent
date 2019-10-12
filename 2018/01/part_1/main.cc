#include<string>
#include<iostream>
#include<fstream>
#include<vector>

int main(int argc, char *argv[]) {

    if (argc != 2) {
        std::cerr << "Bad number of parameters" << std::endl;
        std::cerr << "Usage: " << argv[0] << " <filename>" << std::endl;
        std::exit(-1);
    }

    int64_t pos = 0;

    std::vector<std::string> freq;
    char sign;
    int64_t value;

    std::ifstream file(argv[1]);

    while(!file.eof()) {
        file >> sign >> value;
        if (sign == '+') {
            pos += value;
        } else {
            pos -= value;
        }
    }
    file.close();

    std::cout << pos << std::endl;

    return 0;
}