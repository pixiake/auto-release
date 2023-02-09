package repo

type Repo struct {
	Name     string   `json:"name" yaml:"name"`
	Tag      string   `json:"tag" yaml:"tag"`
	Branches []string `json:"branches" yaml:"branches"`
}
