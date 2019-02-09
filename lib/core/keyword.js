const keywords = {
  ':as': Symbol(':as'),
};

function keyword(name) {
  const nameWithPrefix = `:${name}`;

  if (!keywords[nameWithPrefix]) {
    keywords[nameWithPrefix] = Symbol(nameWithPrefix);
  }

  return keywords[nameWithPrefix];
}

module.exports = ['keyword', keyword];
