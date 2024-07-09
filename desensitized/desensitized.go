package desensitizedUtil

import strUtil "github.com/OrangeDrk/go-toolbox/str"

// Desensitized 数据脱敏
func Desensitized(str string, desensitizedType DesensitizedType) string {
	if strUtil.IsEmpty(str) {
		return ""
	}
	newStr := str
	switch desensitizedType {
	case UserId:
		newStr = userId(str)
		break
	case ChineseName:
		newStr = chineseName(str)
		break
	case IDCard:
		newStr = idCard(str, 1, 2)
		break
	case FixedPhone:
		newStr = fixedPhone(str)
		break
	case MobilePhone:
		newStr = mobilePhone(str)
		break
	case Address:
		newStr = address(str, 8)
		break
	case Email:
		newStr = email(str)
		break
	case Password:
		newStr = password(str)
		break
	case CarLicense:
		newStr = carLicense(str)
		break
	case BankCard:
		newStr = bankCard(str)
	case IPV4:
		newStr = ipv4(str)
		break
	case IPV6:
		newStr = ipv6(str)
		break
	default:

	}

	return newStr
}

// 用户ID脱敏
func userId(id string) string {
	return "0"
}

// 中文名称脱敏
func chineseName(name string) string {
	// 只展示第一个字符
	if strUtil.IsEmpty(name) {
		return ""
	}
	return strUtil.Hide(name, 1, strUtil.Length(name))
}

// 身份证脱敏
func idCard(cardNo string, front int, end int) string {
	if strUtil.IsEmpty(cardNo) {
		return ""
	}
	if front+end > strUtil.Length(cardNo) {
		return ""
	}
	if front < 0 || end < 0 {
		return ""
	}

	return strUtil.Hide(cardNo, front, strUtil.Length(cardNo)-end)
}

// 固定电话脱敏
func fixedPhone(phone string) string {
	if strUtil.IsEmpty(phone) {
		return ""
	}

	return strUtil.Hide(phone, 4, strUtil.Length(phone)-2)
}

// 移动电话脱敏
func mobilePhone(phone string) string {
	if strUtil.IsEmpty(phone) {
		return ""
	}

	return strUtil.Hide(phone, 3, strUtil.Length(phone)-4)
}

// 地址脱敏,后 s 位进行脱敏
func address(addr string, sensitiveSize int) string {
	if strUtil.IsEmpty(addr) {
		return ""
	}
	length := strUtil.Length(addr)
	return strUtil.Hide(addr, length-sensitiveSize, length)
}

// 邮箱脱敏
func email(email string) string {
	if strUtil.IsEmpty(email) {
		return ""
	}
	index := strUtil.IndexOf(email, "@")
	if index <= 1 {
		return email
	}
	return strUtil.Hide(email, 1, index)
}

// 密码脱敏
func password(pass string) string {
	if strUtil.IsEmpty(pass) {
		return ""
	}
	return strUtil.RepeatByLength("*", strUtil.Length(pass))
}

// 车牌号脱敏
func carLicense(carNo string) string {
	if strUtil.IsEmpty(carNo) {
		return ""
	}

	newCarNo := carNo
	// 普通车牌
	if strUtil.Length(carNo) == 7 {
		newCarNo = strUtil.Hide(carNo, 3, 6)
	} else if strUtil.Length(carNo) == 8 { // 新能源
		newCarNo = strUtil.Hide(carNo, 3, 7)
	}
	return newCarNo
}

// 银行卡号脱敏
func bankCard(cardNo string) string {
	if strUtil.IsEmpty(cardNo) {
		return cardNo
	}

	cleanCarNo := strUtil.CleanEmpty(cardNo)

	if strUtil.Length(cleanCarNo) < 9 {
		return cardNo
	}
	length := strUtil.Length(cleanCarNo)
	endLength := length % 4
	if endLength == 0 {
		endLength = 4
	}
	midLength := length - 4 - endLength

	newCardNo := []rune(cleanCarNo[:4])
	for i := 0; i < midLength; i++ {
		if i%4 == 0 {
			newCardNo = append(newCardNo, ' ')
		}
		newCardNo = append(newCardNo, '*')
	}
	newCardNo = append(newCardNo, ' ')
	newCardNo = append(newCardNo, []rune(cleanCarNo[length-endLength:length])...)

	return string(newCardNo)
}

// ipv4 脱敏
func ipv4(ip string) string {
	newIP := strUtil.SubBefore(ip, ".", false)
	return newIP + ".*.*.*"
}

// ipv6 脱敏
func ipv6(ip string) string {
	newIP := strUtil.SubBefore(ip, ":", false)
	return newIP + ":*:*:*:*:*:*:*"
}
