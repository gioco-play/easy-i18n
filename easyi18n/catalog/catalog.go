package catalog

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// init
func init() {
	initEn(language.Make("en"))
	initZhHans(language.Make("zh-Hans"))
	initZhHant(language.Make("zh-Hant"))
}
// initEn will init en support.
func initEn(tag language.Tag) {
	message.SetString(tag, "%s extract [path] [outfile]", "%s extract [path] [outfile]")
	message.SetString(tag, "%s generate [path] [outfile]", "%s generate [path] [outfile]")
	message.SetString(tag, "%s update srcfile destfile", "%s update srcfile destfile")
	message.SetString(tag, "a tool for managing message translations.", "a tool for managing message translations.")
	message.SetString(tag, "destfile cannot be empty", "destfile cannot be empty")
	message.SetString(tag, "extracts strings to be translated from code", "extracts strings to be translated from code")
	message.SetString(tag, "generated go file package name", "generated go file package name")
	message.SetString(tag, "generates code to insert translated messages", "generates code to insert translated messages")
	message.SetString(tag, "merge translations and generate catalog", "merge translations and generate catalog")
	message.SetString(tag, "package name", "package name")
	message.SetString(tag, "srcfile cannot be empty", "srcfile cannot be empty")
}
// initZhHans will init zh-Hans support.
func initZhHans(tag language.Tag) {
	message.SetString(tag, "%s extract [path] [outfile]", "%s extract [path] [outfile]")
	message.SetString(tag, "%s generate [path] [outfile]", "%s generate [path] [outfile]")
	message.SetString(tag, "%s update srcfile destfile", "%s update srcfile destfile")
	message.SetString(tag, "a tool for managing message translations.", "a tool for managing message translations.")
	message.SetString(tag, "destfile cannot be empty", "destfile cannot be empty")
	message.SetString(tag, "extracts strings to be translated from code", "extracts strings to be translated from code")
	message.SetString(tag, "generated go file package name", "generated go file package name")
	message.SetString(tag, "generates code to insert translated messages", "generates code to insert translated messages")
	message.SetString(tag, "merge translations and generate catalog", "merge translations and generate catalog")
	message.SetString(tag, "package name", "package name")
	message.SetString(tag, "srcfile cannot be empty", "srcfile cannot be empty")
}
// initZhHant will init zh-Hant support.
func initZhHant(tag language.Tag) {
	message.SetString(tag, "%s extract [path] [outfile]", "%s ?????? [??????] [????????????]")
	message.SetString(tag, "%s generate [path] [outfile]", "%s ?????? [??????] [????????????]")
	message.SetString(tag, "%s update srcfile destfile", "%s ?????? ????????? ????????????")
	message.SetString(tag, "a tool for managing message translations.", "????????????????????????????????????")
	message.SetString(tag, "destfile cannot be empty", "????????????????????????")
	message.SetString(tag, "extracts strings to be translated from code", "???????????????????????????????????????")
	message.SetString(tag, "generated go file package name", "?????????go???????????????")
	message.SetString(tag, "generates code to insert translated messages", "???????????????????????????????????????")
	message.SetString(tag, "merge translations and generate catalog", "???????????????????????????")
	message.SetString(tag, "package name", "package name")
	message.SetString(tag, "srcfile cannot be empty", "?????????????????????")
}
