import * as React from 'react';
import { createRoot } from 'react-dom/client';
import ScopedCssBaseline from '@mui/material/ScopedCssBaseline';

import routes from '@tagioalisi/routes';
import { RouterProvider, createBrowserRouter } from 'react-router-dom';
import { createTheme, ThemeProvider } from '@mui/material';

(() => {
  const theme = createTheme({
    components: {
      MuiButton: { defaultProps: { variant: 'contained' } },
    },
  });

  return createRoot(document.getElementById('app-container') as HTMLElement).render(
    <React.StrictMode>
      <ScopedCssBaseline enableColorScheme>
        <ThemeProvider theme={theme}>
          <RouterProvider router={createBrowserRouter(routes)} />
        </ThemeProvider>
      </ScopedCssBaseline>
    </React.StrictMode>,
  );
})();
