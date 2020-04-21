//参考：https://my.oschina.net/bravozu/blog/1827254

package code

const (
	CHARSLENGTH = 32    // 字符集长度
	CODELENGTH  = 7     // 码长度
	PRIME1      = 3     // 质数1
	PRIME2      = 5     // 质数2
	SLAT        = 12345 // 随机数
)

var CHARS = []byte{'2', '3', '4', '5', '6', '9', '7', '8', 'A', 'B',
	'C', 'D', 'E', 'F', 'G', 'H', 'J', 'K', 'L', 'M', 'N', 'P', 'Q',
	'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}

var CHARSMAP = map[byte]int{
	'2': 0, '3': 1, '4': 2, '5': 3, '6': 4, '9': 5, '7': 6, '8': 7,
	'A': 8, 'B': 9, 'C': 10, 'D': 11, 'E': 12, 'F': 13, 'G': 14,
	'H': 15, 'J': 16, 'K': 17, 'L': 18, 'M': 19, 'N': 20, 'P': 21,
	'Q': 22, 'R': 23, 'S': 24, 'T': 25, 'U': 26, 'V': 27, 'W': 28,
	'X': 29, 'Y': 30, 'Z': 31}

func Encode(id int) string {
	//补位，并扩大整体
	id = id*PRIME1 + SLAT

	//将 id 转换成32进制的值
	var a = make([]int, CODELENGTH)

	//32进制数
	a[0] = id
	for i := int(0); i < CODELENGTH-1; i++ {
		a[i+1] = a[i] / CHARSLENGTH
		//扩大每一位的差异
		a[i] = (a[i] + i*a[0]) % CHARSLENGTH
	}

	// 校验位
	sum := 0
	for i := 0; i < CODELENGTH-1; i++ {
		sum += a[i]
	}
	a[CODELENGTH-1] = sum * PRIME1 % CHARSLENGTH

	//进行混淆
	var b = make([]byte, CODELENGTH)
	for i := 0; i < CODELENGTH; i++ {
		b[i] = CHARS[a[i*PRIME2%CODELENGTH]]
	}

	return string(b)
}

func Decode(code string) int {
	if len(code) != CODELENGTH {
		return -1
	}

	//将字符还原成对应数字
	var a = make([]int, CODELENGTH)
	for i := 0; i < CODELENGTH; i++ {
		index, ok := CHARSMAP[code[i]]
		if !ok {
			return -1
		}
		a[i*PRIME2%CODELENGTH] = index
	}

	// 校验位
	sum := 0
	for i := 0; i < CODELENGTH-1; i++ {
		sum += a[i]
	}
	if a[CODELENGTH-1] != sum*PRIME1%CHARSLENGTH {
		return -1
	}

	var b = make([]int, CODELENGTH)
	for i := CODELENGTH - 2; i >= 0; i-- {
		b[i] = (a[i] - a[0]*i + CHARSLENGTH*i) % CHARSLENGTH
	}

	var res int
	for i := CODELENGTH - 2; i >= 0; i-- {
		res += b[i]
		if i > 0 {
			res *= CHARSLENGTH
		}
	}
	return (res - SLAT) / PRIME1
}
