package bean

// create table problem
// (
//
//	id        int primary key auto_increment,
//	dfbyid    varchar(255),
//	title     varchar(255),
//	contestid int references contest (id)
//
// )
type Problem struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Contestid int    `json:"contestid"`
}

func (Problem) TableName() string {
	return "problem"
}
