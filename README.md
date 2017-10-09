# subsets

Fast iteration of subsets, at the expense of a rather large amount of memory.

_Q: What's this for?_

A: I wanted to see if a lookup table would be fast than the traditional way of calculating subsets. CPUs have become faster at a greater rate than RAM.

_Q: Is it?_

A: Yes. Although unless your application does nothing except subset iteration, you won't notice it.

And for Go1.9 there is https://github.com/golang/go/issues/18616 ....