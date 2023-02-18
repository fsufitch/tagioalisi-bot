// Common styles/fonts, only import once
import '@tagioalisi/styles/common.scss';

import React from "react";

import { Container } from "@mui/material";
import { BrowserRouter as Router } from "react-router-dom";

import ApplicationBar from "@tagioalisi/components/ApplicationBar";
import TagioalisiRoutes from '@tagioalisi/components/TagioalisiRoutes';
import TagioalisiContextProvider from "@tagioalisi/contexts";


export default () => {

  return (
    <TagioalisiContextProvider>
      <Router>
        <Container maxWidth="md">
          <ApplicationBar />
          <TagioalisiRoutes />
        </Container>
      </Router>
    </TagioalisiContextProvider>
  );
}
