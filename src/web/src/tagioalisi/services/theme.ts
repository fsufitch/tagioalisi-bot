import * as React from 'react';
import useMediaQuery from '@mui/material/useMediaQuery';
import { createTheme, ThemeProvider } from '@mui/material/styles';
import { useSynchronizedState } from 'tagioalisi/services/state';

const MODE_KEY = 'tagioalisi/mode';

type DisplayMode = 'dark' | 'light';

export const useSelectedMode = () => {
  // https://mui.com/material-ui/customization/dark-mode/#system-preference
  const prefersDarkMode = useMediaQuery('(prefers-color-scheme: dark)');
  const [selectedMode, setSelectedMode] = useSynchronizedState<DisplayMode>(
    MODE_KEY, prefersDarkMode ? 'dark' : 'light',
    s => s, s => s == 'dark' ? 'dark' : 'light');

  return [selectedMode, setSelectedMode] as [DisplayMode, (m: DisplayMode) => void];
}

export const useTheme = () => {
  const x = useSelectedMode();
  const [ selectedMode ] = useSelectedMode();
  const theme = React.useMemo(() => createTheme({
    palette: {
      mode: selectedMode,
    },
  }), [ selectedMode ]);
  return theme;
}