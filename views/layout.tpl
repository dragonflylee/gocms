{{define "header"}}
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <title>{{.}}</title>
  <meta content="noarchive" name="robots">
  <meta content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no" name="viewport">
  <link href="//cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/3.3.7/css/bootstrap.min.css" rel="stylesheet">
  <link href="//cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css" rel="stylesheet">

  <link href="//cdnjs.cloudflare.com/ajax/libs/jasny-bootstrap/3.1.3/css/jasny-bootstrap.min.css" rel="stylesheet">
  <link href="//cdnjs.cloudflare.com/ajax/libs/bootstrap-switch/3.3.4/css/bootstrap3/bootstrap-switch.min.css" rel="stylesheet">
  <link href="//cdnjs.cloudflare.com/ajax/libs/iCheck/1.0.2/skins/all.css" rel="stylesheet">
  <link href="//cdnjs.cloudflare.com/ajax/libs/select2/4.0.5/css/select2.min.css" rel="stylesheet">
  <link href="//cdnjs.cloudflare.com/ajax/libs/bootstrap-datepicker/1.8.0/css/bootstrap-datepicker3.min.css" rel="stylesheet">
  <link href="//cdnjs.cloudflare.com/ajax/libs/jstree/3.3.5/themes/default/style.min.css" rel="stylesheet">

  <link href="//cdnjs.cloudflare.com/ajax/libs/admin-lte/2.4.8/css/AdminLTE.min.css" rel="stylesheet">
  <link href="//cdnjs.cloudflare.com/ajax/libs/admin-lte/2.4.8/css/skins/_all-skins.min.css" rel="stylesheet">
  <link href="/static/css/custom.min.css?v=20180422" rel="stylesheet" type="text/css">
  
  {{html "<!--[if lt IE 9]>"}}
  <script src="//cdnjs.cloudflare.com/ajax/libs/html5shiv/3.7.3/html5shiv.min.js"></script>
  <script src="//cdnjs.cloudflare.com/ajax/libs/respond.js/1.4.2/respond.min.js"></script>
  {{html "<![endif]-->"}}
{{end}}

{{define "navbar"}}
  <header class="main-header">
    <a href="/" class="logo">
      <span class="logo-mini"><b>后</b></span>
      <span class="logo-lg"><b>后台管理</b></span>
    </a>
    <nav class="navbar navbar-static-top">
      <a href="#" class="sidebar-toggle" data-toggle="push-menu">
        <span class="sr-only"></span>
        <span class="icon-bar"></span>
        <span class="icon-bar"></span>
        <span class="icon-bar"></span>
      </a>

      <div class="navbar-custom-menu">
        <ul class="nav navbar-nav">
          <li class="dropdown user user-menu">
            <a href="#" class="dropdown-toggle" data-toggle="dropdown">
              <img src="{{.user.Headpic}}" class="user-image" alt="用户头像">
              <span class="hidden-xs">{{.user.Email}}</span>
            </a>
            <ul class="dropdown-menu">
              <li class="user-header">
                  <img src="{{.user.Headpic}}" class="img-circle" alt="用户头像">
                <p>
                    {{.user.Email}}
                    <small>{{.user.Group.Name}}</small>
                    <small>上次登录 {{date .user.LastLogin}}</small>
                </p>
              </li>
              <li class="user-footer">
                <div class="pull-left">
                  <a href="/profile" class="btn btn-default btn-flat">个人中心</a>
                </div>
                <div class="pull-right">
                  <a href="/logout" class="btn btn-default btn-flat">注销</a>
                </div>
              </li>
            </ul>
          </li>
        </ul>
      </div>
    </nav>
  </header>

  <aside class="main-sidebar">
    <section class="sidebar">
      <div class="user-panel">
        <div class="pull-left image">
            <img src="{{.user.Headpic}}" class="img-circle" alt="用户头像">
        </div>
        <div class="pull-left info">
          <P>{{.user.Group.Name}}</P>
          <a href="#" title="{{.user.Email}}"><i class="fa fa-circle text-success"></i> 在线</a>
        </div>
      </div>
      <form class="sidebar-form">
        <div class="input-group">
          <input type="text" name="q" class="form-control" placeholder="搜索...">
          <span class="input-group-btn">
            <button type="submit" id="search-btn" class="btn btn-flat"><i class="fa fa-search"></i></button>
          </span>
        </div>
      </form>
      <ul class="sidebar-menu" data-widget="tree">
        {{template "sidebar" .menu.Assign .user.GroupID .node}}
      </ul>
    </section>
  </aside>
{{end}}

