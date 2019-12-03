package tmpl

var ConventionalCommit = CommitTemplate{
	Questions: []CommitQuestion{
		CommitQuestion{
			PromptType: "conventional",
			Mandatory:  true,
			ValueCode:  "type",
		},
		CommitQuestion{
			PromptType: "text",
			Mandatory:  true,
			Prompt:     "Enter the commit description, with JIRA number at end",
			ValueCode:  "description",
		},
		// TODO: Ask if this is a breaking change
		CommitQuestion{
			PromptType: "text",
			Mandatory:  false,
			Prompt:     "Enter the (optional) commit body",
			ValueCode:  "body",
		},
		CommitQuestion{
			PromptType: "text",
			Mandatory:  false,
			Prompt:     "Enter the (optional) commit footer",
			ValueCode:  "footer",
		},
	},
	Command: "git",
	CommandArgs: []string{
		"commit",
		"-m",
		"{{.type}}: {{.description}}",
		"{{with .body}}-m{{end}}",
		"{{.body}}",
		"{{with .footer}}-m{{end}}",
		"{{.footer}}",
	},
}
