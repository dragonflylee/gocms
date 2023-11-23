// 手机号校验
$.validator.addMethod('mobile', function (value, element) {
  var mobile = /^((\+?86)|(\(\+86\)))?(13[0-9][0-9]{8}|15[0-9][0-9]{8}|18[0-9][0-9]{8}|17[0678][0-9]{8}|147[0-9]{8}|1349[0-9]{7})$/;
  return this.optional(element) || (value.length == 11 && mobile.test(value));
}, 'invalid mobile phone');
// 密码验证正则表达式
$.validator.addMethod('complex', function (value, element) {
  return this.optional(element) || /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[^]{8,}$/.test(value);
}, 'at least 8 characters');

$(function () {
  $('.needs-validation').validate({
    errorElement: 'span', focusInvalid: false,
    errorClass: 'invalid-feedback',
    highlight: function (element) {
      $(element).addClass('is-invalid');
    },
    unhighlight: function (element) {
      $(element).removeClass('is-invalid');
    },
    errorPlacement: function (err, el) {
      var target = el.data('target');
      if (target) {
        $(target).append(err);
      } else if (el.prop('type') === 'checkbox') {
        el.closest('.checkbox-inline').append(err)
      } else if (el.prop('type') === 'radio') {
        el.closest('.radio-inline').append(err);
      } else {
        el.closest('.form-group,.has-validation').append(err);
      }
    }
  });
})

$(document).on('submit', '.needs-validation', function (ev) {
  var url = $(ev.target).attr('action'), data = $(ev.target).serializeArray();
  fetch(url, {method: "POST", mode: "cors", body: data}).
    then(data => data.json()).then(function(resp) {
      if (resp.error !== 'OK') {
        $(document).Toasts('create', {
          class: 'bg-danger', title: document.title, body: resp.message
        })
      } else if (typeof resp.meta === 'string') {
        window.location = resp.meta;
      } else if (typeof resp.message === 'string') {
        $(document).Toasts('create', {
          class: 'bg-success', autohide: true, title: document.title, body: resp.message
        })
      }
    });
  return false;
})