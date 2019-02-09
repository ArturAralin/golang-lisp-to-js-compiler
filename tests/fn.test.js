const chaiAsPromised = require('chai-as-promised');
const {
  expect,
  use: chaiUse,
} = require('chai');
const { execute } = require('./common');

chaiUse(chaiAsPromised);

describe('(fn [args] (body))', () => {
  it('should create function "pow"', async () => {
    const res = await execute({}, `
      (def "pow" (fn [a b] (.Math.pow a b)))
    `);

    return expect(res.pow).to.be.a('function');
  });

  it('should create function "pow" and execute her', async () => {
    const res = await execute({}, `
      (def pow (fn [a b] (.Math.pow a b)))
      (def "result" (pow 2 3))
    `);

    return expect(res.result)
      .to.be.a('number')
      .that.to.equals(8);
  });

  it('should pass agrs deep into', async () => {
    const res = await execute({}, `
      (def curry-pow (fn [a] (fn [b] (.Math.pow b a))))
      (def square (curry-pow 2))
      (def "result" (square 3))
    `);

    return expect(res.result)
      .to.be.a('number')
      .that.to.equals(9);
  });
});
