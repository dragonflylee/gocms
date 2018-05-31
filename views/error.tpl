<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <title>{{.code}}</title>
    <meta content="noarchive" name="robots">
    <link href="/static/css/error.min.css?v=20180422" rel="stylesheet">  
</head>
<body>
  <div class="container">
    <input type="checkbox" id="switch"/>
    <div class="ellipse"></div>
    <div class="ray"></div>
    <div class="head"></div>
    <div class="neck"></div>
    <div class="body">
      <label for="switch"></label>
    </div>
  </div>
  <div class="container">
    <div class="msg msg_1">{{.code}}</div>
    <div class="msg msg_2">{{.msg}}</div>
  </div>
</body>