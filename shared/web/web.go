package web

const RespGreeting = `
<html>
	<head>
		<title>Hello, {{.Name}}
	</head>
	<body>
	<p>
		<b>Hi, {{.Name}}</b>
	</p>
	<p>
		Today is {{.Date}}. The current time is {{.Time}}
	</p>
	</body>
</htm>
`
