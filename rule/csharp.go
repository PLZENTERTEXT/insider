package rule

import (
	"regexp"

	"github.com/insidersec/insider/engine"
)

var CsharpRules []engine.Rule = []engine.Rule{

	Rule{
		ExactMatch:    regexp.MustCompile(`(?i)\.Filter\s*=\s*.*(\+.*|String\.Format\s*\(.*\)|String\.Concat\s*\(.*\))`),
		CWE:           "CWE-90",
		AverageCVSS:   6,
		Description:   "Potential LDAP Injection: LDAP filters are being built dynamically using user-controlled input.",
		Recomendation: "Use safe encoders such as System.DirectoryServices.Protocols.LdapEncoder.LdapFilterEncode before including user input in LDAP queries.",
	},
}
