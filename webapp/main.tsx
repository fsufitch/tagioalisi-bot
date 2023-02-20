/* SPDX-FileCopyrightText: 2014-present Kriasoft */
/* SPDX-License-Identifier: MIT */

import * as React from "react";
import * as ReactDOM from "react-dom/client";
import { BrowserRouter } from "react-router-dom";
import Application from "@tagioalisi/components/Application";

const container = document.getElementById("app-container") as HTMLElement;

// Render the top-level React component
ReactDOM.createRoot(container).render(
  <React.StrictMode>
      <BrowserRouter>
        <Application />
      </BrowserRouter>
  </React.StrictMode>,
);