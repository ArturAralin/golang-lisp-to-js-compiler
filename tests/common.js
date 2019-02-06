const path = require('path');
const fs = require('fs');
const crypto = require('crypto');
const { exec } = require('child_process');
const proxyquire = require('proxyquire');
const core = require('../lib/core');

const TMP_FOLDER = path.resolve(__dirname, './tmp');

const compile = code => new Promise((resolve, reject) => {
  const fjsCompilerPath = path.resolve(__dirname, '../bin/test-fjs-compiler');
  const cmd = `${fjsCompilerPath} "${code}"`;
  const filename = `${TMP_FOLDER}/${crypto.randomBytes(10).toString('hex')}.js`;

  const compiler = exec(cmd);

  compiler.stdout.on('data', (data) => {
    fs.writeFileSync(filename, data);
  });

  compiler.stderr.on('data', (data) => {
    console.log(`stderr: ${data}`);

    reject(data);
  });

  compiler.on('close', () => {
    resolve(filename);
  });
});

const execute = async (ctx, code) => {
  const absoluteFileName = await compile(code.replace(/"/g, '\\"'));

  return proxyquire(absoluteFileName, {
    'fjs-compiler/lib/core/core.js': {
      ...core,
      '@noCallThru': true,
    },
    ...ctx,
  }).CTX;
};

module.exports = {
  compile,
  execute,
};
