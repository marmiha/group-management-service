package payload

type CreateGroupPayload struct {
	Name string `json:"name"`
}

func (cgp CreateGroupPayload) Validate() error {
	panic("implement me")
}
