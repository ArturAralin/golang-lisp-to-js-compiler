const keyword = require('../keywords');

const $as = keyword('as');
const $isGlobal = keyword('is-global');

const requireModule = (modulePath, isGlobal) => {
  const [moduleId, ...path] = modulePath.split('.');

  if (isGlobal) {
    return global[moduleId];
  }

  // eslint-disable-next-line global-require, import/no-dynamic-require
  let v = require(moduleId);

  while (path.length > 0) {
    v = v[path.pop()];

    if (!v) {
      throw new Error(`Invalid require path for "${moduleId}"`);
    }
  }

  return v;
};

const normalize = (item) => {
  if (typeof item === 'string') {
    const isGlobal = item[0] === '.';
    const moduleId = isGlobal ? item.slice(1) : item;

    return [moduleId, {
      [$as]: moduleId,
      [$isGlobal]: isGlobal,
    }];
  }

  // (require ["x" :as "y"]) case
  if (Array.isArray(item) && item[1] === $as) {
    const [path, s, v] = item;

    return [path, { [s]: v }];
  }

  return item;
};

const checkTypes = (item) => {
  const isString = typeof item === 'string';
  const isArray = Array.isArray(item);

  if (!isString && !isArray) {
    throw new Error(`Invalid require declaration "${item}"`);
  }
};

/**
 * @name require
 * @sig (require [args])
 * @analogueOf import, require
 * @example
 *  (require ["fs"])
 */
function customRequire(args) {
  if (!Array.isArray(args)) {
    throw new Error('require args must be an array');
  }

  const l = args.length;
  let i = 0;

  while (i < l) {
    checkTypes(args[i]);

    const [
      path,
      opts,
    ] = normalize(args[i]);

    this.ROOT[opts[$as]] = requireModule(path, opts[$isGlobal]);

    i += 1;
  }
}

module.exports = ['require', customRequire];
