#include "day5.h"
#include "gtest/gtest.h"
#include <string>
#include <vector>

TEST(TESTDAY5, checkOrderFails){
    
    Rule r1 = Rule(97, 75);
    Rule r2 = Rule(29, 13);

    std::vector<int> t1 = {75, 97, 47, 61, 53};
    std::vector<int> t2 = {61, 13, 29};

    RULERESULT res1 = check_order(t1, r1, 0, 1);

    ASSERT_EQ(res1, RULERESULT::FAIL);

    RULERESULT res2 = check_order(t2, r2, 1, 2);

    ASSERT_EQ(res2, RULERESULT::FAIL);

}

TEST(TESTDAY5, test1){
    std::vector<Rule> rules;

    rules.push_back(Rule(47, 53));
    rules.push_back(Rule(97, 13));
    rules.push_back(Rule(97, 61));
    rules.push_back(Rule(97, 47));
    rules.push_back(Rule(75, 29));
    rules.push_back(Rule(61, 13));
    rules.push_back(Rule(75, 53));
    rules.push_back(Rule(29, 13));
    rules.push_back(Rule(97, 29));
    rules.push_back(Rule(53, 29));
    rules.push_back(Rule(61, 53));
    rules.push_back(Rule(97, 53));
    rules.push_back(Rule(61, 29));
    rules.push_back(Rule(47, 13));
    rules.push_back(Rule(75, 47));
    rules.push_back(Rule(97, 75));
    rules.push_back(Rule(47, 61));
    rules.push_back(Rule(75, 61));
    rules.push_back(Rule(47, 29));
    rules.push_back(Rule(75, 13));
    rules.push_back(Rule(53, 13));

    std::vector<std::vector<int>> printing_orders;
    printing_orders.push_back({75, 47, 61, 53, 29});
    printing_orders.push_back({97, 61, 53, 29, 13});
    printing_orders.push_back({75, 29, 13});
    printing_orders.push_back({75, 97, 47, 61, 53});
    printing_orders.push_back({61, 13, 29});
    printing_orders.push_back({97, 13, 75, 29, 47});

    int sum = count_order_middle(printing_orders, rules);

    ASSERT_EQ(sum, 143);
}

TEST(TESTDAY5, test2){
    std::vector<Rule> rules;

    rules.push_back(Rule(47, 53));
    rules.push_back(Rule(97, 13));
    rules.push_back(Rule(97, 61));
    rules.push_back(Rule(97, 47));
    rules.push_back(Rule(75, 29));
    rules.push_back(Rule(61, 13));
    rules.push_back(Rule(75, 53));
    rules.push_back(Rule(29, 13));
    rules.push_back(Rule(97, 29));
    rules.push_back(Rule(53, 29));
    rules.push_back(Rule(61, 53));
    rules.push_back(Rule(97, 53));
    rules.push_back(Rule(61, 29));
    rules.push_back(Rule(47, 13));
    rules.push_back(Rule(75, 47));
    rules.push_back(Rule(97, 75));
    rules.push_back(Rule(47, 61));
    rules.push_back(Rule(75, 61));
    rules.push_back(Rule(47, 29));
    rules.push_back(Rule(75, 13));
    rules.push_back(Rule(53, 13));

    std::vector<std::vector<int>> printing_orders;
    printing_orders.push_back({75, 47, 61, 53, 29});
    printing_orders.push_back({97, 61, 53, 29, 13});
    printing_orders.push_back({75, 29, 13});
    printing_orders.push_back({75, 97, 47, 61, 53});
    printing_orders.push_back({61, 13, 29});
    printing_orders.push_back({97, 13, 75, 29, 47});

    int sum = correct_orders(printing_orders, rules);

    ASSERT_EQ(sum, 123);
}