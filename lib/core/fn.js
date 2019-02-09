const buildArgsCtx = (mapping, args) => mapping
  .reduce((acc, k, idx) => ({
    ...acc,
    [k]: args[idx],
  }), {});

/**
 * @name fn
 * @sig (fn [args] (body))
 * @analogueOf function, arrow function
 * @example
 *  (def inc
 *    (fn [x] (add x 1)))
*/
function fn(argsMapping, f) {
  return (...args) => f.call({ ...this, ...buildArgsCtx(argsMapping, args) });
}

module.exports = ['fn', fn];
