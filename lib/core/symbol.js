const symbols = {};

/**
 * @name symbol
 * @sig (symbol str)
 * @analogueOf Symbol
 * @example
 *  (symbol "my-symbol")
 */
function symbol(name) {
  if (!symbols[name]) {
    symbols[name] = Symbol(name);
  }

  return symbols[name];
}

module.exports = ['symbol', symbol];
