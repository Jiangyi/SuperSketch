var ap = require('css-mqpacker'),
  postcss = require('postcss'),
  map = require('multi-stage-sourcemap'),
  path = require('path')

module.exports = function (opts) {
  if (!opts) { opts = {} }

  var showWarnings = !opts.hideWarnings
  delete opts.hideWarnings

  return function (style) {
    style = this || style
    var filename = style.options.filename

    style.on('end', function (err, css) {
      var process_opts = {
        from: filename,
        to: path.join(
          path.dirname(filename),
          path.basename(filename, path.extname(filename))
        ) + '.css'
      }

      if (style.sourcemap) {
        process_opts.map = { annotation: false }
      }

      // run css-mqpacker
      var res = postcss([ap(opts)]).process(css, process_opts)

      // if sourcemaps are generated, combine the two
      if (res.map && style.sourcemap) {
        var combined_map = map.transfer({
          fromSourceMap: res.map.toString(),
          toSourceMap: style.sourcemap
        })

        // then set the combined result as the new sourcemap
        style.sourcemap = JSON.parse(combined_map)
      }

      if (showWarnings) {
        res.warnings().forEach(console.error)
      }

      // return the css output
      return res.css
    })
  }
}
