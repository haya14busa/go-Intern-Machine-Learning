setup:
	mkdir -p data
	curl https://archive.ics.uci.edu/ml/machine-learning-databases/iris/iris.data | head -n -1 > data/iris.data
	shuf data/iris.data > data/iris.shuf.data
	head -n 100 data/iris.shuf.data > data/iris.training.data
	tail -n 50 data/iris.shuf.data > data/iris.test.data
