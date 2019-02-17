function add(...items) {
  if (items.length < 2) {
    throw Error('add expect greater or equal to 2 args');
  }

  return items.reduce((acc, item) => acc + item, 0);
}

module.exports = ['add', add];
