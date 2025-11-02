# Advent of Code 2024 — Day 3

The following regular expression does most of the work in my solution:

```re
(?P<mul>mul\((?P<left>\d{1,3}),(?P<right>\d{1,3})\))|(?P<do>do\(\))|(?P<dont>don't\(\))
```

It composes three main capture groups that are `OR`ed together:
1. Multiplication
2. Do
3. Don't

I use this regular expression to construct a sequence of `mul`, `do`, and `don't` commands that I then iterate over to calculate the final answer.

## Multiplication

We are interested in `mul(X,Y)` commands, where `X` and `Y` are each 1-3 digit numbers.

```re
(?P<mul>mul\((?P<left>\d{1,3}),(?P<right>\d{1,3})\))
```

I am using named capture groups to make it easier to pull out the left and right hand sides of the command.

## Do

We are interested in `do()` commands. This is just a string literal.

```re
(?P<do>do\(\))
```

## Don't

We are interested in `don't()` commands. This is just a string literal.

```re
(?P<dont>don't\(\))
```