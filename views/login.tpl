<!DOCTYPE html>
<html>
<head>
  {{template "header" "后台登录"}}
  <style type="text/css">
    footer {
      margin-left: 0px !important;
      position: absolute;
      width: 100%;
      bottom: 0;
    }
  </style>
</head>
<body class="hold-transition login-page">
<div class="login-box">
  <div class="login-logo">
    <a href="#"><b>Go</b>CMS</a>
  </div>
  <div class="login-box-body box">
    <p class="login-box-msg">登录系统后台</p>
    <form action="/login?refer={{urlquery .}}" method="post">
      <div class="form-group has-feedback">
        <input name="username" type="email" class="form-control" autocomplete="off" placeholder="请输入管理员邮箱" data-message="{'required':'登录名称不能为空'}" required>
        <span class="glyphicon glyphicon-user form-control-feedback"></span>
      </div>
      <div class="form-group has-feedback">
        <input name="password" type="password" class="form-control" autocomplete="off" placeholder="请输入密码" data-message="{'required':'密码不能为空'}" required>
        <span class="glyphicon glyphicon-lock form-control-feedback"></span>
      </div>
      <div class="row">
        <div class="col-xs-8">
          <div class="checkbox icheck">
            <label>
              <input type="checkbox" name="remember" value="1"> 记住我
            </label>
          </div>
        </div>
        <div class="col-xs-4">
          <button type="submit" class="btn btn-primary btn-block btn-flat">登录</button>
        </div>
      </div>
    </form>
  </div>
</div>
{{template "footer"}}
</body>
</html>
