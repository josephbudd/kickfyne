package root

const (
	backendWorkSpaceFileName  = "backend.code-workspace"
	frontendWorkSpaceFileName = "frontend.code-workspace"

	backendWorkSpaceTemplate = `{
	"folders": [
		{
			"path": "./backend"
		},
		{
			"path": "./shared"
		},
	],
	"settings": {
		"go":{
			"installDependenciesWhenBuilding": false,
		},
	},
}
`

	frontendWorkSpaceTemplate = `{
	"folders": [
		{
			"path": "./frontend"
		},
		{
			"path": "./shared"
		},
	],
	"settings": {
		"go":{
			"installDependenciesWhenBuilding": false,
		},
	},
}
`
)
