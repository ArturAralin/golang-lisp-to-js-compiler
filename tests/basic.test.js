const { expect } = require('chai');
const { execute } = require('./common');


describe('Basic types', () => {
  it('should define a int number', async () => {
    const res = await execute({}, '(def "x" 10)');

    return expect(res)
      .to.have.property('x')
      .that.to.equals(10);
  });

  it('should define a float number', async () => {
    const res = await execute({}, '(def "x" 10.23)');

    return expect(res)
      .to.have.property('x')
      .that.to.equals(10.23);
  });

  it('should define a negative number', async () => {
    const res = await execute({}, '(def "x" -10.23)');

    return expect(res)
      .to.have.property('x')
      .that.to.equals(-10.23);
  });

  it('should define a exponential form number', async () => {
    const res = await execute({}, '(def "x" 10e12)');

    return expect(res)
      .to.have.property('x')
      .that.to.equals(10000000000000);
  });

  it('should define a string', async () => {
    const res = await execute({}, '(def "x" "some value")');

    return expect(res)
      .to.have.property('x')
      .that.to.equals('some value');
  });

  it.skip('should define a string with escape symbol', async () => {
    const res = await execute({}, '(def "x" "some \\" value\\")');

    return expect(res)
      .to.have.property('x')
      .that.to.equals('some "value');
  });

  it('should define a object', async () => {
    const res = await execute({}, `
      (def "x" {})
    `);

    return expect(res)
      .to.have.property('x')
      .that.to.equals('some value');
  });
});
