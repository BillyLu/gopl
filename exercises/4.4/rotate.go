//Package rotate Exercise 4.4: Write a version of rotate that operates in a single pass.
package rotate

// rotate a int left by n element
func rotate(s []int, n int) {
	tmp := make([]int, n)

	copy(tmp, s[:n])
	copy(s, s[n:])
	copy(s[len(s)-n:], tmp)
}
