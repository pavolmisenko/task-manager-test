import resolve from "@rollup/plugin-node-resolve";
import commonjs from "@rollup/plugin-commonjs";
import typescript from "rollup-plugin-typescript2";

export default {
  input: "src/lit-components/file_search.ts",
  output: {
    file: "src/web/static/js/file_search.js",
    format: "iife", // Suitable for <script> tags
    name: "MyComponentBundle", // Global variable name for your bundle
    sourcemap: true,
  },
  plugins: [resolve(), commonjs(), typescript({ tsconfig: "./tsconfig.json" })],
};
