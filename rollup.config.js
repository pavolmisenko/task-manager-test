//import typescript from "@rollup/plugin-typescript";
import { nodeResolve } from "@rollup/plugin-node-resolve";
//import postcss from "rollup-plugin-postcss";

export default {
  input: "dist/file_search.js",
  output: {
    file: "lib/bundle.js",
    format: "iife",
  },
  plugins: [nodeResolve()],
  // plugins: [
  //   typescript(),
  //   postcss({
  //     extract: true,
  //     minimize: true,
  //   }),
  //   nodeResolve(),
  // ],
};
