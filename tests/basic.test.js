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
      (def "x" {
        "key" "value"
        "key2" "value2"
      })
    `);

    return expect(res)
      .to.have.property('x')
      .that.to.eqls({
        key: 'value',
        key2: 'value2',
      });
  });

  it('should define a object', async () => {
    const res = await execute({}, `
      (def "x" ["a" 123 123.2 2e2 {}])
    `);

    return expect(res)
      .to.have.property('x')
      .that.to.eqls(['a', 123, 123.2, 200, {}]);
  });

  it('should define a nil', async () => {
    const res = await execute({}, '(def "x" nil)');

    return expect(res)
      .to.have.property('x')
      .that.to.equals(null);
  });

  it('should define a true', async () => {
    const res = await execute({}, '(def "x" true)');

    return expect(res)
      .to.have.property('x')
      .that.to.equals(true);
  });

  it('should define a true', async () => {
    const res = await execute({}, '(def "x" NaN)');

    // eslint-disable-next-line no-restricted-globals
    return expect(isNaN(res.x)).to.be.true;
  });

  it('should define a Infinity', async () => {
    const res = await execute({}, '(def "x" Infinity)');

    return expect(res)
      .to.have.property('x')
      .that.to.equals(Infinity);
  });

  it('should define a -Infinity', async () => {
    const res = await execute({}, '(def "x" -Infinity)');

    return expect(res)
      .to.have.property('x')
      .that.to.equals(-Infinity);
  });

  it('should define a false', async () => {
    const res = await execute({}, '(def "x" false)');

    return expect(res)
      .to.have.property('x')
      .that.to.equals(false);
  });
});
