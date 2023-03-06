// Necessary so vite.config.ts doesn't cry about missing types
declare module 'vite-plugin-node-stdlib-browser' {
  import { PluginOption } from 'vite';
  const callableModule: () => PluginOption;
  export default callableModule;
}
