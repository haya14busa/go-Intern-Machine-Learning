# intern.ml

go-Intern-Machine-Learning is the template for the task of "[Web開発におけるコンピュータサイエンス - 機械学習編](http://developer.hatenastaff.com/entry/hatena-textbook-machine-learning-01-2016)" from Hatena Co., Ltd  written in Go.

## Setup

```
$ make setup
```

## Test: F value 

```
$ go test ./ml/evaluation
```

## Run classification of Iris command

```
$ ./script/run-iris
# or
$ go build ./cmd/iris
$ ./iris
```

## Plot learning curve

### Execution

```
$ go build ./cmd/iris
$ seq 100 | xargs -i ./iris -training=./data/iris.training.data -test=./data/iris.test.data -i={}
```

### Graph

```
$ seq 100 | xargs -i ./iris -training=./data/iris.training.data -test=./data/iris.test.data -i={} | script/plot
```

It needs [gnuplot](http://www.gnuplot.info/) to use `script/plot`.

## Author

[haya14busa](https://github.com/haya14busa)

## Credit

Copyright 2016 [はてな教科書](https://github.com/hatena/Hatena-Textbook) by [はてな](http://www.hatena.ne.jp/)

- [hatena/scala-Intern-Machine-Learning](https://github.com/hatena/scala-Intern-Machine-Learning)
- [hatena/perl-Intern-Machine-Learning](https://github.com/hatena/perl-Intern-Machine-Learning)
