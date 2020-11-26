<!DOCTYPE html>
<html>
<head>
  {{- template "header" "Login"}}
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
          <input name="username" type="email" class="form-control" autocomplete="off" placeholder="请输入管理员邮箱" data-msg-required="登录名称不能为空" required>
          <span class="glyphicon glyphicon-user form-control-feedback"></span>
        </div>
        <div class="form-group has-feedback">
          <input name="password" type="password" class="form-control" autocomplete="off" placeholder="请输入密码" data-msg-required="密码不能为空" required>
          <span class="glyphicon glyphicon-lock form-control-feedback"></span>
        </div>
        {{- if .Key}}
        <div class="form-group">
          <div class="g-recaptcha" data-callback="onsubmit" data-sitekey="{{.Key}}"></div>
        </div>
        {{- end}}
        <div class="row">
          <div class="col-xs-4 pull-right">
            <button type="submit" class="btn btn-primary btn-block btn-flat">登录</button>
          </div>
        </div>
      </form>
    </div>
  </div>
  {{- template "footer" "login"}}
  <script src="//cdnjs.cloudflare.com/ajax/libs/jquery-backstretch/2.0.4/jquery.backstretch.min.js"></script>
  {{- if .Key}}
  <script src='//recaptcha.net/recaptcha/api.js'></script>
  {{- end}}
  <script type="text/javascript">
    var onsubmit = function () {
      $('form').submit();
    }
    $(document).on('click', '.btn-verify', function (e) {
      var timeout = 90, btn = $(e.target).
        attr('disabled', true).text(timeout + ' s'),
        count = setInterval(function () {
          if (--timeout > 0)
            return btn.text(timeout + ' s')
          clearInterval(count);
          btn.text('点击获取').attr('disabled', false);
        }, 1000);
      // send code
      var form = btn.closest('form').ajaxSubmit({
        url: btn.data('href'),
        dataType: 'json',
        success: function (resp) {
          if (resp.code != 200) return Admin.alert({
            container: form, type: 'danger', message: resp.msg
          });
        }
      })
    })
    $(document).ready(function () {
      $.backstretch('/bingpic');
    })
  </script>
</body>
</html>

{{define "login-verify"}}
<div class="form-group has-feedback">
  {{- if .Phone}}
  <p class="form-control-static">手机号码: {{.Phone}}</p>
  {{- else}}
  <p class="form-control-static">电子邮箱: {{.Email}}</p>
  {{- end}}
  <input name="token" type="hidden" value="{{.Token}}">
</div>
<div class="form-group has-feedback">
  <div class="input-group">
    <input name="code" type="text" class="form-control" placeholder="请输入验证码" data-msg-required="请输入验证码" data-target=".input-group" required>
    <span class="input-group-btn">
      <button class="btn btn-default btn-verify" data-href="/verify?type=login" type="button">点击获取</button>
    </span>
  </div>
</div>
<div class="row">
  <div class="col-xs-4 pull-right">
    <button type="submit" class="btn btn-primary btn-block btn-flat" disabled>Continue</button>
  </div>
</div>
{{- end}}