import type * as Vuetify from 'vuetify';

type VuetifyOptions = NonNullable<Parameters<typeof Vuetify.createVuetify>[0]>;

export const createVuetify = async (dryRun?: boolean) => {
  const options: VuetifyOptions = {};

  options.defaults = {
    VBtn: {
      variant: 'elevated',
    },
  };
  options.theme = {
    themes: {
      discordDark: {
        dark: true,
        colors: {
          // See: https://discord.com/branding
          'background': '#222222',
          'surface': '#444751',
          'primary': '#7289DA', // Blurple
          'secondary': '#EB459E', // Fuchsia
          'success': '#57F287', // Green
          'warning': '#FEE75C', // Yellow
          'error': '#ED4245', // Red
          'info': '#777777',

          'on-background': '#ffffff',
          'on-surface': '#ffffff',
          'on-primary': '#ffffff',
          'on-secondary': '#ffffff',
          'on-success': '#000000',
          'on-warning': '#000000',
          'on-error': '#ffffff',
          'on-info': '#ffffff',
        },
      },
    },
  };

  if (!dryRun) {
    options.components = await import('vuetify/components');
    options.directives = await import('vuetify/directives');

    const { aliases: VuetifyMdiAliases, mdi: VuetifyMdi } = await import('vuetify/iconsets/mdi');
    options.aliases = VuetifyMdiAliases;
    options.icons = {
      defaultSet: 'mdi',
      aliases: VuetifyMdiAliases,
      sets: {
        mdi: VuetifyMdi,
      },
    };
  }

  const Vuetify = await import('vuetify');
  return Vuetify.createVuetify(options);
};
