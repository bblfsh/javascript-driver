export function error({ message }) {
  return responsify({ "status": "error", "errors": [message]});
}

export function fatal({ message }) {
  return responsify({ "status": "fatal", "errors": [message]});
}

export function ok(ast) {
  return responsify({ast, "status": "ok"});
}

function responsify(resp) {
  return resp;
}
