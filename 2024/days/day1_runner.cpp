#include <iostream>
#include <string>
#include <fstream>
#include <sstream>
#include <queue>


#include "day1.h"

using namespace std;

int main(int argc, char* argv[]){
    string fName = argv[1];
    std::deque<int> lis1;
    std::deque<int> lis2;

    string line;
    ifstream inputFile (fName);
    if (inputFile.is_open()){
        while ( getline (inputFile,line) ){
            std::stringstream ss(line);
            string word;

            ss >> word;

            lis1.push_back(stoi(word));

            ss >> word;

            lis2.push_back(stoi(word));

            ss.clear();
            
        }
        inputFile.close();
    } else {
        cout << "Unable to open file";
    }

    int res = check_lists(lis1, lis2);

    cout << "the summed distances is: " << res << "\n";

    res = sim_score(lis1, lis2);

    cout << "the sim score is: " << res << "\n";
}