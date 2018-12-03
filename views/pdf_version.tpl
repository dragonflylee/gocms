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

                <div class="box">
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
            </section>
        </div>
        {{template "modal"}}
        {{template "footer"}}
    </div>
</body>

</html>