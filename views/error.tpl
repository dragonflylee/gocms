<!DOCTYPE html>
<html>

<head>
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <title>{{.Code}}</title>
  <meta content="noarchive" name="robots">
  <meta content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no" name="viewport">
  <link href="//cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/3.4.1/css/bootstrap.min.css" rel="stylesheet">
  <link href="//cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css" rel="stylesheet">
  <link href="//cdnjs.cloudflare.com/ajax/libs/admin-lte/2.4.10/css/AdminLTE.min.css" rel="stylesheet">
  <link href="//cdnjs.cloudflare.com/ajax/libs/admin-lte/2.4.10/css/skins/_all-skins.min.css" rel="stylesheet">
</head>

<body class="hold-transition skin-blue layout-top-nav">
  <div class="wrapper">
    <div class="content-wrapper">
      <section class="content">
        <div class="error-page">
          <h2 class="headline text-red">{{.Code}}</h2>
          <div class="error-content">
            <h3><i class="fa fa-warning text-red"></i> {{.Text}}</h3>
            <p>{{.Msg}}</p>
            <form class="search-form">
              <div class="input-group">
                <input type="text" name="search" class="form-control" placeholder="搜索">
                <div class="input-group-btn">
                  <button type="submit" class="btn btn-danger btn-flat"><i class="fa fa-search"></i></button>
                </div>
              </div>
            </form>
          </div>
        </div>
      </section>
    </div>
  </div>
  <script src="//cdnjs.cloudflare.com/ajax/libs/jquery/1.12.4/jquery.min.js"></script>
  <script src="//cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/3.4.1/js/bootstrap.min.js"></script>
</body>

</html>