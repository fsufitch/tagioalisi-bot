import React from "react";

import { Container } from "@mui/material";
import { BrowserRouter as Router } from "react-router-dom";

import ApplicationBar from "tagioalisi/components/ApplicationBar";
import TagioalisiRoutes from 'tagioalisi/components/TagioalisiRoutes';


export default () =>
  <Router>
    <Container maxWidth="md">
      <ApplicationBar />
      <TagioalisiRoutes />
    </Container>
  </Router>
