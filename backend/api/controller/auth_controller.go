package controller

// func Login(db *gorm.DB, email, password string) (interface{}, error) {
// 	var roles = []string{"USER", "ORGANIZER", "ADMIN"}

// 	for _, role := range roles {
// 		switch role {
// 		case "USER":
// 			var user User
// 			if err := db.Where("email = ?", email).First(&user).Error; err == nil {
// 				if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) == nil {
// 					return user, nil
// 				}
// 			}
// 		case "ORGANIZER":
// 			var organizer domain.Organizer
// 			if err := db.Where("email = ?", email).First(&organizer).Error; err == nil {
// 				if bcrypt.CompareHashAndPassword([]byte(organizer.Password), []byte(password)) == nil {
// 					return organizer, nil
// 				}
// 			}
// 		case "ADMIN":
// 			var admin domain.Admin
// 			if err := db.Where("email = ?", email).First(&admin).Error; err == nil {
// 				if bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password)) == nil {
// 					return admin, nil
// 				}
// 			}
// 		}
// 	}

// 	return nil, errors.New("invalid credentials")
// }
