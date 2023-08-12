package service

type NewCourseRequest struct {
	Name  string `json:"name"`
	Price string `json:"price"`
}

type NewCourseResponse struct {
}
