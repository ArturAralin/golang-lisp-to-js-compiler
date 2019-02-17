/* eslint-disable no-use-before-define */

const handleArray = (ctx, v) => E
  .apply(ctx, [
    (...args) => args,
    ...v,
  ])
  .call(ctx);

// this is executor function
// oh. this function full of pain and tears
function E(f, ...args) {
  return function expression() {
    const computedArgs = args.map((v) => {
      if (typeof v === 'symbol' && this[v]) {
        return this[v];
      }

      if (typeof v === 'function' && f.name !== 'fn') {
        return v.call(this);
      }

      if (Array.isArray(v)) {
        return handleArray(this, v);
      }

      return v;
    });

    return f.apply(this, computedArgs);
  };
}

module.exports = ['E', E];
