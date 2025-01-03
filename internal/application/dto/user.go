package dto

type CreateUserReqBody struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
