import fs from 'fs';
import { guessParsing as parse } from '../lib';
import test from 'ava';

function isDir(path) {
  return fs.statSync(path).isDirectory();
}

function* getAllFiles(path) {
  let files = fs.readdirSync(path);
  for(let filename of files) {
    let completePath = `${path}/${filename}`;
    if (isDir(completePath)) {
      yield *getAllFiles(completePath);
    } else {
      yield completePath;
    }
  }
}

// TODO: Improve once this is working
function snapshotDir(path, cb) {
  for(let filepath of getAllFiles(path)) {
    if (!filepath.endsWith('.gitkeep')) {
      fs.readFile(filepath, 'utf8', function(err, content) {
        test(`could read ${filepath}`, t => t.ifError(err));

        if (!err) {
          cb(filepath, content);
        }
      });
    }
  }
}

test(`snapshot fixtures`, t => {
  snapshotDir(`${__dirname}/fixtures`, function(filepath, content) {
    let ast = parse(content);
    t.snapshot(ast, `snapshot for ${filepath}`)
    t.truthy(ast, `${filepath} produced an ast`);
  });
});
