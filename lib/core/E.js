// this is executor function
// oh. this function full of pain and tears
function E(f, ...args) {
  return function e() {
    const computedArgs = args.map((v) => {
      if (typeof v === 'symbol' && this[v]) {
        return this[v];
      }

      return v;
    });

    return f.apply(this, computedArgs);
  };
}

module.exports = E;
