package rule

import (
	"github.com/dlclark/regexp2"

	"github.com/insidersec/insider/engine"
)

var CsharpRules []engine.Rule = []engine.Rule{

	Rule{
		ExactMatch:    regexp2.MustCompile(`(?i)\.Filter\s*=\s*.*(\+.*|String\.Format\s*\(.*\)|String\.Concat\s*\(.*\))`, regexp2.None),
		CWE:           "CWE-90",
		AverageCVSS:   6,
		Description:   "Potential LDAP Injection: LDAP filters are being built dynamically using user-controlled input. This vulnerability in code could lead to an arbitrary LDAP query execution.",
		Recomendation: "TODO: Use safe encoders such as System.DirectoryServices.Protocols.LdapEncoder.LdapFilterEncode before including user input in LDAP queries.",
	},
}
