/// <reference types="vitest" />
import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import svgr from '@svgr/rollup'
import path from 'path'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    react(),
    svgr()
  ],
  test: {
    globals: true,
    environment: 'happy-dom',
    restoreMocks: true,
    include: ['**/*.{test,spec,type-test}.{js,mjs,cjs,ts,tsx,jsx}'],
    coverage: {
      exclude: ['tests/**', '.eslintrc.cjs'],
      provider: "v8",
      reporter: ['cobertura', 'text', 'html', 'clover', 'json'],
    },
    setupFiles: `${path.resolve(__dirname, 'tests/setup.ts')}`,
    snapshotFormat: {
      printBasicPrototype: true,
    },
    alias: [
      {
        find: new RegExp('^.+\\.(png|jpg|ttf|woff|woff2|svg|gif)$', 'g'),
        replacement: path.resolve(__dirname, '/mocks/jest/file-mock.ts'),
      },
    ],
  },
  server: {
    port: 9000,
    open: true,

    proxy: {
      '/api': {
        target: 'http://127.0.0.1:8989',
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/api/, '')
      }
    }
  },
  resolve: {
    alias: {
      '@': path.resolve(__dirname, './src'),
      '@assets': path.resolve(__dirname, './src/assets'),
      '@views': path.resolve(__dirname, './src/views'),
      '@components': path.resolve(__dirname, './src/components'),
      '@utils': path.resolve(__dirname, './src/utils'),
      '@constants': path.resolve(__dirname, './src/constants'),
      '@api': path.resolve(__dirname, './src/api'),
    },
    extensions: ['.js', '.jsx', '.ts', '.tsx']
  }

})
