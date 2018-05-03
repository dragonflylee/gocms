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
          <h3 class="box-title">角色列表</h3>
          <div class="box-tools">
            {{if .user.Access "/admin/group/add"}}
              <a class="btn bg-purple btn-sm" data-target="#modal-add" data-toggle="modal" title="添加">添加 <i class="fa fa-plus"></i></a>
            {{end}}
            </form>
          </div>
        </div>
      {{if .data}}
        <div class="box-body table-responsive">
          <table class="table table-bordered">
            <tbody>
              <tr>
                <th style="width: 10px;">#</th>
                <th>名称</th>
                <th>操作</th>
              </tr>
            {{range $id, $name := .data}}
              <tr>
                <td>{{$id}}</td>
                <td>{{$name}}</td>
                <td>
                </td>
              </tr>
            {{end}}
            </tbody>
          </table>
        </div>
      {{else}}
        <div class="box-body">无数据</div>
      {{end}}
      </div>
    </section>
  </div>
  {{if .user.Access "/admin/group/add"}}
  <div class="modal" id="modal-add">
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
  {{template "footer"}}
</div>
</body>
</html>