const chaiAsPromised = require('chai-as-promised');
const {
  expect,
  use: chaiUse,
} = require('chai');
const { execute } = require('./common');

chaiUse(chaiAsPromised);

describe('(require [args])', () => {
  it('should require "fs"', async () => {
    const res = await execute({}, `
      (require ["fs"])
    `);

    // eslint-disable-next-line global-require
    return expect(res.fs === require('fs'))
      .to.be.true;
  });

  it('should require global Math', async () => {
    const res = await execute({}, `
      (require [".Math"])
    `);

    return expect(res.Math === Math)
      .to.be.true;
  });

  // Maybe it make sense import one module one string argument
  it.skip('[controversial case] (require "fs")', async () => {
    const res = await execute({}, `
      (require "fs")
    `);

    // eslint-disable-next-line global-require
    return expect(res.fs === require('fs'))
      .to.be.true;
  });

  it('should throw error "require args must be an array"', async () => {
    const res = execute({}, `
      (require 12)
    `);

    return expect(res)
      .to.eventually.be.rejectedWith(Error, 'require args must be an array');
  });

  // Need to make choice about import "fs.readFile"
  // I see two main cases
  // 1) Import this as "readFile"
  // 2) Import this as "fs.readFile"
  it.skip('[controversial case] (require ["fs.readFile"])', async () => {
    await execute({}, `
      (require ["fs.readFile"])
    `);
  });

  it.skip('should require global Math as MATH', async () => {
    const res = await execute({}, `
      (require [".Math" :as "MATH"])
    `);

    return expect(res.MATH === Math)
      .to.be.true;
  });

  it.skip('should require global Math as MaTh', async () => {
    const res = await execute({}, `
      (require [
        "Math" { :is-global true, :as MaTh }
      ])
    `);

    return expect(res.MaTh === Math)
      .to.be.true;
  });
});
