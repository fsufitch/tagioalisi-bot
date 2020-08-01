import React from "react";

import { NavLink } from "react-router-dom";

import styles from "tagioalisi/styles";
import { Auth } from "./Auth";

export function Sidebar() {
  return (
    <div className={styles.sidebar}>
      <h3> Tagioalisi Web API </h3>
      <Auth />
      <nav>
        <ul>
        <li>
            <NavLink exact to="/" activeClassName={styles.active}>Home</NavLink>
          </li>
          <li>
            <NavLink exact to="/sockpuppet" activeClassName={styles.active}>Sockpuppet</NavLink>
          </li>
        </ul>
      </nav>
    </div>
  );
}
