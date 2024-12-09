#include "day1.h"
#include <deque>
#include <limits>
#include <cmath>
#include <iostream>
#include <algorithm>

int check_lists(std::deque<int> lis1, std::deque<int> lis2){
    
    int res = 0;
    int small1 = std::numeric_limits<int>::max();
    int small2 = std::numeric_limits<int>::max();

    while (!lis1.empty() && !lis2.empty()){
        for (int i = 0; i < lis1.size()+1; i++){
            int val = lis1.front();
            lis1.pop_front();

            if (val<small1){
                if (small1 != std::numeric_limits<int>::max()){
                    lis1.push_back(small1);
                }
                small1 = val;
            } else {
                lis1.push_back(val);
            }
        }

        for (int i = 0; i < lis2.size()+1; i++){
            int val = lis2.front();
            lis2.pop_front();

            if (val<small2){
                if (small2 != std::numeric_limits<int>::max()){
                    lis2.push_back(small2);
                }
                small2 = val;
            } else {
                lis2.push_back(val);
            }
        }

        int ans = std::abs(small2 - small1);

    
        res += ans;
        small1 = std::numeric_limits<int>::max();
        small2 = std::numeric_limits<int>::max();
        
    }
    
    return res;
}

int sim_score(std::deque<int> lis1, std::deque<int> lis2){
    int res = 0;

    while (!lis1.empty()){
        int val = lis1.front();
        lis1.pop_front();

        int count = std::count(lis2.begin(), lis2.end(), val);

        res += count * val;
    }

    return res;
}
