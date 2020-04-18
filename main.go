package main

import (
	"bytes"
	"fmt"
	"os"
	"time"

	"github.com/danielkvist/talk"
	"gopkg.in/liderman/text-generator.v1"
)

func main() {
	tg := text_generator.New()
	template := "{Good {morning|evening|day}|Goodnight|Hello}, {friend|brother|sister}! {How are you|What's new with you|What up!}?"

	for {
		f, err := os.Create("/var/htdocs/index.html")
		if err != nil {
			panic(err)
		}
		first := `<html>
	<head>
		<title>Welcome! Hola</title>
		<link rel="stylesheet" href="http://netdna.bootstrapcdn.com/bootstrap/3.1.1/css/bootstrap.min.css"/>
	</head>
	<body>
		<div class="container">
			<div class="jumbotron">
				<h1>Welcome! Hola</h1>
				<p>Image by opsLab</p>
<pre style="background-color:#000000" align="left"><font color=#00FFFF>
`
		txt := tg.Generate(template)
		var b bytes.Buffer
		talk.Gopher(txt, &b)
		cowsayfortune := b.String()
		second := `</font></pre>
	</div>
</div>
</body>
</html>
`
		fmt.Printf("%s  Cow is saying to /var/htdocs/index.html\n", time.Now())
		html := first + cowsayfortune + second
		l, err := f.WriteString(html)
		if err != nil {
			panic(err)
		}
		fmt.Println(l, "bytes written successfully")

		defer f.Close()
		time.Sleep(5 * time.Second)
	}
}
