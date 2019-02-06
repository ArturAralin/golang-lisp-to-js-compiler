function def(k, v) {
  if (k === undefined || v === undefined) {
    throw new Error('def expect to get two arguments');
  }
  if (this[k]) {
    throw new Error(`"${k}" already declared in this scope`);
  }

  this[k] = v;
}

module.exports = def;
