package utils

import (
	"github.com/russross/blackfriday"
	"html/template"
	"regexp"
	"strings"
)

func Html2Str(html string) string {
	src := string(html)

	// Lowercase HTML tags
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)

	// Remove styles
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")

	// Remove scripts
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")

	// Strip all HTML tags
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")

	// Remove continuous `\n`
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")

	return strings.TrimSpace(src)
}

func SubString(str string, begin, length int) (substr string) {
	rs := []rune(str)
	lth := len(rs)
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length
	if end > lth {
		end = lth
	}
	return string(rs[begin:end])
}

func Html2Excerpt(html string, length int) string {
	return SubString(Html2Str(html), 0, length)
}

func Markdown2Html(text string) string {
	return string(blackfriday.MarkdownCommon([]byte(text)))
}

func Markdown2HtmlTemplate(text string) template.HTML {
	return template.HTML(Markdown2Html(text))
}
