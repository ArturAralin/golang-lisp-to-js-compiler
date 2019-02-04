const def = (name, value) => {
  if (ctx[name]) {
    throw Error(`"${name}" already declared in this scope`);
  }

  ctx[name] = value
}

module.exports = def;
