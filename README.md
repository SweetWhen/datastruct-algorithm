data struct and algorithm in golang

branch rand-qsort use time package to assess our Qsort performance.
It maybe in-correct because of the process swithing, we need to use process clock to assess, but i cann't find such api in golang.(c is clock())

root@ubuntu01:mysort# go run qsort.go 
100000 numbers Use 22.622869ms 
1000000 numbers Use 182.775282ms 


