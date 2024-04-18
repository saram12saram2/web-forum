package models

import (
	"web-forum/validator"
)

var (
	LikePost       = "like"
	DislikePost    = "dislike"
	LikeComment    = "likeComm"
	DislikeComment = "dislikeComm"
)

type DataTransfer struct {
	PostId     int
	UserId     int
	UserName   string
	Title      string
	Content    string
	CreateDate string
	Tags       []string
	validator.Validator
}

type Teg struct {
	Teg string
}

type GetOnePost struct {
	Id         int
	Title      string
	Content    string
	CreateDate string
	UserName   string
	Tags       []string
	Comments   []CommentInPost
	validator.Validator
	Likes    int
	Dislikes int
}

type LikeANDDislike struct {
	Id      int
	UserId  int
	Like    int
	Dislike int
}

type CommentInPost struct {
	Id         int
	PostId     int
	UserId     int
	Content    string
	UserName   string
	CreateDate string
	Likes      int
	Dislikes   int
	validator.Validator
}

type GetAllPosts struct {
	Id         int
	CreateDate string
	Title      string
	UserName   string
	Likes      int
	Dislikes   int
	Tags       []string
	ValidTags  string
}

/*
configs.NewConfig() загружает настройки из файла конфигурации.
store.NewSqlite3(config) инициализирует подключение к базе данных SQLite.
Сервисы (service.New(store)) и обработчики (handler.New(service)) создаются для выполнения бизнес-логики и обработки HTTP-запросов соответственно.
server.Run() запускает HTTP-сервер, который обслуживает входящие запросы.
*/

/*
В configs.go, имя функции NewConfig() ->  LoadConfig()  (функция загружает конфигурацию из файла)

В handler.go, функцию New() -> NewHandler(),   (создается экземпляр обработчика)
              Функция getTags() ->  FilterPostsByTags()

В server.go, метод Run()  ->  Start(), (запуска сервера)

В service.go, конструктор New()  ->  NewService(), (создает новый сервис)


В store.go, функцию NewSqlite3()   ->   OpenSqliteDB() (открывает базу данных)
         Функция CreateTables()   ->    InitializeDatabase(), (инициализации структуры базы данных)

В models.go, некоторые структуры и переменные имеют названия, которые можно улучшить для большей ясности:
        DataTransfer может быть переименована в PostData или PostRequest.
        Teg следует исправить на Tag для правильного английского.
        LikeANDDislike можно переименовать в PostReaction для отражения реакций на пост.

*/
