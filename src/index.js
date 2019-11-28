window.Raphael = require('raphael')
require('morris.js/morris.js')
require('admin-lte')

$(document).ready(function () {
  var area = new Morris.Line({
    element: $('#area-chart .chart'),
    xkey: 'y',
    hideHover: 'auto',
    ykeys: ['item1', 'item2'],
    labels: ['类型1', '类型2'],
    lineColors: ['#a0d0e0', '#3c8dbc']
  });
  $('#area-chart').boxRefresh({
    loadInContent: false, responseType: 'json',
    onLoadStart: function () {
      this.get(0).options.params =
        $('#area-chart input,select').serializeArray();
    },
    onLoadDone: function (e) {
      area.setData(e.data);
    }
  })
})