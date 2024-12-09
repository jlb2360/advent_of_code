#include <iostream>
#include <string>
#include <fstream>
#include <sstream>
#include <deque>


#include "day2.h"

using namespace std;

void print_vector(std::vector<int> &lvls){
    for (int i = 0; i < lvls.size(); i++){
        std::cout << lvls[i] << ", ";
    }
    std::cout << "\n";
    
}

int main(int argc, char* argv[]){
    string fName = argv[1];

    int sum = 0;

    string line;
    ifstream inputFile (fName);

    int i = 0;
    if (inputFile.is_open()){
        while ( getline (inputFile,line) ){
            std::deque<int> lvls;
            std::vector<int> lvl_vec;
            std::stringstream ss(line);
            string word;

            while ( getline(ss, word, ' ')){
                lvls.push_back(stoi(word));
            }

            int past = lvls.front();
            lvls.pop_front();
            

            int dampened = 0;
            bool res = check_levels(lvls, past, 0, dampened);

            if (res == false){
                dampened = 1;
                int current = lvls.front();
                lvls.pop_front();
                res = check_levels(lvls, current, 0, dampened);
                
                if (res == false){
                    dampened = 1;
                    res = check_levels(lvls, past, 0, dampened);
                }
            }

            if (res == true){
                std::cout << "SAFE idx: " << i << "\n";
                sum += 1;
            } else {
                std::cout << "UNSAFE idx: " << i << "\n";
            }
            
            i += 1;
        }
        inputFile.close();
    } else {
        cout << "Unable to open file";
    }

    std::cout << "Sum of safe levels: " << sum << "\n";

}