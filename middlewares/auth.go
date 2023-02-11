package middleware

// func Auth(echo.HandlerFunc) echo.HandlerFunc {

// 	authHeader := c.Request.Header.Get("Authorization")

// 	if authHeader == "" {
// 		utils.SendResponse(c, http.StatusUnauthorized, "Authorization header not found")
// 		return
// 	}

// 	userID, err := utils.ValidateToken(authHeader)
// 	if err != nil {
// 		utils.SendResponse(c, http.StatusUnauthorized, "Unauthorized")
// 		return
// 	}

// 	// if admin userID
// 	if userID == 0 {
// 		utils.SendResponse(c, http.StatusUnauthorized, "Unauthorized")
// 		return
// 	}

// 	c.Set("userID", userID)
// 	c.Next()

// 	return func(c echo.Context) error {
// 		return next(c)
// 	}

// }
