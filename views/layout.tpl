{{define "header"}}
<meta charset="utf-8">
<meta http-equiv="X-UA-Compatible" content="IE=edge">
<title>{{.}}</title>
<meta content="noarchive" name="robots">
<meta content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no" name="viewport">
<script src="/static/js/admin.js?v{{version}}" type="text/javascript"></script>

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
{{end}}

{{define "sidebar"}}
  {{range .m}}
  {{if not (.HasGroup $.Group)}}
  {{else if .Child}}
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
{{if and . .page}}
<span style="padding-left: 10px;">共 {{.page.Nums}} 条记录</span>
{{end}}
{{if and . .page .page.HasPages}}
<ul class="pagination pagination-sm no-margin pull-right">
  {{if .page.HasPrev}}
  <li><a href="{{.page.PageLinkFirst}}"><i class="fa fa-angle-double-left"></i></a></li>
  {{end}}
  {{range $index, $page := .page.Pages}}
  <li {{if $.page.IsActive .}} class="active" {{end}}>
    <a href="{{$.page.PageLink $page}}">{{$page}}</a>
  </li>
  {{end}}
  {{if .page.HasNext}}
  <li><a href="{{.page.PageLinkLast}}"><i class="fa fa-angle-double-right"></i></a></li>
  {{end}}
</ul>
{{end}}
{{end}}