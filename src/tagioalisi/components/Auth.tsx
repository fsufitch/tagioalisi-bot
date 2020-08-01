import React, { Fragment } from "react";

import styles from "tagioalisi/styles";
import { useAuthenticatedUserData, useAuthentication } from "tagioalisi/services/auth";

interface WhoAmIData {
  id?: string;
  fullname?: string;
  avatarUrl?: string;
}


export function Auth() {
  const {login, logout} = useAuthentication();
  const [userData] = useAuthenticatedUserData();

  return (<div className={styles.authSection}>
    {userData.authenticated ?
      <Fragment>
        <p>
          {userData.fullname}
        </p>
        <img src={userData.avatarUrl} />
        <p>
          <button onClick={logout}> Logout </button>
        </p>
      </Fragment>
      : userData.authPending ?
        <Fragment>Loading...</Fragment>
        :
        <Fragment>
          <button onClick={login}>Login with Discord</button>
        </Fragment>}
  </div>)
}
