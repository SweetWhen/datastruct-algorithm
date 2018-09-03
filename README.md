data struct and algorithm in golang

test-qsort branch use testing to test our MyQsort function
TestQsortFunction -> test is our function work fine:
  root@ubuntu01:mysort# go test
  init func print...
  sort result: [1 2 3 5 6 10 12 24 32 33 55 60 78]
  PASS
  ok  	_/home/agan/data-struct/struct-algorit/mysort	3.585s


generator rand number slice a and b in init func, both a and has diff number
then use TestQsortPerfomance1 and TestQsortPerfomance2 test our func
We can see that when numbers is 100000 (slice a), consume 0.02s
when numbers is 1000000 (slice b), consume 1.78s:

 root@ubuntu01:mysort# go test -v -run="Perfomance"
 init func print...
 === RUN   TestQsortPerfomance1
 --- PASS: TestQsortPerfomance1 (0.02s)
	qsort_test.go:32: 0~100000 numbers qsor..
 === RUN   TestQsortPerfomance2
 --- PASS: TestQsortPerfomance2 (1.78s)
	qsort_test.go:37: 0~1000000 numbers qsor..
 PASS
 ok  	_/home/agan/data-struct/struct-algorit/mysort	3.171s




