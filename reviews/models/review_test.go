package models

import "testing"

func NewReview(stars int, comment string) *CreateReviewCMD {
	return &CreateReviewCMD{
		Stars:   stars,
		Comment: comment,
	}
}
func Test_withCorrectParams(t *testing.T) {
	r := NewReview(4, "El iphone X me encanto!")

	err := r.validate()

	if err != nil {
		t.Error("the validation will not pass")
		t.Fail()
	}
}

func Test_shouldFailWithWrongStars(t *testing.T) {
	r := NewReview(8, "Great Phone!")

	err := r.validate()

	if err == nil {
		t.Error("should fail with 5 stars")
		t.Fail()
	}
}

func Test_shouldFailWithWrongComments(t *testing.T) {

}
