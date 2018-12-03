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
        <!-- Small boxes (Stat box) -->
        <div class="row">
          <div class="col-lg-3 col-xs-6">
            <!-- small box -->
            <div class="small-box bg-aqua">
              <div class="inner">
                <h3><sup style="font-size: 24px">{{.data.total.InstallEnd}}</sup></h3>
                <p>累计安装</p>
              </div>
              <div class="icon">
                <i class="fa fa-user-plus" aria-hidden="true"></i>
              </div>
            </div>
          </div>
          <!-- ./col -->
          <div class="col-lg-3 col-xs-6">
            <!-- small box -->
            <div class="small-box bg-yellow">
              <div class="inner">
                <h3><sup style="font-size: 24px">{{.data.total.UninstallEnd}}</sup></h3>
                <p>累计卸载</p>
              </div>
              <div class="icon">
                <i class="fa fa-user-times" aria-hidden="true"></i>
              </div>
            </div>
          </div>
          <!-- ./col -->
          <div class="col-lg-3 col-xs-6">
            <!-- small box -->
            <div class="small-box bg-green">
              <div class="inner">
                <h3><sup style="font-size: 24px">{{.data.total.MFShow}}</sup></h3>
                <p>累计活跃</p>
              </div>
              <div class="icon">
                <i class="fa fa-star" aria-hidden="true"></i>
              </div>
            </div>
          </div>
        </div>
        <div class="box">
          <div class="box-header with-border">
            <div class="nav-tabs-custom">
              <ul class="nav nav-tabs">
                <li class="active"><a href="/pdf/install_runs" class="text-purple">安装活跃</a></li>
                <li><a href="/pdf/retentions" class="text-orange">留存率</a></li>
                <li><a href="/pdf/mfshow_versions" class="text-green">版本占比</a></li>
                <li><a href="/pdf/crashs" class="text-red">崩溃统计</a></li>
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
                  <th>卸载 </th>
                  <th>新用户卸载比 </th>
                  <th>Doc加载</th>
                  <th>活跃</th>
                  <th>7日活跃</th>
                  <th>月活跃</th>
                  <th>老用户活跃</th>
                </tr>
                {{range .data.list}}
                <tr>
                  <td>{{.Date}}</td>
                  <td>{{.InstallEnd}}/{{.InstallStart}}</td>
                  <td>{{.UninstallEnd}}/{{.UninstallStart}}</td>
                  <td>{{rate .NewUserUninstallRate}}</td>
                  <td>{{.LoadDocDistinctDevice}}/{{.LoadDoc}}</td>
                  <td>{{.MFShow}}/{{.ServerRun}}</td>
                  <td>{{.MFShow7}}/{{.ServerRun7}}</td>
                  <td>{{.MFShow30}}/{{.ServerRun30}}</td>
                  <td>{{.MFShowOld}}</td>
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