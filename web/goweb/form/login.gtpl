<html>
<head>
<title></title>
</head>
<body>
<input type="checkbox" name="interest" value="football">Football
<input type="checkbox" name="interest" value="basketball">Basketball
<input type="checkbox" name="interest" value="tennis">Tennis
<form action="/login" method="post">
用户名:<input type="text" name="username">
密码:<input type="password" name="password">
<input type="hidden" nam="token" value="{{.}}">
<input type="submit" value="登陆">
</form>
</body>
</html>
