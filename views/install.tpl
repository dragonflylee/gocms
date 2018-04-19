<!DOCTYPE html>
<html>
<head>
  {{template "header" "系统初始化"}}
  <style type="text/css">
    footer {
      margin-left: 0px !important;
      position: absolute;
      width: 100%;
      bottom: 0;
    }
  </style>
</head>
<body class="hold-transition register-page">
<div class="register-box" style="width: 600px">
  <div class="register-logo">
    <a href="#"><b>Go</b>CMS</a>
  </div>
  <div class="register-box-body box">
    <form action="/" method="post">
      <div class="row">
        <div class="col-md-6">
          <p class="login-box-msg">配置数据库参数</p>
          <div class="form-group has-feedback">
            <input type="text" name="host" class="form-control" placeholder="数据库主机" value="127.0.0.1:3306">
            <span class="fa fa-ioxhost form-control-feedback"></span>
          </div>
          <div class="form-group has-feedback">
            <input type="text" name="user" class="form-control" placeholder="数据库用户名" value="root">
            <span class="fa fa-terminal form-control-feedback"></span>
          </div>
          <div class="form-group has-feedback">
            <input type="text" name="pass" class="form-control" placeholder="数据库密码">
            <span class="fa fa-key form-control-feedback"></span>
          </div>
          <div class="form-group has-feedback">
            <input type="text" name="name" class="form-control" placeholder="数据库名称" value="gocms">
            <span class="fa fa-font form-control-feedback"></span>
          </div>
        </div>
        <div class="col-md-6">
          <p class="login-box-msg">配置管理员账户</p>
          <div class="form-group has-feedback">
            <input type="text" name="username" class="form-control" placeholder="请输入用户名" value="admin" data-rule="{'maxlength':32}" required>
            <span class="fa fa-font form-control-feedback"></span>
          </div>
          <div class="form-group has-feedback">
            <input type="email" name="email" class="form-control" placeholder="请输入管理员邮箱" data-rule="{'maxlength':255}" required>
            <span class="fa fa-envelope form-control-feedback"></span>
          </div>
          <div class="form-group has-feedback">
            <input type="password" name="password" class="form-control" id="register_password" placeholder="请输入管理员密码" data-rule="{'minlength':6}" required>
            <span class="fa fa-lock form-control-feedback"></span>
          </div>
          <div class="form-group has-feedback">
            <input type="password" name="rpassword" class="form-control" placeholder="请重新输入密码" data-rule="{'equalTo':'#register_password'}" data-message="{'equalTo':'两次输入的密码不一致'}" required>
            <span class="fa fa-check form-control-feedback"></span>
          </div>
          <div class="row">
            <div class="col-xs-4 pull-right">
              <button type="submit" class="btn btn-primary btn-block btn-flat">安装</button>
            </div>
          </div>
        </div>
      </div>
    </form>
  </div>
</div>
{{template "footer"}}
</body>
</html>
