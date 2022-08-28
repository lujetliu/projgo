package word

import (
	"strings"
	"unicode"
)

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func ToLower(s string) string {
	return strings.ToLower(s)
}

// 下划线转大写驼峰
func UnderscoreToUpperCamelCase(s string) string {
	s = strings.Replace(s, "_", "", -1)
	s = strings.Title(s)
	// strings.Title()函数返回string的副本, 其中包含string的所有Unicode字母,
	// 其起始单词映射到其Unicode标题大小写
	return strings.Replace(s, " ", "", -1)
}

// 下划线转小写驼峰
func UnderscoreToLowerCamelCase(s string) string {
	s = UnderscoreToLowerCamelCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

// 驼峰转下划线
func CamelCaseToUnderscore(s string) string {
	var output []rune
	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
			continue
		}

		if unicode.IsUpper(r) {
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(r))
	}

	return string(output)
}
