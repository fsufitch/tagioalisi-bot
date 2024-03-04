const THEME_STORAGE = 'theme';

export enum AppTheme {
  system = 'browser-preference',
  light = 'light-theme',
  dark = 'dark-theme',
}

export enum VuetifyThemeKey {
  light = 'light',
  dark = 'dark',
}

export const loadThemeFromStorage = () => {
  let value: unknown = null;
  try {
    value = JSON.parse(sessionStorage.getItem(THEME_STORAGE) || 'null');
  } catch (err) {
    console.debug('Could not load theme from storage', err);
  }
  for (const [, v] of Object.entries(AppTheme)) {
    if (value === v) {
      console.debug('Valid theme found in storage:', value);
      return v;
    }
  }
  console.debug('Invalid theme found in storage, using default; found:', value);
  return AppTheme.system;
};

export const saveThemeToStorage = (value: AppTheme) => {
  console.debug('Save theme to storage:', value);
  sessionStorage.setItem(THEME_STORAGE, JSON.stringify(value));
};

export const getVuetifyThemeKey = (theme: AppTheme) =>
  ({
    [AppTheme.system]: getSystemPreferredThemeKey(),
    [AppTheme.light]: VuetifyThemeKey.light,
    [AppTheme.dark]: VuetifyThemeKey.dark,
  }[theme]);

export const getSystemPreferredThemeKey = () =>
  window.matchMedia('(prefers-color-scheme: dark)').matches
    ? VuetifyThemeKey.dark
    : VuetifyThemeKey.light;
