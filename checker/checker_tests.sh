#!/bin/bash

# Define the path to the checker executable
CHECKER="./checker"

# Helper function to run test cases
run_test() {
    input="$1"
    instructions="$2"
    expected="$3"
    
    echo -e "$instructions" | $CHECKER "$input" > result.txt 2>&1
    result=$(cat result.txt)

    if [ "$result" == "$expected" ]; then
        echo "Test passed for input '$input' with instructions '$instructions'"
    else
        echo "Test failed for input '$input' with instructions '$instructions'"
        echo "Expected: $expected, but got: $result"
    fi
    rm -f result.txt
}

# Test cases

# 1. Simple sorted input with no instructions
run_test "1 2 3 4 5" "" "OK"

# 2. Simple unsorted input with no instructions
run_test "5 4 3 2 1" "" "KO"

# 3. Simple unsorted input, sorted with instructions
run_test "3 2 1" "sa\nra\n" "OK"

# 4. Already sorted input with instructions that keep it sorted
run_test "1 2 3" "ra\nrb\nrr\n" "OK"

# 5. Unsorted input that cannot be sorted with given instructions
run_test "3 2 1" "pb\npb\nra\n" "KO"

# 6. Test with duplicates (should return Error)
run_test "3 2 2 1" "" "Error"

# 7. Invalid argument (non-integer, should return Error)
run_test "3 2 one 1" "" "Error"

# 8. Invalid instruction (should return Error)
run_test "3 2 1" "invalid\n" "Error"

# 9. Multiple valid moves resulting in a sorted stack
run_test "4 1 3 2" "pb\npb\nra\nsa\npa\npa\n" "OK"

# 10. Large input test with already sorted numbers
run_test "1 2 3 4 5 6 7 8 9 10" "" "OK"

# 11. Large input test with random unsorted numbers
run_test "10 9 8 7 6 5 4 3 2 1" "ra\nra\nra\nra\nra\nra\nra\nra\nra\n" "KO"

# Run without arguments (should output nothing)
echo -n "Test no arguments: "
$CHECKER > /dev/null 2>&1
if [ $? -eq 0 ]; then
    echo "Passed (no output as expected)"
else
    echo "Failed (unexpected output)"
fi
