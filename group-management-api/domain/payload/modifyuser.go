package payload

type ModifyUserPayload struct {
	Name string `json:"name"`
	Email string `json:"email"`
}

func (mup ModifyUserPayload) Validate() error {
	panic("implement me")
}

