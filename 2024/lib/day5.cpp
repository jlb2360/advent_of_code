#include "day5.h"
#include <cmath>
#include <iostream>

bool Rule::isRule(int p){
    if (this->second_page == p){
        return true;
    }

    return false;
}

RULERESULT Rule::check_rule(int p1, int p2) const{
    if (first_page == p1 && second_page == p2){
        return RULERESULT::PASS;
    } else if (first_page == p2 && second_page == p1){
        return RULERESULT::FAIL;
    }

    return RULERESULT::NOTAPPLICABLE;
}


RULERESULT check_order(const std::vector<int> &order, const Rule &rule, int val_idx, int check_idx){

    if (check_idx >= order.size()){
        return RULERESULT::PASS;
    }

    RULERESULT res = rule.check_rule(order[val_idx], order[check_idx]);

    switch (res)
    {
    case RULERESULT::PASS:
        res = check_order(order, rule, val_idx, check_idx+1);
        break;
    case RULERESULT::FAIL:
        return res;
        break;
    case RULERESULT::NOTAPPLICABLE:
        res = check_order(order, rule, val_idx, check_idx+1);
        break;
    default:
        std::cout << "(val_idx, check_idx): (" << order[val_idx] << ", " << order[check_idx] << ")";
        break;
    }

    return res;

}

int find_idx_to_switch(const std::vector<int> &order, const Rule &rule, int val_idx, int check_idx){
    int ret_idx = -1;

    RULERESULT res = rule.check_rule(order[val_idx], order[check_idx]);

    switch (res){
        case RULERESULT::PASS:
            ret_idx = find_idx_to_switch(order, rule, val_idx, check_idx+1);
            break;
        case RULERESULT::FAIL:
            return check_idx;
            break;
        case RULERESULT::NOTAPPLICABLE:
            ret_idx = find_idx_to_switch(order, rule, val_idx, check_idx+1);
            break;
        default:
            std::cout << "(val_idx, check_idx): (" << order[val_idx] << ", " << order[check_idx] << ")";
            break;
    }

    return ret_idx;

}

int count_order_middle(const std::vector<std::vector<int>> &orders, const std::vector<Rule> &rules){

    int sum = 0;

    for (auto order : orders){

        bool print = true;

        for (int i = 0; i < order.size(); i++){

            for (auto rule : rules){

                if (rule.isRule(order[i])){

                    RULERESULT res = check_order(order, rule, i, i+1);

                    if (res == RULERESULT::FAIL){
                        print = false;
                    }

                }

                if (print == false){
                    break;
                }

            }

            if (print==false){
                break;
            }

        }

        if (print == true){
            int idx = floor(order.size()/2);

            sum += order[idx];
        }

    }

    return sum;

}

int correct_orders(const std::vector<std::vector<int>> &orders, const std::vector<Rule> &rules){

    int sum = 0;

    for (auto order : orders){
        int fixes = 0;

        for (int i = 0; i < order.size(); i++){
            bool no_errors = false;

            while (no_errors == false){
                bool print = true;
            
                for (auto rule : rules){
                    if (rule.isRule(order[i])){
                    
                        RULERESULT res = check_order(order, rule, i, i+1);

                        if (res == RULERESULT::FAIL){
                            print = false;
                        }

                    }

                    if (print == false){
                        int switch_idx = find_idx_to_switch(order, rule, i, i+1);

                        int temp = order[i];
                        order[i] = order[switch_idx];
                        order[switch_idx] = temp;
                        fixes += 1;
                        break;
                    }

                }

                if (print == true){

                    no_errors = true;
                }
            }

        }

        if (fixes != 0){
            int idx = floor(order.size()/2);

            sum += order[idx];
        }

    }

    return sum;

}