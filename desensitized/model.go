package desensitizedUtil

type DesensitizedType int

const (
	UserId      DesensitizedType = 1  // 用户ID
	ChineseName DesensitizedType = 2  // 中文名称
	IDCard      DesensitizedType = 3  // 身份证号
	FixedPhone  DesensitizedType = 4  // 固定电话号码
	MobilePhone DesensitizedType = 5  // 移动电话号码
	Address     DesensitizedType = 6  // 地址
	Email       DesensitizedType = 7  // 邮箱
	Password    DesensitizedType = 8  // 密码
	CarLicense  DesensitizedType = 9  // 车牌号：油车、电车
	BankCard    DesensitizedType = 10 // 银行卡号
	IPV4        DesensitizedType = 11 // ipv4
	IPV6        DesensitizedType = 12 // ipv6
)
