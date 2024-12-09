#include "day3.h"
#include "gtest/gtest.h"
#include <string>
#include <vector>

TEST(DAY3TEST, test1){
    std::string line = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))";

    std::vector<Token*> toks; 

    int idx = 0;

    tokenize_string(line, toks, idx);

    int modif = 1;

    int res = process_tokens(toks, modif);

    ASSERT_EQ(res, 161);
    

}

TEST(DAY3TEST, test2){
    std::string line = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))";

    std::vector<Token*> toks; 

    int idx = 0;

    tokenize_string(line, toks, idx);

    int modif = 1;

    int res = process_tokens(toks, modif);

    ASSERT_EQ(res, 48);
    

}