import React, { FC } from 'react';
import { Outlet } from 'react-router-dom';
import { Container, Paper, Stack, Card, CardContent } from '@mui/material';
import { NavigationBar } from './NavigationBar';

export {};
export const foo = {};

export const ApplicationLayout: FC = () => {
  return (
    <Container maxWidth="md">
      <NavigationBar />
      <Stack direction="column">
        <Outlet />
      </Stack>
    </Container>
  );
};
