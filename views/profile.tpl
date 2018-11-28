<!DOCTYPE html>
<html>
<head>
  {{template "header" .node}}
</head>
<body class="hold-transition skin-blue sidebar-mini">
<div class="wrapper">
  {{template "navbar" .}}
  <div class="content-wrapper">
    {{template "title" .}}
    
    <section class="content">
      <div class="row">
        <div class="col-md-3">
          <div class="box box-primary">
            <div class="box-body box-profile">
              <img class="profile-user-img img-responsive img-circle" src="{{.user.Headpic}}" alt="用户头像">
              <h3 class="profile-username text-center"></h3>
              <p class="text-muted text-center"></p>

              <ul class="list-group list-group-unbordered">
                <li class="list-group-item">
                  <b>注册时间</b> <a class="pull-right">{{date .user.CreatedAt}}</a>
                </li>
                <li class="list-group-item">
                  <b>上次登录</b> <a class="pull-right">{{date .user.LastLogin}}</a>
                </li>
                <li class="list-group-item">
                  <b>上次IP</b> <a class="pull-right">{{.user.LastIP}}</a>
                </li>
              </ul>

              <a href="#" class="btn btn-primary btn-block"><b>签到</b></a>
            </div>
          </div>
        </div>

        <div class="col-md-9">
          <div class="box box-warning">
            <div class="box-header with-border">
              <h3 class="box-title">密码安全</h3>
            </div>
            <div class="box-body">
              <form class="form-horizontal" action="/password" method="post">
              {{if not .user.Status}}
                <div class="form-group">
                  <div class="col-sm-6 col-sm-offset-2">
                    <div class="callout callout-warning">
                      <i class="icon fa fa-warning"></i> 请修改密码激活管理员账户
                    </div>
                  </div>
                </div>
              {{end}}
                <div class="form-group">
                  <label class="col-sm-3 control-label">新密码</label>
                  <div class="col-sm-5">
                    <input name="password" type="password" id="register_password" class="form-control input-medium" placeholder="请输入新密码" data-rule="{'regexPasswd':true}" required>
                  </div>
                </div>
                <div class="form-group">
                  <label class="col-sm-3 control-label">确认密码</label>
                  <div class="col-sm-5">
                    <input name="rpasswd" type="password" class="form-control input-medium" data-rule="{'equalTo':'#register_password'}" data-message="{'equalTo':'两次输入的密码不一致'}" placeholder="请再次输入新密码" required>
                  </div>
                </div>
                <div class="form-group">
                  <div class="col-sm-offset-2 col-sm-8">
                    <button type="submit" class="btn btn-danger">修改</button>
                  </div>
                </div>
              </form>
            </div>
          </div>
        </div>
      </div>
    </section>
  </div>

  {{template "footer"}}
</div>
</body>
</html>

