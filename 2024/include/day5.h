#pragma once

#include <vector>
#include <iostream>

enum RULERESULT{
    PASS = 0,
    FAIL,
    NOTAPPLICABLE
};

class Rule {
    public:
        int first_page;
        int second_page;

        Rule(int p1, int p2): first_page(p1), second_page(p2){}


        bool isRule(int p);
        RULERESULT check_rule(int p1, int p2) const;

        void print_rule(){
            std::cout << first_page << "|" << second_page << "\n";
        }
        
};

RULERESULT check_order(const std::vector<int> &, const Rule &, int val_idx, int check_idx);

int count_order_middle(const std::vector<std::vector<int>> &, const std::vector<Rule> &);

int correct_orders(const std::vector<std::vector<int>> &, const std::vector<Rule> &);