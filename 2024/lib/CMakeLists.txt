set(LIBRARY_SRC
    day1.cpp
    day2.cpp
    day3.cpp
    day4.cpp
    day5.cpp
    day6.cpp
)

set(LIBRARY_INCLUDES
    "./"
)

add_library(${LIBRARY_NAME} STATIC
    ${LIBRARY_SRC}
    ${LIBRARY_HEADERS}
)

target_include_directories(${LIBRARY_NAME} PUBLIC
    ${LIBRARY_INCLUDES}
    ${CMAKE_SOURCE_DIR}/include
)