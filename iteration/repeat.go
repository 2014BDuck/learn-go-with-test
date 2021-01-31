// @Author: 2014BDuck
// @Date: 2021/1/31

package iteration

const repeatCount = 5

func Repeat(character string, repeatTime int) string {
	result := ""
	if repeatTime < 0 {
		repeatTime = repeatCount
	}
	for i := 0; i < repeatTime; i++ {
		result += character
	}
	return result
}
