var reduce = fn(arr, initial, f) {
	var iter = fn(arr, result) {
		if (len(arr) == 0) {
			result
		} else {
			iter(rest(arr), f(result, first(arr)));
		}
	};

	iter(arr, initial);
};

var sum = fn(arr) {
	reduce(arr, 0, fn(initial, el) { initial + el });
};

puts(sum([1, 2, 3, 4, 5]));