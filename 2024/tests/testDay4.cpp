#include "day4.h"
#include "gtest/gtest.h"
#include <string>
#include <vector>


TEST(TESTDAY4, test1){
    std::vector<std::string> lines;

    lines.push_back("MMMSXXMASM");
    lines.push_back("MSAMXMSMSA");
    lines.push_back("AMXSXMAAMM");
    lines.push_back("MSAMASMSMX");
    lines.push_back("XMASAMXAMM");
    lines.push_back("XXAMMXXAMA");
    lines.push_back("SMSMSASXSS");
    lines.push_back("SAXAMASAAA");
    lines.push_back("MAMMMXMMMM");
    lines.push_back("MXMXAXMASX");

    int sum=0;

    count_xmas(sum, lines);

    ASSERT_EQ(sum, 18);
}

TEST(TESTDAY4, test2){
    std::vector<std::string> lines;

    lines.push_back("MMMSXXMASM");
    lines.push_back("MSAMXMSMSA");
    lines.push_back("AMXSXMAAMM");
    lines.push_back("MSAMASMSMX");
    lines.push_back("XMASAMXAMM");
    lines.push_back("XXAMMXXAMA");
    lines.push_back("SMSMSASXSS");
    lines.push_back("SAXAMASAAA");
    lines.push_back("MAMMMXMMMM");
    lines.push_back("MXMXAXMASX");

    int sum=0;

    count_x_mas(sum, lines);

    ASSERT_EQ(sum, 9);
}