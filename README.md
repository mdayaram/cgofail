cgofail
=======
Benchmark test to see the speed differences between Cgo and native Go when trying to combine two strings.

TODO: comparison is not apples to apples because the C code needs to calculate the length of the string, whereas in go, the length of the string is saved as part of the datastructure.  Should redo the experiment where the lenght of the strings are passed in.
