// This package contains the base logic for the creation and rendering of field widgets. Base widgets are defined for most input fields,
// both in classic and Bootstrap3 style; custom widgets can be defined and associated to a field, provided that they implement the
// WidgetInterface interface.
package widgets

import (
	"bytes"
	"html/template"
	"path/filepath"

	formcommon "github.com/kirves/go-form-it/common"
)

// Simple widget object that gets executed at render time.
type Widget struct {
	template *template.Template
}

// WidgetInterface defines the requirements for custom widgets.
type WidgetInterface interface {
	Render(data interface{}) string
}

// Render executes the internal template and returns the result as a template.HTML object.
func (w *Widget) Render(data interface{}) string {
	var s string
	buf := bytes.NewBufferString(s)
	w.template.ExecuteTemplate(buf, "main", data)
	return buf.String()
}

// BaseWidget creates a Widget based on style and inputType parameters, both defined in the common package.
func BaseWidget(style, inputType string) *Widget {
	var templatesPrefix = filepath.Join("templates", style)
	var urls []string = []string{"generic.tmpl"}

	switch inputType {
	case formcommon.BUTTON,
		formcommon.RESET,
		formcommon.SUBMIT:
		urls = append(urls, "button.html")
	case formcommon.CHECKBOX:
		urls = append(urls, filepath.Join("options", "checkbox.html"))
	case formcommon.TEXTAREA:
		urls = append(urls, filepath.Join("text", "textareainput.html"))
	case formcommon.SELECT:
		urls = append(urls, filepath.Join("options", "select.html"))
	case formcommon.PASSWORD:
		urls = append(urls, filepath.Join("text", "passwordinput.html"))
	case formcommon.RADIO:
		urls = append(urls, filepath.Join("options", "radiobutton.html"))
	case formcommon.TEXT:
		urls = append(urls, filepath.Join("text", "textinput.html"))
	case formcommon.RANGE:
		urls = append(urls, filepath.Join("number", "range.html"))
	case formcommon.NUMBER:
		urls = append(urls, filepath.Join("number", "number.html"))
	case formcommon.DATE:
		urls = append(urls, filepath.Join("datetime", "date.html"))
	case formcommon.DATETIME:
		urls = append(urls, filepath.Join("datetime", "datetime.html"))
	case formcommon.TIME:
		urls = append(urls, filepath.Join("datetime", "time.html"))
	case formcommon.DATETIME_LOCAL:
		urls = append(urls, filepath.Join("datetime", "datetime.html"))
	case formcommon.STATIC:
		urls = append(urls, filepath.Join("static.html"))
	case formcommon.SEARCH,
		formcommon.TEL,
		formcommon.URL,
		formcommon.WEEK,
		formcommon.COLOR,
		formcommon.EMAIL,
		formcommon.FILE,
		formcommon.HIDDEN,
		formcommon.IMAGE,
		formcommon.MONTH:
		urls = append(urls, filepath.Join(templatesPrefix, "input.html"))
	default:
		urls = append(urls, filepath.Join(templatesPrefix, "input.html"))
	}
	// resolve paths
	for i := range urls {
		urls[i] = formcommon.CreateUrl(filepath.Join(templatesPrefix, urls[i]))
	}
	templ := template.Must(template.ParseFiles(urls...))
	return &Widget{templ}
}
