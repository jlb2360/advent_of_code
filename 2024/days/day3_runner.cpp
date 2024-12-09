#include <iostream>
#include <string>
#include <fstream>
#include <sstream>
#include <vector>


#include "day3.h"

using namespace std;

int main(int argc, char* argv[]){
    string fName = argv[1];

    int sum = 0;

    string line;
    ifstream inputFile (fName);


    int modif = 1;
    if (inputFile.is_open()){
        while ( getline (inputFile,line) ){
            std::vector<Token*> toks;

            int idx = 0;

            tokenize_string(line, toks, idx);

            sum += process_tokens(toks, modif);

        }
        inputFile.close();
    } else {
        cout << "Unable to open file";
    }

    std::cout << "Sum of mults: " << sum << "\n";

}