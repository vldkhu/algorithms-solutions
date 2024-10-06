/*
Учитывая целое число x, верните true, если x является
палиндромом, и false в противном случае.
Input: x = 121
Output: true
Input: x = -121
Output: false

Как решить:
-для начала проверим не является ли x < 0
-если число заканчивается на 0 и не равно 0, оно также не может быть палиндромом.
-нужна переменная для хранения обратного числа
-цикл для получения обратного числа
-если длина числа четная, то x и reversedStr должны быть равны и если длина нечетная, то нужно убрать среднюю цифру из reversedStr
*/
package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x%10 == 0 && x != 0 {
		return false
	}

	reversedStr := 0
	for x > reversedStr {
		reversedStr = reversedStr*10 + x%10
		x /= 10
	}
	return x == reversedStr || x == reversedStr/10
}
