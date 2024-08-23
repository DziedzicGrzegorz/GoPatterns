# Variables
CXX = clang++
CXXFLAGS = -std=c++17 -O3 -Wall -Wextra -fPIC
LDFLAGS = -shared
SOURCES = internal/Leetcode/592/592.cpp
TARGET = internal/Leetcode/592/lib592.so

# Default target
all: $(TARGET)

# Rule to build the shared library
$(TARGET): $(SOURCES)
	$(CXX) $(CXXFLAGS) $(LDFLAGS) -o $@ $^

# Clean up build artifacts
clean:
	rm -f $(TARGET)
