import React from "react";
import { HashRouter as Router, Switch, Route } from "react-router-dom";

import styles from "tagioalisi/styles";
import { Sidebar } from "./Sidebar";
import { Home } from "./Home";
import { Sockpuppet } from "./Sockpuppet";
import { Auth } from "./Auth";
import { useOnLoadAuthenticationEffect, useUpdateAuthenticatedUserDataEffect } from "tagioalisi/services/auth";

export function Root() {
  // This MUST come first or the onload will not affect it
  useUpdateAuthenticatedUserDataEffect();
  useOnLoadAuthenticationEffect();
  
  return (
    <Router>
      <div className={styles.rootContainer}>
        <div className={styles.row}>
          <Sidebar />
          <div className={styles.rootContent}>
            <Switch>
              <Route path="/sockpuppet">
                <Sockpuppet />
              </Route>
              <Route path="/">
                <Home />
              </Route>
            </Switch>
          </div>
        </div>
      </div>
    </Router>
  );
}
