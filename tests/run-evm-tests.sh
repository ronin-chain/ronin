#!/usr/bin/env bash
EXCLUDED_TESTS="TestState|TestExecutionSpecState|TestTransaction|TestRLP"
TESTS_TO_RUN=$(go test -list . | grep -vE "$EXCLUDED_TESTS" | grep -E "^Test")

rm -f evm_test.log
touch evm_test.log

for test in $TESTS_TO_RUN; do
    echo "Running test: $test"
    go test -run "$test" -v -short >>evm_test.log
done

cat evm_test.log | grep FAIL >evm_test_fail.log

TEST_PASS_COUNT=$(cat evm_test.log | grep "PASS" | wc -l)
TEST_FAIL_COUNT=$(cat evm_test_fail.log | wc -l)

if [ $TEST_FAIL_COUNT -ne 0 ]; then
    cat evm_test_fail.log
fi

echo "Passed test cases: $TEST_PASS_COUNT"
echo "Failed test cases: $TEST_FAIL_COUNT"

if [ $TEST_FAIL_COUNT -ne 0 ]; then
    exit 1
fi
