# perl-Httpclient-vs-golang-Httpclient
use Anyevent::Http VS use Golang goroutine, both 1000 concurrency request 50000 connections to nginx server


##benchmark
```
server:nginx  client : aeperl  600 concurrency 50000 connection 26s 50000 success

bash-4.1$ time perl aehttp.pl >tb.log

real    0m26.217s
user    0m25.289s
sys     0m0.911s

server:nginx  client : golang  600 concurrency 50000 connection 16s 50000 success

-bash-4.1$  time go run httptest.go >tb.log

real    0m16.861s
user    0m14.042s
sys     0m3.159s

-bash-4.1$ cat tb.log |grep " 200 |wc -l
50000

```

##AUTHOR
zhangqi Jason asd1986_n@126.com
