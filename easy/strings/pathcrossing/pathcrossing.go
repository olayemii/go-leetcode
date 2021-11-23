/*
1496. Path Crossing

https://leetcode.com/problems/path-crossing/

Given a string path, where path[i] = 'N', 'S', 'E' or 'W', each representing moving one unit north, south, east, or west, respectively. You start at the origin (0, 0) on a 2D plane and walk on the path specified by path.

Return true if the path crosses itself at any point, that is, if at any time you are on a location you have previously visited. Return false otherwise.



Example 1:


Input: path = "NES"
Output: false
Explanation: Notice that the path doesn't cross any point more than once.
Example 2:


Input: path = "NESWW"
Output: true
Explanation: Notice that the path visits the origin twice.


Constraints:

1 <= path.length <= 104
path[i] is either 'N', 'S', 'E', or 'W'.
*/
package main

import (
	"strconv"
	"strings"
)

func isPathCrossing(path string) bool {
    curr := "0,0"
    paths := map[string]bool{
    }
    paths[curr] = true
    
    for _, char := range path {
        currInt := strings.Split(curr, ",")
        x, _ := strconv.Atoi(currInt[0])
        y, _ := strconv.Atoi(currInt[1])
        if char == 'N' {
            curr = strconv.Itoa(x) + "," + strconv.Itoa(y-1)
            if _, ok := paths[curr]; ok {
                return true
            }else{
                 paths[curr] = true   
            }
        }else
        if char == 'S' {
            curr = strconv.Itoa(x) + "," + strconv.Itoa(y+1)
            if _, ok := paths[curr]; ok {
                return true
            }else{
                 paths[curr] = true   
            }
        }else
        if char == 'W' {
            curr = strconv.Itoa(x-1) + "," + strconv.Itoa(y)
            if _, ok := paths[curr]; ok {
                return true
            }else{
                 paths[curr] = true   
            }
        }else
        if char == 'E' {
            curr = strconv.Itoa(x+1) + "," + strconv.Itoa(y)
            if _, ok := paths[curr]; ok {
                return true
            }else{
                 paths[curr] = true   
            }
        }
    }
    
    return false
}

// Time Complexity: O(n)
// Space Complexity: O(n)