# Polynomial Problem
[Notion Link for input format](https://honorable-marimba-e75.notion.site/Hashira-Placements-Assignment-Online-16c880a946cf805f9ab6c3aacd180f8d)

I didn't attend the process where this problem was asked but I found the problem statement very interesting.

So here is my attempt to solve it with [Test Driven Development (TDD)](https://en.wikipedia.org/wiki/Test-driven_development) as I am also learning [Go with tests](https://github.com/RunAtTekky/Learn-GO-with-tests)

> Basically you are given system of linear equations. Solve it.

> I am solving this problem with [Gaussian Elimination Method](https://www.geeksforgeeks.org/dsa/gaussian-elimination/) which uses [Row Echelon Form](https://en.wikipedia.org/wiki/Row_echelon_form) and Back Substitution

### Input
```json
{
    "keys": {
        "n": 4,
        "k": 3
    },
    "1": {
        "base": "10",
        "value": "4"
    },
    "2": {
        "base": "2",
        "value": "111"
    },
    "3": {
        "base": "10",
        "value": "12"
    },
    "6": {
        "base": "4",
        "value": "213"
    }
}
```

### JSON Parser
Input: input.json

Output: Unmarshal data into keys `Keys` and entries `map[string]Entry`

Also we get n, the number of entries given

We get k, which is equations required to solve polynomial.

We get m, which is degree of polynomial.

### Conversion
Convert bases (say from base 2 to base 6)

Input: base, value

Output: Value in decimal

### Create equation
Input: x, y, m

Output: Linear equation {expression, y}

### Create Augmented Matrix
Input: Linear equations and their values y

Output: Matrix with linear equations and y


### Gaussian Elimination Method
#### Row Echelon Form
Input: Augmented matrix

Output: Convert to row echelon form

#### Back substitution
Input: Row echelon form

Output: Give the value of variables in a map {a: X, b: X, c: X, ....}


> We got our answer


