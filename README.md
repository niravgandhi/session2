To test if your code works as expected, run the test_script.sh shell script. 

On bash, to add a total of 10 tasks with a QPS of 1 second, run the following command :-

for (( n = 1 ; n < 10 ; n++ )) ; do ./test_script.sh ; sleep 1 ; done
