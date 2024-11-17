cmake_minimum_required(VERSION 3.15)
project(Advent2024 CXX)

set(CXX_STANDARD 17)
set(CXX_STANDARD_REQUIRED ON)
set(CXX_EXTENSIONS OFF)

add_subdirectory(lib)
add_subdirectory(days)


# Add the tests subdirectory (for running tests)
add_subdirectory(tests)

