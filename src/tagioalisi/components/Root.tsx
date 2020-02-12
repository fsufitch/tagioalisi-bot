import React, { useState } from "react";
import { HashRouter as Router, Switch, Route } from "react-router-dom";

import styles from "tagioalisi/styles";
import { Sidebar } from "./Sidebar";
import { Home } from "./Home";
import { Sockpuppet } from "./Sockpuppet";

export function Root() {
  const [endpoint, setEndpoint] = useState(process.env.BOT_BASE_URL);
  const [authToken, setAuthToken] = useState("");
  return (
    <Router>
      <div className={styles.rootContainer}>
        <div className={styles.row}>
          <Sidebar />
          <div className={styles.rootContent}>
            <Switch>
              <Route path="/sockpuppet">
                <Sockpuppet
                  endpoint={endpoint}
                  onEndpointChanged={setEndpoint}
                  authToken={authToken}
                  onAuthTokenChanged={setAuthToken}
                />
              </Route>
              <Route path="/">
                <Home endpoint={endpoint} onEndpointChanged={setEndpoint} />
              </Route>
            </Switch>
          </div>
        </div>
      </div>
    </Router>
  );
}
