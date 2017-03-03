import { guessParsing } from "./parser";
import { error, ok } from "./response";

function parse(data) {
  try {
    let { content } = JSON.parse(data)
    let ast = guessParsing(content);
    return  ok(ast);
  } catch (ex) {
    return error(ex);
  }
}

function jsonify(data) {
  return `${JSON.stringify(data)}
`;
}

export function handler(input) {
  return console.log(jsonify(parse(input)));
}
