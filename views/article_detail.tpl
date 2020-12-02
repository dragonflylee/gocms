<!DOCTYPE html>
<html>
<head>
  {{- template "header" .Node}}
  <link rel="stylesheet" href="//cdnjs.cloudflare.com/ajax/libs/prism/1.22.0/themes/prism-tomorrow.min.css" />
</head>
<body class="hold-transition skin-blue sidebar-mini">
  <div class="wrapper">
    {{- template "navbar" .}}
    <div class="content-wrapper">
      {{- template "title" .}}
      <section class="content">
        {{- with .Data }}
        <div class="box">
          <div class="box-body table-responsive">
            {{.Render}}
          </div>
        </div>
        {{- end }}
      </section>
    </div>
    {{- template "footer"}}
  </div>
  <script src="//cdnjs.cloudflare.com/ajax/libs/prism/1.22.0/prism.min.js"></script>
  <script src="//cdnjs.cloudflare.com/ajax/libs/prism/1.22.0/plugins/autoloader/prism-autoloader.min.js"></script>
  <script type="text/javascript">
    $(document).ready(function () {
      $('.box-body table').addClass('table');
      $('.box-body img').addClass('img-responsive');
    })
  </script>
</body>
</html>