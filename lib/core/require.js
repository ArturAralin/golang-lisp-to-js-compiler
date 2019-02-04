const symbols = require('../symbols');

// @example
// (require [
//   "module",
//   "module.sub",
//   ["module" :as "other-name"],
//   ["module.sub" :as "sub"],
// ])
const $as = symbols[':as'];

const requireModule = (s) => {
  const [moduleId, ...path] = s.split('.');
  let v = require(moduleId)

  while (path.length > 0) {
    v = v[path.pop()];

    if (!v) {
      throw new Error(`Invalid require path for "${moduleId}"`);
    }
  }

  return v;
};

const normalize = item => {
  if (typeof item === 'string') {
    return [item, { [$as]: item }]
  }

  if (Array.isArray(item) && typeof item[1] === 'symbol') {
    const [path, s, v] = item;

    return [path, { [s]: v }]
  }

  return item;
}

const checkTypes = item => {
  const isString = typeof item === 'string';
  const isArray = Array.isArray(item);

  if (!isString && !isArray) {
    throw new Error(`Invalid require declaration "${item}"`);
  }
}

const __require = ctx => (args) => {
  const l = args.length;
  let i = 0;

  while (i < l) {
    checkTypes(args[i]);

    const [
      path,
      opts,
    ] = normalize(args[i]);
    
    ctx[opts[$as]] = requireModule(path);

    i += 1;
  }
}

module.exports = __require;
