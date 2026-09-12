package forms

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

// Form 구조체는 폼 데이터와 검증 에러 정보를 포함합니다.
type Form struct {
	url.Values
	Errors errors
}

// Valid는 폼에 에러가 없는지 확인합니다.
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// New 함수는 새로운 Form 구조체 인스턴스를 생성하고 초기화합니다.
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Required는 폼에 특정 필드들이 모두 존재하는지 확인합니다.
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be empty")
		}
	}
}

// Has 메서드는 HTTP 요청의 폼 데이터에 특정 필드가 존재하고 값이 비어있지 않은지 확인합니다.
func (f *Form) Has(field string) bool {
	x := f.Get(field)

	if x == "" {
		return false
	}

	return true
}

// MinLength는 폼 필드의 최소 길이를 검증합니다.
func (f *Form) MinLength(field string, length int) bool {
	value := f.Get(field)
	if len(value) < length {
		f.Errors.Add(field, fmt.Sprintf("This field must be at least %d characters long", length))
		return false
	}
	return true
}

// IsEmail는 폼 필드가 유효한 이메일 형식인지 검증합니다.
func (f *Form) IsEmail(field string) {
	value := f.Get(field)
	if !govalidator.IsEmail(value) {
		f.Errors.Add(field, "Invalid email address")
	}
}
