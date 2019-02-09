function def(k, v) {
  const type = typeof k;

  if (type !== 'string' && type !== 'symbol') {
    throw new Error('def key must be a string or symbol');
  }

  if (v === undefined) {
    throw new Error('def expect to get two arguments');
  }

  if (this[k]) {
    throw new Error(`"${k}" already declared in this scope`);
  }

  this[k] = typeof v === 'function' ? v.call(this) : v;
}

module.exports = ['def', def];
