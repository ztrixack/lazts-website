package models

type Book struct {
	Number      string   `json:"number"`
	Title       string   `json:"title"`
	Subtitle    string   `json:"subtitle"`
	Description string   `json:"description"`
	Authors     []string `json:"authors"`
	Translator  string   `json:"translator"`
	Publisher   string   `json:"publisher"`
	Catalog     string   `json:"catalog"`
	Status      string   `json:"status"`
	Review      string   `json:"review"`
	Cover       string   `json:"cover"`
}

type BookStats struct {
	Total     int `json:"total"`
	Completed int `json:"completed"`
	Reading   int `json:"reading"`
	Unread    int `json:"unread"`
}
