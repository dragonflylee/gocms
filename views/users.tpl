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
      <div class="row">
        <div class="col-md-2">
          {{if .user.Access "/admin/group/add"}}
            <a class="btn bg-purple btn-block margin-bottom" data-target="#group-add" data-toggle="modal" title="添加">添加角色 <i class="fa fa-plus"></i></a>
          {{end}}
          <div class="box box-solid">
            <div class="box-header with-border">
              <h3 class="box-title">组</h3>
              <div class="box-tools">
                <a class="btn btn-box-tool" data-widget="collapse"><i class="fa fa-minus"></i></a>
              </div>
            </div>
          {{if .data.group}}
            <div class="box-body no-padding">
              <ul class="nav nav-pills nav-stacked">
                <li {{if not (.form.Get "group")}}class="active"{{end}}><a href="?">所有</a></li>
            {{range $id, $name := .data.group}}
              {{if eq (print $id) ($.form.Get "group")}}
                <li class="active">
                  <a>{{$name}}
                    <span class="btn btn-xs bg-navy pull-right" data-href="/admin/group/{{$id}}" data-target="#modal-edit" data-toggle="modal"><i class="fa fa-edit"></i></span>
                  </a>
                </li>
              {{else}}
                <li><a href="?group={{$id}}">{{$name}}</a></li>
              {{end}}
            {{end}}
              </ul>
            </div>
          {{else}}
            <div class="box-body">
              <p class="lead text-center">无数据</p>
            </div>
          {{end}}
          </div>
        </div>

        <div class="col-md-10">
          <div class="box">
            <div class="box-header with-border">
              <h3 class="box-title">管理员列表</h3>
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
                  <a class="btn bg-purple btn-sm" data-target="#user-add" data-toggle="modal" title="添加">添加 <i class="fa fa-plus"></i></a>
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
                    <td>{{index $.data.group .GroupID}}</td>
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
            <div class="box-body">
              <p class="lead text-center">无数据</p>
            </div>
          {{end}}
          </div>
        </div>
      </div>
    </section>
  </div>
  {{if .user.Access "/admin/user/add"}}
  <div class="modal modal-add" id="user-add">
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
              {{range $id, $name := .data.group}}
                {{if lt $.user.GroupID $id}}
                  <option value="{{$id}}" {{if eq (print $id) ($.form.Get "group")}}selected{{end}}>{{$name}}</option>
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
  {{if .user.Access "/admin/group/add"}}
  <div class="modal modal-add" id="group-add">
    <div class="modal-dialog">
      <div class="modal-content">
        <form action="/admin/group/add" method="post" class="form-horizontal">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal" aria-label="Close">
              <span aria-hidden="true">×</span></button>
            <h4 class="modal-title">添加角色</h4>
          </div>
          <div class="modal-body">
            <div class="form-group">
              <label class="col-sm-3 control-label">名称</label>
              <div class="col-sm-5">
                <input type="text" class="form-control" name="name" required>
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
  <div class="modal" id="modal-edit">
    <div class="modal-dialog">
      <div class="modal-content box">
      </div>
    </div>
  </div>
  {{template "footer"}}
  <script type="text/javascript">
    $(document).on('click', '#modal-edit button[type="submit"]', function(e) {
      var nodes = $('.jstree').jstree('get_selected');
      console.log(nodes);
    })
  </script>
</div>
</body>
</html>