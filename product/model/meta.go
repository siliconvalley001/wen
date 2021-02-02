package model

type MenusList struct {
	Id int64 `gorm:"primary_key;auto_increment;not null" json:"id"`
	Authname string `gorm:"unique_index;not null" json:"authname"`
	Path string `json:"path"`
	//Children []ChildrenList `gorm:"ForeignKey:ChildrenList" json:"children"`
	ChildrenId int64 `json:"children_id"`
}

type ChildrenList struct {
	
	Id int64 `gorm:"primary_key;auto_increment;not null" json:"id"`
	Uid int64 `json:"uid"`
	Authname string `gorm:"unique_index;not null" json:"authname"`
	Path string `json:"path"`
	//Children []ChildrenList `gorm:"ForeignKey:ChildrenList" json:"children"`
}
type List struct {
	Id int64 `gorm:"primary_key;auto_increment;not null" json:"id"`
	Authname string `gorm:"unique_index;not null" json:"authname"`
	Path string `json:"path"`
	//Children []ChildrenList `gorm:"ForeignKey:ChildrenList" json:"children"`
	Children []ChildrenList `gorm:"ForeignKey:ChildrenList" json:"children"`

}
