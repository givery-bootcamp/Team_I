package repositories

const FOREIGN_KEY_CONSTRAINTS_ERROR_NUMBER = 1452

type PostIdNotFound struct {
}

func (err *PostIdNotFound) Error() string {
	return "Cannot create comment. Post id not found."
}
