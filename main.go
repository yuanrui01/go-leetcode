package main

import (
	"fmt"
	"leetcode/daily"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {
	// s := "aabaaaacaabc"
	k := 3
	nums := []int{1, 2, 4, 5}
	fmt.Println(daily.MinimumDifference(nums, k))
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
