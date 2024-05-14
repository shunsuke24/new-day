package presenter

type UserInformation struct {
	// ユーザID
	ID string `json:"id"`
	// ユーザ名
	Name string `json:"name"`
	// 年齢
	Age int `json:"age"`
}
