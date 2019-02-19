/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 *
 * https://leetcode-cn.com/problems/palindrome-number/description/
 *
 * algorithms
 * Easy (55.90%)
 * Total Accepted:    66.8K
 * Total Submissions: 119.4K
 * Testcase Example:  '121'
 *
 * 判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
 *
 * 示例 1:
 *
 * 输入: 121
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: -121
 * 输出: false
 * 解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
 *
 *
 * 示例 3:
 *
 * 输入: 10
 * 输出: false
 * 解释: 从右向左读, 为 01 。因此它不是一个回文数。
 *
 *
 * 进阶:
 *
 * 你能不将整数转为字符串来解决这个问题吗？
 *
 */
// package main

// import "strconv"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0{
		return true
	}
	var i int
	s := strconv.Itoa(x)
	rs := ""
	for i = len(s) - 1; i >= 0; i-- {
		rs += string(s[i])
	}
	for i=0;i<len(rs);i++{
		if string(rs[i])=="0"{
			continue
		}
		break
	}
	if i!=0{
		return false
	}
	if s == rs {
		return true
	}
	return false
}

// func main() {

// }
