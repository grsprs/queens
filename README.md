# N-Queens Problem Solver

[![Go Version](https://img.shields.io/github/go-mod/go-version/grsprs/queens)](https://go.dev/)
[![License](https://img.shields.io/github/license/grsprs/queens)](LICENSE)
[![CI](https://github.com/grsprs/queens/workflows/Go/badge.svg)](https://github.com/grsprs/queens/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/grsprs/queens)](https://goreportcard.com/report/github.com/grsprs/queens)

Author: Spiros Nikoloudakis (@grsprs)

A backtracking algorithm implementation in Go that solves the N-Queens problem.

## Problem

Place N queens on an N×N chessboard such that no two queens attack each other. Queens attack along rows, columns, and diagonals.

## Solution

This implementation uses backtracking with constraint propagation:
- Places one queen per row
- Tracks occupied columns and diagonals using hash maps
- Prunes invalid branches early
- Finds all distinct valid arrangements

## Usage

```bash
go run queens.go
```

Input N when prompted. The program outputs all solutions and the total count.

### Example

Input:
```
4
```

Output:
```
Solution 1:
.Q..
...Q
Q...
..Q.

Solution 2:
..Q.
Q...
...Q
.Q..

Total solutions: 2
```

## Build

```bash
go build
```

## Test

```bash
go test -v
```

All tests validate correctness against known solution counts and verify no queens attack each other.

## Performance

- N=8: 92 solutions in <100ms
- N=12: completes within seconds
- Constraint: 1 ≤ N ≤ 12

## Implementation

- `queens.go`: Main solver with backtracking algorithm
- `queens_test.go`: Comprehensive test suite

## Repository

https://github.com/grsprs/queens

## License

MIT License - see LICENSE file
