package common

import "encoding/json"

// SwapTo 通过 json tag 进行结构体赋值
func SwapTo(request, category interface{}) (err error) {
	dataByte, err := json.Marshal(request)
	if err != nil {
		return err
	}
	return json.Unmarshal(dataByte, category)
}
