package entities

type Group struct {
	ID   int    `json:"id"`
	Name string `json:"name" validate:"max=250"`
}

func (g *Group) Validate() error {
	validate := validator.New()
	err := validate.Struct(g)
	if err != nil {
		return err
	}
	return nil
}
