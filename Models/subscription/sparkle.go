package subscription

type Category struct {
	CategoryID int    `json:"categoryId"`
	Name       string `json:"name"`
	Typeid     int    `json:"typeId"`
	SortOrder  int    `json:"sortOrder"`
	Status     string `json:"status"`
}
type SubCategory struct {
	CategoryID    int `json:"categoryId"`
	SubCategoryID int `json:"subcategoryId"`
    Name      string `json:"name"`
	Typeid    int    `json:"typeId"`
	SortOrder int    `json:"sortOrder"`
	Status    string `json:"status"`
}
type Module struct {
	CategoryID int `json:"categoryId"`
	ModuleID int `json:"moduleId"`
   Name string `json:"name"`
   Content string `json:"content"`
    Imageurl string `json:"imageurl"`
	Logourl  string `json:"logourl"`
}
type Packages struct {
	ModuleID  int    `json:"moduleId"`
	PackageID int    `json:"packageId"`
	Name      string `json:"name"`
    Status string `json:"status"`
}
type CompanyDetails struct {
	Name string `json:"name"`
	Regno string `json:"regno"`
	Email       string `json:"email"`
	Mobile      string `json:"mobile"`
	CategoryID    int `json:"categoryId"`
	Typeid    int    `json:"typeId"`
}
type CompanyLocation struct {
	Street string `json:"street"`
	Suburb string `json:"suburb"`
	State string `json:"state"`
	Zip string `json:"zip"`
}