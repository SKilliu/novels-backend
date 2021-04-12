package middlewares

type Middleware struct {
	auth string
}

func New(authKey string) *Middleware {
	return &Middleware{
		auth: authKey,
	}
}
