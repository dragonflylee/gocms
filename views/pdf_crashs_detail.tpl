<!DOCTYPE html>
<html>

<head>
	{{template "header" .node.Name}}
	<style>
		.select2 {
			min-width: 120px;
		}
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
								<li><a href="/pdf/install_runs" class="text-purple">安装活跃</a></li>
								<li><a href="/pdf/retentions" class="text-orange">留存率</a></li>
								<li><a href="/pdf/mfshow_versions" class="text-green">版本占比</a></li>
								<li class="active"><a href="/pdf/crashs" class="text-red">崩溃统计</a></li>
							</ul>
						</div>
					</div>
					{{if .data.crash_rate}}
					<div class="box-body left">
						{{range .data.crash_rate}}
						<p>{{.Version}}({{.Rate}}%)
							<div class="progress progress-xs">
								<div class="progress-bar progress-bar-red" role="progressbar" aria-valuenow="40" aria-valuemin="10%"
								 aria-valuemax="100" style="width: {{.Rate}}%; min-width:5em"></div>
							</div>
							{{end}}
					</div>
					{{else}}
					<div class="box-body">
						<p class="lead text-center">无数据</p>
					</div>
					{{end}}
					<!------------------------>
					<div class="box">
						{{if .data.crash_list}}
						<div class="box-body table-responsive right">
							<table class="table table-bordered">
								<tbody>
									<tr>
										<th>日期</th>
										<th>版本号</th>
										<th>系统</th>
										<th>硬件ID</th>
										<th>客户端IP</th>
									</tr>
									{{range .data.crash_list}}
									<tr>
										<td>{{date .LogTime}}</td>
										<td>{{.Version}}</td>
										<td>{{.OS}}</td>
										<td>{{.DeviceID}}</td>
										<td>{{.ClientIP}}</td>
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
				</div>
			</section>
		</div>
		{{template "modal"}}
		{{template "footer"}}
	</div>
</body>

</html>