import type { ThemeDefinition } from 'vuetify/lib/framework.mjs';

declare global {
  type VuetifyThemes = Record<string, ThemeDefinition>;
}
