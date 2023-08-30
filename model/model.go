package model

import "time"

// 定义数据库中存储的模型

// User 用户实体
type User struct {
	ID             int64  `gorm:"type:BIGINT;primaryKey"`
	Name           string `gorm:"type:varchar(65);not null;unique;uniqueIndex:idx_user_name" json:"name"`
	Password       string `gorm:"type:varchar(65);not null" json:"password"`
	Signature      string `gorm:"type:varchar(200);not null" json:"signature"`
	FollowCount    int64  `gorm:"type:BIGINT;unsigned;not null;default:0" json:"follow_count"`
	FollowerCount  int64  `gorm:"type:BIGINT;unsigned;not null;default:0" json:"follower_count"`
	FavoriteCount  int64  `gorm:"type:BIGINT;unsigned;not null;default:0" json:"favorite_count"`
	TotalFavorited int64  `gorm:"type:BIGINT;unsigned;not null;default:0" json:"total_favorited"`
	WorkCount      int64  `gorm:"type:BIGINT;unsigned;not null;default:0" json:"work_count"`
}

// Video 视频实体
type Video struct {
	ID            int64 `gorm:"type:BIGINT;primaryKey"`
	CreatedAt     time.Time
	AuthorID      int64  `gorm:"type:BIGINT;not null;index:idx_author_id" json:"author_id"`
	CommentCount  int64  `gorm:"type:BIGINT;not null;default:0" json:"comment_count"`
	CoverURL      string `gorm:"type:varchar(100);not null" json:"cover_url"`
	FavoriteCount int64  `gorm:"type:BIGINT;not null;default:0" json:"favorite_count"`
	PlayURL       string `gorm:"type:varchar(100);not null" json:"play_url"`
	Title         string `gorm:"type:varchar(40);not null" json:"title"`
}

// Favorite 点赞实体
type Favorite struct {
	ID      int64 `gorm:"type:BIGINT;primaryKey"`
	UserID  int64 `gorm:"type:BIGINT;not null;index:idx_user_id" json:"user_id"`
	VideoID int64 `gorm:"type:BIGINT;not null;index:idx_video_id" json:"video_id" `
}

// Comment  评论实体
type Comment struct {
	ID        int64 `gorm:"type:BIGINT;primaryKey"`
	CreatedAt time.Time
	UserID    int64  `gorm:"type:BIGINT;not null" json:"user_id"`
	VideoID   int64  `gorm:"type:BIGINT;not null;index:idx_video_id" json:"video_id"`
	Content   string `gorm:"type:varchar(100);not null" json:"content"`
}

// Follow 关注实体
type Follow struct {
	ID       int64 `gorm:"type:BIGINT;primaryKey"`
	UserID   int64 `gorm:"type:BIGINT;not null;index:idx_user_id" json:"user_id" validate:""`
	ToUserID int64 `gorm:"type:BIGINT;not null;index:idx_to_user_id" json:"to_user_id" validate:""`
}

// Message 消息实体
type Message struct {
	ID         int64 `gorm:"type:BIGINT;primaryKey"`
	CreatedAt  time.Time
	FromUserID int64  `gorm:"type:BIGINT;not null" json:"from_user_id"`
	ToUserID   int64  `gorm:"type:BIGINT;not null" json:"to_user_id"`
	Content    string `gorm:"type:varchar(100);not null" json:"content"`
}
