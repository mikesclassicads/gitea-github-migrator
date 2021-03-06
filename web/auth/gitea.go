package auth

import (
	"code.gitea.io/sdk/gitea"
	"git.jonasfranz.software/JonasFranzDEV/gitea-github-migrator/web/context"
)

// GiteaLoginForm represents the data required for logging in into gitea
type GiteaLoginForm struct {
	Username    string `form:"username"`
	Password    string `form:"password"`
	AccessToken string `form:"access-token"`
	GiteaURL    string `form:"gitea-url"`
	Type        string `form:"use" binding:"Required;In(token,password)"`
}

// LoginToGitea handles the POST request for signing in with a Gitea account
func LoginToGitea(ctx *context.Context, form GiteaLoginForm) {
	var token string
	if form.Type == "password" {
		client := gitea.NewClient(form.GiteaURL, "")
		tkn, err := client.CreateAccessToken(form.Username, form.Password, gitea.CreateAccessTokenOption{
			Name: "gitea-github-migrator",
		})
		if err != nil {
			ctx.Flash.Error("Cannot create access token please check your credentials!")
			ctx.Redirect("/")
			return
		}
		token = tkn.Sha1
	} else {
		token = form.AccessToken
	}
	client := gitea.NewClient(form.GiteaURL, token)
	usr, err := client.GetMyUserInfo()
	if err != nil {
		ctx.Flash.Error("Invalid Gitea credentials.")
		ctx.Redirect("/")
		return
	}
	ctx.Session.Set("gitea_user", &context.User{
		ID:        usr.ID,
		Username:  usr.UserName,
		Token:     token,
		AvatarURL: usr.AvatarURL,
	})
	ctx.Session.Set("gitea", form.GiteaURL)
	ctx.Redirect("/")
	return
}
