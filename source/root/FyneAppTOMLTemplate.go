package root

type fyneAppTOMLData struct {
	WebSiteURL string // "https://github.com/josephbudd/okp"
	AppName    string // "OKP"
	AppID      string // "com.github.josephbudd.okp"
}

var dyneAppTOMLTemplate = `# Website = "{{ .WebSiteURL }}"

[Details]
# Icon = "Icon.png"
Name = "{{ .AppName }}"
ID = "{{ .AppID }}"
Version = "0.1.0"
Build = 1

[FrontEnd]
Landing = "home"
`
