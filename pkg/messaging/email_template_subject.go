package messaging

// EmailTemplateSubject stores email subject templates.
var EmailTemplateSubject = map[string]string{
	"en/registration_confirmation": `Registration Confirmation Required`,
	"en/registration_ready":        `Review User Registration`,
	"en/registration_verdict": `{{- if eq .verdict "approved" -}}
User Registration Approved
{{- else -}}
User Registration Declined
{{- end -}}`,
}
