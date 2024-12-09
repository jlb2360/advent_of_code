cmake_minimum_required(VERSION 3.15)
project(PackageTest CXX)

find_package(GTest REQUIRED CONFIG)

add_executable(test_hello test_helloworld.cpp)
target_link_libraries(test_hello GTest::gtest GTest::gtest_main)

add_executable(test_day1 testDay1.cpp)
target_link_libraries(test_day1 PRIVATE GTest::gtest GTest::gtest_main ${LIBRARY_NAME})

add_executable(test_day2 testDay2.cpp)
target_link_libraries(test_day2 PRIVATE GTest::gtest GTest::gtest_main ${LIBRARY_NAME})

add_executable(test_day3 testDay3.cpp)
target_link_libraries(test_day3 PRIVATE GTest::gtest GTest::gtest_main ${LIBRARY_NAME})


add_executable(test_day4 testDay4.cpp)
target_link_libraries(test_day4 PRIVATE GTest::gtest GTest::gtest_main ${LIBRARY_NAME})

add_executable(test_day5 testDay5.cpp)
target_link_libraries(test_day5 PRIVATE GTest::gtest GTest::gtest_main ${LIBRARY_NAME})

add_executable(test_day6 testDay6.cpp)
target_link_libraries(test_day6 PRIVATE GTest::gtest GTest::gtest_main ${LIBRARY_NAME})