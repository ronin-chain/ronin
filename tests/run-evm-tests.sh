#!/usr/bin/env bash
rm -f evm_test.log
touch evm_test.log

go test . -test.v -short >evm_test.log

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
