#include <iostream>
#include <string>
#include <fstream>
#include <sstream>
#include <vector>


#include "day6.h"

using namespace std;

int main(int argc, char* argv[]){
    string fName = argv[1];
    string line;
    ifstream inputFile (fName);

    std::vector<std::vector<MapObjects>> map;

    Guard g;
    int i = 0;
    if (inputFile.is_open()){
        while ( getline (inputFile,line) ){

            std::vector<MapObjects> map_line;
            int j = 0;
            for (auto c : line){
                
                if (c == '.'){
                    map_line.push_back(MapObjects(false));
                } else if (c == '#'){
                    map_line.push_back(MapObjects(true));
                } else if (c == '^'){
                    g.pos = std::vector<int>({i, j});
                    g.dir = std::vector<int>({-1, 0});
                    map_line.push_back(MapObjects(false));
                }
                j += 1;
            }
            i += 1;
            map.push_back(map_line);

        }
        inputFile.close();
    } else {
        cout << "Unable to open file";
    }


    int sum = FindPath(map, g);

    std::cout << "Number of objects visitid: " << sum << "\n";

}