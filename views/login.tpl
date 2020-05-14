<!DOCTYPE html>
<html>
<head>
  {{template "header" "Login"}}
  <style>
    .backstretch { opacity: .5 }
  </style>
</head>
<body class="hold-transition login-page bg-black">
  <div class="login-box">
    <div class="login-logo">
      <a href="#" class="text-gray"><b>Go</b>CMS</a>
    </div>
    <div class="login-box-body box">
      <p class="login-box-msg">登录系统后台</p>
      <form action="?refer={{urlquery .Ref}}" method="post">
        <div class="form-group has-feedback">
          <input name="username" type="email" class="form-control" autocomplete="off" placeholder="请输入管理员邮箱" data-rule="{'messages':{'required':'登录名称不能为空'}}" required>
          <span class="glyphicon glyphicon-user form-control-feedback"></span>
        </div>
        <div class="form-group has-feedback">
          <input name="password" type="password" class="form-control" autocomplete="off" placeholder="请输入密码" data-rule="{'messages':{'required':'密码不能为空'}}" required>
          <span class="glyphicon glyphicon-lock form-control-feedback"></span>
        </div>
        <div class="form-group">
          <div class="g-recaptcha" data-callback="onsubmit" data-sitekey="{{.Key}}"></div>
        </div>
        <div class="row">
          <div class="col-xs-4 pull-right">
            <button id="submit" type="submit" class="btn btn-primary btn-block btn-flat" disabled>登录</button>
          </div>
        </div>
      </form>
    </div>
  </div>
  <script src="//cdnjs.cloudflare.com/ajax/libs/jquery/1.12.4/jquery.min.js"></script>
  <script src="//cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/3.4.1/js/bootstrap.min.js"></script>
  <script src="//cdnjs.cloudflare.com/ajax/libs/jquery.form/4.2.2/jquery.form.min.js"></script>
  <script src="//cdnjs.cloudflare.com/ajax/libs/jquery-validate/1.19.0/jquery.validate.min.js"></script>
  <script src="//cdnjs.cloudflare.com/ajax/libs/blueimp-md5/2.10.0/js/md5.min.js"></script>
  <script src="//cdnjs.cloudflare.com/ajax/libs/jquery-backstretch/2.0.4/jquery.backstretch.min.js"></script>
  <script src='//recaptcha.net/recaptcha/api.js?onload=onload'></script>
  <script src="/static/js/global.js?v{{version}}" type="text/javascript"></script>
  <script type="text/javascript">
    var onsubmit = function() {
      document.getElementById('submit').disabled = false;
    }
    $(document).ready(function () {
      $.backstretch('/bingpic');
    })
  </script>
</body>
</html>