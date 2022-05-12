import React from "react";

import { NavLink } from "react-router-dom";

// import styles from "tagioalisi/styles";
import { Auth } from "./Auth";
import { usePromiseEffect } from 'tagioalisi/services/async';

export function Sidebar() {

  const [styles] = usePromiseEffect(() => import('./Sidebar.module.scss').then(m => m.default));

  return (
    <div className={styles?.sidebar}>
      <h3> Tagioalisi Web API </h3>
      <Auth />
      <nav>
        <ul>
        <li>
            <NavLink end to="/" className={({isActive}) => isActive ? styles?.active : ''}>Home</NavLink>
          </li>
          <li>
            <NavLink end to="/sockpuppet" className={({isActive}) => isActive ? styles?.active : ''}>Sockpuppet</NavLink>
          </li>
        </ul>
      </nav>
    </div>
  );
}
