/**
 * @name def
 * @sig (fn key value)
 * @analogueOf const
 * @example
 *  (def a 20)
 *  (.console.log a)
 */

function def(k, v) {
  const type = typeof k;

  if (type !== 'string' && type !== 'symbol') {
    throw new Error('def key must be a string or symbol');
  }

  if (v === undefined) {
    throw new Error('def expect to get two arguments');
  }

  if (this.ROOT[k]) {
    throw new Error(`"${k}" already declared in this scope`);
  }

  this.ROOT[k] = v;
}

module.exports = ['def', def];
