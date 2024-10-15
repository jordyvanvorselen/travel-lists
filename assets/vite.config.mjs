import { defineConfig } from 'vite';
import path from 'path';

export default defineConfig({
  build: {
    outDir: 'dist',
    rollupOptions: {
      input: {
        js: path.resolve(__dirname, 'index.js'),
        css: path.resolve(__dirname, 'index.css'),
      },
      output: {
        manualChunks: undefined,
        entryFileNames: 'index.min.js',
        assetFileNames: 'index.min.[ext]',
      }
    },
    minify: 'esbuild',
  }
});

