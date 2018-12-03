<!DOCTYPE html>
<html>

<head>
  {{template "header" "系统安装"}}
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
      <a href="#">系统安装</a>
    </div>
    <div class="register-box-body box">
      <form action="/" method="post">
        <div class="row">
          <div class="col-md-6">
            <div class="form-group has-feedback">
              <select name="type" class="form-control"></select>
            </div>
            <div class="form-group has-feedback">
              <input type="text" name="host" class="form-control" placeholder="数据库主机" value="localhost" required>
              <span class="fa fa-ioxhost form-control-feedback"></span>
            </div>
            <div class="form-group has-feedback">
              <input type="text" name="port" class="form-control" placeholder="数据库端口" data-rule="{'digits':true}"
                required>
              <span class="fa fa-server form-control-feedback"></span>
            </div>
            <div class="form-group has-feedback">
              <input type="text" name="user" class="form-control" placeholder="数据库用户名" required>
              <span class="fa fa-terminal form-control-feedback"></span>
            </div>
            <div class="form-group has-feedback">
              <input type="password" name="pass" class="form-control" placeholder="数据库密码">
              <span class="fa fa-key form-control-feedback"></span>
            </div>
            <div class="form-group has-feedback">
              <input type="text" name="name" class="form-control" placeholder="数据库名称" value="gocms" required>
              <span class="fa fa-font form-control-feedback"></span>
            </div>
          </div>
          <div class="col-md-6">
            <p class="login-box-msg">配置管理员账户</p>
            <div class="form-group has-feedback">
              <input type="email" name="email" class="form-control" placeholder="请输入邮箱" data-rule="{'maxlength':255}"
                required>
              <span class="fa fa-envelope form-control-feedback"></span>
            </div>
            <div class="form-group has-feedback">
              <input type="password" name="password" class="form-control" placeholder="请输入密码" id="register_password"
                data-rule="{'minlength':6}" required>
              <span class="fa fa-lock form-control-feedback"></span>
            </div>
            <div class="form-group has-feedback">
              <input type="password" class="form-control" placeholder="请重新输入密码" data-rule="{'equalTo':'#register_password'}"
                data-message="{'equalTo':'两次输入的密码不一致'}" required>
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
  <script type="text/javascript">
    var options = {
      'MySQL': {
        'port': 3306,
        'user': 'root'
      },
      'Postgres': {
        'port': 5432,
        'user': 'postgres'
      }
    };
    $(document).ready(function () {
      var select = $('select[name="type"]');
      $.each(options, function (item) {
        $('<option/>').text(item).appendTo(select);
      })
      select.on('change', function () {
        var type = $(this).val();
        $.each(options[type], function (i, e) {
          $('input[name="' + i + '"]').val(e);
        })
      }).change();
    })
  </script>
</body>

</html>