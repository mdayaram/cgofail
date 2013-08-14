#!/bin/sh

# Test 1A
#echo "Test 1A"
#time ./cgofail -cgo -c=100 -j=1 -w=1 -t=10
#echo "\n\n"
# Test 2A
#echo "Test 2A"
#time ./cgofail -cgo -c=100 -j=10 -w=1 -t=10
#echo "\n\n"
# Test 3A
#echo "Test 3A"
#time ./cgofail -cgo -c=100 -j=100 -w=1 -t=10
#echo "\n\n"
# Test 4A
#echo "Test 4A"
#time ./cgofail -cgo -c=100 -j=1000 -w=1 -t=10
#echo "\n\n"
# Test 1B
echo "Test 1B"
time ./cgofail -cgo -c=100 -j=1 -w=100 -t=4
echo "\n\n"
# Test 2B
echo "Test 2B"
time ./cgofail -cgo -c=100 -j=10 -w=100 -t=4
echo "\n\n"
# Test 3B
echo "Test 3B"
time ./cgofail -cgo -c=100 -j=100 -w=100 -t=4
echo "\n\n"
# Test 4B
echo "Test 4B"
time ./cgofail -cgo -c=100 -j=1000 -w=100 -t=4
echo "\n\n"



