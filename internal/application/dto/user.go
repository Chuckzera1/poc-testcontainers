package dto

type CreateUserReqBody struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UserListDTO struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
