package controllers

import (
	"fmt"
	"i-shop/models"
	"i-shop/pkg/db/rediss"
	"i-shop/storage"
	"i-shop/utils"
	"i-shop/utils/password"
	"i-shop/validation"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)



type AuthController struct {
	Storage *storage.UserStorage
	Redis *rediss.RedisDb
}

func NewAuthController(storage *storage.UserStorage, redis *rediss.RedisDb) (*AuthController, error) {
	return &AuthController{
		Storage: storage,
		Redis: redis,
	}, nil
}

var validateObj *validator.Validate

func init() {
	validateObj = validator.New()
	validation.RegisterCustomValidators(validateObj)
	fmt.Println("-----------============")
	
}

// CreateUser godoc
// @Summary      Create a new user
// @Description  Registers a new user by providing phone number, email, and password. 
//               Validates phone number, email, and password, then generates a verification code for phone number validation.
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user body models.UserRegister true "User Registration Data"
// @Success      201 {object} models.UsersResponse "Successfully created the user"
// @Failure      400 {object} map[string]interface{} "Bad request"
// @Failure      500 {object} map[string]interface{} "Internal server error"
// @Router       /auth/register [post]
func (ac *AuthController) CreateUser(c *gin.Context) {
	var user models.UserRegister

	err := c.ShouldBindJSON(&user)
	if err != nil {
		HandleResponse(c, http.StatusBadRequest, err.Error())
		log.Println("failed to parse request user", err)
		return
	}

	if err := validateObj.Struct(&user); err != nil {
		HandleResponse(c, http.StatusBadRequest, err.Error())
		log.Println("Validatsiya error", err)
		return
	}

	hashedPassword, err := password.HashPassword(user.Password)
	if err != nil {
		HandleResponse(c, http.StatusInternalServerError, "Password hashing failed: "+err.Error())
		log.Println("Password hashing error", err)
		return
	}

	code := utils.GenerateCode(6)
	log.Println("Tasdiqlash kodi >>> ", code)

	if err := ac.Redis.SetEx(c, user.PhoneNumber, code, 5*time.Minute); err != nil {
		HandleResponse(c, http.StatusInternalServerError, "Failed to save verification code")
		return
	}

	info := models.Users{
		FirstName: user.FirstName,
		Email: user.Email,
		PhoneNumber: user.PhoneNumber,
		Password: hashedPassword,
		Role: user.Role,
	}

	if err := ac.Storage.Create(&info); err != nil {
		HandleResponse(c, http.StatusInternalServerError, "Failed to create user"+err.Error())
		log.Printf("Failed to create user: %v", info)
		return
	}

	HandleResponse(c, http.StatusCreated, info)
}

// VerifyCode godoc
// @Summary      Verify phone number with the code
// @Description  Verifies the user's phone number using the code sent earlier. If valid, activates the user.
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body models.Request true "Phone number and code verification data"
// @Success      200 {string} string "User verified and activated successfully"
// @Failure      400 {object} map[string]interface{} "Invalid or expired code"
// @Failure      500 {object} map[string]interface{} "Internal server error"
// @Router       /auth/verify [post]
func (ac *AuthController) VerifyCode(c *gin.Context) {
	var request models.Request

	if err := c.ShouldBindJSON(&request); err != nil {
		HandleResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	storedCode, err := ac.Redis.Get(c, request.PhoneNumber)
	if err != nil || storedCode != request.Code {
		HandleResponse(c, http.StatusBadRequest, "Invalid or expired code")
		return
	}

	if err := ac.Storage.ActivateUser(request.PhoneNumber); err != nil {
		HandleResponse(c, http.StatusInternalServerError, "Failed to activate user")
		return
	}

	_ = ac.Redis.Delete(c, request.PhoneNumber)

	HandleResponse(c, http.StatusOK, "User verified and activated successfully")
}

// LoginUser godoc
// @Summary      Login a user
// @Description  Allows a user to log in using email and password. If valid, returns a JWT token.
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        login body models.LoginRequest true "Login Credentials"
// @Success      200 {object} map[string]interface{} "Successfully restored the category"
// @Failure      400 {object} map[string]interface{} "Invalid request"
// @Failure      401 {object} map[string]interface{} "Unauthorized: Invalid credentials"
// @Failure      404 {object} map[string]interface{} "User not found"
// @Failure      500 {object} map[string]interface{} "Internal server error"
// @Router       /auth/login [post]
func (ac *AuthController) LoginUser(c *gin.Context) {
	var login models.LoginRequest

	if err := c.ShouldBindJSON(&login); err != nil {
		HandleResponse(c, http.StatusBadRequest, err.Error())
		log.Println("Failed to parse request login", err)
		return
	}

	user, err := ac.Storage.FindByEmail(login.Email)
	if err != nil {
		HandleResponse(c, http.StatusNotFound, "User not found")
		log.Printf("User with Email %s not found: %v", login.Email, err)
		return
	}

	if !user.IsActive {
		HandleResponse(c, http.StatusUnauthorized, "User is not verified yet")
		log.Printf("User with Email %s is not active", login.Email)
		return
	}

	if !password.CheckPasswordHash(login.Password, user.Password) {
		HandleResponse(c, http.StatusUnauthorized, "invalid email or password")
		log.Printf("Login failed: Incorrect password for user with email: %s", user.Email)
		return
	}
 
	token, err := utils.GenerateToken(user.Email, user.Role)
	if err != nil {
    	HandleResponse(c, http.StatusInternalServerError, "Failed to generate token")
    	log.Println("Token generation failed", err)
    	return
	}

	HandleResponse(c, http.StatusOK, gin.H{
		"token": token,
	})
	
}