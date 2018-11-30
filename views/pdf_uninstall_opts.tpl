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
              <li><a href="/pdf/feedbacks" class="text-purple">用户反馈</a></li>
              <li class="active"><a href="/pdf/uninstall_opts" class="text-orange">卸载反馈</a></li>
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
                <th>联系</th>
                <th>IP</th>
                <th>系统</th>
				<th>版本</th>
				<th>内容</th>
				<th>反馈</th>
              </tr>
            {{range .data.list}}
              <tr>
                <td>{{date .LogTime}}</td>
                <td>{{.Contact}}</td>
                <td>{{.ClientIP}}</td>
                <td>{{.OS}}</td>
				<td>{{.Version}}</td>
				<td>{{.Content}}</td>
				<td>{{.Feedback}}</td>
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
	  <div class="box box-info">
	  {{if .data.results}}
	  <div class="box box-solid">
	    <div class="box-header with-border">
          <h3 class="box-title">卸载原因</h3>
      </div>
	  {{range .data.results}}	  
		<p>{{.Result}}
        <div class="progress">
          <div class="progress-bar progress-bar-green" role="progressbar" aria-valuenow="40" aria-valuemin="10%" aria-valuemax="100" style="width: {{rate .Rate}}; min-width:5em">
            <span >{{.Count}} ({{rate .Rate}})</span>
          </div>
        </div>
	  {{end}}
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

