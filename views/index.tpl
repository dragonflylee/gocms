<!DOCTYPE html>
<html>
<head>
  {{template "header" .Node}}
</head>
<body class="hold-transition skin-blue sidebar-mini">
  <div class="wrapper">
    {{- template "navbar" .}}
    <div class="content-wrapper">
      {{- template "title" .}}
      <section class="content">
        <div class="row">
          <div class="col-lg-3 col-xs-6">
            <div class="small-box bg-aqua">
              <div class="inner">
                <h3>150</h3>
                <p>新订单</p>
              </div>
              <div class="icon">
                <i class="ion ion-bag"></i>
              </div>
              <a href="#" class="small-box-footer">更多 <i class="fa fa-arrow-circle-right"></i></a>
            </div>
          </div>
          <div class="col-lg-3 col-xs-6">
            <div class="small-box bg-green">
              <div class="inner">
                <h3>53<sup style="font-size: 20px">%</sup></h3>
                <p>增长率</p>
              </div>
              <div class="icon">
                <i class="ion ion-stats-bars"></i>
              </div>
              <a href="#" class="small-box-footer">更多 <i class="fa fa-arrow-circle-right"></i></a>
            </div>
          </div>
          <div class="col-lg-3 col-xs-6">
            <div class="small-box bg-yellow">
              <div class="inner">
                <h3>44</h3>
                <p>用户注册</p>
              </div>
              <div class="icon">
                <i class="ion ion-person-add"></i>
              </div>
              <a href="#" class="small-box-footer">更多 <i class="fa fa-arrow-circle-right"></i></a>
            </div>
          </div>
          <div class="col-lg-3 col-xs-6">
            <div class="small-box bg-red">
              <div class="inner">
                <h3>65</h3>
                <p>访问量</p>
              </div>
              <div class="icon">
                <i class="ion ion-pie-graph"></i>
              </div>
              <a href="#" class="small-box-footer">更多 <i class="fa fa-arrow-circle-right"></i></a>
            </div>
          </div>
        </div>
        <div class="row">
          {{ $from := .Data.AddDate 0 0 -15}}
          <section class="col-lg-7">
            <div class="box box-warning" id="area-chart" data-source="?action=user">
              <div class="box-header with-border">
                <h3 class="box-title">趋势</h3>
                <div class="box-tools pull-right form-inline visible-lg">
                  <div class="form-group">
                    <div class="input-group input-group-sm input-daterange" data-provide="datepicker" data-date-language="zh-CN" data-date-format="yyyy-mm-dd" data-date-end-date="0d" data-date-autoclose="true" data-date-orientation="bottom">
                      <input type="text" class="form-control" placeholder="请选择开始" name="from" value="{{$from.Format "2006-01-02"}}" readonly>
                      <span class="input-group-addon"><i class="fa fa-calendar"></i></span>
                      <input type="text" class="form-control" placeholder="请选择结束" name="to" value="{{.Data.Format "2006-01-02"}}" readonly>
                    </div>
                  </div>
                  <div class="form-group">
                    <a class="btn btn-box-tool refresh-btn"><i class="fa fa-refresh"></i></a>
                  </div>
                </div>
              </div>
              <div class="box-body">
                <div class="chart" style="height: 250px;"></div>
              </div>
            </div>
          </section>
          <section class="col-lg-5">
            <form class="box box-info" action="/upload" method="post" data-target=".box-footer" enctype="multipart/form-data">
              <div class="box-header with-border">
                <h3 class="box-title">上传文件</h3>
                <div class="box-tools pull-right">
                  <a class="btn btn-box-tool" data-widget="collapse"><i class="fa fa-minus"></i></a>
                </div>
              </div>
              <div class="box-body form-horizontal">
                <div class="form-group">
                  <label class="col-sm-3 control-label">方法</label>
                  <div class="col-sm-6 icheck">
                    <label class="radio-inline">
                      <input type="radio" name="type" value="1" checked>
                      MD5
                    </label>
                    <label class="radio-inline">
                      <input type="radio" name="type" value="2">
                      QRCode
                    </label>
                  </div>
                </div>
                <div class="form-group">
                  <label class="col-sm-3 control-label">上传</label>
                  <div class="col-sm-8">
                    <input type="file" class="filestyle" name="file" data-buttonText="浏览" required>
                  </div>
                </div>
              </div>
              <div class="box-footer">
                <button type="submit" class="btn bg-purple pull-right">上传</button>
              </div>
            </form>
          </section>
        </div>
      </section>
    </div>
    {{template "footer"}}
    <script src="/static/js/index.js?v{{version}}" type="text/javascript"></script>
  </div>
</body>
</html>