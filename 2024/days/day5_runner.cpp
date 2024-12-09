#include <iostream>
#include <string>
#include <fstream>
#include <sstream>
#include <vector>


#include "day5.h"

using namespace std;

int main(int argc, char* argv[]){
    string fName = argv[1];
    string line;
    ifstream inputFile (fName);

    std::vector<Rule> rules;

    std::vector<std::vector<int>> print_orders;


    int modif = 1;
    if (inputFile.is_open()){
        while ( getline (inputFile,line) ){
            if (line == ""){
                break;
            }

            size_t pos = line.find('|');
    
            int first = std::stoi(line.substr(0, pos));
            int second = std::stoi(line.substr(pos + 1));

            rules.push_back(Rule(first, second));

        }

        while (getline(inputFile, line)){

            std::stringstream ss(line);
    
            std::string temp;
            std::vector<int> order;
    
            while (std::getline(ss, temp, ',')) {
                order.push_back(std::stoi(temp));
            }

            print_orders.push_back(order);
        }
        

        inputFile.close();
    } else {
        cout << "Unable to open file";
    }

    int sum = count_order_middle(print_orders, rules);
    int sum2 = correct_orders(print_orders, rules);

    std::cout << "the sum of middles is: " << sum << "\n";
    std::cout << "the sum of middles after correcting orders: " << sum2 << "\n";

}