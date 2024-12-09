#pragma once

#include <vector>
#include <deque>

bool check_levels(const std::vector<int> &);
bool check_levels(std::deque<int> &);
bool check_levels(std::deque<int>, int past, int modif, int &dampened);