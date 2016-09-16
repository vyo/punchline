package model

type Project struct {
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func NewProject(key string, name string, description string) *Project {
	return &Project{key, name, description}
}

func (p *Project) GetKey() *string {
	return &p.Key
}

func (p *Project) GetName() *string {
	return &p.Name
}

func (p *Project) GetDescription() *string {
	return &p.Description
}
