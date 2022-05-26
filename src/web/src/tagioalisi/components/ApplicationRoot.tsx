import React from "react";

import { Container } from "@mui/material";
import { BrowserRouter as Router } from "react-router-dom";

import ApplicationBar from "tagioalisi/components/ApplicationBar";
import TagioalisiRoutes from 'tagioalisi/components/TagioalisiRoutes';
import TagioalisiTheme from 'tagioalisi/Theme';


export default () => {

  return (
    <TagioalisiTheme>
      <Router>
        <Container maxWidth="md">
          <ApplicationBar />
          <TagioalisiRoutes />
        </Container>
      </Router>
    </TagioalisiTheme>
  );
}
