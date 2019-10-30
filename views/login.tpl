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
  <script src="//cdnjs.cloudflare.com/ajax/libs/jquery/1.12.4/jquery.min.js"></script>
  <script src="//cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/3.4.1/js/bootstrap.min.js"></script>
  <script src="//cdnjs.cloudflare.com/ajax/libs/jquery.form/4.2.2/jquery.form.min.js"></script>
  <script src="//cdnjs.cloudflare.com/ajax/libs/jquery-validate/1.19.0/jquery.validate.min.js"></script>
  <script src="//cdnjs.cloudflare.com/ajax/libs/blueimp-md5/2.10.0/js/md5.min.js"></script>
  <script src="//cdnjs.cloudflare.com/ajax/libs/jquery-backstretch/2.0.4/jquery.backstretch.min.js"></script>
  <script src="/static/js/global.js?v{{version}}" type="text/javascript"></script>
  <script type="text/javascript">
    $(document).ready(function () {
      $('img.img-responsive').click(function (e) {
        var src = e.target.src;
        e.target.src = src.substr(0, src.indexOf('?') + 1) +
          'reload=' + new Date().getTime();
      });
      $.backstretch('/bingpic');
      // 页面表单
      Admin.validate('form:has([name="pass"])', {
        beforeSubmit: function (arr) {
          for (var i = 0; i < arr.length; i++)
            if (arr[i].name == 'pass')
              arr[i].value = md5(arr[i].value);
        },
        success: function (resp) {
          if (resp.code == 200) {
            window.location = resp.data;
            return;
          }
          if (typeof resp.data === 'string') {
            $('input[name="id"]').val(resp.data);
            $('img.img-responsive').attr('src', '/captcha/' + resp.data + '.png?')
          }
          Admin.alert({container: $('form'), type: 'danger', message: resp.msg });
        }
      });
    })
  </script>
</body>

</html>