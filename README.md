# go-pg-orm
Wrapper that simplifies use of Golang ORM with focus on PostgreSQL found at https://github.com/go-pg/pg

NOTE: Work in progress!

### Installation
To use this tool, install it from code with make install, go install directly, or 
``` go get -u github.com/AndrewDonelson/go-pg-orm ```

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
	// Create models object
	models := *pgorm.Model.NewModel()
	err = models.Open()
	if err != nil {
		return err
	}
	
	// Register models
	models.Register(
		&models.User{},
		&models.Story{},
		// ... Register More models here ...
	)
	
  	//Do something with models
}
```
