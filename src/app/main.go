package main

import (
	"strconv"
	"time"

	//"github.com/happierall/l"
	soso "github.com/happierall/soso-server"
	"github.com/happierall/soso-auth"
	"math/rand"
	"views"
	"github.com/happierall/l"
	"os"
)


var lastID int64 = 0

type MyUser struct {
	ID    int64
	Name  string
	Email string
}

var Users = []*MyUser{}



func main() {
	Admins := []string{
		os.Getenv("firstAdmin"),
	}

	rand.Seed(time.Now().Unix())

	Router := soso.Default()

	authObj := auth.New("super_duper_key", false, Router)

	googleAuth := auth.UseGoogleAuth(
		authObj,
		os.Getenv("googleClientId"),
		os.Getenv("googleSecret"),
		[]string{"email"},
		"http://localhost:8000", // Then in github settings set: "http://localhost:8000/oauth/callback/github"
	)



	googleAuth.OnSuccess = func(userData *auth.User, session soso.Session, authType string) {
		for _, user := range Admins  {
			if userData.Email == user {
				user := &MyUser{
					Name:  userData.Name,
					Email: userData.Email,
				}

				// Try auth current
				for _, u := range Users {
					if u.Email == user.Email {
						successResponce(u, authObj.Sign, session)
						return
					}
				}

				// Or register
				lastID++
				user.ID = lastID
				Users = append(Users, user)

				successResponce(user, authObj.Sign, session)
			}
		}
	}


	// Read token from every request (other.token)
	Router.Middleware.Before(func(m *soso.Msg, start time.Time) {
		token, tokenData, err := auth.ReadToken(m, authObj.Sign)
		if err != nil {
			return
		}

		for _, u := range Users {
			if u.ID == tokenData.UID {

				strID := strconv.FormatInt(tokenData.UID, 10)
				m.User.ID = strID
				m.User.Token = token
				m.User.IsAuth = true
				m.User.IsAnonymous = tokenData.IsAnonymous

				// Register session
				soso.Sessions.Push(m.Session, strID)

			}
		}

	})

	Router.SEARCH("user", func(m *soso.Msg) {
		if m.User.IsAuth {

			uid, _ := strconv.Atoi(m.User.ID)

			for _, user := range Users {
				if user.ID == int64(uid) {

					l.Logf("User email: %s", user.Email)

					m.Success(map[string]interface{}{
						"auth": true,
						"email": user.Email,
					})


				}
			}
		}
	})



	Router.HandleRoutes(views.Routes)
	l.Print("Running app at localhost:8000")
	Router.Run(8000)
}

func successResponce(user *MyUser, sign string, session soso.Session) {
	authToken := auth.CreateToken(map[string]interface{}{
		"uid":         user.ID,
		"isAnonymous": false,
	}, sign)
	soso.SendMsg("auth", "SUCCESS", session, map[string]interface{}{
		"token": authToken,
		"user":  user,
	})
}
