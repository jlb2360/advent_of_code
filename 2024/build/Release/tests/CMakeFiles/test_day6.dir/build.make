# CMAKE generated file: DO NOT EDIT!
# Generated by "Unix Makefiles" Generator, CMake Version 3.22

# Delete rule output on recipe failure.
.DELETE_ON_ERROR:

#=============================================================================
# Special targets provided by cmake.

# Disable implicit rules so canonical targets will work.
.SUFFIXES:

# Disable VCS-based implicit rules.
% : %,v

# Disable VCS-based implicit rules.
% : RCS/%

# Disable VCS-based implicit rules.
% : RCS/%,v

# Disable VCS-based implicit rules.
% : SCCS/s.%

# Disable VCS-based implicit rules.
% : s.%

.SUFFIXES: .hpux_make_needs_suffix_list

# Command-line flag to silence nested $(MAKE).
$(VERBOSE)MAKESILENT = -s

#Suppress display of executed commands.
$(VERBOSE).SILENT:

# A target that is always out of date.
cmake_force:
.PHONY : cmake_force

#=============================================================================
# Set environment variables for the build.

# The shell in which to execute make rules.
SHELL = /bin/sh

# The CMake executable.
CMAKE_COMMAND = /usr/bin/cmake

# The command to remove a file.
RM = /usr/bin/cmake -E rm -f

# Escaping for special characters.
EQUALS = =

# The top-level source directory on which CMake was run.
CMAKE_SOURCE_DIR = /home/jlb1694/dev/advent_of_code/2024

# The top-level build directory on which CMake was run.
CMAKE_BINARY_DIR = /home/jlb1694/dev/advent_of_code/2024/build/Release

# Include any dependencies generated for this target.
include tests/CMakeFiles/test_day6.dir/depend.make
# Include any dependencies generated by the compiler for this target.
include tests/CMakeFiles/test_day6.dir/compiler_depend.make

# Include the progress variables for this target.
include tests/CMakeFiles/test_day6.dir/progress.make

# Include the compile flags for this target's objects.
include tests/CMakeFiles/test_day6.dir/flags.make

tests/CMakeFiles/test_day6.dir/testDay6.cpp.o: tests/CMakeFiles/test_day6.dir/flags.make
tests/CMakeFiles/test_day6.dir/testDay6.cpp.o: ../../tests/testDay6.cpp
tests/CMakeFiles/test_day6.dir/testDay6.cpp.o: tests/CMakeFiles/test_day6.dir/compiler_depend.ts
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/jlb1694/dev/advent_of_code/2024/build/Release/CMakeFiles --progress-num=$(CMAKE_PROGRESS_1) "Building CXX object tests/CMakeFiles/test_day6.dir/testDay6.cpp.o"
	cd /home/jlb1694/dev/advent_of_code/2024/build/Release/tests && /usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -MD -MT tests/CMakeFiles/test_day6.dir/testDay6.cpp.o -MF CMakeFiles/test_day6.dir/testDay6.cpp.o.d -o CMakeFiles/test_day6.dir/testDay6.cpp.o -c /home/jlb1694/dev/advent_of_code/2024/tests/testDay6.cpp

tests/CMakeFiles/test_day6.dir/testDay6.cpp.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing CXX source to CMakeFiles/test_day6.dir/testDay6.cpp.i"
	cd /home/jlb1694/dev/advent_of_code/2024/build/Release/tests && /usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -E /home/jlb1694/dev/advent_of_code/2024/tests/testDay6.cpp > CMakeFiles/test_day6.dir/testDay6.cpp.i

tests/CMakeFiles/test_day6.dir/testDay6.cpp.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling CXX source to assembly CMakeFiles/test_day6.dir/testDay6.cpp.s"
	cd /home/jlb1694/dev/advent_of_code/2024/build/Release/tests && /usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -S /home/jlb1694/dev/advent_of_code/2024/tests/testDay6.cpp -o CMakeFiles/test_day6.dir/testDay6.cpp.s

# Object files for target test_day6
test_day6_OBJECTS = \
"CMakeFiles/test_day6.dir/testDay6.cpp.o"

# External object files for target test_day6
test_day6_EXTERNAL_OBJECTS =

tests/test_day6: tests/CMakeFiles/test_day6.dir/testDay6.cpp.o
tests/test_day6: tests/CMakeFiles/test_day6.dir/build.make
tests/test_day6: lib/libadvent_lib.a
tests/test_day6: /home/jlb1694/.conan2/p/b/gtest2dac020849679/p/lib/libgtest_main.a
tests/test_day6: /home/jlb1694/.conan2/p/b/gtest2dac020849679/p/lib/libgtest.a
tests/test_day6: tests/CMakeFiles/test_day6.dir/link.txt
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --bold --progress-dir=/home/jlb1694/dev/advent_of_code/2024/build/Release/CMakeFiles --progress-num=$(CMAKE_PROGRESS_2) "Linking CXX executable test_day6"
	cd /home/jlb1694/dev/advent_of_code/2024/build/Release/tests && $(CMAKE_COMMAND) -E cmake_link_script CMakeFiles/test_day6.dir/link.txt --verbose=$(VERBOSE)

# Rule to build all files generated by this target.
tests/CMakeFiles/test_day6.dir/build: tests/test_day6
.PHONY : tests/CMakeFiles/test_day6.dir/build

tests/CMakeFiles/test_day6.dir/clean:
	cd /home/jlb1694/dev/advent_of_code/2024/build/Release/tests && $(CMAKE_COMMAND) -P CMakeFiles/test_day6.dir/cmake_clean.cmake
.PHONY : tests/CMakeFiles/test_day6.dir/clean

tests/CMakeFiles/test_day6.dir/depend:
	cd /home/jlb1694/dev/advent_of_code/2024/build/Release && $(CMAKE_COMMAND) -E cmake_depends "Unix Makefiles" /home/jlb1694/dev/advent_of_code/2024 /home/jlb1694/dev/advent_of_code/2024/tests /home/jlb1694/dev/advent_of_code/2024/build/Release /home/jlb1694/dev/advent_of_code/2024/build/Release/tests /home/jlb1694/dev/advent_of_code/2024/build/Release/tests/CMakeFiles/test_day6.dir/DependInfo.cmake --color=$(COLOR)
.PHONY : tests/CMakeFiles/test_day6.dir/depend

