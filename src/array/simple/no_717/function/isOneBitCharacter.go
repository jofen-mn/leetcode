package function

/**
有两种特殊字符。第一种字符可以用一比特0来表示。第二种字符可以用两比特(10 或 11)来表示。
现给一个由若干比特组成的字符串。问最后一个字符是否必定为一个一比特字符。给定的字符串总是由0结束。
示例 1:

输入:
bits = [1, 0, 0]
输出: True
解释:
唯一的编码方式是一个两比特字符和一个一比特字符。所以最后一个字符是一比特字符。
示例 2:

输入:
bits = [1, 1, 1, 0]
输出: False
解释:
唯一的编码方式是两比特字符和两比特字符。所以最后一个字符不是一比特字符。
注意:

1 <= len(bits) <= 1000.
bits[i] 总是0 或 1.
 */
/**
我们可以对bits数组从左到右扫描来判断最后一位是否为一比特字符。当扫描到第 i 位时，如果 bits[i]=1
，那么说明这是一个两比特字符，将i的值增加 2。如果bits[i]=0，那么说明这是一个一比特字符，将i的值增加 1。
如果i最终落在了bits.length−1 的位置，那么说明最后一位一定是一比特字符。
 */
func IsOneBitCharacter(bits []int) bool {
	start := 0
	for ; start < len(bits) - 1; {
		//if bits[start] == 1 {
		//	start += 2
		//} else {
		//	start ++
		//}

		start += bits[start] + 1 // bit[start]只能为0或1
	}

	return start == len(bits) - 1
}
