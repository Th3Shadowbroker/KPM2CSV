package kpm

type Website struct {
	Group     string
	Name      string
	URL       string
	LoginName string
	Login     string
	Password  string
	Comment   string
}

func (w *Website) AsSlice() []string {
	return []string{
		"Kaspersky Password Manager",
		w.Name,
		w.Login,
		w.Password,
		w.URL,
		w.Comment,
	}
}
