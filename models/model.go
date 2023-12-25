package models

type Lenders struct {
	ID        int    `json:"id"`
	FULL_NAME string `json:"full_name"`
}

type Books struct {
	ID             int    `json:"id"`
	BOOK_NAME      string `json:"book_name"`
	BOOK_AUTHOR    string `json:"book_author"`
	AVAILABLE_TILL string `json:"available_till"`
	GENRE          string `json:"genre"`
	HOSTED_BY      int    `json:"hosted_by"`
	AVAILABLE      int    `json:"available"`
}

type Borrows struct {
	ID                int    `json:"id"`
	BOOK              int    `json:"book"`
	TAKEN_BY          int    `json:"taken_by"`
	BORROW_START_TIME string `json:"borrow_start_time"`
	BORROW_END_TIME   string `json:"borrow_end_time"`
	BORROW_COMPLETE   int    `json:"borrow_complete"`
}
