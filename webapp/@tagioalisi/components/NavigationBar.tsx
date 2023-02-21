import React from 'react';
import { AppBar, Toolbar, Typography, Stack, Button } from '@mui/material';

export const NavigationBar = () => (
  <AppBar position="static">
    <Stack component={Toolbar} direction="row" justifyContent="space-between">
      <Typography variant="h6">Tagioalisi &mdash;</Typography>
      <Button>Login</Button>
    </Stack>
  </AppBar>
);
