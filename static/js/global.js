if (jQuery().validate) {
  // 表单验证提示中文化
  $.extend($.validator.messages,{required:"该项不能为空",remote:"请修正此字段",email:"请输入有效的电子邮件地址",url:"请输入有效的网址",date:"请输入有效的日期",dateISO:"请输入有效的日期 (YYYY-MM-DD)",number:"请输入有效的数字",digits:"只能输入数字",creditcard:"请输入有效的信用卡号码",equalTo:"你的输入不相同",extension:"请输入有效的后缀",maxlength:$.validator.format("最多可以输入 {0} 个字符"),minlength:$.validator.format("最少要输入 {0} 个字符"),rangelength:$.validator.format("请输入长度在 {0} 到 {1} 之间的字符串"),range:$.validator.format("请输入范围在 {0} 到 {1} 之间的数值"),max:$.validator.format("请输入不大于 {0} 的数值"),min:$.validator.format("请输入不小于 {0} 的数值")});
  // 手机号校验
  $.validator.addMethod("mobile", function(value, element) {
    var mobile = /^((\+?86)|(\(\+86\)))?(13[0-9][0-9]{8}|15[0-9][0-9]{8}|18[0-9][0-9]{8}|17[0678][0-9]{8}|147[0-9]{8}|1349[0-9]{7})$/;
    return this.optional(element) || (value.length == 11 && mobile.test(value)); 
  }, "请填写正确的手机号码");
  // 密码验证正则表达式
  $.validator.addMethod("regexPasswd", function(value, element) {
    return this.optional(element) || /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[^]{8,}$/.test(value);
  }, "密码至少包大小写字母及数字，长度至少8位");
}

