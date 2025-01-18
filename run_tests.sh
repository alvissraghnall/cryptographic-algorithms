#!/bin/bash

# Initialize results
test_results=""

# Find all directories with `go.mod` and test them
for dir in $(find . -name "go.mod" -exec dirname {} \;); do
  echo "Running tests in $dir"
  pushd $dir
  if ls *_test.go > /dev/null 2>&1; then
    # Run tests and capture output
    test_output=$(go test -v)
    echo "$test_output"
    # Collect results
    test_passed=$(echo "$test_output" | grep -c "PASS")
    test_failed=$(echo "$test_output" | grep -c "FAIL")
    test_results="$test_results\n$dir: Passed=$test_passed, Failed=$test_failed"
  else
    echo "No tests found in $dir"
  fi
  popd
done

# Display test results
echo "Test Results:"
echo "$test_results"
