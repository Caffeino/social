package mailer

import "embed"

const (
	FromName        = "GopherSocial"
	maxRetries      = 3
	UserWelcomeTmpl = "user_invitation.tmpl"
)

//go:embed "templates"
var FS embed.FS

type Client interface {
	Send(tmplFile, username, email string, data any, isSandbox bool) (int, error)
}
