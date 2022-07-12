package basics

import (
	"encoding/json"
	"fmt"
)

type CompleteUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type omit *struct{}
type PublicUser struct {
	*CompleteUser
	Password omit `json:"password,omitempty"`
}

type BlogPost struct {
	URL   string `json:"url"`
	Title string `json:"title"`
}

type Analytics struct {
	Visitors  int `json:"visitors"`
	PageViews int `json:"page_views"`
}

func MarshalUnmarshalLearn() {
	removingAFieldExample()
	bt := combiningMultipleStruct()
	splittingIntoMultipleStruct(bt)
}

func removingAFieldExample() {
	user := CompleteUser{
		Email:    "arnob@gmail.com",
		Password: "1234",
		Name:     "Arnob",
	}

	b1, _ := json.Marshal(PublicUser{
		CompleteUser: &user,
	})
	fmt.Println(string(b1))
}

func combiningMultipleStruct() []byte {
	blogObj := BlogPost{
		URL:   "attilaolah@gmail.com",
		Title: "Attila's Blog",
	}
	anaObj := Analytics{
		Visitors:  7,
		PageViews: 45,
	}
	bt, _ := json.Marshal(struct {
		*BlogPost
		*Analytics
	}{&blogObj, &anaObj})
	fmt.Println(string(bt))
	return bt
}

func splittingIntoMultipleStruct(bt []byte) {
	var newBlogObj BlogPost
	var newAnaObj Analytics
	_ = json.Unmarshal(bt, &struct {
		*BlogPost
		*Analytics
	}{&newBlogObj, &newAnaObj})

	fmt.Println(newBlogObj, " ++ ", newAnaObj)
}
