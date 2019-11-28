'use strict';

const gulp = require('gulp');
// const uglify = require('gulp-uglify');
const browserify = require('browserify');

gulp.task('build-admin', function (cb) {
  console.log('buiding admin')
  return browserify({
    entries: 'src/js/admin.js',
    debug: false,
    standalone: 'admin'
  }).bundle()
    .on('error', function (error) {
      console.log(error.toString())
    })
    .pipe(dest('static/js'))
})