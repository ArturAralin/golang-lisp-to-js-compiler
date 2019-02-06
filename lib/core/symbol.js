const symbols = {
  ':as': Symbol(':as'),
};

function symbol(name) {
  if (!symbols[name]) {
    symbols[name] = Symbol(name);
  }

  return symbols[name];
}

module.exports = symbol;
