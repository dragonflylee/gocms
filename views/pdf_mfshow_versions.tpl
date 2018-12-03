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
        <div class="box">
          <div class="box-header with-border">
            <div class="nav-tabs-custom">
              <ul class="nav nav-tabs">
                <li><a href="/pdf/install_runs" class="text-purple">安装活跃</a></li>
                <li><a href="/pdf/retentions" class="text-orange">留存率</a></li>
                <li class="active"><a href="/pdf/mfshow_versions" class="text-green">版本占比</a></li>
                <li><a href="/pdf/crashs" class="text-red">崩溃统计</a></li>
              </ul>
            </div>
            <div class="box-tools">
              <form class="form-inline">
                <a class="btn bg-olive btn-sm btn-export" href="?export=xls" title="导出">导出 <i class="fa fa-file-excel-o"></i></a>
              </form>
            </div>
          </div>
        </div>
        {{if .data.list}}
        <div class="box-body">
          {{range .data.list}}
          <p>{{.Version}}({{.Rate}}%)
            <div class="progress progress-xs">
              <div class="progress-bar progress-bar-green" role="progressbar" aria-valuenow="40" aria-valuemin="10%"
                aria-valuemax="100" style="width: {{.Rate}}%; min-width:5em">
              </div>
            </div>
            {{end}}
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