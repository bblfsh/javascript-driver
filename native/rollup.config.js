import babel from 'rollup-plugin-babel';
import nodeResolve from 'rollup-plugin-node-resolve';
import commonJS from 'rollup-plugin-commonjs';
import multientry from 'rollup-plugin-multi-entry';

export default {
    input: 'src/*.js',

    output: {
	file: 'lib/index.js',
	format: 'cjs',
	banner: '#!/usr/bin/env node',
	exports: 'named',
    },

    plugins: [
	multientry(),
	babel(),
	nodeResolve({
	    mainFields: ['jsnext', 'main'],
	}),
	commonJS({
	}),
    ],

    treeshake: false,
    external: ['readline'],
};
