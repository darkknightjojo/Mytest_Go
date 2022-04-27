package Web

import (
	"net/http"
)

func HttpDemo() {
	http.Handle("/", &ThisHandler{})
	// 注意，这里是HandleFunc，下面是HandlerFunc
	http.HandleFunc("/hi", sayHi)
	http.Handle("/hello", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("<html><body><H1>Hello!</H1></body></html>"))
	}))

	http.ListenAndServe(":8080", nil)
}

type ThisHandler struct{}

func (m *ThisHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("<html><body><H1>ThisHandler's ServerHttp</H1></body></html>"))
}

func sayHi(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("<html><head><style type=\"text/css\">\n\t\tdiv{\n\t\t\theight: 200px;\n\t\t\twidth:200px;\n\t\t\tbackground-color: #dea46b;\n\t\t\ttext-align: center; \n\t\t\tline-height: 200px;/*文字水平居中*/\n\t\t\tmargin:auto;/*div水平居中*/\n\t\t}\n\t</style></head><body><div>Hi!</div></body></html>"))
}
