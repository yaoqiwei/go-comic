package success

type DataBase struct {
	ErrorCode int    `json:"code" example:"0"`
	ErrorMsg  string `json:"msg" example:"ok"`
}

type InfoData struct {
	DataBase
	Data map[string]interface{} `json:"data"`
}

type ListData struct {
	DataBase
	Data struct {
		List  []map[string]interface{} `json:"list"`
		Count int64                    `json:"count"`
	} `json:"data"`
}

type ListNoCountData struct {
	DataBase
	Data struct {
		List []map[string]interface{} `json:"list"`
	} `json:"data"`
}
