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
          <input name="user" type="email" class="form-control" autocomplete="off" placeholder="请输入管理员邮箱" data-rule="{'messages':{'required':'登录名称不能为空'}}" required>
          <span class="glyphicon glyphicon-user form-control-feedback"></span>
        </div>
        <div class="form-group has-feedback">
          <input name="pass" type="password" class="form-control" autocomplete="off" placeholder="请输入密码" data-rule="{'messages':{'required':'密码不能为空'}}" required>
          <span class="glyphicon glyphicon-lock form-control-feedback"></span>
        </div>
        <div class="row">
          <div class="col-xs-6 form-group has-feedback">
            <input name="code" type="text" class="form-control" autocomplete="off" placeholder="验证码" data-rule="{'maxlength':6,'digits':true,'messages':{'required':'验证码不能为空'}}" required>
          </div>
          <div class="col-xs-4 pull-right">
            <input type="hidden" name="id" value="{{.Captcha}}">
            <button type="submit" class="btn btn-primary btn-block btn-flat">登录</button>
          </div>
        </div>
        <img class="img-responsive" src="/captcha/{{.Captcha}}.png?">
      </form>
    </div>
  </div>
</body>
<script src="/static/js/login.js?v{{version}}" type="text/javascript"></script>
</html>