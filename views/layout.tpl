{{define "header"}}
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <title>{{.}}</title>
  <meta content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no" name="viewport">
  <link href="http://cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/3.3.7/css/bootstrap.min.css" rel="stylesheet">
  <link href="http://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css" rel="stylesheet">
  <link href="http://cdnjs.cloudflare.com/ajax/libs/ionicons/2.0.1/css/ionicons.min.css" rel="stylesheet">

  <link href="http://cdnjs.cloudflare.com/ajax/libs/jasny-bootstrap/3.1.3/css/jasny-bootstrap.min.css" rel="stylesheet">
  <link href="http://cdnjs.cloudflare.com/ajax/libs/bootstrap-switch/3.3.4/css/bootstrap3/bootstrap-switch.min.css" rel="stylesheet">
  <link href="http://cdnjs.cloudflare.com/ajax/libs/iCheck/1.0.2/skins/all.css" rel="stylesheet">
  <link href="http://cdnjs.cloudflare.com/ajax/libs/select2/4.0.5/css/select2.min.css" rel="stylesheet">
  <link href="http://cdnjs.cloudflare.com/ajax/libs/lightbox2/2.10.0/css/lightbox.min.css" rel="stylesheet">

  <link href="http://cdnjs.cloudflare.com/ajax/libs/admin-lte/2.4.3/css/AdminLTE.min.css" rel="stylesheet">
  <link href="http://cdnjs.cloudflare.com/ajax/libs/admin-lte/2.4.3/css/skins/_all-skins.min.css" rel="stylesheet">
  <link href="/static/css/custom.min.css" rel="stylesheet" type="text/css">
  
  {{html "<!--[if lt IE 9]>"}}
  <script src="http://cdnjs.cloudflare.com/ajax/libs/html5shiv/3.7.3/html5shiv.min.js"></script>
  <script src="http://cdnjs.cloudflare.com/ajax/libs/respond.js/1.4.2/respond.min.js"></script>
  {{html "<![endif]-->"}}
{{end}}

{{define "navbar"}}
  <header class="main-header">
    <a href="/admin/" class="logo">
      <span class="logo-mini"><b>后</b></span>
      <span class="logo-lg"><b>后台管理</b></span>
    </a>
    <nav class="navbar navbar-static-top">
      <a href="#" class="sidebar-toggle" data-toggle="push-menu" role="button">
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
                    <small>上次登录 {{.user.LastLogin.Format "2006-01-02 15:04:05"}}</small>
                </p>
              </li>
              <li class="user-footer">
                <div class="pull-left">
                  <a href="/admin/profile" class="btn btn-default btn-flat">个人中心</a>
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
      <form action="#" method="get" class="sidebar-form">
        <div class="input-group">
          <input type="text" Username="q" class="form-control" placeholder="搜索...">
              <span class="input-group-btn">
                <button type="submit" Username="search" id="search-btn" class="btn btn-flat"><i class="fa fa-search"></i>
                </button>
              </span>
        </div>
      </form>
      <ul class="sidebar-menu" data-widget="tree">
{{range .menu}}
  {{if .Child}}
        <li class="treeview {{if ($.node.HasParent .ID)}}active{{end}}">
          <a href="#">
            <i class="{{.Icon}}"></i> <span>{{.Name}}</span>
            <span class="pull-right-container">
              <i class="fa fa-angle-left pull-right"></i>
            </span>
          </a>
          <ul class="treeview-menu">
    {{range .Child}}
            <li {{if ($.node.HasParent .ID)}}class="active"{{end}}>
              <a href="{{.Path}}"><i class="fa fa-circle-o"></i> {{.Name}}</a>
            </li>
    {{end}}
          </ul>
        </li>
  {{else}}
        <li {{if ($.node.HasParent .ID)}}class="active"{{end}}>
          <a href="{{.Path}}">
            <i class="{{.Icon}}"></i>
            <span>{{.Name}}</span>
          </a>
        </li>
  {{end}}  
{{end}}
      </ul>
    </section>
  </aside>
{{end}}

{{define "footer"}}
  <footer class="main-footer">
    <div class="pull-right hidden-xs">
      Powered by <b>{{version}}</b>
    </div>
    <strong>版权所有 &copy; 2018 <a href="http://github.com/dragonflylee/gocms">Gocmd</a>.</strong>
  </footer>
  <script src="http://cdnjs.cloudflare.com/ajax/libs/jquery/1.12.4/jquery.min.js"></script>
  <script src="http://cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/3.3.7/js/bootstrap.min.js"></script>
  <script src="http://cdnjs.cloudflare.com/ajax/libs/jQuery-slimScroll/1.3.8/jquery.slimscroll.min.js"></script>
  <script src="http://cdnjs.cloudflare.com/ajax/libs/fastclick/1.0.6/fastclick.min.js"></script>
  <script src="http://cdnjs.cloudflare.com/ajax/libs/admin-lte/2.4.3/js/adminlte.min.js"></script>
  <script src="http://cdnjs.cloudflare.com/ajax/libs/jquery.form/4.2.2/jquery.form.min.js"></script>
  <script src="http://cdnjs.cloudflare.com/ajax/libs/jquery-validate/1.17.0/jquery.validate.min.js"></script>
  <script src="http://cdnjs.cloudflare.com/ajax/libs/blueimp-md5/2.10.0/js/md5.min.js"></script>

  <script src="http://cdnjs.cloudflare.com/ajax/libs/jasny-bootstrap/3.1.3/js/jasny-bootstrap.min.js"></script>
  <script src="http://cdnjs.cloudflare.com/ajax/libs/bootstrap-switch/3.3.4/js/bootstrap-switch.min.js"></script>
  <script src="http://cdnjs.cloudflare.com/ajax/libs/iCheck/1.0.2/icheck.min.js"></script>
  <script src="http://cdnjs.cloudflare.com/ajax/libs/select2/4.0.5/js/select2.min.js"></script>
  <script src="http://cdnjs.cloudflare.com/ajax/libs/select2/4.0.5/js/i18n/zh-CN.js"></script>
  <script src="http://cdnjs.cloudflare.com/ajax/libs/lightbox2/2.10.0/js/lightbox.min.js"></script>
 
  <script src="/static/js/global.min.js?v=20180422" type="text/javascript"></script>
{{end}}

{{define "title"}}
  <section class="content-header">
    <h1>
      {{.node.Name}}
      <small>{{.node.Remark}}</small>
    </h1>
    <ol class="breadcrumb">
      <li><a href="/admin/"><i class="fa fa-dashboard"></i> 首页</a></li>
    {{range .node.Parents}}
      <li><a href="{{.Path}}"> {{.Name}}</a></li>
    {{end}}
      <li class="active">{{.node.Name}}</li>
    </ol>
  </section>
{{end}}

{{define "modal"}}
  <div class="modal modal-remote" id="modal-edit">
    <div class="modal-dialog">
      <div class="modal-content">
      </div>
    </div>
  </div>

  <div class="modal modal-href" id="modal-confirm">
    <div class="modal-dialog modal-sm">
      <div class="modal-content">
        <div class="modal-header">
          <button type="button" class="close" data-dismiss="modal" aria-hidden="true"></button>
          <h4 class="modal-title">确定？</h4>
        </div>
        <div class="modal-footer">
          <a class="btn btn-default" data-dismiss="modal">取消</a>
          <button type="submit" class="btn btn-danger">确定</button>
        </div>
      </div>
    </div>
  </div>
{{end}}

{{define "paginator"}}
  {{if and .page .page.HasPages}}
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