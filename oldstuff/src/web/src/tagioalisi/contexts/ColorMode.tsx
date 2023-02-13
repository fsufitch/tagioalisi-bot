import React, { ReactNode } from 'react';
import useMediaQuery from '@mui/material/useMediaQuery';
import { createTheme } from '@mui/material/styles';
import { ThemeProvider } from '@emotion/react';
import CssBaseline from '@mui/material/CssBaseline';
import { StorageContext } from './Storage';


type ColorMode = 'dark' | 'light';

interface ColorModeContextValue {
    colorMode?: ColorMode;
    setColorMode?: (m: ColorMode) => void;
}

export const ColorModeContext = React.createContext<ColorModeContextValue>({});


export default (props: {children: ReactNode}) => {
    const defaultMode = useMediaQuery('(prefers-color-scheme: dark)') ? 'dark': 'light';
    const { state } = React.useContext(StorageContext);
    const [storedColorMode, setStoredColorMode] = state?.useJSON<'light' | 'dark' | 'default'>('color-mode') || [];
    const actualColorMode = React.useMemo(() => (
        {light: 'light', dark: 'dark', default: defaultMode}[storedColorMode || 'default'] as 'light' | 'dark'
    ), [storedColorMode]);

    const theme = React.useMemo(() => createTheme({
        palette: {
            mode: actualColorMode,
        },
    }), [actualColorMode]);


    return (
        <ColorModeContext.Provider value={{colorMode: actualColorMode, setColorMode: setStoredColorMode}}>
            <ThemeProvider theme={theme}>
                <CssBaseline />
                {props.children}
            </ThemeProvider>
        </ColorModeContext.Provider>);
}