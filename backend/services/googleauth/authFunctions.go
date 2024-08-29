package googleauth

import (
	"HangAroundBackend/config"

	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const userInfoAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="
const refreshTokenAPI = "https://www.googleapis.com/oauth2/v4/token"

var oauth2Config = oauth2.Config{
	ClientID:     "YOUR_CLIENT",
	ClientSecret: "YOUR_SECRET",
	RedirectURL:  "",
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
	Endpoint:     google.Endpoint,
}

// Provide the URL to the frontend for google login
func GetURL() (url string, state string) {
	b := make([]byte, 16)
	rand.Read(b)
	state = base64.URLEncoding.EncodeToString(b)
	return oauth2Config.AuthCodeURL(state, oauth2.AccessTypeOffline), state
}

func Exchange(code string, c *gin.Context) (*oauth2.Token, error) {
	return oauth2Config.Exchange(c, code)
}

func GetUserInfo(accessToken string) (map[string]interface{}, error) {
	resp, err := http.Get(userInfoAPI + accessToken)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	//unmarshal Json
	var userData map[string]interface{}
	err = json.Unmarshal(data, &userData)
	return userData, err
}

func ResetToken(c *gin.Context, refreshToken string, accessToken string) (*oauth2.Token, error) {
	token := &oauth2.Token{
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
	}
	tokenSource := oauth2Config.TokenSource(c, token)
	return tokenSource.Token()
}

func init() {

	oauth2Config.RedirectURL = config.GetEnv("GOOGLE_CLIENT_SECRET")
	oauth2Config.ClientID = config.GetEnv("GOOGLE_CLIENT_ID")
	oauth2Config.ClientSecret = config.GetEnv("GOOGLE_CLIENT_SECRET")
}
