import * as React from 'react';
import { createRoot } from 'react-dom/client';

import routes from '@tagioalisi/routes';
import { RouterProvider, createBrowserRouter } from 'react-router-dom';
import { createTheme, CssBaseline, ThemeProvider } from '@mui/material';

(() => {
  const theme = createTheme({
    components: {
      MuiButton: { defaultProps: { variant: 'contained' } },
    },
  });

  return createRoot(document.getElementById('app-container') as HTMLElement).render(
    <React.StrictMode>
      <CssBaseline enableColorScheme />
      <ThemeProvider theme={theme}>
        <RouterProvider router={createBrowserRouter(routes)} />
      </ThemeProvider>
    </React.StrictMode>,
  );
})();
