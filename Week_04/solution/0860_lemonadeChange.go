package solution

// At a lemonade stand, each lemonade costs $5.
// Customers are standing in a queue to buy from you, and order one at a time (in the order specified by bills).
// Each customer will only buy one lemonade and pay with either a $5, $10, or $20 bill.
// You must provide the correct change to each customer, so that the net transaction is that the customer pays $5.
// Note that you don't have any change in hand at first.
// Return true if and only if you can provide every customer with correct change.
//
// Examples:
//
// Input: [5,5,10]
// Output: true
//
// Input: [10,10]
// Output: false
//
// Input: [5,5,5,10,20]
// Output: true
// Explanation:
//  - From the first 3 customers, we collect three $5 bills in order.
//  - From the fourth customer, we collect a $10 bill and give back a $5.
//  - From the fifth customer, we give a $10 bill and a $5 bill.
//  - Since all customers got correct change, we output true.
//
// Input: [5,5,10,10,20]
// Output: false
// Explanation:
//  - From the first two customers in order, we collect two $5 bills.
//  - For the next two customers in order, we collect a $10 bill and give back a $5 bill.
//  - For the last customer, we can't give change of $15 back because we only have two $10 bills.
//  - Since not every customer received correct change, the answer is false.

// 分析: 零钱的面值成倍数关系, 可以贪心.
func lemonadeChange(bills []int) bool {
	p5, p10 := 0, 0

	for _, bill := range bills {
		if bill == 5 {
			p5 += 1
		}

		if bill == 10 {
			if p5 >= 1 {
				p5 -= 1
				p10 += 1
			} else {
				return false
			}
		}

		if bill == 20 {
			if p10 >= 1 && p5 >= 1 {
				p10 -= 1
				p5 -= 1
			} else if p10 == 0 && p5 >= 3 {
				p5 -= 3
			} else {
				return false
			}
		}
	}

	return true
}
