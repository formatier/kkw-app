package valueobject

import "net/mail"

type StudentCredentialType int8

const (
	StudentCredentialTypeEmail StudentCredentialType = iota
	StudentCredentialTypeSchoolId
	StudentCredentialTypeUsername
)

type StudentCredentialsUnion string

func (u StudentCredentialsUnion) read() (StudentCredentialType, string) {
	if email, isEmail := u.readAsEmail(); isEmail {
		return StudentCredentialTypeEmail, email
	} else if schoolId, isSchoolId := u.readAsSchoolId(); isSchoolId {
		return StudentCredentialTypeSchoolId, schoolId
	} else {
		return StudentCredentialTypeUsername, string(u)
	}
}

func (u StudentCredentialsUnion) readAsEmail() (string, bool) {
	address, err := mail.ParseAddress(string(u))
	if err != nil {
		return "", false
	}

	return address.String(), true
}

func (u StudentCredentialsUnion) readAsSchoolId() (string, bool) {
	if len(u) == 5 {
		return string(u), true
	} else {
		return "", false
	}
}

type StudentLevel uint8

func NewStudentLevel(val uint8) (StudentLevel, bool) {
	if 0 < val && val <= 6 {
		return StudentLevel(val), true
	} else {
		return StudentLevel(0), true
	}
}
