# gobloom

Bloom Filter and Counting Bloom Filter implement in Go,

inspired by willf's implementation of Bloom Filter, see https://github.com/willf/bloom


## Getting Started

#### Bool Bloom Filter

~~~ go
package main

import (
	"fmt"
	"github.com/yusong/gobloom"
)

func main() {
	bf := gobloom.NewBloomFilter(1<<24, 3)
	bf.Add([]byte("hello gobloom"))
	if bf.Contains([]byte("hello gobloom")) {
		fmt.Println("\"hello gobloom\" is in bf")
	}
}
~~~

#### Counting Bloom Filter

~~~ go
package main

import (
	"fmt"
	"github.com/yusong/gobloom"
)

func main() {
	bf := gobloom.NewCountingBloomFilter(1<<24, 3)
	bf.Add([]byte("hello gobloom"))
	if bf.Contains([]byte("hello gobloom")) {
		fmt.Println("\"hello gobloom\" is in bf")
	}
	bf.Remove([]byte("hello gobloom"))
	if bf.Contains([]byte("hello gobloom")) == false {
		fmt.Println("\"hello gobloom\" is not in bf")
	}
}
~~~

## Resources

Bloom Filter: http://en.wikipedia.org/wiki/Bloom_filter#Counting_filters

FNV Hash Function: http://en.wikipedia.org/wiki/Fowler%E2%80%93Noll%E2%80%93Vo_hash_function

## License

gobloom is released under the BSD License. See LICENSE for more information.
