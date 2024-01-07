# Golang Exercises

Data structures

- [x] *LRU Cache* — Least Recently Used (LRU) cache.
- [x] *Matrix Operations* — Matrix operations like addition, subtraction, multiplication, and transpose.
- [ ] *Lists / collections*
  - [ ] *Linked List* —  Singly linked list with methods to add an element, delete an element, and display the list.
  - [ ] *Stack Using Queues* — Implement a stack using queues.
  - [ ] *Queue Using Stacks* — Implement a queue using stacks.
  - [ ] *Circular Queue:* Implement a circular queue using an array.
  - [ ] *Priority Queue:* Implement a priority queue.
  - [x] *Hash table*
- [ ] *Trees*
  - [x] *Binary tree*
  - [x] *Binary search tree (BST)* — A binary tree where for each node, all elements in the left subtree are less than the node, and all elements in the right subtree are greater.
  - [ ] *AVL tree* — A self-balancing binary search tree, where the difference of heights of left and right subtrees cannot be more than one for all nodes.
  - [x] *B-Tree* — A self-balancing tree that maintain its height to be logarithmic of the number of entries, ensuring optimal performance.
  - [ ] *Binary Heap* — A special tree-based data structure that satisfies the heap property. If P is a parent node of C, then the key (the value) of P is either greater than or equal to (in a max heap) or less than or equal to (in a min heap) the key of C.
  - [ ] *Red-Black Tree* — Another self-balancing binary search tree, where each node stores an extra bit for denoting the color of the node, either red or black.
  - [ ] *Trie (Prefix Tree)* — A tree-like data structure that proves to be very efficient for solving problems related to strings. Each string is represented by a path from the root to the leaf.
  - [ ] *Suffix Tree* — A compressed trie containing all the suffixes of the given text as their keys and positions in the text as their values. It's a powerful data structure for text processing.
  - [ ] *Segment Tree* — A tree data structure for storing intervals, or segments. It allows querying which of the stored segments contain a given point.
  - [ ] *Fenwick Tree* — (Binary Indexed Tree): A data structure providing efficient methods for calculation and manipulation of the prefix sums of a table of values.


Databases

- [x] *Hash Tables* — Hash tables, also known as hash maps, are used for fast data retrieval. They are key-value stores that allow for O(1) average complexity for search, insert, and delete operations.
- [ ] *Trees* — Particularly, Binary Search Trees (BST), AVL Trees, and B-Trees. These are used in databases for indexing purposes to speed up data retrieval. B-Trees and AVL Trees are self-balancing trees that maintain their height to be logarithmic of the number of entries, ensuring optimal performance.
- [ ] *Disk I/O* — Paging, buffering, etc.
- [ ] *Concurrency control* — Databases often handle multiple concurrent requests, so understanding concepts like locks, deadlocks, and transactions is important.
- [ ] *Caching and buffering*


Compiling, parsing and interpreting

- [x] *Simple arithmetic evaluator* — Simple operations such as addition.
- [ ] *Simple interpreter* — Simple language with limited instructions like variable assignment and arithmetic.
- [ ] *Function interpreter* — Add functions to the simple language.
- [ ] *Byte code compiler* — Compile source into byte code that can be executed, include a simple VM to execute it.
- [ ] *Static type checker* — Checks for type errors in source code.


Networking

- [x] *HTTP Server* — Write a simple HTTP server from scratch at the socket layer.
  - [ ] *Routing*
  - [ ] *Body parsing*
  - [ ] *Error handling*
  - [ ] *Connection pooling*
  - [ ] *Static file serving*
- [ ] *Concurrent Web Crawler* — Write a concurrent web crawler using Go's goroutines and channels.

Algorithms

- [ ] *Depth-First Search and Breadth-First Search* — DFS and BFS algorithms on a graph data structure that you create.
- [ ] *Sorting Algorithms* — Such as quicksort, mergesort, and heapsort.
- [ ] *Hash Table* — Hash table with separate chaining.
- [ ] *Dijkstra's Algorithm* — Dijkstra's algorithm to find the shortest path in a graph.
- [ ] *Knapsack Problem* — Solve the 0/1 Knapsack problem using dynamic programming.
- [ ] *Longest Common Subsequence* — Solve the Longest Common Subsequence (LCS) problem using dynamic programming.
- [ ] *Merge Two Sorted* Lists: Merge two sorted linked lists into a new sorted list.
- [ ] *Reverse a Stack* — Reverse a stack using only push and pop operations.

Games

- [ ] *Chess*
- [x] *Tic tac toe*
