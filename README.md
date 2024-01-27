
# Doubly Linked List Implementation

## Overview

This repository contains a basic implementation of a doubly linked list in [language]. It provides core functionality for creating, manipulating, and iterating over a doubly linked list data structure.

## Key Features

Doubly Linked List Structure:

* Nodes with forward and backward pointers.
* Efficient insertion and deletion at both ends.

Basic Operations:

* Create a new list: NewList(elements ...interface{}) *List
* Add elements:
  * Unshift(v interface{}) to add at the beginning.
  * Push(v interface{}) to add at the end.
* Remove elements:
  * Shift() (interface{}, error) to remove from the beginning.
  * Pop() (interface{}, error) to remove from the end.
* Access elements:
  * First() *Node to get the first node.
  * Last() *Node to get the last node.
* Iterate over the list using the Next() method of node

## Running Tests

```go test -v --bench . --benchmem```
