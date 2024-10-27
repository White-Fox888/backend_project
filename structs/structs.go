package structs

import (
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	jwt.RegisteredClaims
	Login string `json:"login"`
}

type Identification struct {
	Id       int    `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Token struct {
	Token string `json:"token"`
}

type Grant struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	SourceURL    string `json:"source_url"`
	FilterValues struct {
		CuttingOffCriteria []int `json:"cutting_off_criteria"`
		ProjectDirection   []int `json:"project_direction"`
		Amount             int   `json:"amount"`
		LegalForm          []int `json:"legal_form"`
		Age                int   `json:"age"`
	} `json:"filter_values"`
}

type FilterMapping struct {
	Age struct {
		Title   string `json:"title"`
		Mapping struct {
		} `json:"mapping"`
	} `json:"age"`
	ProjectDirection struct {
		Title   string `json:"title"`
		Mapping struct {
			Num0 struct {
				Title string `json:"title"`
			} `json:"0"`
			Num1 struct {
				Title string `json:"title"`
			} `json:"1"`
			Num2 struct {
				Title string `json:"title"`
			} `json:"2"`
			Num3 struct {
				Title string `json:"title"`
			} `json:"3"`
		} `json:"mapping"`
	} `json:"project_direction"`
	LegalForm struct {
		Title   string `json:"title"`
		Mapping struct {
			Num0 struct {
				Title string `json:"title"`
			} `json:"0"`
			Num1 struct {
				Title string `json:"title"`
			} `json:"1"`
			Num2 struct {
				Title string `json:"title"`
			} `json:"2"`
		} `json:"mapping"`
	} `json:"legal_form"`
	CuttingOffCriteria struct {
		Title   string `json:"title"`
		Mapping struct {
			Num0 struct {
				Title string `json:"title"`
			} `json:"0"`
			Num1 struct {
				Title string `json:"title"`
			} `json:"1"`
			Num2 struct {
				Title string `json:"title"`
			} `json:"2"`
			Num3 struct {
				Title string `json:"title"`
			} `json:"3"`
		} `json:"mapping"`
	} `json:"cutting_off_criteria"`
	Amount struct {
		Title   string `json:"title"`
		Mapping struct {
		} `json:"mapping"`
	} `json:"amount"`
}

type MetaPages struct {
	CurrentPage int `json:"current_page"`
	TotalPages  int `json:"total_pages"`
}

type FilterOrder []string

type DataFilters struct {
	Data struct {
		ProjectDirection   []int `json:"project_direction"`
		Amount             int   `json:"amount"`
		LegalForm          []int `json:"legal_form"`
		Age                int   `json:"age"`
		CuttingOffCriteria []int `json:"cutting_off_criteria"`
	} `json:"data"`
}
