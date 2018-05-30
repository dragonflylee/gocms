<form method="POST" action="/admin/group/{{.group.ID}}" class="form-horizontal">
  <div class="modal-header">
    <button type="button" class="close" data-dismiss="modal" aria-label="Close">
      <span aria-hidden="true">×</span></button>
    <h4 class="modal-title">编辑角色</h4>
  </div>
  <div class="modal-body">
    <div class="form-group">
      <label class="col-sm-3 control-label">名称</label>
      <div class="col-sm-5">
        <input type="text" class="form-control" name="name" value="{{.group.Name}}" required>
      </div>
    </div>
    <div class="form-group">
      <label class="col-sm-3 control-label">名称</label>
      <div class="col-sm-5 jstree">
        <ul>
          {{template "nodetree" .node}}
        </ul>
      </div>
    </div>
  </div>
  <div class="modal-footer">
    <a class="btn btn-default" data-dismiss="modal">取消</a>
    <button type="submit" class="btn btn-danger">确定</button>
  </div>
</form>

{{define "nodetree"}}
  {{range .}}
    <li data-jstree='{"opened":true,"selected":true,"icon":"{{.Icon}}"}'>{{.Name}}
    {{if .Child}}
      <ul>{{template "nodetree" .Child}}</ul>
    {{end}}
    </li>
  {{end}}
{{end}}
