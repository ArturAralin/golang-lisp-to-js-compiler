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
  }

  return v;
};

const __require = ctx => (args) => {
  const l = args.length;
  let i = 0;

  while (i < l) {
    const item = args[i];
    const isString = typeof item === 'string';
    const isArray = Array.isArray(item);

    if (!isString && !isArray) {
      throw new Error(`Invalid require declaration "${item}"`);
    }

    const v = isString
      ? requireModule(item)
      : requireModule(item[0]);

    if (!v) {
      throw new Error(`Invalid require path "${isString ? item : item[0]}"`)
    }

    
    console.log(v);

    

    i += 1;
  }
  // if (ctx[moduleName]) {
  //   throw new Error(`"${moduleName}" already declared in this scope`);
  // }

  // ctx[moduleName] = require(moduleName);
}

const C = {};
__require(C)([
  // "fs",
  "fs.readdir",
  ["fs.readdir", $as, "other-name"],
  ["fs.readdir", $as, "sub"],
])

module.exports = __require;
