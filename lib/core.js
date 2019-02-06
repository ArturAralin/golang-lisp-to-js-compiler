const rt = require('require-tree');

const fns = rt('./core');

const core = Object
  .keys(fns)
  .reduce((acc, fnName) => {
    const fn = fns[fnName];

    return {
      ...acc,
      [fnName]: fn,
    };
  }, {});

module.exports = core;
