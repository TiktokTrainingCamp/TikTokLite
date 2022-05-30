package dao

import "time"

type User struct {
	UserId   int `gorm:"primaryKey"`
	Name     string
	Username string
	Password string
}

type Video struct {
	VideoId    int `gorm:"primaryKey"`
	UserId     int
	Title      string
	PlayUrl    string
	CoverUrl   string
	UpdateTime time.Time
}

type Comment struct {
	CommentId  int `gorm:"primaryKey"`
	UserId     int
	VideoId    int
	Content    string
	CreateTime time.Time
}

type Relation struct {
	Id         int `gorm:"primaryKey"`
	FollowId   int
	FollowerId int
}

type Favorite struct {
	Id      int `gorm:"primaryKey"`
	UserId  int
	VideoId int
}

type Token struct {
	TokenId    int `gorm:"primaryKey"`
	UserToken  string
	UserId     int
	ExpireTime time.Time
}
