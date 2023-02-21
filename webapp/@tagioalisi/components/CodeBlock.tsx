import * as React from 'react';
import { Paper, useTheme } from '@mui/material';

export const CodeBlock: React.FC<React.PropsWithChildren> = ({ children }) => {
  const theme = useTheme();
  theme.palette;

  return (
    <Paper
      css={{
        padding: theme.spacing(3),
        backgroundColor: theme.palette.grey[200],
        overflow: 'auto',
      }}
      component="pre"
      variant="outlined"
      color="disabled"
    >
      {children}
    </Paper>
  );
};
