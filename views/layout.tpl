{{define "header"}}
<meta charset="utf-8">
<meta http-equiv="X-UA-Compatible" content="IE=edge">
<title>{{.}}</title>
<meta content="noarchive" name="robots">
<meta content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no" name="viewport">
<link href="//cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/3.4.1/css/bootstrap.min.css" rel="stylesheet">
<link href="//cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css" rel="stylesheet">

<link href="//cdnjs.cloudflare.com/ajax/libs/bootstrap-switch/3.3.4/css/bootstrap3/bootstrap-switch.min.css" rel="stylesheet">
<link href="//cdnjs.cloudflare.com/ajax/libs/iCheck/1.0.2/skins/all.css" rel="stylesheet">
<link href="//cdnjs.cloudflare.com/ajax/libs/select2/4.0.5/css/select2.min.css" rel="stylesheet">
<link href="//cdnjs.cloudflare.com/ajax/libs/bootstrap-datepicker/1.8.0/css/bootstrap-datepicker3.min.css" rel="stylesheet">

<link href="//cdnjs.cloudflare.com/ajax/libs/admin-lte/2.4.10/css/AdminLTE.min.css" rel="stylesheet">
<link href="//cdnjs.cloudflare.com/ajax/libs/admin-lte/2.4.10/css/skins/_all-skins.min.css" rel="stylesheet">
<link href="/static/css/custom.min.css?v{{version}}" rel="stylesheet" type="text/css">

{{html "<!--[if lt IE 9]>"}}
  <script src="//cdnjs.cloudflare.com/ajax/libs/html5shiv/3.7.3/html5shiv.min.js"></script>
  <script src="//cdnjs.cloudflare.com/ajax/libs/respond.js/1.4.2/respond.min.js"></script>
{{html "<![endif]-->"}}
{{end}}

{{define "navbar"}}
<header class="main-header">
  <a href="#" class="logo">
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
            <img src="{{.User.Headpic}}" class="user-image" alt="用户头像">
            <span class="hidden-xs">{{.User.Email}}</span>
          </a>
          <ul class="dropdown-menu">
            <li class="user-header">
              <img src="{{.User.Headpic}}" class="img-circle" alt="用户头像">
              <p>
                {{.User.Email}}
                <small>{{.User.Group.Name}}</small>
                <small>上次登录 {{date .User.LastLogin}}</small>
              </p>
            </li>
            <li class="user-footer">
              <div class="pull-left">
                <a href="{{urlfor "Profile"}}" class="btn btn-default btn-flat">个人中心</a>
              </div>
              <div class="pull-right">
                <a href="{{urlfor "Logout"}}" class="btn btn-default btn-flat">注销</a>
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
        <img src="{{.User.Headpic}}" class="img-circle" alt="用户头像">
      </div>
      <div class="pull-left info">
        <P>{{.User.Group.Name}}</P>
        <a href="#" title="{{.User.Email}}"><i class="fa fa-circle text-success"></i> 在线</a>
      </div>
    </div>
    <form class="sidebar-form">
      <div class="input-group">
        <input type="text" name="q" class="form-control" placeholder="搜索..." value="{{.Form.Get "q"}}">
        <span class="input-group-btn">
          <button type="submit" id="search-btn" class="btn btn-flat"><i class="fa fa-search"></i></button>
        </span>
      </div>
    </form>
    <ul class="sidebar-menu" data-widget="tree">
      {{template "sidebar" .Menu.Assign .User.GroupID .Node}}
    </ul>
  </section>
</aside>
{{end}}

{{define "footer"}}
<footer class="main-footer">
  <div class="pull-right hidden-xs">
    Powered by <b>{{version}}</b>
  </div>
  <strong>版权所有 &copy; 2019 <a href="https://github.com/dragonflylee/gocms">GoCMS</a>.</strong>
</footer>
<script src="//cdnjs.cloudflare.com/ajax/libs/jquery/1.12.4/jquery.min.js"></script>
<script src="//cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/3.4.1/js/bootstrap.min.js"></script>
<script src="//cdnjs.cloudflare.com/ajax/libs/jQuery-slimScroll/1.3.8/jquery.slimscroll.min.js"></script>
<script src="//cdnjs.cloudflare.com/ajax/libs/fastclick/1.0.6/fastclick.min.js"></script>
<script src="//cdnjs.cloudflare.com/ajax/libs/admin-lte/2.4.10/js/adminlte.min.js"></script>

