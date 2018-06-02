# pbf - Read protocol buffers using your own implementation
[![GoDoc](https://img.shields.io/badge/api-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/murphy214/pbf)

This project was designed to serialize protocol buffers structures from their type primitives. There are quite a few performance advantages to this including reduced allocation cost say if you protocol buffer structure is immediately tranformed into another data struture that whole protocol buffer allocation is sort of a waste. 

Their our some functions in here that don't quite fit the scope of the project which I need to change but other than that, the functions for primitives (at least the ones I use) were labored over quite meticulously in regards to performance. If you would like to see an example of a reader I've implemented myself there is one [here](https://github.com/murphy214/geobuf/blob/master/geobuf_raw/read_feature.go). 

As you can see whats you understand the patters of pbfs its pretty straight forward, but more importantly if you have a really large structure of which 1000s need to be read, you can peak in and grab just the context or information you need in a custom read function. 
