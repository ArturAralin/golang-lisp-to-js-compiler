const buildArgsCtx = (mapping, args) => mapping
  .reduce((acc, k, idx) => ({
    ...acc,
    [k]: args[idx],
  }), {});

// (fn args body)
function fn(argsMapping, f) {
  return (...args) => f.call({ ...this, ...buildArgsCtx(argsMapping, args) });
}

module.exports = fn;
