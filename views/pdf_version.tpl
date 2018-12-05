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
                <div class="row">
                    <div class="col-md-6">
                        <div class="box box-solid">
                            <div class="box-header with-border">
                                <h3 class="box-title">官网版本</h3>
                            </div>
                            <div class="box-body box-profile">
                                <ul class="list-group list-group-unbordered">
                                    <li class="list-group-item">
                                        <b>版本号</b> <a class="pull-right">1.0.1</a>
                                    </li>
                                    <li class="list-group-item">
                                        <b>更新类型</b> <a class="pull-right">普通更新</a>
                                    </li>
                                    <li class="list-group-item">
                                        <b>版本类型</b> <a class="pull-right">Release</a>
                                    </li>
                                    <li class="list-group-item">
                                        <b>发布说明</b> <a class="pull-right">1112222</a>
                                    </li>
                                    <li class="list-group-item">
                                        <b>发布时间</b> <a class="pull-right">2018-01-02</a>
                                    </li>
                                    <li class="list-group-item">
                                        <b>包大小</b> <a class="pull-right">12323423</a>
                                    </li>
                                    <li class="list-group-item">
                                        <b>MD5</b> <a class="pull-right">a23ac232eassdddddd</a>
                                    </li>
                                    <li class="list-group-item">
                                        <b>发布人</b> <a class="pull-right">Wans</a>
                                    </li>
                                    <li class="list-group-item">
                                        <b>下载地址</b> <a class="pull-right">http://archive.xundupdf.com/XDInstaller.exe</a>
                                    </li>
                                </ul>
                            </div>
                            <!--
                            {{if .data.list}}
                            <div class="box-body table-responsive">
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
                            -->
                        </div>
                    </div>
                    <div class="col-md-6">
                        <div class="box box-solid">
                            <div class="box-header with-border">
                                <h3 class="box-title">更新接口版本</h3>
                            </div>
                            {{if .data.list}}
                            <div class="box-body table-responsive">
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
                </div>
                <div class="box box-info">
                    <div class="box-header with-border">
                        <h3 class="box-title">版本列表</h3>
                    </div>
                    {{if .data.list}}
                    <div class="box-body table-responsive">
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
    </div>
    </section>
    </div>
    {{template "modal"}}
    {{template "footer"}}
    </div>
</body>

</html>