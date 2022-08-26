package shared

const (
	FyneAppTOMLFileName = "fyneapp.toml"
)

type FyneAppTOMLData struct {
	WebSiteURL string // "https://github.com/josephbudd/okp"
	AppName    string // "OKP"
	AppID      string // "com.example.kickfyne"
}

var FyneAppTOML = `# Website = "{{ .WebSiteURL }}"

[Details]
# Icon = "Icon.png"
Name = "{{ .AppName }}"
ID = "{{ .AppID }}"
Version = "0.0.1"
Build = 1

`
