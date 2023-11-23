# GoCMS
基于 [AdminLTE](https://adminlte.io)、[Gin](https://gin-gonic.com/) 和 [Gorm](https://gorm.io) 实现的内容管理系统

## 获取安装

```bash
GO111MODULE=on go get -v github.com/dragonflylee/gocms
```

## 目录结构

```conf
├─handler             # 控制器
├─model               # 数据模型
├─pkg
│  ├─captcha          # 行为验证
│  ├─config           # 配置中心
│  ├─mail             # 邮件发送
│  ├─web
│  └─util
└─themes
    ├─admin
    └─front
```

## 前端框架

1. 表单校验

使用 [jQeury Validate](https://jqueryvalidation.org/documentation/) 校验表单，支持使用 `data-rule` 标签配置规则，示例如下
```html
<form method="post">
    <input name="username" type="text" data-msg-required="登录名称不能为空" required>
    <input name="password" type="password" id="register_password" placeholder="请输入新密码" data-rule-passwd="true" required>
    <input name="rpasswd" type="password" data-rule-equalTo="#register_password" data-msg-equalTo="两次输入的密码不一致" placeholder="请再次输入新密码" required>
    <button type="submit">保存</button>
</form>
```

2. 模态框

使用的 [Bootstrap](https://v3.bootcss.com/javascript/#modals) 的 modal 组件。

```html
<span class="btn btn-xs bg-navy pull-right" data-href="/group/edit/1" data-target="#modal-node" data-toggle="modal"><i class="fa fa-edit"></i></span>
```

```bash
docker run --name postgres --restart=always --network host \
  -e POSTGRES_PASSWORD=123456 -e POSTGRES_DB=registry -e TZ=UTC \
  -v postgres:/var/lib/postgresql/data -d postgres:14-alpine

docker run --rm --net host -it dbcliorg/pgcli -h 127.0.0.1 -p 5432 -u postgres -W
```