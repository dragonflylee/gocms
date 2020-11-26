# GoCMS
基于 [AdminLTE](https://adminlte.io)、[Gorilla](http://www.gorillatoolkit.org) 和 [Gorm](http://gorm.io) 实现的内容管理系统

## 获取安装

```bash
git clone https://github.com/dragonflylee/gocms.git && cd gocms && make
```

## 目录结构

 ├── handler    Web业务逻辑  
 ├── model      数据操作层  
 ├── static     前端静态资源  
 ├── util       工具函数  
 ├── views      模板页面  
 ├── src        js源文件  
 ├── main.go    路由入口  
 └── nodes.json 节点初始化数据  

## 前端框架

1. 表单校验

使用 [jQeury Validate](https://jqueryvalidation.org/documentation/) 校验表单，支持使用 `data-rule` 标签配置规则，示例如下
```html
<form method="post">
    <input name="username" type="text" data-msg-required="登录名称不能为空" required>
    <input name="password" type="password" id="register_password" placeholder="请输入新密码" data-rule-regexPasswd="true" required>
    <input name="rpasswd" type="password" data-rule-equalTo="#register_password" data-msg-equalTo="两次输入的密码不一致" placeholder="请再次输入新密码" required>
    <button type="submit">保存</button>
</form>
```

2. 模态框

使用的 [Bootstrap](https://v3.bootcss.com/javascript/#modals) 的 modal 组件。

```html
<span class="btn btn-xs bg-navy pull-right" data-href="/group/edit/1" data-target="#modal-node" data-toggle="modal"><i class="fa fa-edit"></i></span>
```