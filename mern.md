# mac@home

Node.js version = v14.16.1 (v12.x is maintenace LTS, v14.x is active LTS, v16.x is current)

# setup webpack, typescript, react

* https://www.skcript.com/svr/using-webpack-with-react-typescript/
* https://magiclen.org/webpack-typescript/

`npx ts-node-dev src/index.ts`

```
// no
// import express from 'express';
//
// Problem occurs when we want to import CommonJS module into ES6 module codebase.

// ok
// import * as express from 'express';

// ok
const express = require('express');
 
const app = express();
 
app.listen(3000, () =>
  console.log('Example app listening on port 3000!'),
);
```
