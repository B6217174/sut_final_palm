package entity

import (
	"testing"
	//"time"

	"gorm.io/gorm"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

//entyty

type Disinfection struct {
	gorm.Model

	Amount 		int 			`valid:"required~Not Zero, range(1|99)~Not Zero"`
	Note		string			`valid:"required~Not Blank"`
	SutID		string			`gorm:"uniqueIndex" valid:"matches(^[BMD]\\d{7}$)~Not SutID"`
	Url 		string			`gorm:"uniqueIndex" valid:"url~Not URL"`
}

// test 

func TestCorrectAll(t *testing.T){  //ถูกหมด
	g := NewGomegaWithT(t)

	dis := Disinfection{
		Amount: 10,
		Note: "พรุ่งนี้สอบได้ ครึ่งนึงก็ยังดี",
		Url: "https://www.facebook.com/",
	}

	ok, err := govalidator.ValidateStruct(dis)

	g.Expect(ok).To(BeTrue())
	g.Expect(err).To(BeNil())
}
func TestAmountNotZoro(t *testing.T){
	g := NewGomegaWithT(t)

	Num := []int { 0,100,}

	for _, num := range Num {

		did := Disinfection{
			Amount: num, //ผิด
			Note: "no",
			Url: "https://www.facebook.com/",
			//SutP:	"B6217174",
		}

		ok, err := govalidator.ValidateStruct(did)

		g.Expect(ok).NotTo(BeTrue())  // OK -> NOT TRUE

		g.Expect(err).NotTo(BeNil())  // ERR -> NOT NULL

		g.Expect(err.Error()).To(Equal("Not Zero"))
	}
}

func TestNoteNotBlack(t *testing.T) {
	g := NewGomegaWithT(t)

	dis := Disinfection {
		Amount: 10,
		Note: "",  //ผิด
		Url: "https://www.facebook.com/",
	}

	ok, err := govalidator.ValidateStruct(dis)

	g.Expect(ok).NotTo(BeTrue())
	g.Expect(err).NotTo(BeNil())
	g.Expect(err.Error()).To(Equal("Not Blank"))
}

func TestUrlNotTrue(t *testing.T) {
	g := NewGomegaWithT(t)

	dis := Disinfection {
		Amount: 10,
		Note: "no",
		Url: "face",
	}

	ok, err := govalidator.ValidateStruct(dis)

	g.Expect(ok).NotTo(BeTrue())
	g.Expect(err).NotTo(BeNil())
	g.Expect(err.Error()).To(Equal("Not URL"))
}

func TestSutId(t * testing.T) {
	g := NewGomegaWithT(t)

	sut := []string {"A6217174", "B62171670"}

	for _, sutID := range sut {

		dis := Disinfection {
			Amount: 10,
			Note: "no",
			Url: "https://www.facebook.com/",
			SutID: sutID,
		}

		ok, err := govalidator.ValidateStruct(dis)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Not SutID"))
	}
}


