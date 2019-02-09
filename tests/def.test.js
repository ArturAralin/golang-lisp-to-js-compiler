const chaiAsPromised = require('chai-as-promised');
const {
  expect,
  use: chaiUse,
} = require('chai');
const { execute } = require('./common');

chaiUse(chaiAsPromised);

describe('(def k v)', () => {
  it('should define "x"', async () => {
    const res = await execute({}, `
      (def "x" 10)
    `);

    return expect(res)
      .to.have.property('x')
      .that.to.equals(10);
  });

  it('should throw error when property already declared', async () => {
    const res = execute({}, `
      (def "x" 10)
      (def "x" 10)
    `);

    return expect(res)
      .to.eventually.be.rejectedWith(Error, '"x" already declared in this scope');
  });

  it('should throw error when less then two arguments received', async () => {
    const res = execute({}, `
      (def "x")
    `);

    return expect(res)
      .to.eventually.be.rejectedWith(Error, 'def expect to get two arguments');
  });

  it('should inherits value from "x"', async () => {
    const res = await execute({}, `
      (def x 1)
      (def "y" x)
    `);

    return expect(res.y).to.equal(1);
  });

  it('should define x as symbol', async () => {
    const res = await execute({}, `
      (def x 1)
    `);

    return expect(res[res.$('x')]).to.equals(1);
  });

  it('should throw error when first arg not a string', async () => {
    const res = execute({}, `
      (def {} 1)
    `);

    return expect(res).to.be.rejectedWith(Error, 'def key must be a string or symbol');
  });
});
