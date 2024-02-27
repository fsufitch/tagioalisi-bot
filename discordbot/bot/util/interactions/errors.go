package interactions

type InteractionError struct {
	Title       string
	Description string
}

func (err InteractionError) Error() string {
	return err.Title
}
