package helper

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// Claims là cấu trúc để lưu thông tin người dùng trong token JWT
type Claims struct {
	ID bson.ObjectID `json:"id"` // Lưu thông tin ID của người dùng
	jwt.RegisteredClaims  // Các trường chuẩn của JWT như exp, iat
}

// CreateJWT tạo một JWT mới với ID của người dùng và thời gian hết hạn là 24 giờ
func CreateJWT(id bson.ObjectID) string {
	// Tạo một JWT mới với các claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Hết hạn sau 24 giờ
		},
	})

	// Ký và trả về token dưới dạng chuỗi
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		// In lỗi nếu có sự cố khi tạo token
		fmt.Println("Error creating token:", err)
		return ""
	}
	return tokenString
}
// RemoveBearerPrefix để xóa định dạng bearer đầu chuỗi jwt
func RemoveBearerPrefix(tokenString string) string {
	// Kiểm tra nếu token bắt đầu bằng "Bearer "
	if strings.HasPrefix(tokenString, "Bearer ") {
			// Loại bỏ phần "Bearer " (7 ký tự đầu)
			return strings.TrimSpace(strings.TrimPrefix(tokenString, "Bearer "))
	}
	return tokenString
}
// ParseJWT phân tích và xác thực JWT, trả về Claims nếu hợp lệ
func ParseJWT(tokenString string) (*Claims, error) {
	tokenString = RemoveBearerPrefix(tokenString)
	// Phân tích token và kiểm tra tính hợp lệ của nó
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		// Kiểm tra Signature của token
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		// Trả về secret key để xác thực token
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}

	// Kiểm tra nếu token hợp lệ và trả về claims
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}


