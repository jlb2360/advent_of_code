cmake_minimum_required(VERSION 3.15)
project(PackageTest CXX)

find_package(GTest REQUIRED CONFIG)

add_executable(test_hello test_helloworld.cpp)
target_link_libraries(test_hello GTest::gtest GTest::gtest_main)
