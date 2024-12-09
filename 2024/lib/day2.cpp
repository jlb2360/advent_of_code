#include "day2.h"
#include <vector>
#include <deque>
#include <iostream>

bool check_levels(std::deque<int> lvls, int past, int modif, int &dampened){
    bool res = false;
    // * reached the end of the deque
    if (lvls.size() == 0){
        return true;
    }

    int current = lvls.front();


    if (current == past){
        return false;
    }

    int new_modif = 0;

    if (current > past){
        new_modif = 1;
    } else if (current < past){
        new_modif = -1;
    }

    if (modif != new_modif && modif != 0){
        return false;
    }

    int diff = (current - past)*new_modif;

    if (diff < 1){
       return false;
    } else if (diff > 3){
        return false;
    }

    lvls.pop_front();

    res = check_levels(lvls, current, new_modif, dampened);

    if (res == false && dampened == 0 && lvls.size() > 1){
        dampened += 1;
        std::cout << "Dampened: " << current << "\n";
        //int cur_size = lvls.size();
        res = check_levels(lvls, past, new_modif, dampened);
        if (res == false){
            std::cout << "Wrong try Dampened: " << lvls.front() << "\n";
            lvls.pop_front();
            res = check_levels(lvls, current, new_modif, dampened);
        }

        return res;
    } else if (lvls.size() == 1 && dampened == 0){
        return true;
    }

    return res;

}

bool check_levels(std::deque<int> &lvls){
    int past = lvls.front();
    lvls.pop_front();
    int current = lvls.front();
    lvls.pop_front();

    int modif = 0;

    if (current > past){
        modif = 1;
    } else if (current < past){
        modif = -1;
    } else {
        return false;
    }

    int diff = (current - past)*modif;

    if (diff < 1){
        return false;
    } else if (diff > 3){
        return false;
    }

    past = current;


    for (int i = 0; i < lvls.size()+1; i++){
        current = lvls.front();
        lvls.pop_front();
        int diff = (current - past)*modif;


        if (diff < 1){
            return false;
        } else if (diff > 3){
            return false;
        }

        past = current;
    }

    return true;
    
}

bool check_levels(const std::vector<int> &lvls){
    int current = lvls[1];
    int past = lvls[0];
    int modif = 0;

    if (current > past){
        modif = 1;
    } else if (current < past){
        modif = -1;
    } else {
        return false;
    }

    int diff = (current - past)*modif;

    if (diff < 1){
        return false;
    } else if (diff > 3){
        return false;
    }

    past = lvls[1];
    
    for (int i = 2; i < lvls.size(); i++){
        current = lvls[i];
        int diff = (current - past)*modif;


        if (diff < 1){
            return false;
        } else if (diff > 3){
            return false;
        }

        past = lvls[i];
    }

    return true;
}