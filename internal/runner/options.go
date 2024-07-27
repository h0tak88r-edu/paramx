package runner


// Options represents the configuration options for the runner.
type Options struct {
	URLs         []string // URLs is a list of target URLs.
	TempletesPath string   // TempletesPath is the path to the templates directory.
	BugType      string   // BugType is the type of bug to be injected.
	FileInput    string   // FileInput is the path to the input file.
	ReplaceWith  string   // ReplaceWith is the string to replace the bug with.
}