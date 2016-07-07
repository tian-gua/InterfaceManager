package entity


type Interfaces struct {
	Id          int
	Name        string
	Description string
	Url         string
	ModuleId    int
	Method      int
	Example     string
}

func (this *Interfaces) MethodString() string {
	if this.Method == 1 {
		return "Put"
	}
	return "Post"
}