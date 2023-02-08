package xmltargets

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/beevik/etree"
)

var XmlStr string = `<bookstore xmlns:p="urn:schemas-books-com:prices">
<books type="chinese">
	<title lang="en">all book</title>

<book category="COOKING">
  <title lang="en">Everyday Italian</title>
  <author>Giada De Laurentiis</author>
  <year>2005</year>
  <p:price>30.00</p:price>
</book>

<book category="CHILDREN">
  <title lang="en">Harry Potter</title>
  <author>J K. Rowling</author>
  <year>2005</year>
  <p:price>29.99</p:price>
</book>

<book category="WEB">
  <title lang="en">XQuery Kick Start</title>
  <author>James McGovern</author>
  <author>Per Bothner</author>
  <author>Kurt Cagle</author>
  <author>James Linn</author>
  <author>Vaidyanathan Nagarajan</author>
  <year>2003</year>
  <p:price>49.99</p:price>
</book>

<book category="WEB">
  <title lang="en">Learning XML</title>
  <author>Erik T. Ray</author>
  <year>2003</year>
  <p:price>39.95</p:price>
</book>
</books>
</bookstore>`

func XPathUnSafe(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	input := r.PostForm.Get("input")
	path := fmt.Sprintf("//books/book[@category='%s']/*", input)

	w.Header().Add("Content-Type", "application/json; charset=UTF-8")
	w.Write(FindElement(path))
}

func XPathSafe(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	input := r.PostForm.Get("input")
	input = replacerPath(input)
	path := fmt.Sprintf("//books/book[@category='%s']/*", input)
	w.Header().Add("Content-Type", "application/json; charset=UTF-8")
	w.Write(FindElement(path))
}

func FindElement(path string) []byte {
	doc := etree.NewDocument()
	doc.ReadFromString(XmlStr)
	etreePaht, err := etree.CompilePath(path)
	if err != nil {
		return []byte(fmt.Sprintf(`{"err":"%s"}`, err.Error()))
	}

	elment := doc.FindElementsPath(etreePaht)
	resultMap := map[string]interface{}{}
	for _, t := range elment {
		resultMap[t.GetPath()] = t.Text()
	}
	resultByte, err := json.Marshal(resultMap)
	if err != nil {
		return []byte(fmt.Sprintf(`{"err":"%s"}`, err.Error()))
	}
	return resultByte
}

func replacerPath(path string) string {
	replacer := strings.NewReplacer(
		"..", ".",
		"[", "",
		"[", "",
		"@", "",
		"*", "",
	)
	path = replacer.Replace(path)
	return path
}
