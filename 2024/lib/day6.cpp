#include "day6.h"
#include <vector>
#include <iostream>


void Guard::change_dir(){
    // rotate 90 degrees

    std::vector<int> newD(2);


    newD[0] = this->dir[1];
    newD[1] = -1*this->dir[0];


    this->dir = newD;

}

void Guard::change_pos(){
    this->pos[0] += this->dir[0];
    this->pos[1] += this->dir[1];
}

void Guard::revert_pos(){
    this->pos[0] -= this->dir[0];
    this->pos[1] -= this->dir[1];
}

bool Guard::isBlocked(const std::vector<std::vector<MapObjects>> &Map){
    this->change_pos();

    bool isBlock = false;

     if (this->pos[0] < 0 || this->pos[0] > Map.size())
    {
        return false;
    }

    if (this->pos[1] < 0 || this->pos[1] > Map[this->pos[0]].size())
    {
        return false;
    }

    if (Map[this->pos[0]][this->pos[1]].isBlock){
        isBlock = true;
    }

    this->revert_pos();

    return isBlock;
    
}

int count_visited(std::vector<std::vector<MapObjects>> &map){
    int sum = 0;
    
    for (int i = 0; i < map.size(); i++)
    {
        for (int j = 0; j < map[i].size(); j++)
        {
            if (map[i][j].Visited)
            {
                sum += 1;
            }
            
        }
        
    }

    return sum;
    
}

int FindPath(std::vector<std::vector<MapObjects>> &map, Guard &g){
    bool walking = true;

    int step = 0;

    while (walking){
        step += 1;

        if (!map[g.pos[0]][g.pos[1]].Visited){
            map[g.pos[0]][g.pos[1]].Visited = true;
        }

        if (g.pos[0] + g.dir[0] < 0 || g.pos[0] + g.dir[0] >= map.size())
        {
            walking = false;
            break;
        }

        if (g.pos[1] + g.dir[1] < 0 || g.pos[1]+g.dir[1] >= map[g.pos[0]].size())
        {
            walking = false;
            break;
        }

        if (map[g.pos[0]+g.dir[0]][g.pos[1]+g.dir[1]].isBlock){
            g.change_dir();
        }

        g.change_pos();


        //if (step == 4000){
        //    break;
        //}

    }

    int sum = count_visited(map);

    return sum;
    
}

void print_map(std::vector<std::vector<MapObjects>> &map){
    for (int i = 0; i < map.size(); i++)
    {
        for (int j = 0; j < map[i].size(); j++)
        {
            if (map[i][j].isBlock)
            {
                std::cout << "#";
            } else{
                std::cout << ".";
            }
            
        }

        std::cout << "\n";
        
    }
    
}