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
      },
      onwarn: (entry, next) => {
        if(entry.loc?.file && /htmx\.min\.js$/.test(entry.loc.file) && /Use of eval in/.test(entry.message))
          return;
        return next(entry);
      }
    },
    minify: 'esbuild',
  }
});

