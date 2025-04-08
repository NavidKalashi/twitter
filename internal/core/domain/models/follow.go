package models

type Follow struct {
	ID            uint   `gorm:"primaryKey" json:"id"`
	FollowerName  string `gorm:"index" json:"follower_name"`
	FollowingName string `gorm:"index" json:"following_name"`
}
