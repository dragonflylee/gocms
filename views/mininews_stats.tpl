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
        <!--
        <div class="box-header with-border">
		 <div class="nav-tabs-custom">
            <ul class="nav nav-tabs">
              <li class="active"><a href="/admin/pdf/feedbacks" class="text-purple">用户反馈</a></li>
              <li ><a href="/admin/pdf/uninstall_opts" class="text-orange">卸载反馈</a></li>
            </ul>
          </div>
          <div class="box-tools">
            <form class="form-inline">
              <a class="btn bg-olive btn-sm btn-export" href="?export=xls" title="导出">导出 <i class="fa fa-file-excel-o"></i></a>
            </form>
          </div>
        </div>
        -->
      {{if .data.list}}
        <div class="box-body table-responsive">
          <table class="table table-bordered">
            <tbody>
              <tr>
                <th>日期</th>
                <th>用户禁用</th>
                <th>常驻进程启动</th>
                <th>下载组件(end/start)</th>
                <th>加载组件(end/start)</th>
                <th>云端禁用</th>
                <th>服务调起常驻进程
              </tr>
            {{range .data.list}}
              <tr>
                <td>{{.Date}}</td>
                <td>{{.ForbidMiniNews}}</td>
                <td>{{.SpeedupRun}}</td>
                <td>{{.LoaderDownloadEnd}}/{{.LoaderDownloadStart}}</td>
				<td>{{.LoaderLoadEnd}}/{{.LoaderLoadStart}}</td>
				<td>{{.CloudDisable}}</td>
				<td>{{.ServerCallSpeedup}}</td>
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

