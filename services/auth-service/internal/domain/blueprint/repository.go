package blueprint

import (
	"auth-service/internal/domain/valueobject"
	"context"
)

type Repository interface {
	StudentRepository
}

type StudentRepository interface {
	GetStudentByCredentials(ctx context.Context, studentCredential valueobject.StudentCredentialsUnion)
}
