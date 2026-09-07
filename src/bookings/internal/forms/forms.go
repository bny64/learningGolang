package forms

import (
	"net/http"
	"net/url"
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

// Has 메서드는 HTTP 요청의 폼 데이터에 특정 필드가 존재하고 값이 비어있지 않은지 확인합니다.
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)

	if x == "" {
		f.Errors.Add(field, "This field cannot be empty")
		return false
	}

	return true
}
