<!DOCTYPE html>
<html>
<head>
  {{template "header" .node.Name}}
</head>
<body class="hold-transition skin-blue sidebar-mini">
<div class="wrapper">
  {{template "navbar" .}}
  <div class="content-wrapper">
    {{template "title" .}}
    <section class="content">
      <div class="box">
        <div class="box-header with-border">
          <h3 class="box-title">用户列表</h3>
          <div class="box-tools">
            <form class="form-inline">
              <div class="form-group">
                <div class="input-group input-group-sm">
                  <input type="email" class="form-control" placeholder="请输入管理员邮箱" name="email" value="{{.form.Get "email"}}" required>
                  <span class="input-group-btn">
                    <button type="submit" class="btn btn-info btn-sm" title="搜索"><i class="fa fa-search"></i></button>
                  </span>
                </div>
              </div>
            {{if .user.Access "/admin/user/add"}}
              <a class="btn bg-purple btn-sm" data-target="#modal-add" data-toggle="modal" title="添加">添加 <i class="fa fa-plus"></i></a>
            {{end}}
            </form>
          </div>
        </div>
      {{if .data.list}}
        <div class="box-body table-responsive">
          <table class="table table-bordered">
            <tbody>
              <tr>
                <th>#</th>
                <th>邮箱</th>
                <th>用户组</th>
                <th>创建时间</th>
                <th>最后登录</th>
                <th>状态</th>
              </tr>
            {{range .data.list}}
              <tr>
                <td>{{.ID}}</td>
                <td>{{.Email}}</td>
                <td>{{.Group.Name}}</td>
                <td>{{date .CreatedAt}}</td>
                <td>{{date .LastLogin}} / {{.LastIP}}</td>
                <td>
                {{if .Status}}
                  <span class="text-maroon">已激活</span>
                {{else}}
                  <span class="text-navy">未激活</span>
                {{end}}
                </td>
              </tr>
            {{end}}
            </tbody>
          </table>
        </div>
        <div class="box-footer clearfix">
          {{template "paginator" .data}}
        </div>
      {{else}}
        <div class="box-body">无数据</div>
      {{end}}
      </div>
    </section>
  </div>
  {{if .user.Access "/admin/user/add"}}
  <div class="modal" id="modal-add">
    <div class="modal-dialog">
      <div class="modal-content">
        <form action="/admin/user/add" method="post" class="form-horizontal">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal" aria-label="Close">
              <span aria-hidden="true">×</span></button>
            <h4 class="modal-title">添加管理员</h4>
          </div>
          <div class="modal-body">
            <div class="form-group">
              <label class="col-sm-3 control-label">邮箱</label>
              <div class="col-sm-5">
                <input type="email" class="form-control" name="email" required>
              </div>
            </div>
            <div class="form-group">
              <label class="col-sm-3 control-label">用户组</label>
              <div class="col-sm-5">
                <select name="group" class="form-control">
              {{range $id, $name := .data.groups}}
                {{if lt $.user.GroupID $id}}
                  <option value="{{$id}}">{{$name}}</option>
                {{end}}
              {{end}}
                </select>
              </div>
            </div>
          </div>
          <div class="modal-footer">
            <a class="btn btn-default" data-dismiss="modal">取消</a>
            <button type="submit" class="btn bg-purple">新增</button>
          </div>
        </form>
      </div>
    </div>
  </div>
  {{end}}
  {{template "footer"}}
</div>
</body>
</html>