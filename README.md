# Game of Life

## Problem Statement

It is a one time initialization game, where player initializes live cells at the start and waits for next generation.
The next generation depends on following conditions

* Any live cell with fewer than two live neighbours dies, as if by underpopulation.
* Any live cell with two or three live neighbours lives on to the next generation.
* Any live cell with more than three live neighbours dies, as if by overpopulation.
* Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.

## Environment

* go version 1.19

## Test Instructions

To execute test cases, open terminal in lib directory and run below command

    $ go test

To execute test cases and also see verbose, run below command

    $ go test -v