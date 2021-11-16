/*
1812. Determine Color of a Chessboard Square

https://leetcode.com/problems/determine-color-of-a-chessboard-square/

You are given coordinates, a string that represents the coordinates of a square of the chessboard. Below is a chessboard for your reference.



Return true if the square is white, and false if the square is black.

The coordinate will always represent a valid chessboard square. The coordinate will always have the letter first, and the number second.



Example 1:

Input: coordinates = "a1"
Output: false
Explanation: From the chessboard above, the square with coordinates "a1" is black, so return false.
Example 2:

Input: coordinates = "h3"
Output: true
Explanation: From the chessboard above, the square with coordinates "h3" is white, so return true.
Example 3:

Input: coordinates = "c7"
Output: false


Constraints:

coordinates.length == 2
'a' <= coordinates[0] <= 'h'
'1' <= coordinates[1] <= '8'
*/

package main

func squareIsWhite(coordinates string) bool {
    return (int(coordinates[0] - 'a') + int(coordinates[1] - '0')) % 2 == 0
}

// Space Complexity: O(1)
// Time Complexity: O(1)