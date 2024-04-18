package service

import (
	"time"
	"web-forum/internal/store"
	"web-forum/models"
)

type Post interface {
	CreatePost(form *models.DataTransfer) (int, error)
	CreateComment(form *models.CommentInPost) error
	GetPost(id int) (models.GetOnePost, error)
	GetPosts() ([]models.GetAllPosts, error)
	GetMyCreatedPost(userId int) ([]models.GetAllPosts, error)
	GetMyLikesPost(userId int) ([]models.GetAllPosts, error)
	ReactionPost(postId, userId int, reactionType string) error
	ReactionComment(postId, userId, commetId int, reactionType string) error
	CheckLastPost() (int, error)
	CheckLastComment() (int, error)
}

type Autorization interface {
	InsertUser(form *models.UserSignUp) error
	Authenticate(form *models.UserSignIn) (int, error)
	UserSessionsAdd(userId int, sessionToken string, expiresAt time.Time) error
	DeleteToken(sessionToken string) error
	GetIdInSessions(sessionToken string) (int, error)
	CheckSessions(userId int) (bool, error)
	UpdateToken(sessionToken string, user_id int) error
	GetTokenSession(cookieToken string) (bool, error)
}

type Service struct {
	Post
	Autorization
}

func New(store *store.Store) *Service {
	return &Service{
		Post:         NewPostService(store.Post),
		Autorization: NewAuthService(store.Autorization),
	}
}
