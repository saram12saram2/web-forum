package handler

import "net/http"

func (h *Handler) InitRouters() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	dynamic := aliceNew(h.requireAuthentication)

	mux.HandleFunc("/", h.home)
	mux.Handle("/my-posts", dynamic.ThenFunc(h.MyPosts))
	mux.Handle("/my-likes", dynamic.ThenFunc(h.MyLikes))

	mux.HandleFunc("/post/view", h.postView)
	mux.Handle("/post/create", dynamic.ThenFunc(h.postCreate))
	mux.Handle("/post/coment-post", dynamic.ThenFunc(h.createPostComment))

	mux.Handle("/post/reaction", dynamic.ThenFunc(h.reaction))
	mux.Handle("/post/reaction-form-home", dynamic.ThenFunc(h.reactionFromHome))
	mux.Handle("/post/reaction_comment", dynamic.ThenFunc(h.reactionComment))

	mux.HandleFunc("/auth/sign-up", h.signUp)
	mux.HandleFunc("/auth/sign-in", h.signIn)
	mux.Handle("/logout", dynamic.ThenFunc(h.logout))

	standard := aliceNew(h.recoverPanic, h.logRequest)

	return standard.Then(mux)
}
