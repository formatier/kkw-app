package entity

import (
	"auth-service/internal/domain/valueobject"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type CommonModel struct {
	Role           valueobject.Role // Kind
	Name, Lastname string
	SchoolId       string
}

type StudentModel struct {
	Level        valueobject.StudentLevel
	Class        valueobject.StudentClass
	PhoneNumber  string
	Email        string
	PasswordHash string
}

type TeacherModel struct {
	Level      *valueobject.StudentLevel
	Class      *valueobject.StudentClass
	JobTitle   string
	CitizenId  string
	Department string
}

type StudentTeacherUnion struct {
	CommonModel
	Student *StudentModel
	Teacher *TeacherModel
}

func UnmarshalJSON(raw []byte) error {
	commonModel := &CommonModel{}
	err := bson.Unmarshal(raw, commonModel)
	if err != nil {
		return err
	}

	switch commonModel.Role {
	case valueobject.RoleStudent:
		studentModel := &StudentModel{}
		err := bson.Unmarshal(raw, studentModel)
		if err != nil {
			return err
		}

	case valueobject.RoleTeacher:
		teacherModel := &TeacherModel{}
		err := bson.Unmarshal(raw, teacherModel)
		if err != nil {
			return err
		}

	default:
	}
}
