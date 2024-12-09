#include "day2.h"
#include "gtest/gtest.h"
#include <vector>
#include <map>
#include <string>

namespace {
    TEST(Day2Tests, Test1){
        std::multimap<std::string, std::vector<int>> tests;

        tests.insert({"safe", std::vector<int> {7, 6, 4, 2, 1}});
        tests.insert({"unsafe", std::vector<int> {1, 2, 7, 8, 9}});
        tests.insert({"unsafe", std::vector<int> {9, 7, 6, 2, 1}});
        tests.insert({"unsafe", std::vector<int> {1, 3, 2, 4, 5}});
        tests.insert({"unsafe", std::vector<int> {8, 6, 4, 4, 1}});
        tests.insert({"safe", std::vector<int> {1, 3, 6, 7, 9}});

        int i = 0;
        for (const auto& pair : tests ){
            std::cout << "Test " << i << "\n";
            bool res = check_levels(pair.second);

            if (pair.first=="safe")
            {
                ASSERT_EQ(res, true);
            } else {
                ASSERT_EQ(res, false);
            }
            i += 1;
        }
    }
}