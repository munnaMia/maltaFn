package patterns

import "fmt"

var (
	print   = fmt.Print
	println = fmt.Println
)

// Print a square on n.
func StarSquare(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			print("*")
		}
		println()
	}
}

// Output:
// *****
// *****
// *****
// *****
// *****

// Prnt Right Triangle of n
func RightTringle(n int) {
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			print("*")
		}
		println()
	}
}

// Output:
// *
// **
// ***
// ****
// *****

// Print a Inverse Right Tringle of n
func RightTringleInverse(n int) {
	for i := 1; i <= n; i++ {
		for j := i; j <= n; j++ {
			print("*")
		}
		println()
	}
}

// Output :
// *****
// ****
// ***
// **
// *

// Print a right tringle of number where numbers are present as columns
func NumberRightTringleCol(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			print(j)
		}
		println()
	}
}

// Output:
// 1
// 12
// 123
// 1234
// 12345

// Print a right tringle of number where numbers are present as rows
func NumberRightTringleRow(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			print(i)
		}
		println()
	}
}

//Output :
// 1
// 22
// 333
// 4444
// 55555

// Print a inverse right tringle of number where numbers are present as columns
func NumberRightTringleColInverse(n int) {
	for i := n; i > 0; i-- {
		for j := 1; j <= i; j++ {
			print(j)
		}
		println()
	}
}

// Output
// 12345
// 1234
// 123
// 12
// 1

// Print a Binary Right tringle & format : 010101... like this
func BinaryTringle(n int) {
	var start int
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			start = 0
		} else {
			start = 1
		}

		for j := 1; j <= i; j++ {
			print(start)
			start = 1 - start
		}
		println()
	}
}

// output :
// 1
// 01
// 101
// 0101
// 10101

// Print a Number right tringle & format : 1 2 3 4 ... n like this
func NumberTringle(n int) {
	start := 1
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			print(start, " ")
			start++
		}
		println()
	}
}

// Output:
// 1
// 2 3
// 4 5 6
// 7 8 9 10
// 11 12 13 14 15

// Print double tringle with in between spaces like : 1___1 and 12__21 ...
func DualTringleBetSpc(n int) {
	space := 2 * (n - 1)
	for i := 1; i <= n; i++ {
		//number
		for j := 1; j <= i; j++ {
			print(j)
		}
		//space
		for k := 1; k <= space; k++ {
			print(" ")
		}
		//number
		for j := i; j >= 1; j-- {
			print(j)
		}
		println()
		space -= 2
	}
}

// Output
// 1        1
// 12      21
// 123    321
// 1234  4321
// 1234554321

// Print tringle of n
func Tringle(n int) {
	for i := 0; i < n; i++ {
		for j := n - i - 1; j > 0; j-- {
			print(" ")
		}
		for k := 0; k < i*2+1; k++ {
			print("*")
		}
		println()
	}
}

// Output:
//     *
//    ***
//   *****
//  *******
// *********

// Print inverse tringle of n
func TringleInverse(n int) {
	for i := n - 1; i >= 0; i-- {
		for j := n - i - 1; j > 0; j-- {
			print(" ")
		}
		for k := i*2 + 1; k > 0; k-- {
			print("*")
		}
		println()
	}
}

// Output:
// *********
//  *******
//   *****
//    ***
//     *

// Print a Side wall tringle of n
func SideWallTringle(n int) {
	for i := 1; i <= 2*n-1; i++ {
		star := i
		if i > n {
			star = 2*n - i
		}
		for j := 1; j <= star; j++ {
			print("*")
		}
		println()
	}
}

// Output:
// *
// **
// ***
// ****
// *****
// ****
// ***
// **
// *

// Print alphabet right tringle alphabet as column
func AlpaRightTringle(n int) {
	for i := 1; i <= n; i++ {
		for j := 'A'; j < 'A'+rune(i); j++ {
			print(string(j))
		}
		println()
	}
}

// Output
// A
// AB
// ABC
// ABCD
// ABCDE

// Print inverse alphabet right tringle alphabet as column
func ReverseAlpaTringle(n int) {
	for i := n; i > 0; i-- {
		for j := 'A'; j < 'A'+rune(i); j++ {
			print(string(j))
		}
		println()
	}
}

// Output :
// ABCDE
// ABCD
// ABC
// AB
// A

// Print alphabet right tringle alphabet as row
func AlpaTringleRow(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			print(string('A' + rune(i)))
		}
		println()
	}
}

// Output
// A
// BB
// CCC
// DDDD
// EEEEE

// Print alphabet  tringle
func AlpaPiramid(n int) {
	for i := 0; i < n; i++ {
		for j := n - i - 1; j > 0; j-- {
			print(" ")
		}
		ch := 'A'
		breakPoint := int((i*2 + 1) / 2)
		for k := 1; k <= i*2+1; k++ {
			print(string(ch))
			if breakPoint >= k {
				ch++
			} else {
				ch--
			}
		}
		println()
	}
}

// Output
//     A
//    ABA
//   ABCBA
//  ABCDCBA
// ABCDEDCBA

// Print backward alphabet right tringle
func BackWardAplaTringle(n int) {
	for i := 0; i < n; i++ {
		for ch := 'E' - rune(i); ch <= 'E'; ch++ {
			print(string(ch))
		}
		println()
	}
}

// output
// E
// DE
// CDE
// BCDE
// ABCDE
