package patreon

type Error struct {
	Code     int    `json:"code"`
	CodeName string `json:"code_name"`
	Detail   string `json:"detail"`
	Id       string `json:"id"`
	Status   string `json:"status"`
	Title    string `json:"title"`
}

type ErrorResponse struct {
	Errors []Error `json:"errors"`
}

func (e ErrorResponse) Error() string {
	// In most cases there is only one error
	if len(e.Errors) > 0 {
		return e.Errors[0].Detail
	}

	return ""
}
