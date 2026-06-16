# go-dsa

My personal collection of Data Structures & Algorithms problems, solved in Go.

## Purpose

This repo is my DSA practice journal — a single place to log questions, solutions, and notes as I prepare for technical interviews. The goal is to grow this into a comprehensive reference that I can look back on when an interview comes up.

## Topics Covered

| Topic | Problems |
|---|---|
| Two Pointers | Longest Palindromic Substring |
| Dynamic Programming (1D) | House Robber |

## Structure

```
go-dsa/
├── two-pointers/       # Two-pointer technique problems
│   └── longest-palindromic-substr.go
├── dp1d/               # 1D Dynamic Programming problems
│   └── robber2.go
└── main.go             # Quick entry point to run solutions
```

## How I Use This

- **By topic** — when an interview focuses on a specific pattern, I review the relevant folder.
- **By problem name** — if I've seen a question before and want to refresh my approach.
- **As a growth tracker** — the more problems I add, the more confident I feel walking into interviews.

## Running Solutions

```bash
go run main.go
```

Or run individual files:

```bash
go run two-pointers/longest-palindromic-substr.go
```
