package tests

import (
	"errors"
	"fmt"
	"go_server/helper"
	"regexp"
	"testing"
)

func TestGetCurrentDateTime(t *testing.T) {
	resultDate := helper.GetCurrentDateTime()

	// Regular expression to match the "dd-mm-yyyy hh:mm:ss" format.
	re := regexp.MustCompile(`^\d{2}-\d{2}-\d{4} \d{2}:\d{2}:\d{2}$`)

	match := re.MatchString(resultDate)

	if !match {
		t.Errorf("GetCurrentDateTime() = %v, want match", resultDate)
	}

	fmt.Println(resultDate)
}

func TestValidateEmailAddressTrue(t *testing.T) {
	isEmail, err := helper.ValidateEmailAddress("test@test.com")
	if err != nil {
		t.Fatalf("ValidateEmailAddress returned an error: %v", err)
	}

	if !isEmail {
		t.Errorf("ValidateEmailAddress(\"test@test.com\") = %v, want true", isEmail)
	}
}

func TestValidateEmailAddressFalse(t *testing.T) {
	isEmail, err := helper.ValidateEmailAddress("invalid email format")
	if err != nil {
		t.Fatalf("ValidateEmailAddress returned an error: %v", err)
	}

	if isEmail {
		t.Errorf("ValidateEmailAddress(\"invalid_email\") = %v, want false", isEmail)
	}
}

func TestValidateEmailAddressEdgeCases(t *testing.T) {
	testCases := []struct {
		email string
		want  bool
	}{
		{"", false},
		{"@", false},
		{"test@", false},
		{"@test.com", false},
		{"test@.com", false},
		{"test.test@test.com", true},
		{"test+filter@test.com", true},
		{"test.test@test.co.uk", true},
		// Add more edge cases as needed
	}

	for _, tc := range testCases {
		isEmail, err := helper.ValidateEmailAddress(tc.email)
		if errors.Is(err, fmt.Errorf("invalid email format")) {
			t.Fatalf("ValidateEmailAddress returned an error: %v", err)
		}
		if isEmail != tc.want {
			t.Errorf("ValidateEmailAddress(%q) = %v, want %v", tc.email, isEmail, tc.want)
		}
	}
}
