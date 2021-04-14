package content

const ChangePasswordRequestEmailContent = `<!DOCTYPE HTML PULBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
<html>
<head>
    <meta http-equiv="content-type" content="text/html" ; charset="ISO-8859-1">
</head>
<body>
<p>Dear {{.Name}},<br>
<p>You requested for reset your password.</p>
<p>For reset your password follow this <a href="{{.URL}}">link</a>.</p><br>
<b>
With gratitude,<br>
The Novels Team
</b>
</body>
</html>`
