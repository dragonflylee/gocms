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
              <li class="active"><a href="/p2w/all" class="text-purple">总体统计</a></li>
              <li ><a href="/p2w/qd" class="text-orange">渠道统计</a></li>
            </ul>
          </div>
          <div class="box-tools">
            <form class="form-inline">
              <a class="btn bg-olive btn-sm btn-export" href="?export=xls" title="导出">导出 <i class="fa fa-file-excel-o"></i></a>
            </form>
          </div>
        </div>
      {{if .data.list}}
        <div class="box-body table-responsive">
          <table class="table table-bordered">
            <tbody>
              <tr>
                <th>日期</th>
                <th>下载</th>
                <th>安装</th>
                <th>运行</th>
              </tr>
            {{range .data.list}}
              <tr>
                <td>{{.Date}}</td>
                <td>{{.DownloadStart}}/{{.DownloadEnd}}</td>
                <td>{{.InstallEnd}}/{{.InstallStart}}</td>
                <td>{{.Run}}</td>
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

