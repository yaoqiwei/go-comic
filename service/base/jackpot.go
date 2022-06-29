package base

import "fehu/common/lib/mysql_lib"

type JackpotInfo struct {
	Total int64 `json:"total" db:"total"`
	Level int   `json:"level" db:"level"`
}

func GetJackpotInfo() JackpotInfo {
	info := JackpotInfo{}
	mysql_lib.FetchOne(&info, "SELECT total,level FROM cmf_jackpot WHERE id=1")
	return info
}

type JackpotLevel struct {
	Id int   `json:"levelid" db:"levelid"`
	Up int64 `json:"level_up" db:"level_up"`
}

type JackpotLevelList []*JackpotLevel

func GetJackpotLevel() JackpotLevelList {
	info := JackpotLevelList{}
	mysql_lib.Select(&info, "SELECT levelid,level_up FROM cmf_jackpot_level ORDER BY levelid")
	return info
}

func GetJackpotLevelInfo(total int64) int {
	i := GetJackpotLevel()
	for _, v := range i {
		if v.Up <= total {
			return v.Id
		}
	}
	return 0
}
