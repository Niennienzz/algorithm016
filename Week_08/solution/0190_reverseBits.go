package solution

// Reverse bits of a given 32 bits unsigned integer.
//
// Examples:
//
// Input: n = 00000010100101000001111010011100
// Output: 964176192 (00111001011110000010100101000000)
// Explanation:
//  - The input binary string 00000010100101000001111010011100 represents the unsigned integer 43261596.
//  - Return 964176192 which its binary representation is 00111001011110000010100101000000.
//
// Input: n = 11111111111111111111111111111101
// Output: 3221225471 (10111111111111111111111111111111)
// Explanation:
//  - The input binary string 11111111111111111111111111111101 represents the unsigned integer 4294967293.
//  - Return 3221225471 which its binary representation is 10111111111111111111111111111111.

// We iterate through the bit string of the input integer, from right to left (i.e. n = n >> 1).
// To retrieve the right-most bit of an integer, we apply the bit AND operation (n & 1).
//
// For each bit, we reverse it to the correct position (i.e. (n & 1) << power).
// Then we accumulate this reversed bit to the final result.
//
// When there is no more bits of one left (i.e. n == 0), we terminate the iteration.
func reverseBits(num uint32) uint32 {
	res := uint32(0)
	pow := uint32(31)
	for num != 0 {
		res += (num & 1) << pow
		num = num >> 1
		pow -= 1
	}
	return res
}
