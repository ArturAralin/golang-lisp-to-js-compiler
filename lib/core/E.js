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

      return v;
    });

    return f.apply(this, computedArgs);
  };
}

module.exports = ['E', E];
