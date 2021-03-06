<!DOCTYPE html>
<html>
<head>
  {{- template "header" "系统安装"}}
</head>
<body class="hold-transition register-page">
  <div class="register-box">
    <div class="register-logo">
      <a href="#">系统安装</a>
    </div>
    <div class="register-box-body box">
      <form method="post" class="tab-content">
        <div class="tab-pane active" id="database">
          <p class="login-box-msg">配置数据库连接</p>
          <div class="form-group has-feedback">
            <select name="type" class="form-control"></select>
          </div>
          <div class="form-group has-feedback">
            <input type="text" name="host" class="form-control" placeholder="数据库主机" value="localhost" data-msg-required="数据库地址不能为空" required>
            <span class="fa fa-ioxhost form-control-feedback"></span>
          </div>
          <div class="form-group has-feedback">
            <input type="text" name="port" class="form-control" placeholder="数据库端口" value="5432" data-rule-digits="true" data-msg-digits="端口号必须是数字" required>
            <span class="fa fa-server form-control-feedback"></span>
          </div>
          <div class="form-group has-feedback">
            <input type="text" name="user" class="form-control" placeholder="数据库用户名" value="postgres" required>
            <span class="fa fa-terminal form-control-feedback"></span>
          </div>
          <div class="form-group has-feedback">
            <input type="password" name="pass" class="form-control" placeholder="数据库密码" data-msg-required="数据库密码不能为空" required>
            <span class="fa fa-key form-control-feedback"></span>
          </div>
          <div class="form-group has-feedback">
            <input type="text" name="name" class="form-control" placeholder="数据库名称" value="gocms" required>
            <span class="fa fa-font form-control-feedback"></span>
          </div>
          <div class="row">
            <div class="col-xs-4 pull-right">
              <a class="btn btn-primary btn-block btn-flat" href="#admin" data-toggle="tab">下一步</a>
            </div>
          </div>
        </div>
        <div class="tab-pane" id="admin">
          <p class="login-box-msg">配置管理员账户</p>
          <div class="form-group has-feedback">
            <input type="email" name="email" class="form-control" placeholder="请输入邮箱" data-rule-maxlength="255" required>
            <span class="fa fa-envelope form-control-feedback"></span>
          </div>
          <div class="form-group has-feedback">
            <input type="password" name="password" class="form-control" placeholder="请输入密码" id="register_password" data-rule-minlength="6" required>
            <span class="fa fa-lock form-control-feedback"></span>
          </div>
          <div class="form-group has-feedback">
            <input type="password" class="form-control" placeholder="请重新输入密码" data-rule-equalTo="#register_password" data-msg-equalTo="两次输入的密码不一致" required>
            <span class="fa fa-check form-control-feedback"></span>
          </div>
          <div class="row">
            <div class="col-xs-4 pull-right">
              <button type="submit" class="btn btn-primary btn-block btn-flat">安装</button>
            </div>
          </div>
        </div>
      </form>
    </div>
  </div>
  {{- template "footer" "install"}}
  <script type="text/javascript">
    var options = {
      'MySQL': { 'host': 'localhost', 'port': 3306, 'user': 'root' },
      'Postgres': { 'host': 'localhost', 'port': 5432, 'user': 'postgres' },
      'SQLite3': { 'host': 'gocms.db', 'port': null, 'user': 'admin', 'name': null },
    };
    $(document).ready(function () {
      var opts = $.map(options, function (el, i) {
        return new Option(i)
      })
      $('select[name="type"]').on('change', function () {
        $.each(options[this.value], function (i, e) {
          var v = $('input[name="' + i + '"]').val(e).
            attr('disabled', e ? false : true).parent();
          e ? v.show() : v.hide();
        })
      }).trigger('change').append(opts);

      $('a[data-toggle="tab"]').on('show.bs.tab', function(e) {
        return $('input', e.relatedTarget).valid()
      })
    })
  </script>
</body>
</html>