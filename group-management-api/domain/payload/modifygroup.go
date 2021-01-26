package payload

type ModifyGroupPayload struct {
	Name string `json:"name"`
}

func (mgp ModifyGroupPayload) Validate() error {
	panic("implement me")
}
