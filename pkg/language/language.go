package language

// Language is a struct that holds the long and short name of a language
type Language struct {
	Long  string
	Short string
}

var (
	Arabic  = Language{Long: "العربية", Short: "ar"}
	English = Language{Long: "English", Short: "en"}
	Russian = Language{Long: "Русский", Short: "ru"}
	Tatar   = Language{Long: "Татарча", Short: "tt"}
	Uzbek   = Language{Long: "O'zbekcha", Short: "uz"}
	Turkmen = Language{Long: "Türkmençe", Short: "tk"}
)

// IsValidLang returns true if the given language is valid
func IsValidLang(l string) bool {
	switch l {
	case Arabic.Short, English.Short, Russian.Short, Tatar.Short, Uzbek.Short, Turkmen.Short:
		return true
	}
	return false
}

// AvaliableLanguages returns all the avaliable languages for the application
func AvaliableLanguages() []Language {
	return []Language{Arabic, English, Russian, Tatar, Uzbek, Turkmen}
}

// DefaultLang returns the default language for the application
func DefaultLang() Language {
	return English
}