<script src="//cdnjs.cloudflare.com/ajax/libs/jquery.form/4.2.2/jquery.form.min.js"></script>
<script src="//cdnjs.cloudflare.com/ajax/libs/jquery-validate/1.19.0/jquery.validate.min.js"></script>
<script src="//cdnjs.cloudflare.com/ajax/libs/Sortable/1.8.3/Sortable.min.js"></script>
<script src="//cdnjs.cloudflare.com/ajax/libs/blueimp-md5/2.10.0/js/md5.min.js"></script>

<script src="//cdnjs.cloudflare.com/ajax/libs/bootstrap-switch/3.3.4/js/bootstrap-switch.min.js"></script>
<script src="//cdnjs.cloudflare.com/ajax/libs/iCheck/1.0.2/icheck.min.js"></script>
<script src="//cdnjs.cloudflare.com/ajax/libs/select2/4.0.5/js/select2.min.js"></script>
<script src="//cdnjs.cloudflare.com/ajax/libs/select2/4.0.5/js/i18n/zh-CN.js"></script>
<script src="//cdnjs.cloudflare.com/ajax/libs/bootstrap-datepicker/1.8.0/js/bootstrap-datepicker.min.js"></script>
<script src="//cdnjs.cloudflare.com/ajax/libs/bootstrap-datepicker/1.8.0/locales/bootstrap-datepicker.zh-CN.min.js"></script>

<script src="/static/js/global.js?v{{version}}" type="text/javascript"></script>
{{end}}

{{define "sidebar"}}
  {{range .m}}
  {{if and .Status (.HasGroup $.Group)}}
  {{if and .Child (not .Path)}}
  <li class="treeview {{if $.Node.HasParent .ID}}active{{end}}">
    <a href="#">
      <i class="{{.Icon}}"></i> <span>{{.Name}}</span>
      <span class="pull-right-container">
        <i class="fa fa-angle-left pull-right"></i>
      </span>
    </a>
    <ul class="treeview-menu">
      {{template "sidebar" .Child.Assign $.Group $.Node}}
    </ul>
  </li>
  {{else}}
  <li {{if $.Node.HasParent .ID}} class="active" {{end}}>
    <a href="{{urlfor .Path}}">
      <i class="{{.Icon}}"></i>
      <span>{{.Name}}</span>
    </a>
  </li>
  {{end}}
  {{end}}
  {{end}}
{{end}}

{{define "title"}}
<section class="content-header">
  <h1>
    {{.Node}}
    <small>{{.Node.Remark}}</small>
  </h1>
  <ol class="breadcrumb">
    <li><a href="#"><i class="fa fa-dashboard"></i> 首页</a></li>
    {{range .Node.Parents}}
    <li><a href="{{urlfor .Path}}"> {{.Name}}</a></li>
    {{end}}
    <li class="active">{{.Node}}</li>
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
    <form class="modal-content">
      <div class="modal-header">
        <a class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></a>
        <h4 class="modal-title">确定？</h4>
      </div>
      <div class="modal-footer">
        <a class="btn btn-default" data-dismiss="modal">取消</a>
        <button type="submit" class="btn btn-danger">确定</button>
      </div>
    </form>
  </div>
</div>
{{end}}

{{define "paginator"}}
{{if and . .Page}}
<span style="padding-left: 10px;">共 {{.Page.Nums}} 条记录</span>
{{end}}
{{if and . .Page .Page.HasPages}}
<ul class="pagination pagination-sm no-margin pull-right">
  {{if .Page.HasPrev}}
  <li><a href="{{.Page.PageLinkFirst}}"><i class="fa fa-angle-double-left"></i></a></li>
  {{end}}
  {{range $index, $page := .Page.Pages}}
  <li {{if $.Page.IsActive .}} class="active" {{end}}>
    <a href="{{$.Page.PageLink $page}}">{{$page}}</a>
  </li>
  {{end}}
  {{if .Page.HasNext}}
  <li><a href="{{.Page.PageLinkLast}}"><i class="fa fa-angle-double-right"></i></a></li>
  {{end}}
</ul>
{{end}}
{{end}}