{{define "footer"}}
  <footer class="main-footer">
    <div class="pull-right hidden-xs">
      Powered by <b>{{version}}</b>
    </div>
    <strong>版权所有 &copy; 2018 <a href="https://github.com/dragonflylee/gocms">GoCMS</a>.</strong>
  </footer>
  <script src="//cdnjs.cloudflare.com/ajax/libs/jquery/1.12.4/jquery.min.js"></script>
  <script src="//cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/3.3.7/js/bootstrap.min.js"></script>
  <script src="//cdnjs.cloudflare.com/ajax/libs/jQuery-slimScroll/1.3.8/jquery.slimscroll.min.js"></script>
  <script src="//cdnjs.cloudflare.com/ajax/libs/fastclick/1.0.6/fastclick.min.js"></script>
  <script src="//cdnjs.cloudflare.com/ajax/libs/admin-lte/2.4.8/js/adminlte.min.js"></script>

  <script src="//cdnjs.cloudflare.com/ajax/libs/jquery.form/4.2.2/jquery.form.min.js"></script>
  <script src="//cdnjs.cloudflare.com/ajax/libs/jquery-validate/1.17.0/jquery.validate.min.js"></script>
  <script src="//cdnjs.cloudflare.com/ajax/libs/blueimp-md5/2.10.0/js/md5.min.js"></script>
  <script src="//cdnjs.cloudflare.com/ajax/libs/jstree/3.3.5/jstree.min.js"></script>

  <script src="//cdnjs.cloudflare.com/ajax/libs/jasny-bootstrap/3.1.3/js/jasny-bootstrap.min.js"></script>
  <script src="//cdnjs.cloudflare.com/ajax/libs/bootstrap-switch/3.3.4/js/bootstrap-switch.min.js"></script>
  <script src="//cdnjs.cloudflare.com/ajax/libs/iCheck/1.0.2/icheck.min.js"></script>
  <script src="//cdnjs.cloudflare.com/ajax/libs/select2/4.0.5/js/select2.min.js"></script>
  <script src="//cdnjs.cloudflare.com/ajax/libs/select2/4.0.5/js/i18n/zh-CN.js"></script>
  <script src="//cdnjs.cloudflare.com/ajax/libs/bootstrap-datepicker/1.8.0/js/bootstrap-datepicker.min.js"></script>
  <script src="//cdnjs.cloudflare.com/ajax/libs/bootstrap-datepicker/1.8.0/locales/bootstrap-datepicker.zh-CN.min.js"></script>

  <script src="/static/js/global.js?v=20181114" type="text/javascript"></script>
{{end}}

{{define "sidebar"}}
  {{range .m}}
    {{if not (.HasGroup $.group)}}
    {{else if .Child}}
      <li class="treeview {{if $.node.HasParent .ID}}active{{end}}">
        <a href="#">
          <i class="{{.Icon}}"></i> <span>{{.Name}}</span>
          <span class="pull-right-container">
            <i class="fa fa-angle-left pull-right"></i>
          </span>
        </a>
        <ul class="treeview-menu">
          {{template "sidebar" .Child.Assign $.group $.node}}
        </ul>
      </li>
    {{else}}
      <li {{if $.node.HasParent .ID}}class="active"{{end}}>
        <a href="{{.Path}}">
          <i class="{{.Icon}}"></i>
          <span>{{.Name}}</span>
        </a>
      </li>
    {{end}}  
  {{end}}
{{end}}

{{define "title"}}
  <section class="content-header">
    <h1>
      {{.node}}
      <small>{{.node.Remark}}</small>
    </h1>
    <ol class="breadcrumb">
      <li><a href="/"><i class="fa fa-dashboard"></i> 首页</a></li>
    {{range .node.Parents}}
      <li><a href="{{.Path}}"> {{.Name}}</a></li>
    {{end}}
      <li class="active">{{.node}}</li>
    </ol>
  </section>
{{end}}

{{define "modal"}}
  <div class="modal" id="modal-edit">
    <div class="modal-dialog">
      <div class="modal-content box">
      </div>
    </div>
  </div>
  <div class="modal" id="modal-detail">
    <div class="modal-dialog modal-lg">
      <div class="modal-content box">
      </div>
    </div>
  </div>
  <div class="modal" id="modal-confirm">
    <div class="modal-dialog modal-sm">
      <div class="modal-content">
        <div class="modal-header">
          <a class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></a>
          <h4 class="modal-title">确定？</h4>
        </div>
        <div class="modal-footer">
          <a class="btn btn-default" data-dismiss="modal">取消</a>
          <a type="submit" class="btn btn-danger">确定</a>
        </div>
      </div>
    </div>
  </div>
{{end}}

{{define "paginator"}}
  {{if and . .page}}
  <span style="padding-left: 10px;">共 {{.page.Nums}} 条记录</span>
  {{end}}
  {{if and . .page .page.HasPages}}
  <ul class="pagination pagination-sm no-margin pull-right">
    {{if .page.HasPrev}}
      <li><a href="{{.page.PageLinkFirst}}"><i class="fa fa-angle-double-left"></i></a></li>
    {{end}}
    {{range $index, $page := .page.Pages}}
      <li{{if $.page.IsActive .}} class="active"{{end}}>
        <a href="{{$.page.PageLink $page}}">{{$page}}</a>
      </li>
    {{end}}
    {{if .page.HasNext}}
      <li><a href="{{.page.PageLinkLast}}"><i class="fa fa-angle-double-right"></i></a></li>
    {{end}}
  </ul>
  {{end}}
{{end}}