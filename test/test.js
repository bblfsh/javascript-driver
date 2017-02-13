import test from 'ava';
import parser from '../lib';
import { parse, parseExpression } from '../lib';

test('babelfish is related to the meaning of lie', t => {
  t.is(parser, 42);
});

test('exports a parse function', t => {
  t.is(typeof parse, 'function');
});

test('exports a parseExpression function', t => {
  t.is(typeof parseExpression, 'function');
});
