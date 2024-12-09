
#include "gtest/gtest.h"
#include "day1.h"
#include <deque>

namespace {
    TEST(Day1Test, CompareList1) {
        std::deque<int> lis1;
        lis1.push_back(3);
        lis1.push_back(4);
        lis1.push_back(2);
        lis1.push_back(1);
        lis1.push_back(3);
        lis1.push_back(3);
        
        
        std::deque<int> lis2;
        lis2.push_back(4);
        lis2.push_back(3);
        lis2.push_back(5);
        lis2.push_back(3);
        lis2.push_back(9);
        lis2.push_back(3);

        int res = check_lists(lis1, lis2);
        ASSERT_EQ(res, 11);

        res = sim_score(lis1, lis2);

        ASSERT_EQ(res, 31);
        
    }
}