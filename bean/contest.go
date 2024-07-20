package bean

//create table contest
//(
//    id    int primary key auto_increment,
//    title varchar(255)
//);

type Contest struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

func (Contest) TableName() string {
	return "contest"
}
