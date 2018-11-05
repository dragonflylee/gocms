<!DOCTYPE html>
<html>
<head>
  {{template "header" .node.Name}}
  <style>
    .select2 {min-width:120px;}
  </style>
</head>
<body class="hold-transition skin-blue sidebar-mini">
<div class="wrapper">
  {{template "navbar" .}}
  <div class="content-wrapper">
    {{template "title" .}}
    <section class="content">
      <div class="box">
        <div class="box-header with-border">
          <div class="nav-tabs-custom">
            <ul class="nav nav-tabs">
              <li><a href="/admin/pdf/install_runs" class="text-purple">安装活跃</a></li>
              <li><a href="/admin/pdf/retentions" class="text-orange">留存率</a></li>
              <li><a href="/admin/pdf/mfshow_versions" class="text-green">版本占比</a></li>
              <li class="active"><a href="/admin/pdf/crashs" class="text-red">崩溃统计</a></li>
            </ul>
          </div>
	    </div>
      {{if .data.list}}
        <div class="box-body table-responsive">
          <table class="table table-bordered">
            <tbody>
              <tr>
                <th>日期</th>
                <th>活跃</th>
                <th>崩溃</th>
                <th>崩溃率</th>
              </tr>
            {{range .data.list}}
              <tr>
                <td>
                  <a href="/admin/pdf/crashs/detail?date={{ .Date }}">{{ .Date }}</a>
                </td>
                <td>{{.MFShow}}</td>
                <td>{{.Crash}}</td>
                <td>{{.CrashRate}}</td>
              </tr>
            {{end}}
            </tbody>
          </table>
        </div>
        <div class="box-footer clearfix">
          <a href="javascript:history.go(-1);" class="btn btn-sm bg-navy">返回</a>
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
  {{template "modal"}}
  {{template "footer"}}
</div>
</body>
</html>

