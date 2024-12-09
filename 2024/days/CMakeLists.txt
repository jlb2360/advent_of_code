set(SOURCE_EXE hello.cpp)

add_executable(hello ${SOURCE_EXE})
add_executable(day1 day1_runner.cpp)
add_executable(day2 day2_runner.cpp)
add_executable(day3 day3_runner.cpp)

add_executable(day4 day4_runner.cpp)

target_include_directories(day1 PRIVATE ${CMAKE_SOURCE_DIR}/include)
target_include_directories(day2 PRIVATE ${CMAKE_SOURCE_DIR}/include)
target_include_directories(day3 PRIVATE ${CMAKE_SOURCE_DIR}/include)
target_include_directories(day4 PRIVATE ${CMAKE_SOURCE_DIR}/include)

target_link_libraries(day1 PRIVATE
    ${LIBRARY_NAME}
)

target_link_libraries(day2 PRIVATE
    ${LIBRARY_NAME}
)

target_link_libraries(day3 PRIVATE
    ${LIBRARY_NAME}
)

target_link_libraries(day4 PRIVATE
    ${LIBRARY_NAME}
)

add_executable(day5 day5_runner.cpp)
target_include_directories(day5 PRIVATE ${CMAKE_SOURCE_DIR}/include)

target_link_libraries(day5 PRIVATE
    ${LIBRARY_NAME}
)

add_executable(day6 day6_runner.cpp)
target_include_directories(day6 PRIVATE ${CMAKE_SOURCE_DIR}/include)

target_link_libraries(day6 PRIVATE
    ${LIBRARY_NAME}
)