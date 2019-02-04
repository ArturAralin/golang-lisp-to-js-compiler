const rt = require('require-tree');
const core = rt('./core');


// const attachToCTX = ctx => (o) => {
//   Object
//     .keys(o)
//     .forEach(k => {
//       const f = o[k];

//       if (ctx[k]) {
//         throw Error('err attachToCTX');
//       }

//       ctx[k] = f(ctx);
//     })
// }
 
// const core = {
//   def,
//   require: __require,
// };
// const CTX = {};
// const addToCTX = attachToCTX(CTX);

// addToCTX(core);

// CTX.require('fs')

// console.log(CTX)


