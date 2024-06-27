package usecases

type CommentNotFound struct{}

func (err *CommentNotFound) Error() string {
	return "Cannot edit comment. Comment not found."
}

type NoPermission struct{}

func (err *NoPermission) Error() string {
	return "cannot edit other user's comment"
}
