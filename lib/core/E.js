/* eslint-disable no-use-before-define */
const [, type] = require('./type');
const [, identity] = require('./identity');

const handleArray = (ctx, v) => E
  .apply(ctx, [
    (...args) => args,
    ...v,
  ])
  .call(ctx);

const handleObject = (ctx, obj) => Object
  .keys(obj)
  .reduce((acc, key) => ({
    ...acc,
    [key]: E(
      identity,
      obj[key],
    ).call(ctx),
  }), {});

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

      if (type(v) === 'Object') {
        return handleObject(this, v);
      }

      return v;
    });

    return f.apply(this, computedArgs);
  };
}

module.exports = ['E', E];
