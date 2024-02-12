package models

type User struct {
	User_Name     *string `json:"user_name" db:"user_name"`
	User_bio      *string `json:"user_bio,omitempty" db:"user_bio,omitempty"`
	User_profile  *string `json:"user_profile,omitempty" db:"user_profile,omitempty"`
	User_email    *string `json:"user_email" db:"user_email"`
	User_password string  `json:"user_password" db:"user_password"`
	Gender        *string `json:"gender" db:"gender"`
	Created_at    *int64  `json:"created_at" db:"created_at"`
	Updated_at    *int64  `json:"updated_at,omitempty" db:"updated_at,omitempty"`
	Deleted_at    *int64  `json:"deleted_at,omitempty" db:"deleted_at,omitempty"`
}

type LoginUserReq struct {
	User_email    *string `json:"user_email" db:"user_email"`
	User_password string  `json:"user_password" db:"user_password"`
}

type UpdateUserReq struct {
	UserId     int64   `json:"id" db:"id"`
	User_Name  *string `json:"user_name,omitempty" db:"user_name,omitempty"`
	User_bio   *string `json:"user_bio,omitempty" db:"user_bio,omitempty"`
	User_email *string `json:"user_email,omitempty" db:"user_email,omitempty"`
	Updated_at *int64  `json:"updated_at,omitempty" db:"updated_at,omitempty"`
}

type BlogsReq struct {
	Id         *int64  `json:"id" db:"id"`
	UserId     *int64  `json:"user_id" db:"user_id"`
	Title      *string `json:"title" db:"title"`
	Content    *string `json:"content" db:"content"`
	BlogImage  *string `json:"blog_image,omitempty" db:"blog_image,omitempty"`
	Created_at *int64  `json:"created_at" db:"created_at"`
}

type DeleteBlogReq struct {
	BlogId *int64 `json:"id" db:"id"`
	UserId int64  `json:"user_id" db:"user_id"`
}

type UpdateBlogContentReq struct {
	Id         *int64  `json:"id" db:"id"`
	UserId     int64   `json:"user_id" db:"user_id"`
	Content    *string `json:"content" db:"content"`
	Updated_at *int64  `json:"updated_at,omitempty" db:"updated_at,omitempty"`
}
type UpdateBlogTitleReq struct {
	Id         *int64  `json:"id" db:"id"`
	UserId     int64   `json:"user_id" db:"user_id"`
	Title      *string `json:"title" db:"title"`
	Updated_at *int64  `json:"updated_at,omitempty" db:"updated_at,omitempty"`
}
type UpdateBlogImgReq struct {
	Id         *int64  `json:"id" db:"id"`
	UserId     int64   `json:"user_id" db:"user_id"`
	Img        *string `json:"blog_image" db:"blog_image"`
	Updated_at *int64  `json:"updated_at,omitempty" db:"updated_at,omitempty"`
}

type Question struct {
	UserId      int64   `json:"user_id" db:"user_id"`
	Title       *string `json:"title" db:"title"`
	Content     *string `json:"content" db:"content"`
	QuestionLan *string `json:"question_language" db:"question_language"`
	Picture     *string `json:"picture" db:"picture"`
	Created_at  *int64  `json:"created_at" db:"created_at"`
}
