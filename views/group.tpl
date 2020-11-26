<form method="POST" action="/group/{{.Group.ID}}" class="form-horizontal">
  <div class="modal-header">
    <a class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></a>
    <h4 class="modal-title">编辑角色</h4>
  </div>
  <div class="modal-body">
    <div class="form-group">
      <label class="col-sm-3 control-label">名称</label>
      <div class="col-sm-5">
        <input type="text" class="form-control" name="name" value="{{.Group.Name}}" required>
      </div>
    </div>
    <div class="form-group">
      <label class="col-sm-3 control-label">名称</label>
      <div class="col-sm-5 jstree" name="node">
        <ul>
          {{- template "nodetree" .Node.Assign .Group.ID nil}}
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
{{- range .m}}
<li data-jstree='{"icon":"{{.Icon}}","selected":{{.HasGroup $.Group|print}},"opened":false{{- if .Type}},"disabled":true{{- end}}}' id="{{.ID}}">{{.Name}}
  {{- if .Child}}
  <ul>{{- template "nodetree" .Child.Assign $.Group nil}}</ul>
  {{- end}}
</li>
{{- end}}
{{- end}}