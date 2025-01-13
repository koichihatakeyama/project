package validator

import (
    "project/internal/model"
    "regexp"
)

type UserValidator struct{}

func NewUserValidator() *UserValidator {
    return &UserValidator{}
}

func (v *UserValidator) ValidateCreate(req *model.CreateUserRequest) map[string]string {
    errors := make(map[string]string)

    if req.Name == "" {
        errors["name"] = "名前は必須です"
    }
    if len(req.Name) > 100 {
        errors["name"] = "名前は100文字以内で入力してください"
    }

    if req.Email == "" {
        errors["email"] = "メールアドレスは必須です"
    }
    if !isValidEmail(req.Email) {
        errors["email"] = "有効なメールアドレスを入力してください"
    }

    return errors
}

func isValidEmail(email string) bool {
    pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
    return regexp.MustCompile(pattern).MatchString(email)
}
