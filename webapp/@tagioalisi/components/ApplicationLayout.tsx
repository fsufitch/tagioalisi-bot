import React, { FC } from 'react';
import { Outlet } from 'react-router-dom';

export {};
export const foo = {};

export const ApplicationLayout: FC = () => {
  return (
    <>
      before
      <Outlet />
      after
    </>
  );
};
