import { AppBar, Box, Button, Container, IconButton, Typography } from "@mui/material";
import { Menu } from '@mui/icons-material';

import React from "react";
import { HashRouter as Router, Routes, Route } from "react-router-dom";

import { TopBar } from "./TopBar.component";
// import { Home } from "./Home";
// import { Sockpuppet } from "./Sockpuppet";

import { useOnLoadAuthenticationEffect, useUpdateAuthenticatedUserDataEffect } from "tagioalisi/services/auth";
import { usePromiseEffect } from 'tagioalisi/services/async';


export function Root() {
  // This MUST come first or the onload will not affect it
  
  // useUpdateAuthenticatedUserDataEffect();
  // useOnLoadAuthenticationEffect();

  // const [styles] = usePromiseEffect(() => import('./Root.module.scss').then(m => m.default));

  return (
    <Router>
      <Container maxWidth="md">
        <TopBar />

             <h1>hello world</h1>
           <p>the quick brown fox jumped over the lazy dog.</p>

      </Container>
    </Router>
    );

}
      // {/* <div className={styles?.page}>
      //   <div className={styles?.appContent}>
      //     <TopBar />

      //     <h1>hello world</h1>
      //     <p>the quick brown fox jumped over the lazy dog.</p>
      //   </div>
      // </div> */}

    // <Router>
    //   <div className={styles?.rootContainer}>
    //     <div className={styles?.row}>
    //       {/* <Sidebar /> */}
    //       <div className={styles?.rootContent}>
    //         {/* <Routes>
    //           <Route path="/sockpuppet" element={<Sockpuppet />} />
    //           <Route path="/" element={<Home />} />
    //         </Routes> */}

    //         The quick brown fox jumped over the lazy dog.
    //         The quick brown fox jumped over the lazy dog.
    //         The quick brown fox jumped over the lazy dog.
    //         The quick brown fox jumped over the lazy dog.
    //         The quick brown fox jumped over the lazy dog.
    //         The quick brown fox jumped over the lazy dog.
    //         The quick brown fox jumped over the lazy dog.
    //         The quick brown fox jumped over the lazy dog.
    //         The quick brown fox jumped over the lazy dog.
    //       </div>
    //     </div>
    //   </div>
    // </Router>
  // );
// }
