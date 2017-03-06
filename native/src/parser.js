const babylon = require('babylon');

export const ALL_PLUGINS = [
  'jsx',
  'flow',
  'doExpressions',
  'objectRestSpread',
  //'decorators' // leaving them out since they are based on an old standard
  'classProperties',
  'exportExtensions',
  'asyncGenerators',
  'functionBind',
  'functionSend',
  'dynamicImport',
];

export function parse(code, options) {
  return babylon.parse(code, Object.assign({ plugins: ALL_PLUGINS }, options));
}

export function parseExpression(code, option) {
  return babylon.parseExpression(code, Object.assign({ plugins: ALL_PLUGINS }, options));
}

const GUESSING_ORDER = [
  [parse, { sourceType: 'module', allowImportExportEverywhere: true }],
  [parse, { sourceType: 'script', allowReturnOutsideFunction: true }],
];

export function guessParsing(code) {
  let exceptions = [];
  for (let [fn, opts] of GUESSING_ORDER) {
    try {
      return fn(code, opts);
    } catch (ex) {
      exceptions.push(ex)
    }
  }
  throw new Error(exceptions.map((x) => x.message).join(","));
}

