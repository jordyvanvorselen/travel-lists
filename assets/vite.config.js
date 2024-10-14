import { defineConfig } from 'vite';
import path from 'path';

export default defineConfig({
  build: {
    outDir: 'dist',
    rollupOptions: {
      input: path.resolve(__dirname, 'index.js'),
      output: {
        manualChunks: undefined,
        entryFileNames: '[name].min.js',
      }
    },
    minify: 'esbuild',
  }
});