var Admin = {
  // 通用表单验证
  validate: function(form, options) {
    if (jQuery().validate) {
      var rules = {}, messages = {};
      form.find('[data-rule]').each(function(i, o) {
        rules[$(o).attr('name')] = $.parseJSON($(o).data('rule').replace(/'/g,'"'));
      });
      form.find('[data-message]').each(function(i, o) {
        messages[$(o).attr('name')] = $.parseJSON($(o).data('message').replace(/'/g,'"'));
      });
      form.validate({
        errorElement: 'span',
        errorClass: 'help-inline help-block',
        focusInvalid: false,
        rules: rules,
        messages: messages,
        highlight: function (element) { 
          $(element).closest('.form-group').addClass('has-error');
        },
        success: function (label) {
          label.closest('.form-group').removeClass('has-error');
          label.remove();
        },
        errorPlacement: function (error, element) {
          var target = element.data("target");
          if (target) {
            error.insertAfter($(target));
          } else if (element.prop("type") === "checkbox") {
            error.appendTo(element.closest(".checkbox-inline"));
          } else if (element.prop("type") === "radio") {
            error.appendTo(element.closest(".radio-inline"));
          } else {
            error.insertAfter(element);
          }
        },
        submitHandler: function(form) {
          Admin.form(form, options)
        }
      });
    } else {
      $(form).submit(function() {
        Admin.form(form, options);
        return false;
      });
    }
  },
  // 表单提示
  alert: function(options) {
    options = $.extend(true, {
      container: "", // alerts parent container(by default placed after the page breadcrumbs)
      place: "prepend", // "append" or "prepend" in container 
      type: 'success', // alert's type
      message: "", // alert's message
      close: true, // make alert closable
      closeInSeconds: 0, // auto close after defined seconds
      icon: "" // put icon before the message
    }, options);

    var id = 'prefix_' + Math.floor(Math.random() * (new Date()).getTime());;
    var html = '<div id="' + id + '" class="custom-alerts text-left alert alert-' + options.type + ' fade in">' + (options.close ? '<button type="button" class="close" data-dismiss="alert" aria-hidden="true">×</button>' : '') + (options.icon !== "" ? '<i class="fa-lg fa fa-' + options.icon + '"></i>  ' : '') + options.message + '</div>';

    $('.custom-alerts').remove();

    if (!options.container) {
      $('.content-header').after(html);
    } else {
      if (options.place == "append") {
        $(options.container).append(html);
      } else {
        $(options.container).prepend(html);
      }
    }
    if (options.closeInSeconds > 0) {
      setTimeout(function() {
        $('#' + id).remove();
      }, options.closeInSeconds * 1000);
    }
    return id;
  },
  form: function(form, options) {
    if (jQuery().ajaxSubmit) {
      var modal = $(form).parent().hasClass('modal-content');
      options = $.extend(true, {
        dataType: 'json',
        beforeSubmit: function(arr, $form, options) {
          for (var i = 0; i < arr.length; i++) {
            if (arr[i].name == 'password') {
              arr[i].value = md5(arr[i].value);
            }
          }
          $(form).parents('.box').
            append('<div class="overlay"><i class="fa fa-refresh fa-spin"></i></div>');
        },
        success: function(data) {
          if (typeof data.data == 'string') {
            window.location = data.data;
          } else if (data.code != 200) {
            $('.box .overlay').remove();
            Admin.alert({
              container: modal ? $(form).find('.modal-footer') : form, 
              type: 'danger', 
              message: data.msg 
            });
          } else if (modal) {
            window.location.reload();
          } else {
            $('.box .overlay').remove();
            Admin.alert({ container: form, message: data.msg, closeInSeconds: 3 });
          }
        },
        error: function(xhr, status) {
          $('.box .overlay').remove();
          Admin.alert({
            container: modal ? $(form).find('.modal-footer') : form, 
            type: 'danger', 
            message: status
          });
        }
      }, options);
      $(form).ajaxSubmit(options)
    }
  },
  init: function($container) {
    if (jQuery().tooltip) {
      $('[data-toggle="tooltip"]', $container).tooltip();
    }
    if (jQuery().select2) {
      $(".select2", $container).select2({
        language: "zh-CN"
      });
    }
    if (jQuery().iCheck) {
      $('.icheck :checkbox, .icheck :radio', $container).iCheck({
        checkboxClass: 'icheckbox_minimal-blue',
        radioClass: 'iradio_minimal-blue',
        increaseArea: '20%' // optional
      });
      // 表格复选框
      $('th :checkbox', $container).on('ifChanged', function () {
        var set = $(this).attr("data-set");
        var checked = $(this).is(":checked");
        $(set).each(function () {
          if (checked) {
            $(this).iCheck('check');
            $(this).parents('tr').addClass("active");
          } else {
            $(this).iCheck('uncheck');
            $(this).parents('tr').removeClass("active");
          }
        });
      });
      $('tr :checkbox').on('ifChanged', function () {
        $(this).parents('tr').toggleClass("active");
      });
    }
    if (jQuery().bootstrapSwitch) {
      $('.make-switch', $container).bootstrapSwitch({
        onSwitchChange: function(e, state) {
          var target = $(e.currentTarget).data('target');
          if (target) {
            $(target).modal('show', e.currentTarget).one('hide.bs.modal', function() {
              $(e.currentTarget).bootstrapSwitch('toggleState', 'skip');
            });
          }
        }
      });
    }
    if (jQuery().jstree) {	
      $('.jstree', $container).jstree({	
        "core" : {	
          "themes" : { "variant" : "large" }	
        },	
        "checkbox": {	
          "cascade": "undetermined",	
          "three_state" : false	
        },	
        "plugins" : ["checkbox"]	
      });	
    }
  }
}

// 模态框内分页
$(document).on("click", ".modal-content .pagination a,.modal-content .nav-tabs-custom a", function(e) {
  e.preventDefault();
  $(this).parents('.modal-content').load($(this).attr('href'), function() {
    Admin.init($(this));
    $(this).find('form').each(function(i, form) {
      Admin.validate($(form));
    })
  });
})

$(document).ready(function() {
  // 初始化插件
  Admin.init();
  // 表单验证
  $('.login-box form,.register-box form,.box-body form').each(function(index, form) {
    Admin.validate($(form));
  })
  // 模态框请求
  $('.modal:has(a[type="submit"])').on('show.bs.modal', function (e) {
    var config = {
      url: $(e.relatedTarget).attr('data-href'),
      dataType: 'json',
      success: function(data) {
        if (data.data) {
          window.location = data.data;
        } else if (data.code == 200) {
          window.location.reload();
        } else {
          $('.box .overlay').remove();
          Admin.alert({ 
            container: $(e.currentTarget).find('.modal-footer'), 
            type: 'danger',
            message: data.msg 
          });
        }
      },
      error: function(xhr, status) {
        $('.box .overlay').remove();
        Admin.alert({ 
          container: $(e.currentTarget).find('.modal-footer'), 
          type: 'danger',
          message: status 
        });
      }
    }
    $(this).find('.modal-title').text('确定'+$(e.relatedTarget).attr('title')+'?');
    $(this).find('a[type="submit"]').on('click', function () {
      if ($(e.relatedTarget).hasClass('btn-batch')) {
        config['type'] = 'POST';
        config['data'] = { 'id': $('td :checkbox:checked').map(function(){
          return $(this).val();
        }).get() }
      }
      $(this).parents('.box').append('<div class="overlay"><i class="fa fa-refresh fa-spin"></i></div>');
      $.ajax(config);
    });
  }).on('hide.bs.modal', function() {
    $(this).find('a[type="submit"]').off('click');
    $(this).find('.custom-alerts').remove();
  });
  // 新增对话框
  $('.modal:has(form)').on('show.bs.modal', function (e) {
    $(this).find('.modal-title').text($(e.relatedTarget).attr('title'));
    $(this).find('form').each(function(i, form) {
      Admin.validate($(form), { 'url': $(e.relatedTarget).attr('data-href') });
    })
  }).on("hidden.bs.modal", function() {
    $(this).find('form').each(function(i, form) {
      $(form).off('.validate').removeData('validator').resetForm();
    })
    $(this).find('.custom-alerts').remove();
  });
  // 处理远端加载模态框
  $('#modal-edit, #modal-detail').on('show.bs.modal', function (e) {
    if (e.namespace === 'bs.modal') {
      $(this).find('.modal-content').load($(e.relatedTarget).attr('data-href'), function(resp, status) {
        if (status == 'success') {
          Admin.init($(e.currentTarget));
          $(e.currentTarget).find('form').each(function(i, form) {
            Admin.validate($(form));
          })
        } else {
          $(e.currentTarget).find('.modal-content')
            .html('<div class="jumbotron"><h2 class="text-center">加载失败</h2></div>');
        }
      });
    }
  }).on("hide.bs.modal", function() {
    $(this).removeData("bs.modal");
  });
})