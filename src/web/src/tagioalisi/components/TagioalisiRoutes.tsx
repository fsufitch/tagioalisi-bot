import React from 'react';
import { Routes, Route } from 'react-router-dom';
import { ROUTES } from '../routes';


export const TagioalisiRoutes = () => <Routes>
{
  Object.keys(ROUTES)
  .map((id) => ({id, route: ROUTES[id]}))
  .map(({id, route}) => 
    <Route key={id} path={route.path} element={<route.component />}/>
  )
}
</Routes>