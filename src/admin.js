const $ = require('jquery')
window.jQuery = $;
window.$ = $;

require('bootstrap')
require('admin-lte')
require('jquery-form')
require('jquery-validation')
require('jquery-validation/dist/localization/messages_zh')

require('bootstrap-switch')
require('icheck')
require('select2')
require('select2/dist/js/i18n/zh-CN')
require('bootstrap-filestyle')
require('bootstrap-datepicker')
require('bootstrap-datepicker/dist/locales/bootstrap-datepicker.zh-CN.min.js')

window.md5 = require('blueimp-md5')

import './admin.css'

// 手机号校验
$.validator.addMethod('mobile', function (value, element) {
  var mobile = /^((\+?86)|(\(\+86\)))?(13[0-9][0-9]{8}|15[0-9][0-9]{8}|18[0-9][0-9]{8}|17[0678][0-9]{8}|147[0-9]{8}|1349[0-9]{7})$/;
  return this.optional(element) || (value.length == 11 && mobile.test(value));
}, '请填写正确的手机号码');
// 密码验证正则表达式
$.validator.addMethod('regexPasswd', function (value, element) {
  return this.optional(element) || /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[^]{8,}$/.test(value);
}, '密码至少包大小写字母及数字，长度至少8位');

window.Admin = {
  // 通用表单验证
  validate: function (selector, options) {
    $(selector).each(function (i, form) {
      $(form).validate({
        errorElement: 'span', focusInvalid: false,
        errorClass: 'help-inline help-block',
        highlight: function (el) {
          $(el).closest('.form-group').removeClass('has-success').addClass('has-error');
        },
        unhighlight: function (el) {
          $(el).closest('.form-group').removeClass('has-error').addClass('has-success');
        },
        errorPlacement: function (err, el) {
          var target = el.data('target');
          if (target) {
            err.insertAfter($(target));
          } else if (el.prop('type') === 'checkbox') {
            err.appendTo(el.closest('.checkbox-inline'));
          } else if (el.prop('type') === 'radio') {
            err.appendTo(el.closest('.radio-inline'));
          } else {
            err.insertAfter(el);
          }
        },
        submitHandler: function () {
          var $overlay = $('<div class="overlay"><i class="fa fa-refresh fa-spin"></i></div>');
          $(form).hasClass('box') ? $(form).append($overlay) : $(form).closest('.box').append($overlay);

          var target = $(form).data('target');
          if (typeof target !== 'undefined') target = $(form).find(target)
          else if ($(form).hasClass('box')) target = $(form).find('.box-footer')
          else if ((target = $(form).find('.modal-footer')).length == 0) target = $(form);

          $(form).ajaxSubmit($.extend(true, {
            dataType: 'json',
            complete: function (xhr, resp) {
              $overlay.remove();
            },
            success: function (resp) {
              if (resp.code != 200) Admin.alert({
                container: target, type: 'danger', message: resp.msg
              });
              else if (typeof resp.data === 'string')
                window.location = resp.data;
              else if (typeof resp.msg !== 'string')
                location.reload();
              else Admin.alert({
                container: target,
                message: resp.msg, closeInSeconds: 3
              });
            },
            error: function () {
              Admin.alert({
                container: target,
                type: 'danger', message: '请求失败'
              });
            }
          }, options));
        }
      });
      $('[data-rule]', $(form)).each(function (i, o) {
        $(o).rules('add', $.parseJSON($(o).data('rule').replace(/'/g, '"')));
      });
    })
  },
  // 表单提示
  alert: function (options) {
    options = $.extend(true, {
      container: '', // alerts parent container(by default placed after the page breadcrumbs)
      place: 'prepend', // "append" or "prepend" in container 
      type: 'success', // alert's type
      message: '', // alert's message
      close: true, // make alert closable
      closeInSeconds: 0, // auto close after defined seconds
      icon: '' // put icon before the message
    }, options);

    var id = 'prefix_' + Math.floor(Math.random() * (new Date()).getTime());;
    var html = '<div id="' + id + '" class="custom-alerts text-left alert alert-' + options.type + ' fade in">' + (options.close ? '<button type="button" class="close" data-dismiss="alert" aria-hidden="true">&times;</button>' : '') + (options.icon !== '' ? '<i class="fa-lg fa fa-' + options.icon + '"></i>  ' : '') + options.message + '</div>';

    $('.custom-alerts').remove();

    if (!options.container) {
      $('.content-header').after(html);
    } else {
      if (options.place == 'append') {
        $(options.container).append(html);
      } else {
        $(options.container).prepend(html);
      }
    }
    if (options.closeInSeconds > 0) {
      setTimeout(function () {
        $('#' + id).remove();
        window.location.reload();
      }, options.closeInSeconds * 1000);
    }
    return id;
  },
  modal: function ($selector, options) {
    $($selector).on('show.bs.modal', function (e) {
      if (e.namespace === 'bs.modal') {
        $(this).find('.modal-content').load($(e.relatedTarget).data('href'),
          function (resp, status) {
            if (status == 'success') {
              Admin.init($(this));
              $(this).find('form').each(function (i, el) {
                Admin.validate(el, options);
              })
            }
          }).empty();
      }
    }).on('hide.bs.modal', function () {
      $(this).removeData('bs.modal');
    });
  },
  init: function ($container) {
    $('[data-toggle="tooltip"]', $container).tooltip();

    $('[data-toggle="popover"]', $container).popover();

    $('.select2', $container).select2({
      language: 'zh-CN'
    });

    if (jQuery().jstree) {
      $('.jstree', $container).jstree({
        "core": { "themes": { "variant": "large" } },
        "checkbox": { "cascade": "undetermined", "three_state": false },
        "plugins": ["checkbox"]
      });
    }

    $('.icheck :checkbox, .icheck :radio', $container).iCheck({
      checkboxClass: 'icheckbox_minimal-blue',
      radioClass: 'iradio_minimal-blue',
      increaseArea: '20%' // optional
    });
    // 表格复选框
    $('th :checkbox', $container).on('ifChanged', function () {
      var set = $(this).data('set');
      var checked = $(this).is(':checked');
      $(set).each(function () {
        if (checked) {
          $(this).iCheck('check');
          $(this).parents('tr').addClass('active');
        } else {
          $(this).iCheck('uncheck');
          $(this).parents('tr').removeClass('active');
        }
      });
    });
    $('tr :checkbox').on('ifChanged', function () {
      $(this).parents('tr').toggleClass('active');
    });

    $('.make-switch', $container).bootstrapSwitch({
      onSwitchChange: function (e, state) {
        var target = $(e.currentTarget).data('target');
        if (target) $(target).modal('show', e.currentTarget).one('hide.bs.modal', function () {
          $(e.currentTarget).bootstrapSwitch('toggleState', 'skip');
        });
      }
    });
  }
}

// 模态框内分页
$(document).on('click', '.modal-content .pagination a,.modal-content .nav-tabs-custom a', function (e) {
  if (e.target.target == '_blank') return;
  e.preventDefault();
  $(this).closest('.modal-content').load(e.target.href, function () {
    Admin.init($(this));
    $(this).find('form').each(function (i, el) {
      Admin.validate(el);
    })
  });
})

if (typeof Storage !== 'undefined') {
  // 侧边栏
  $(document).on('expanded.pushMenu', function () {
    localStorage.removeItem('sidebar');
  }).on('collapsed.pushMenu', function (e) {
    localStorage.setItem('sidebar', 'sidebar-collapse');
  })
  $(document).ready(function () {
    var sidebar = localStorage.getItem('sidebar');
    if (sidebar) $(document.body).addClass(sidebar);
  })
}

$(document).ready(function () {
  // 初始化插件
  Admin.init();
  // 模态框表单
  $('.modal:has(form)').on('show.bs.modal', function (e) {
    $(this).find('.modal-title').text(e.relatedTarget.title);
    $(this).find('form').each(function (i, el) {
      Admin.validate(el, {
        'url': $(e.relatedTarget).data('href')
      });
    })
  }).on('hidden.bs.modal', function () {
    $(this).find('.custom-alerts').remove();
    $(this).find('form').each(function (i, el) {
      $(el).off('.validate').removeData('validator').resetForm();
    })
  });
  // 远端模态框
  Admin.modal('#modal-edit, #modal-detail');
  // 页面表单
  Admin.validate('form:has([name="password"])', {
    beforeSubmit: function (arr) {
      for (var i = 0; i < arr.length; i++)
        if (arr[i].name == 'password')
          arr[i].value = md5(arr[i].value);
    }
  });
})