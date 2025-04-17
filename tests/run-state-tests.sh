#!/usr/bin/env bash
go test -run "(TestState|TestExecutionSpecState|TestTransaction|TestRLP)" -v -short >state_test.log

cat state_test.log | grep FAIL >state_test_fail.log

TEST_PASS_COUNT=$(cat state_test.log | grep "PASS" | wc -l)
TEST_FAIL_COUNT=$(cat state_test_fail.log | wc -l)

if [ $TEST_FAIL_COUNT -ne 0 ]; then
    cat state_test_fail.log
fi

echo "Passed test cases: $TEST_PASS_COUNT"
echo "Failed test cases: $TEST_FAIL_COUNT"

if [ $TEST_FAIL_COUNT -ne 0 ]; then
    exit 1
fi
