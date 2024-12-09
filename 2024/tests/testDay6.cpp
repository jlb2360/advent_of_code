#include "day6.h"
#include "gtest/gtest.h"
#include <string>
#include <vector>


TEST(TestDay6, test1){
    std::vector<std::vector<MapObjects>> map;

    Guard g;

    std::vector<std::string> lines;

    lines.push_back("....#.....");
    lines.push_back(".........#");
    lines.push_back("..........");
    lines.push_back("..#.......");
    lines.push_back(".......#..");
    lines.push_back("..........");
    lines.push_back(".#..^.....");
    lines.push_back("........#.");
    lines.push_back("#.........");
    lines.push_back("......#...");
    int i = 0;
    for (auto line : lines){
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
            }
            j += 1;
        }
        i += 1;
        map.push_back(map_line);
    }

    int sum = FindPath(map, g);


    ASSERT_EQ(sum, 41);

}