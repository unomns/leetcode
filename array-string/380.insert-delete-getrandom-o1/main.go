package main

import (
	"fmt"
	"math/rand"
)

/**
Implement the RandomizedSet class:
	- RandomizedSet() Initializes the RandomizedSet object.
	- bool insert(int val) Inserts an item val into the set if not present. Returns true if the item was not present, false otherwise.
	- bool remove(int val) Removes an item val from the set if present. Returns true if the item was present, false otherwise.
	- int getRandom() Returns a random element from the current set of elements (it's guaranteed that at least one element exists when this method is called).
	Each element must have the same probability of being returned.

You must implement the functions of the class such that each function works in average O(1) time complexity.

Example 1:
	Input
	["RandomizedSet", "insert", "remove", "insert", "getRandom", "remove", "insert", "getRandom"]
	[[], [1], [2], [2], [], [1], [2], []]
	Output
	[null, true, false, true, 2, true, false, 2]

	Explanation
		RandomizedSet randomizedSet = new RandomizedSet();
		randomizedSet.insert(1); // Inserts 1 to the set. Returns true as 1 was inserted successfully.
		randomizedSet.remove(2); // Returns false as 2 does not exist in the set.
		randomizedSet.insert(2); // Inserts 2 to the set, returns true. Set now contains [1,2].
		randomizedSet.getRandom(); // getRandom() should return either 1 or 2 randomly.
		randomizedSet.remove(1); // Removes 1 from the set, returns true. Set now contains [2].
		randomizedSet.insert(2); // 2 was already in the set, so return false.
		randomizedSet.getRandom(); // Since 2 is the only number in the set, getRandom() will always return 2.


Constraints:
	-2^31 <= val <= 2^31 - 1
	At most 2 * 10^5 calls will be made to insert, remove, and getRandom.
	There will be at least one element in the data structure when getRandom is called.
*/

type RandomizedSet struct {
	set    map[int]int
	values []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		set: make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.set[val]; ok {
		return false
	}
	this.values = append(this.values, val)
	this.set[val] = len(this.values) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.set[val]; !ok {
		return false
	}

	idx, last := this.set[val], len(this.values)-1
	this.set[this.values[last]] = idx
	this.values[idx], this.values[last] = this.values[last], this.values[idx]
	this.values = this.values[:last]

	delete(this.set, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.values[rand.Intn(len(this.values))]
}

func main() {
	randomizedSet := Constructor()
	fmt.Println(randomizedSet.Insert(1))   // true
	fmt.Println(randomizedSet.Remove(2))   // false
	fmt.Println(randomizedSet.Insert(2))   // true
	fmt.Println(randomizedSet.GetRandom()) // N
	fmt.Println(randomizedSet.Remove(1))   // true
	fmt.Println(randomizedSet.Insert(2))   // false
	fmt.Println(randomizedSet.GetRandom()) // 2
}
