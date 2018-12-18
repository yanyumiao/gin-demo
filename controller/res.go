package controller

type Res struct {
	Ret  int         `json:"ret"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func res(code int, desc string, data interface{}) Res {
	var res Res
	res.Ret = code
	res.Msg = desc
	res.Data = data
	return res
}
