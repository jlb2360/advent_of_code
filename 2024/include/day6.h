#pragma once

#include<vector>

class MapObjects{
public:
    bool isBlock = false;
    bool Visited = false;

    MapObjects(bool ib): isBlock(ib){}
    ~MapObjects(){}
};

class Guard{
public:
    std::vector<int> pos = std::vector<int>({0, 0});
    std::vector<int> dir = std::vector<int>({1, 0});

    Guard(std::vector<int> p, std::vector<int> d) : pos(p), dir(d){}
    Guard(){}
    ~Guard(){}

    void change_dir();
    void revert_pos();
    void change_pos();
    bool isBlocked(const std::vector<std::vector<MapObjects>> &);
};

int FindPath(std::vector<std::vector<MapObjects>> &, Guard &g);

void print_map(std::vector<std::vector<MapObjects>> &);