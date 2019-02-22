# go-pg-orm
Wrapper that simplifies use of Golang ORM with focus on PostgreSQL found at https://github.com/go-pg/pg

NOTE: Work in progress!

### Example:
```
type User struct {
    Id     int64
    Name   string
    Emails []string
}

func (u User) String() string {
    return fmt.Sprintf("User<%d %s %v>", u.Id, u.Name, u.Emails)
}

type Story struct {
    Id       int64
    Title    string
    AuthorId int64
    Author   *User
}

func (s Story) String() string {
    return fmt.Sprintf("Story<%d %s %s>", s.Id, s.Title, s.Author)
}


func main() {

  //Create models object
	models := *pgorm.Model.NewModel()
	err = models.Open()
	if err != nil {
		return err
	}
  
  //Do something with models
}
```
