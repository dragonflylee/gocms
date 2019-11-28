require('jquery-backstretch')

window.onsubmit = function () {
  $('form').trigger('submit');
}

$(document).on('click', '.btn-verify', function (e) {
  var timeout = 90, btn = $(e.target).
    attr('disabled', true).text(timeout + ' s'),
    count = setInterval(function () {
      if (--timeout > 0)
        return btn.text(timeout + ' s')
      clearInterval(count);
      btn.text('点击获取').attr('disabled', false);
    }, 1000);
  // send code
  var form = btn.closest('form').ajaxSubmit({
    url: btn.data('href'), dataType: 'json',
    success: function (resp) {
      if (resp.code != 200) return Admin.alert({
        container: form, type: 'danger', message: resp.msg
      });
    }
  })
})

$(document).ready(function () {
  $.backstretch('/bingpic');
})

