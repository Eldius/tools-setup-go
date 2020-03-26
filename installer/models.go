package installer

/*
ToolSpec is an unit of tool spec
*/
type ToolSpec struct {
	SetupType  string
	URL        string
	Version    string
	VersionCmd string
}
