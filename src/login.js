require('jquery-backstretch')

$(document).ready(function () {
    $('img.img-responsive').click(function (e) {
      var src = e.target.src;
      e.target.src = src.substr(0, src.indexOf('?') + 1) +
        'reload=' + new Date().getTime();
    });
    if (jQuery().backstretch) {
      $.backstretch('/bingpic');
    }
    // 页面表单
    Admin.validate('form:has([name="pass"])', {
      beforeSubmit: function (arr) {
        for (var i = 0; i < arr.length; i++)
          if (arr[i].name == 'pass')
            arr[i].value = md5(arr[i].value);
      },
      success: function (resp) {
        if (resp.code == 200) {
          window.location = resp.data;
          return;
        }
        if (typeof resp.data === 'string') {
          $('input[name="id"]').val(resp.data);
          $('img.img-responsive').attr('src', '/captcha/' + resp.data + '.png?')
        }
        Admin.alert({container: $('form'), type: 'danger', message: resp.msg });
      }
    });
  })

