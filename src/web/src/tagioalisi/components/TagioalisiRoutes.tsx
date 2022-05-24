import React, { Suspense } from 'react';
import { Routes, Route } from 'react-router-dom';
import { ROUTES } from 'tagioalisi/routes';


export default () =>
  <Suspense fallback={<Loading />}>
    <Routes>
      {
        Object.keys(ROUTES)
          .map((id) => ({ id, route: ROUTES[id] }))
          .map(({ id, route }) =>
            <Route key={id} path={route.path} element={<><route.component /></>} />
          )
      }
    </Routes>
  </Suspense>


const Loading = () => <>Loading...</>