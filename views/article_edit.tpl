<!DOCTYPE html>
<html>
<head>
  {{- template "header" .Node}}
  <link rel="stylesheet" href="//cdnjs.cloudflare.com/ajax/libs/simplemde/1.11.2/simplemde.min.css" />
</head>
<body class="hold-transition skin-blue sidebar-mini">
  <div class="wrapper">
    {{- template "navbar" .}}
    <div class="content-wrapper">
      {{- template "title" .}}
      <section class="content box-profile">
        {{- with .Data }}
        <form class="box" action="/article/edit/{{.ID}}" method="post">
          <div class="box-footer form-inline">
            <div class="form-group">
              <label class="control-label">标题</label>
              <input name="title" class="form-control" value="{{.Title}}" required>
            </div>
            <div class="form-group">
              <label class="control-label">副标题</label>
              <input name="short_title" class="form-control" value="{{.Remark}}" required>
            </div>
            <button type="submit" class="btn bg-orange pull-right">发布</button>
          </div>
          <div class="box-body">
            <textarea name="content" id="md" rows="20" required>{{.Content}}</textarea>
          </div>
        </form>
        {{- end }}
      </section>
    </div>
    {{- template "footer"}}
  </div>
  <script src="//cdnjs.cloudflare.com/ajax/libs/simplemde/1.11.2/simplemde.min.js"></script>
  <script>
    var simplemde = new SimpleMDE({
      element: document.getElementById('md'),
      autoDownloadFontAwesome: false,
      status: false
    });
  </script>
</body>
</html>