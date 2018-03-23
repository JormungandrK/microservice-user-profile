// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "user-profile": Application User Types
//
// Command:
// $ goagen
// --design=github.com/Microkubes/microservice-user-profile/design
// --out=$(GOPATH)/src/github.com/Microkubes/microservice-user-profile
// --version=v1.3.0

package client

import (
	"github.com/goadesign/goa"
)

// UserProfile data
type userProfilePayload struct {
	// Email of user
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// Full name of the user
	FullName *string `form:"fullName,omitempty" json:"fullName,omitempty" xml:"fullName,omitempty"`
}

// Validate validates the userProfilePayload type instance.
func (ut *userProfilePayload) Validate() (err error) {
	if ut.FullName == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "fullName"))
	}
	if ut.Email == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "email"))
	}
	if ut.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *ut.Email); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`request.email`, *ut.Email, goa.FormatEmail, err2))
		}
	}
	return
}

// Publicize creates UserProfilePayload from userProfilePayload
func (ut *userProfilePayload) Publicize() *UserProfilePayload {
	var pub UserProfilePayload
	if ut.Email != nil {
		pub.Email = *ut.Email
	}
	if ut.FullName != nil {
		pub.FullName = *ut.FullName
	}
	return &pub
}

// UserProfile data
type UserProfilePayload struct {
	// Email of user
	Email string `form:"email" json:"email" xml:"email"`
	// Full name of the user
	FullName string `form:"fullName" json:"fullName" xml:"fullName"`
}

// Validate validates the UserProfilePayload type instance.
func (ut *UserProfilePayload) Validate() (err error) {
	if ut.FullName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "fullName"))
	}
	if ut.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "email"))
	}
	if err2 := goa.ValidateFormat(goa.FormatEmail, ut.Email); err2 != nil {
		err = goa.MergeErrors(err, goa.InvalidFormatError(`type.email`, ut.Email, goa.FormatEmail, err2))
	}
	return
}
