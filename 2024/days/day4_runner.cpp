#include <iostream>
#include <string>
#include <fstream>
#include <sstream>
#include <vector>


#include "day4.h"

using namespace std;

int main(int argc, char* argv[]){
    string fName = argv[1];

    int sum = 0;
    int sum2 = 0;

    string line;
    ifstream inputFile (fName);

    std::vector<std::string> lines;


    int modif = 1;
    if (inputFile.is_open()){
        while ( getline (inputFile,line) ){
            lines.push_back(line);
        }
        inputFile.close();
    } else {
        cout << "Unable to open file";
    }

    count_xmas(sum, lines);

    count_x_mas(sum2, lines);

    std::cout << "Sum of xmas: " << sum << "\n";

    std::cout << "Sum of x-mas: " << sum2 << "\n";

}