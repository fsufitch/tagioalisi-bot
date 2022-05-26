import React from "react";

import { Container } from "@mui/material";
import { BrowserRouter as Router } from "react-router-dom";
import { ThemeProvider } from '@mui/material/styles';
import CssBaseline from '@mui/material/CssBaseline';

import ApplicationBar from "tagioalisi/components/ApplicationBar";
import TagioalisiRoutes from 'tagioalisi/components/TagioalisiRoutes';
import { useTheme } from 'tagioalisi/services/theme';


export default () => {
  const theme = useTheme();
  return (
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <Router>
        <Container maxWidth="md">
          <ApplicationBar />
          <TagioalisiRoutes />
        </Container>
      </Router>
    </ThemeProvider>
  );
}
