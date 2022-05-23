import React from "react";

import { Container } from "@mui/material";
import { BrowserRouter as Router } from "react-router-dom";

import { TopBar } from "tagioalisi/components/ApplicationBar";
import { TagioalisiRoutes } from 'tagioalisi/components/TagioalisiRoutes';


export function ApplicationRoot() {
  return (
    <Router>
      <Container maxWidth="md">
        <TopBar />
        <TagioalisiRoutes />
      </Container>
    </Router>
  );

}