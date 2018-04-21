<!DOCTYPE html>
<html>
<head>
  {{template "header" "个人中心"}}
</head>
<body class="hold-transition skin-blue sidebar-mini">
<div class="wrapper">
  {{template "navbar" .}}
  <div class="content-wrapper">
    <section class="content-header">
      <h1>
        个人中心
        <small>密码修改</small>
      </h1>
      <ol class="breadcrumb">
        <li><a href="#"><i class="fa fa-dashboard"></i> 首页</a></li>
        <li class="active">个人中心</li>
      </ol>
    </section>

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
                  <b>注册时间</b> <a class="pull-right">{{.user.CreatedAt.Format "2006-01-02 15:04:05"}}</a>
                </li>
                <li class="list-group-item">
                  <b>上次登录</b> <a class="pull-right">{{.user.LastLogin.Format "2006-01-02 15:04:05"}}</a>
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
          <div class="nav-tabs-custom">
            <ul class="nav nav-tabs">
              <li class="active"><a href="#password" data-toggle="tab">密码安全</a></li>
              <li><a href="#avatar" data-toggle="tab">修改头像</a></li>
            </ul>
            <div class="tab-content">
              <div class="active tab-pane" id="password">
                <form class="form-horizontal" action="#" method="post">
                  <div class="form-group">
                    <label class="col-sm-2 control-label">原密码</label>
                    <div class="col-sm-8">
                      <input name="password" type="password" class="form-control input-medium" placeholder="请输入原密码" required>
                    </div>
                  </div>
                  <div class="form-group">
                    <label class="col-sm-2 control-label">新密码</label>
                    <div class="col-sm-8">
                      <input name="npassword" type="password" id="register_password" class="form-control input-medium" placeholder="请输入新密码" required>
                    </div>
                  </div>
                  <div class="form-group">
                    <label class="col-sm-2 control-label">确认新密码</label>
                    <div class="col-sm-8">
                      <input name="rppassword" type="password" class="form-control input-medium" data-rule="{'equalTo':'#register_password'}" data-message="{'equalTo':'两次输入的密码不一致'}" placeholder="请再次输入新密码" required>
                    </div>
                  </div>
                  <div class="form-group">
                    <div class="col-sm-offset-2 col-sm-8">
                      <button type="submit" class="btn btn-danger">修改</button>
                    </div>
                  </div>
                </form>
              </div>
              <div class="tab-pane" id="avatar">
                <form>
                  <div class="form-group">
                    <div class="fileinput fileinput-new" data-provides="fileinput">
                      <div class="fileinput-new thumbnail">
                        <img src="{{.user.Headpic}}" alt="">
                      </div>
                      <div class="fileinput-preview fileinput-exists thumbnail"></div>
                      <div>
                        <span class="btn btn-default btn-file">
                        <span class="fileinput-new">浏览</span>
                        <span class="fileinput-exists">更换</span>
                        <input type="file" name="avatar"></span>
                        <a href="#" class="btn btn-default fileinput-exists" data-dismiss="fileinput">删除</a>
                      </div>
                    </div>
                  </div>
                </form>
              </div>
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

