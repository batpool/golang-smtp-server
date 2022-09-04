package smtpsrvr

type HandlerFunc func(*Context) error
type AuthFunc func(username, password string) error
