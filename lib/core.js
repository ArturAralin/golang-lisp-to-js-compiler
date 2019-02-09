const rt = require('require-tree');

const fns = rt('./core');
const symbol = fns.symbol[1];

const core = Object
  .values(fns)
  .reduce((acc, [name, f]) => ({
    ...acc,
    [symbol(name)]: f,
  }), {});

core.$ = symbol;

module.exports = core;
