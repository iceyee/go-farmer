package fweb

import (
	"net/http"
	"sort"
	"strconv"
	"strings"
	//
)

const apiTemplate string = `
    <table class="table table-dark mb-5">
        <thead></thead>
        <tbody>
            <tr>
                <th scope="row">Description</th>
                <td>$$Description</td>
            </tr>
            <tr>
                <th scope="row">Url</th>
                <td>$$Url</td>
            </tr>
            <tr>
                <th scope="row">Method</th>
                <td>$$Method</td>
            </tr>
            <tr>
                <th scope="row">Parameters</th>
                <td>$$Parameters</td>
            </tr>
            <tr>
                <th scope="row">Response</th>
                <td>$$Response</td>
            </tr>
            <tr>
                <th scope="row">Remarks</th>
                <td>$$Remarks</td>
            </tr>
        </tbody>
    </table>
    <br class="mb-5">
`

const apiTemplate2 string = `
<!doctype html>
<html>
<head>
    <title>Api Document</title>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, user-scalable=no, max-scale=1">
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.0.0-beta2/dist/css/bootstrap.min.css"
      rel="stylesheet" integrity="sha384-BmbxuPwQa2lc/FVzBcNJ7UAyJxM6wuqIj61tLrc4wSX0szH/Ev+nYRRuWlolflfl"
      crossorigin="anonymous">
    <style></style>
</head>
<body class="container py-5">
    $$
    <script src="https://code.jquery.com/jquery-3.6.0.min.js"></script>
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.0.0-beta2/dist/js/bootstrap.bundle.min.js"
          integrity="sha384-b5kHyXgcpbZJO/tY9Ul7kGkf1S0CWuKcCD38l8YkeH8z8QjE0GmW1gYU5S9FOnJ0"
          crossorigin="anonymous"></script>
</body>
</html>
`

func init() {
	RegistryController(new(defaultController))
	return
}

type defaultController struct {
}

func (*defaultController) Api_F() string {
	return `
    @Url /0/api

    @MapTo Api
    `
}

var document string

func (*defaultController) Api(session *Session, w http.ResponseWriter, r *http.Request) {
	if "" == document {
		var a001 map[string]t184
		a001 = make(map[string]t184, 0xfff)
		var a002 []string
		a002 = make([]string, 0, 0xfff)
		for _, x := range x251 {
			a001[x.SortKey] = x
		}
		delete(a001, "/0/api")
		delete(a001, "/0/status")
		for x, _ := range a001 {
			a002 = append(a002, x)
		}
		sort.Strings(a002)
		for _, key := range a002 {
			var value t184
			value = a001[key]
			var a003 string
			a003 = apiTemplate
			a003 = strings.Replace(a003, "$$Description", value.Description, -1)
			a003 = strings.Replace(a003, "$$Url", value.Url, -1)
			a003 = strings.Replace(a003, "$$Method", value.Method, -1)
			a003 = strings.Replace(a003, "$$Response", value.Response, -1)
			a003 = strings.Replace(a003, "$$Remarks", value.Remarks, -1)
			var a004 string
			for _, x := range value.Parameters {
				a004 += x.Name
				a004 += ", "
				a004 += x.Type
				if x.Required {
					a004 += ", 必填"
				} else {
					a004 += ", 可选"
				}
				if "" != x.Default {
					a004 += ", 默认"
					a004 += x.Default
				}
				if nil != x.Min {
					a004 += ", 最小值"
					a004 += strconv.FormatFloat(*x.Min, 'f', -1, 64)
				}
				if nil != x.Max {
					a004 += ", 最大值"
					a004 += strconv.FormatFloat(*x.Max, 'f', -1, 64)
				}
				if "" != x.Not {
					a004 += ", 不能是"
					a004 += x.Not
				}
				if "" != x.Description {
					a004 += ", "
					a004 += x.Description
				}
				a004 += "<br>"
			}
			a003 = strings.Replace(a003, "$$Parameters", a004, -1)
			document += a003
		}
		document = strings.Replace(apiTemplate2, "$$", document, -1)
	}
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(document))
	return
}

func (*defaultController) Status_F() string {
	return `
    @Url            /0/status

    @MapTo          Status
    `
}

func (*defaultController) Status(session *Session, w http.ResponseWriter, r *http.Request) {
	R200(w)
	return
}
