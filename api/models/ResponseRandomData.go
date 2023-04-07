package models

type ResponseRandomData struct {
	Id         int32     `json:"id"`
	Uid        string    `json:"uid"`
	Hex_value  string    `json:"hex_value"`
	Color_name string    `json:"color_name"`
	Hsl_value  []float32 `json:"hsl_value"`
	Hsla_value []float32 `json:"hsla_value"`
}
