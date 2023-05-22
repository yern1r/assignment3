package entities

type Fullname struct {
	Firstname  string `json:"firstname"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}

type Contact struct {
	ID       int      `json:"id"`
	fullname Fullname `json:"fullname"`
	Phone    string   `json:"phone" validate:"number"`
}

func (c *Contact) Value() Fullname {
	return c.fullname
}

func (c *Contact) Validate() error {
	validate := validator.New()
	err := validate.Struct(c)
	if err != nil {
		return err
	}
	return nil
}
