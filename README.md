To run in Go:
`go run pancakes.go`

To run in JavaScript:
`node pancakes.js`

The problem...
* Each pancake has a happy-faced side and a blank side.
* Each stack of pancakes is randomly mixed, with some facing up and some facing down.
* The goal is to end up with all the pancakes happy side up.
* The happy side is indicated with a `+` and the blank side with a `-`.
* You can pick up a group of pancakes and flip them over.
* The top of the pancake stack is the beginning of the array or slice.
* The bottom of the pancake stack is the ending of the array or slice.
* You have to pick up the pancakes from the top of the stack, in a group. You can't take a group from the bottom or from the middle. That means that every operation on the array must include the head.
* Look for the SMALLEST number of operations possible to get a result of all positives `++++`, such that all of the pancakes are happy-side up.
* The most efficient way to handle the problem is to examine each pancake from the bottom up, and only execute a flip if we reach one that is in a negative `-` position. Each flip results in some combination of positive and negative outcomes for each group that is flipped, but so long as we count from the bottom and only make a flip each time we encounter a negative, we'll reach the fastest solution.
* Solve both with Node.js and with Go. With Go, we can do this even faster by processing each stack concurrently.
* For the example of the stack starting out at `--+-`, complete the process like so...
   * Look at the tail first. Because it's `-`, the whole stack needs to be flipped, which results in `++-+`.
   * Now look at the next pancake up from the bottom, which WAS a `+`, but now that we flipped the whole stack, is now a `-`, which means we need to flip all the pancakes from that one up to the top, resulting in `--++`.
   * Next, look at the third pancake from the bottom, which also WAS a `+`, but is now a `-`, so it also needs to be flipped in a group from itself to the top of the stack, resulting in `++++`.
   * Finally, look at the top pancake, which is already a `+`, and so we can end the operation without any more flips.