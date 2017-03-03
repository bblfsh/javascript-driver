import { guessParsing } from "./parser";
import { error, ok } from "./response";
import split from 'split';
import map from 'map-stream';

function parse(data, cb) {
  try {
    let { content } = JSON.parse(data)
    let ast = guessParsing(content);
    cb(null, ok(ast));
  } catch (ex) {
    cb(null, error(ex));
  }
}

function jsonify(data, cb) {
  cb(null, `${JSON.stringify(data)}
`);
}

export function handler(readStream, writeStream) {
  readStream
    .pipe(split())
    .pipe(map(parse))
    .pipe(map(jsonify))
    .pipe(writeStream);
}
