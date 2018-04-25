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
                <td></td>
                <td>{{.CreatedAt.Format "2006-01-02 15:04:05"}}</td>
                <td>{{.LastLogin.Format "2006-01-02 15:04:05"}} / {{.LastIP}}</td>
                <td>
                {{if .Status}}
                  <label class="btn btn-xs btn-success">已激活</label>
                {{else}}
                  <label class="btn btn-xs btn-danger">未激活</label>
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
  {{template "footer"}}
</div>
</body>
</html>