import React, { ReactNode } from 'react';
import useMediaQuery from '@mui/material/useMediaQuery';
import { createTheme } from '@mui/material/styles';
import { ThemeProvider } from '@emotion/react';
import CssBaseline from '@mui/material/CssBaseline';
import { useSynchronizedState } from '../services/state';


type ColorMode = 'dark' | 'light';

interface ColorModeContextValue {
    getColorMode: () => ColorMode,
    setColorMode: (m: ColorMode) => void,
}

export const ColorModeContext = React.createContext<ColorModeContextValue>({
    getColorMode: () => 'light',
    setColorMode: (m: ColorMode) => { },
});



export default (props: {children: ReactNode}) => {
    const prefersDarkMode = useMediaQuery('(prefers-color-scheme: dark)');
    const defaultMode = prefersDarkMode ? 'dark' : 'light';
    const [mode, setMode] = useSynchronizedState<'light' | 'dark'>(
        'color-mode', 
        defaultMode,
        s => s,
        s => s === 'dark' ? s : s === 'light' ? s : defaultMode)

    React.useEffect(() => {
        setMode('dark');
    }, [prefersDarkMode]);

    const colorModeContextProviderValue: ColorModeContextValue = {
        setColorMode: (m: ColorMode) => setMode(m),
        getColorMode: () => mode,
    }

    const theme = React.useMemo(() => createTheme({
        palette: {
            mode: mode,
        },
    }), [mode]);


    return (
        <ColorModeContext.Provider value={colorModeContextProviderValue}>
            <ThemeProvider theme={theme}>
                <CssBaseline />
                {props.children}
            </ThemeProvider>
        </ColorModeContext.Provider>);
}