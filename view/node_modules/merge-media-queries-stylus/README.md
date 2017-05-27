CSS MQPacker Stylus
-------------------

An [node-css-mqpacker](https://github.com/hail2u/node-css-mqpacker) plugin for stylus.

### Installation

You can install through npm as such: `npm i daliborgogic/merge-media-queries-stylus`

### Usage

You can include merge-media-queries-stylus as a normal stylus plugin. Basic example below:

```js
var stylus = require('stylus');
var mergeMediaQueries = require('merge-media-queries-stylus');

stylus(css)
  .use(mergeMediaQueries())
  .render(function(err, output) {
    console.log(output);
  });
```

If you'd like to install globally and run from the command line, you can do it like this:

```js
npm i -g daliborgogic/merge-media-queries-stylus
stylus -u merge-media-queries-stylus -c foo.styl
```
