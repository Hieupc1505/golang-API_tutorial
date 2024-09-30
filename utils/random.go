package utils

import (
	"fmt"
	"math/rand"
	"reflect"
	"strings"
)

const alphabet = "abcdefghiklmnopqrstuvwxyz"

// RandomString generates a random string of length n
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// randomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomOwner generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomDescription generates a random description
func RandomDescription() string {
	return RandomString(10)
}

// RandomEnum nhận vào một mảng bất kỳ kiểu dữ liệu và trả về một giá trị ngẫu nhiên từ mảng đó.
func RandomEnum(arr interface{}) interface{} {
	// Sử dụng reflect để làm việc với mảng
	v := reflect.ValueOf(arr)

	// Kiểm tra nếu đầu vào không phải là mảng
	if v.Kind() != reflect.Slice {
		return nil // hoặc có thể trả về một lỗi
	}

	// Chọn ngẫu nhiên một chỉ số
	randomIndex := rand.Intn(v.Len())

	// Trả về giá trị tại chỉ số ngẫu nhiên
	return v.Index(randomIndex).Interface()
}

// radom email
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}
