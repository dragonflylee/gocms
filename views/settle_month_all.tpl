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
              <li> <a href="/qd/day" class="text-purple">安装活跃</a></li>
              <li class="active"><a href="/qd/month" class="text-orange">月度结算</a></li>
            </ul>
          </div>
          <div class="box-tools">
            <form class="form-inline">
              <div class="form-group">
                <label>渠道</label>
                <select class="form-control select2 input-group-sm" name="qd" data-ajax--url="/qd/list" data-ajax--cache="true">
                  <option selected>{{.form.Get "qd"}}</option>
                </select>
              </div>
              <button type="submit" class="btn bg-purple btn-sm" title="筛选">筛选 <i class="fa fa-filter"></i></button>
              <a class="btn bg-olive btn-sm btn-export" href="?export=xls" title="导出">导出 <i class="fa fa-file-excel-o"></i></a>
            </form>
          </div>
        </div>
      {{if .data.list}}
        <div class="box-body table-responsive">
          <table class="table table-bordered">
            <tbody>
              <tr>
                <th>月度</th>
                <th>渠道</th>
                <th>安装</th>
                <th>卸载</th>
                <th>前台活跃</th>
                <th>后台活跃</th>
                <th>Rate</th>
                <th>单价(元)</th>
                <th>结算金额(元)</th>
              </tr>
            {{range .data.list}}
              <tr>
                <td>{{.Month}}</td>
                <td>{{.QD}}</td>
                <td>{{.RealInstallEnd}}</td>
                <td>{{.RealUninstallEnd}}</td>
                <td>{{.RealMFShow}}</td>
                <td>{{.RealServerRun}}</td>
                <td>{{rate .Coefficient}}%</td>
                <td>{{price .Price 4}}</td>
                <td>{{price .Total 2}}</td>
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
