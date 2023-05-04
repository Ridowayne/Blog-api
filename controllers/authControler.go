package controllers

import (
	"example/Go/helpers"
	"example/Go/initializers"
	"example/Go/models"
	"net/http"
	"os"
	"time"
	"unicode"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)


func CreateUser(c *gin.Context) {
	var body struct {
		FirstName string 
		LastName string
		Email string 
		Password string
		Phone string
		Role string
	}
	
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error Message":"KIndly fill in all the details to sign up"})
		return
	}
	if len(body.Phone) != 10 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid phone number"})
        return 
    }

	for _, r := range body.Phone {
        if !unicode.IsDigit(r) {
			c.JSON(http.StatusBadRequest, gin.H{"message": "phone number can only contain numeric characters"})
			
            return 
        }
    }

	var existingUser models.User
	
	 initializers.DB.First(&existingUser, "email = ?", body.Email)
	 

	if existingUser.ID !=0 {
		c.JSON(http.StatusBadRequest, gin.H{"Error Message":"User already registered"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error Message":"Failed to hash password"})
		return
	}
	
	user := models.User{FirstName: body.FirstName, LastName: body.LastName, Email: body.Email, Password: string(hash)}
	newUser:= initializers.DB.Create(&user)

	if newUser.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error Message":"failed to create accoount"})
		return
	}


	c.JSON(http.StatusCreated, gin.H{"message": "User signed up successfully"})


}
func LoginUser(c *gin.Context) {
	var body struct {		
		Email string 
		Password string		
	}
	
	if c.Bind(&body)!= nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error Message":"KIndly fill in all the details to login"})
		return

	}
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)
	
	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"Error Message":"Inavalid email address or password"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error Message":"Inavalid email address or password"})
		return

	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),		

	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message":"failed to create token"})
		return	
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString,3600 * 24* 30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})

}

func ForgotPassword(c *gin.Context) {
	var body struct{
		Email string 
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Kindly provide an email address"})
		return
	}
	
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"Error Message":"Kindly confirm your email address and try again"})
		return
	}
	helpers.Sendmail([]string{"hassanrilwan12@gmail.com"},)
	c.JSON(http.StatusOK, gin.H{"message":"check your mail"})


}
 func ResetPassword(c *gin.Context) {

 }
 func ViewMyProfile(c *gin.Context){
	user, _ :=c.Get("user")

	id:= user.(models.User).ID
	var profile models.User

	 initializers.DB.First(&profile, id)
	 c.JSON(http.StatusOK, gin.H{"profile": profile})
	
	
 }
 func EditProfile(c *gin.Context){
	user, _ := c.Get("user")
	id:= user.(models.User).ID
	var newInfo struct{
		FirstName *string
		LastName *string
		Phone *string
	}
	c.Bind(newInfo)

	
	var existingUser models.User
	initializers.DB.First(&existingUser, id)
	existingUser.FirstName =*newInfo.FirstName
	existingUser.LastName =*newInfo.LastName
	existingUser.Phone =*newInfo.Phone
	initializers.DB.Save(&existingUser)
	
 }