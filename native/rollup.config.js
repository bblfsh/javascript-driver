import babel from "rollup-plugin-babel";
import nodeResolve from "rollup-plugin-node-resolve";
import commonJS from 'rollup-plugin-commonjs';

export default {
  entry: "src/parser.js",
  dest: "lib/index.js",
  exports: 'named',
  plugins: [
    babel(),
    nodeResolve({
      jsnext: true,
      main: true,
    }),
    commonJS({
    }),
  ],
  format: "cjs",
  treeshake: false,
};
