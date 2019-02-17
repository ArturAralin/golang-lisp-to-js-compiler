const E = require('./E')[1];

const buildArgsCtx = (mapping, args) => mapping
  .reduce((acc, k, idx) => ({
    ...acc,
    [k]: args[idx],
  }), {});

const identity = v => v;
/**
 * @name fn
 * @sig (fn [args] (body))
 * @analogueOf function, arrow function
 * @example
 *  (def inc
 *    (fn [x] (add x 1)))
*/
function fn(argsMapping, v) {
  return (...args) => {
    const ctx = { ...this, ...buildArgsCtx(argsMapping, args) };

    if (typeof v !== 'function') {
      return E(identity, v).call(ctx);
    }

    return v.call(ctx);
  };
}

module.exports = ['fn', fn];
