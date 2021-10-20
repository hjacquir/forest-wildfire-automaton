# forest-wildfire-automaton

## Description

### Input

The automaton takes : 
* l [int] (length), 
* h [int] (height) and 
* a list of trees [array] already burning 
with coordinates [string] (example : XY = ["3-1", "2-2"])

### Process
The automaton selects a tree at random from the list of burning trees, 
it searches for its nearest green neighbors and burns one.

Then he takes the closest neighbor that is still green to the one that 
has just burned down and continues.

The automaton stops when all the 
trees step by step are burnt

### Output

For input : l = 3 (X) and h = 2 (Y) and with initial trees burning
["3-1", "2-2"]

```
GreenTrees : [{1-1 false} {1-2 false} {2-1 false} {3-2 false}]
RedTrees : [{3-1 true} {2-2 true}]
one red random : 2-2
-- burned : 2-1
-- burned : 1-1
-- burned : 1-2

GreenTrees : [{1-1 false} {1-2 false} {2-1 false} {3-2 false}]
RedTrees : [{3-1 true} {2-2 true}]
one red random : 3-1
-- burned : 3-2

GreenTrees : [{1-1 false} {1-2 false} {2-1 false} {3-2 false}]
RedTrees : [{3-1 true} {2-2 true}]
one red random : 3-1
-- burned : 2-1
-- burned : 1-1
-- burned : 1-2
```
## Execute

`go run main.go` OR copy-paste the main.go code on the Go playground
https://play.golang.org/

## To do

* parallelism with one go routine per random red tree
* functional tests
* configuration file with : size, red tree XY