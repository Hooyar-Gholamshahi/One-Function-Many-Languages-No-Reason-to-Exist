const koffi = require('koffi');
const path = require('path');

const dllPath = path.resolve(
    __dirname,
    '..',
    '..',
    'target',
    'debug',
    'rust_native_bridge.dll'
);

const library = koffi.load(dllPath);
const hello_world = library.func('void hello_world()');

hello_world();
