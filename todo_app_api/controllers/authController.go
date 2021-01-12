package controllers

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"

	"github.com/amesinp/learning_go/todo_app_api/dto"
	"github.com/amesinp/learning_go/todo_app_api/models"
	"github.com/amesinp/learning_go/todo_app_api/repositories"
	"github.com/amesinp/learning_go/todo_app_api/utils"
)

var userRepository = repositories.UserRepository{}
var refreshTokenRepository = repositories.RefreshTokenRepository{}

// AuthController to categorize controller functions
type AuthController struct{}

type contextKey struct {
	name string
}

type authTokenResponse struct {
	Token        string     `json:"token"`
	ExpiresAt    *time.Time `json:"expiresAt"`
	RefreshToken string     `json:"refreshToken"`
}

type authResponse struct {
	AuthToken *authTokenResponse `json:"authToken"`
	User      *models.User       `json:"user"`
}

// Login function
func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	var loginData dto.LoginDTO

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&loginData)

	validationMsg := utils.ValidateDTO(loginData)
	if validationMsg != "" {
		utils.SendErrorResponse(utils.ResponseParams{Writer: w, Message: validationMsg})
		return
	}

	user := userRepository.GetByUsername(loginData.Username)
	if user == nil {
		utils.SendErrorResponse(utils.ResponseParams{Writer: w, Message: "Invalid username or password", StatusCode: http.StatusUnauthorized})
		return
	}

	if !isHashEqual(user.PasswordHash, loginData.Password) {
		utils.SendErrorResponse(utils.ResponseParams{Writer: w, Message: "Invalid username or password", StatusCode: http.StatusUnauthorized})
		return
	}

	token, err := generateAuthToken(user.ID, r.UserAgent())
	if err != nil {
		utils.HandleServerError(w, err)
		return
	}

	utils.SendSuccessResponse(utils.ResponseParams{Writer: w, Message: "Login successful!", Data: authResponse{token, user}})
}

// Register controller function
func (c *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	var registerData dto.RegisterDTO
	json.NewDecoder(r.Body).Decode(&registerData)

	validationMsg := utils.ValidateDTO(registerData)
	if validationMsg != "" {
		utils.SendErrorResponse(utils.ResponseParams{Writer: w, Message: validationMsg})
		return
	}

	var existingUser = userRepository.GetByUsername(registerData.Username)
	if existingUser != nil {
		utils.SendErrorResponse(utils.ResponseParams{Writer: w, Message: "Username already exists"})
		return
	}

	hashedPassword, err := hashPassword(registerData.Password)
	if err != nil {
		utils.HandleServerError(w, err)
		return
	}

	var user = models.User{Name: registerData.Name, UserName: registerData.Username, PasswordHash: hashedPassword}

	err = userRepository.Create(&user)
	if err != nil {
		utils.HandleServerError(w, err)
		return
	}

	token, err := generateAuthToken(user.ID, r.UserAgent())
	if err != nil {
		utils.HandleServerError(w, err)
		return
	}

	utils.SendSuccessResponse(utils.ResponseParams{Writer: w, Message: "Registration completed successfully!", Data: authResponse{token, &user}})
}

// RefreshToken refreshes an expired access token
func (c *AuthController) RefreshToken(w http.ResponseWriter, r *http.Request) {
	var tokenData dto.RefreshTokenDTO
	json.NewDecoder(r.Body).Decode(&tokenData)

	validationMsg := utils.ValidateDTO(tokenData)
	if validationMsg != "" {
		utils.SendErrorResponse(utils.ResponseParams{Writer: w, Message: validationMsg})
		return
	}

	claims := r.Context().Value(utils.TokenClaimsKey).(utils.TokenClaims)

	token := refreshTokenRepository.Get(claims.RefreshTokenID)

	// Ensure token is valid and has not expired
	if token == nil || !isHashEqual(token.TokenHash, tokenData.Token) || token.ExpiresAt.Sub(time.Now()) < 0 {
		utils.SendErrorResponse(utils.ResponseParams{Writer: w, Message: "Refresh token is not valid"})
		return
	}

	// If token has been used before then most likely an attacker has acquired the refresh token
	// Invalidate all refresh tokens for this user to force the user to login again
	if token.IsUsed {
		log.Println("Multiple refresh token usage for user: ", claims.UserID)
		refreshTokenRepository.DeleteByUserID(claims.UserID)

		utils.SendErrorResponse(utils.ResponseParams{Writer: w, Message: "Refresh token is not valid"})
		return
	}

	// Generate new access token
	authToken, err := generateAuthToken(claims.UserID, r.UserAgent())
	if err != nil {
		utils.HandleServerError(w, err)
		return
	}

	// Set token to used
	refreshTokenRepository.UpdateToUsed(token.ID)

	utils.SendSuccessResponse(utils.ResponseParams{Writer: w, Data: authToken})
}

// Logout invalidates the refresh token
func (c *AuthController) Logout(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value(utils.TokenClaimsKey).(utils.TokenClaims)

	if r.URL.Query().Get("all") == "true" {
		refreshTokenRepository.DeleteByUserID(claims.UserID)
		utils.SendSuccessResponse(utils.ResponseParams{Writer: w, Message: "Logout successful! It might take a few minutes to reflect on all devices"})
		return
	}

	token := refreshTokenRepository.Get(claims.RefreshTokenID)
	if token != nil {
		// If refresh token has already been used then it is likely that token has been hijacked
		if token.IsUsed {
			refreshTokenRepository.DeleteByUserID(token.UserID)
		} else {
			refreshTokenRepository.Delete(token.ID)
		}
	}

	utils.SendSuccessResponse(utils.ResponseParams{Writer: w, Message: "Logout successful!"})
}

func generateAuthToken(userID int, userAgent string) (*authTokenResponse, error) {
	var authToken authTokenResponse

	refreshToken, err := createRefreshToken(userID, userAgent)
	if err != nil {
		return &authToken, err
	}

	tokenExpiry := time.Now().Add(time.Minute * 15)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"ref": refreshToken.ID,
		"exp": tokenExpiry.Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		// Delete generated refresh token
		refreshTokenRepository.Delete(refreshToken.ID)

		return &authToken, err
	}

	authToken = authTokenResponse{
		Token:        tokenString,
		ExpiresAt:    &tokenExpiry,
		RefreshToken: refreshToken.ClearToken,
	}
	return &authToken, nil
}

func createRefreshToken(userID int, userAgent string) (*models.RefreshToken, error) {
	var refreshToken models.RefreshToken

	tokenID, err := uuid.NewUUID()
	if err != nil {
		return &refreshToken, err
	}

	token := base64.StdEncoding.EncodeToString([]byte(tokenID.String()))
	hashedToken, err := hashPassword(token)
	if err != nil {
		return &refreshToken, err
	}

	expiryDate := time.Now().AddDate(0, 0, 14)
	refreshToken = models.RefreshToken{
		UserID:    userID,
		UserAgent: userAgent,
		TokenHash: hashedToken,
		ExpiresAt: &expiryDate,
	}

	err = refreshTokenRepository.Create(&refreshToken)
	if err != nil {
		return &refreshToken, err
	}

	refreshToken.ClearToken = token
	return &refreshToken, nil
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func isHashEqual(hashedValue, clearValue string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedValue), []byte(clearValue))
	return err == nil
}
