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
	 <!-- Small boxes (Stat box) -->
      <div class="row">
        <div class="col-lg-3 col-xs-6">
          <!-- small box -->
          <div class="small-box bg-aqua">
            <div class="inner">
              <h3><sup style="font-size: 20px">{{rate .data.avg.RMFShow}}/{{rate .data.avg.RServerRun}}</sup></h3>
              <p>Avg次日留存</p>
            </div>
            <div class="icon">
              <i class="fa fa-bar-chart" aria-hidden="true"></i>
            </div>
          </div>
        </div>
        <!-- ./col -->
        <div class="col-lg-3 col-xs-8">
          <!-- small box -->
          <div class="small-box bg-green">
            <div class="inner">
              <h3><sup style="font-size: 20px">{{rate .data.avg.RMFShow3}}/{{rate .data.avg.RServerRun3}}</sup></h3>
              <p>Avg三日留存</p>
            </div>
            <div class="icon">
              <i class="fa fa-area-chart" aria-hidden="true"></i>
            </div>
          </div>
        </div>
        <!-- ./col -->
        <div class="col-lg-3 col-xs-6">
          <!-- small box -->
          <div class="small-box bg-yellow">
            <div class="inner">
              <h3><sup style="font-size: 20px">{{rate .data.avg.RMFShow7}}/{{rate .data.avg.RServerRun7}}</sup></h3>
              <p>Avg七日留存</p>
            </div>
            <div class="icon">
              <i class="fa fa-bar-chart" aria-hidden="true"></i>
            </div>
          </div>
        </div>
        <!-- ./col -->
        <div class="col-lg-3 col-xs-6">
          <!-- small box -->
          <div class="small-box bg-red">
            <div class="inner">
              <h3><sup style="font-size: 20px">{{rate .data.avg.RMFShow30}}/{{rate .data.avg.RServerRun30}}</sup></h3>
              <p>Avg三十日留存</p>
            </div>
            <div class="icon">
              <i class="fa fa-line-chart" aria-hidden="true"></i>
            </div>
          </div>
        </div>
      </div>
      <div class="box">
        <div class="box-header with-border">
          <div class="nav-tabs-custom">
            <ul class="nav nav-tabs">
              <li ><a href="/admin/pdf/install_runs" class="text-purple">安装活跃</a></li>
              <li class="active"><a href="/admin/pdf/retentions" class="text-orange">留存率</a></li>
              <li ><a href="/admin/pdf/mfshow_versions" class="text-green">版本占比</a></li>
              <li ><a href="/admin/pdf/crashs" class="text-red">崩溃统计</a></li>
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
                <th>安装</th>
                <th>新用户卸载 </th>
                <th>次日留存</th>
                <th>3日留存</th>
                <th>7日留存</th>
                <th>月留存</th>
              </tr>
            {{range .data.list}}
              <tr>
                <td>{{.Date}}</td>
                <td>{{.Install}}</td>
                <td>{{.NewUserUninstall}}</td>
                <td>{{rate .RMFShow}}/{{rate .RServerRun}}</td>
                <td>{{rate .RMFShow3}}/{{rate .RServerRun3}}</td>
                <td>{{rate .RMFShow7}}/{{rate .RServerRun7}}</td>
                <td>{{rate .RMFShow30}}/{{rate .RServerRun30}}</td>
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


