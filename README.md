# Golang DataStructure

[![N|Solid](https://unit42.paloaltonetworks.com/wp-content/uploads/2019/07/golang-hacker.jpg)]

Efort was made to keep the data structures as Generic as possible with my limited knowledge of golang.
This repository was created for me to understand golang and whats better way to start with DataStructures and algorithms isn't!!

There will be many ongoing changes to structure outline as my knowledge in this field will grow!

# Array
An array data structure, or simply array, is a data structure consisting of a collections of elements
(values or variables), each identified by at least one array index or key. It is the simplest type of
linear data structure.

| Operations | Big O |
| ---------- | ----- |
| lookup     | O(1)  |
| push       | O(1)  |
| insert     | O(n)  |
| delete     | O(n)  |

> Try to implement stack, queue .. etc through array

## Type of Arrays
> [Array](https://en.wikipedia.org/wiki/Array_data_structure)
> [Bit Array](https://en.wikipedia.org/wiki/Array_data_structure)
> [Bit field](https://en.wikipedia.org/wiki/Array_data_structure)
> [BitBoard](https://en.wikipedia.org/wiki/Array_data_structure)
> [Bitmap](https://en.wikipedia.org/wiki/Array_data_structure)
> [Circular buffer](https://en.wikipedia.org/wiki/Array_data_structure)
> [Control table](https://en.wikipedia.org/wiki/Array_data_structure)
> [Dope vector](https://en.wikipedia.org/wiki/Array_data_structure)
> [Dynamic Array](https://en.wikipedia.org/wiki/Dynamic_array)
> [Gap Array](https://en.wikipedia.org/wiki/Dynamic_array)
> [Hashed Array tree](https://en.wikipedia.org/wiki/Hashed_array_tree)
> [Lookup table](https://en.wikipedia.org/wiki/Lookup_table)
> [Matrix](https://en.wikipedia.org/wiki/Matrix_(computer_science))
> [Parallel Array](https://en.wikipedia.org/wiki/Parallel_array)

# LinkList
A linklist is a linear collection of data elements, whose order is not given by their physical placement in memory. Instead each element points to the next. It is a data structure consisting of a collection of nodes which together represents a sequence.

| Operations | Big O |
| ---------- | ----- |
| Access   | O(n)  |
| Insert/Delete at End | O(1) when last element is known O(n) When last element is not known |
| insert/delete at Begining     | O(1)  |
| search | O(n) |