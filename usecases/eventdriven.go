package usecases

import (
	"fmt"
	"time"

	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/go-toolkit/jwtkit"
	"github.com/ilhammhdd/kudaki-entities/user"

	"github.com/golang/protobuf/proto"
)

const (
	DateTimeFormat = "%d-%02d-%02d %02d:%02d:%02d"
)

func GetUserFromKudakiToken(kudakiToken string) *user.User {
	jwt, err := jwtkit.GetJWT(jwtkit.JWTString(kudakiToken))
	errorkit.ErrorHandled(err)

	userClaim := jwt.Payload.Claims["user"].(map[string]interface{})
	usr := &user.User{
		AccountType: user.AccountType(user.AccountType_value[userClaim["account_type"].(string)]),
		Email:       userClaim["email"].(string),
		PhoneNumber: userClaim["phone_number"].(string),
		Role:        user.Role(user.Role_value[userClaim["role"].(string)]),
		Uuid:        userClaim["uuid"].(string),
	}

	return usr
}

func TimeNowToDateTime() string {
	now := time.Now()
	return fmt.Sprintf(DateTimeFormat, now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
}

type EventDrivenUsecase interface {
	Handle(in proto.Message) (out proto.Message)
}
