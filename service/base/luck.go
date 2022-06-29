package base

import "fehu/common/lib/mysql_lib"

type Luck struct {
	GiftId int     `db:"giftid"` // 礼物ID
	Nums   int     `db:"nums"`   // 数量
	Times  int     `db:"times"`  // 倍数
	Rate   float64 `db:"rate"`   // 中奖概率
	IsAll  byte    `db:"isall"`  // 是否全站，0否1是
}
type LuckList []*Luck

func GetLuckRate() LuckList {
	lucks := LuckList{}
	mysql_lib.Select(&lucks, "SELECT giftid,nums,times,rate,isall FROM cmf_luck_rate ORDER BY id desc")
	return lucks
}

func CheckLuck(giftId int, count int) LuckList {
	lucks := GetLuckRate()
	olucks := LuckList{}
	for _, v := range lucks {
		if v.GiftId == giftId && v.Nums == count {
			olucks = append(olucks, v)
		}
	}
	return olucks
}

type JackpotRate struct {
	GiftId int    `db:"giftid"`       // 礼物ID
	Nums   int    `db:"nums"`         // 数量
	Data   string `db:"rate_jackpot"` // 数据
}
type JackpotRateList []*JackpotRate

func GetJackpotRate() JackpotRateList {
	v := JackpotRateList{}
	mysql_lib.Select(&v, "SELECT giftid,nums,rate_jackpot FROM cmf_jackpot_rate ORDER BY id desc")
	return v
}

func CheckJackpot(giftId int, count int) *JackpotRate {
	r := GetJackpotRate()
	for _, v := range r {
		if v.GiftId == giftId && v.Nums == count {
			return v
		}
	}
	return nil
}
