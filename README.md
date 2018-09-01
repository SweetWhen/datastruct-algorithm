data struct and algorithm in golang

branch rand-qsort use time package to assess our Qsort performance.
It maybe in-correct because of the process swithing, we need to use process clock to assess, but i cann't find such api in golang.(c is clock())

(10 0000 numbers)Use 7.156µs 
(100 0000 numbers)Use 6.896µs 

