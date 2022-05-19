import { Container } from "@mui/material";

import React from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";

import { TopBar } from "tagioalisi/components/ApplicationBar";
import { ROUTES, asyncLoadRoute } from '../../routes';
import { usePromiseEffect } from '../../services/async';


export function ApplicationRoot() {
  // This MUST come first or the onload will not affect it

  // useUpdateAuthenticatedUserDataEffect();
  // useOnLoadAuthenticationEffect();

  const [loadedRoutes] = usePromiseEffect(() => Promise.all(
    ROUTES.map(route => asyncLoadRoute(route))
  ))

  return (
    <Router>
      <Container maxWidth="md">
        <TopBar />

        <Routes>
          {
            loadedRoutes?.map((route, idx) =>
              <Route key={idx} path={route.path} element={<route.component />}/>
            )
          }
        </Routes>
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
