//参考1：https://huzb.me/2018/03/23/%E7%AE%80%E5%8D%95%E7%9A%84%E5%AF%86%E7%A0%81%E5%AD%A6%E7%94%9F%E6%88%90%E5%94%AF%E4%B8%80%E9%82%80%E8%AF%B7%E7%A0%81/
//参考2：https://my.oschina.net/bravozu/blog/1827254
package code

const (
	chars   = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	codeLen = 7     // 码长度
	slat    = 12345 // 随机数 - 用于扩大ID
	prime1  = 3     // 质数1 - 用于加密ID
	prime2  = 5     // 质数2 - 用于混淆码 (选择与code长度互质的数)
	prime3  = 17    // 质数3 - 用于混淆字符集 (选择与chars长度互质的数)

	//混淆其实就是将数字洗牌。比如把 1234567 洗成 5237641。
	//这样处理之后可以隐藏密钥和密文之间的关系。
	//洗牌的方式也很简单，选择一个和数组长度互质的数 prime2，和数组角标相乘取余即可
)

var (
	_chars    []byte
	_charsMap map[byte]int
	_charsLen int
)

func init() {
	_charsLen = len(chars)

	_chars = make([]byte, _charsLen)
	for i := 0; i < _charsLen; i++ {
		_chars[i] = chars[i*prime3%_charsLen]
	}

	_charsMap = make(map[byte]int)
	for i, v := range _chars {
		_charsMap[v] = i
	}
}

func Encode(id int) string {
	//补位，并扩大整体
	id = id*prime1 + slat

	//将 id 转换成32进制的值
	var a = make([]int, codeLen)

	//32进制数
	a[0] = id
	for i := int(0); i < codeLen-1; i++ {
		a[i+1] = a[i] / _charsLen
		//扩大每一位的差异
		a[i] = (a[i] + i*a[0]) % _charsLen
	}

	// 校验位
	sum := 0
	for i := 0; i < codeLen-1; i++ {
		sum += a[i]
	}
	a[codeLen-1] = sum * prime1 % _charsLen

	//进行混淆
	var b = make([]byte, codeLen)
	for i := 0; i < codeLen; i++ {
		b[i] = _chars[a[i*prime2%codeLen]]
	}

	return string(b)
}

func Decode(code string) int {
	if len(code) != codeLen {
		return -1
	}

	//将字符还原成对应数字
	var a = make([]int, codeLen)
	for i := 0; i < codeLen; i++ {
		index, ok := _charsMap[code[i]]
		if !ok {
			return -1
		}
		a[i*prime2%codeLen] = index
	}

	// 校验位
	sum := 0
	for i := 0; i < codeLen-1; i++ {
		sum += a[i]
	}
	if a[codeLen-1] != sum*prime1%_charsLen {
		return -1
	}

	var b = make([]int, codeLen)
	for i := codeLen - 2; i >= 0; i-- {
		b[i] = (a[i] - a[0]*i + _charsLen*i) % _charsLen
	}

	var res int
	for i := codeLen - 2; i >= 0; i-- {
		res += b[i]
		if i > 0 {
			res *= _charsLen
		}
	}
	return (res - slat) / prime1
}
