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
          <h3 class="box-title">日志列表</h3>
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
            </form>
          </div>
        </div>
      {{if .data.list}}
        <div class="box-body table-responsive">
          <table class="table table-bordered">
            <tbody>
              <tr>
                <th>#</th>
                <th>管理员</th>
                <th>访问内容</th>
                <th>注释</th>
                <th>操作时间</th>
                <th>IP</th>
              </tr>
            {{range .data.list}}
              <tr>
                <td>{{.ID}}</td>
                <td>{{.Admin.Email}}</td>
                <td>{{.Path}}</td>
                <td>{{.Commit}}</td>
                <td>{{date .CreatedAt}}</td>
                <td>{{.IP}}</td>
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
    </section>
  </div>
  {{template "footer"}}
</div>
</body>
</html>