import React, { useState, useEffect, Fragment } from "react";
import url from "url";

import styles from "tagioalisi/styles";
import { Inputs } from "./Inputs";

interface WhoAmIData {
  id?: string;
  fullname?: string;
  avatarUrl?: string;
}

interface AuthStatus {
  ok: boolean;
  errCode: number;
  errMessage: string;
  whoAmI: WhoAmIData;
}

export function Auth(props: {
  endpoint: string;
  onEndpointChanged: (endpoint: string) => void;
  authToken: string;
  onAuthTokenChanged: (authToken: string) => void;
}) {
  const [queryStatus, setQueryStatus] = useState<"ok" | "error" | "pending">(
    "error"
  );
  const [error, setError] = useState<{ code: number; message: string }>({
    code: 0,
    message: "no query yet"
  });
  const [whoAmI, setWhoAmI] = useState<WhoAmIData>({});
  const [refreshCounter, setRefreshCounter] = useState(0);

  useEffect(() => {
    console.log("auth");
    setQueryStatus("pending");
    (async () => {
      let response: Response;
      try {
        response = await fetch(`${props.endpoint}/whoami`, {
          headers: {
            Authorization: `Bearer ${props.authToken}`
          }
        });
      } catch (e) {
        setQueryStatus("error");
        setError({ code: 0, message: `${e}` });
        return;
      }

      if (response.status >= 200 && response.status < 300) {
        setQueryStatus("ok");
        const data = await response.json();
        setWhoAmI({
          id: data.id,
          fullname: data.fullname,
          avatarUrl: data.avatar_url
        });
      } else {
        setQueryStatus("error");
        setError({ code: response.status, message: await response.text() });
      }
    })();
  }, [refreshCounter, props.authToken, props.endpoint]);

  return (
    <Fragment>
      <h3> Discord API Authentication </h3>
      {refreshCounter}

      <Inputs
        fields={[
          {
            name: "Endpoint",
            value: props.endpoint,
            setter: props.onEndpointChanged
          },
          {
            name: "Auth Token",
            value: props.authToken,
            setter: props.onAuthTokenChanged
          }
        ]}
      />

      {queryStatus == "ok" ? (
        <WhoAmI
          data={whoAmI}
          onRefresh={() => setRefreshCounter(refreshCounter + 1)}
          onLogout={() => {
            props.onAuthTokenChanged("");
            fetch(`${props.endpoint}/logout`).then(() => {
              setRefreshCounter(refreshCounter + 1);
            });
          }}
        />
      ) : (
        <Fragment></Fragment>
      )}

      {queryStatus == "error" ? (
        <NotAuthed
          code={error.code}
          message={error.message}
          onRefresh={() => setRefreshCounter(refreshCounter + 1)}
          onLogin={() => {
            window.location.href = `${
              props.endpoint
            }/login?return_url=${encodeURIComponent(window.location.href)}`;
          }}
        />
      ) : (
        <Fragment></Fragment>
      )}

      {queryStatus == "pending" ? <p>Loading...</p> : <Fragment />}
    </Fragment>
  );
}

function WhoAmI(props: {
  data: WhoAmIData;
  onRefresh: () => void;
  onLogout: () => void;
}) {
  return (
    <div className={styles.media}>
      <img
        src={props.data.avatarUrl}
        className={styles.authMediaAvatar}
        alt="avatar"
      />
      <div className={styles["media-body"]}>
        <h5> {props.data.fullname} </h5>
        <p> You are authenticated with the Discord API.</p>
        <div className={styles.authMediaActions}>
          <button
            className={`${styles.btn} ${styles["btn-secondary"]}`}
            onClick={props.onRefresh}
          >
            Refresh
          </button>
          <button
            className={`${styles.btn} ${styles["btn-danger"]}`}
            onClick={props.onLogout}
          >
            Logout
          </button>
        </div>
      </div>
    </div>
  );
}

function NotAuthed(props: {
  code: number;
  message: string;
  onRefresh: () => void;
  onLogin: () => void;
}) {
  return (
    <div className={styles.media}>
      <div className={styles["media-body"]}>
        <h5> Not authenticated! </h5>
        <p> You do not have a valid authentication cookie set.</p>
        <p>
          <strong>Code: </strong>
          <code>{props.code}</code>; <strong>Message: </strong>
          <code>{props.message}</code>
        </p>
        <div className={styles.authMediaActions}>
          <button
            className={`${styles.btn} ${styles["btn-secondary"]}`}
            onClick={props.onRefresh}
          >
            Refresh
          </button>
          <button
            className={`${styles.btn} ${styles["btn-primary"]}`}
            onClick={props.onLogin}
          >
            Login
          </button>
        </div>
      </div>
    </div>
  );
}

export function AuthURLHandling(props: {
  authToken: string;
  onAuthTokenChanged: (authToken: string) => void;
}) {
  useEffect(() => {
    const u = url.parse(document.location.href);
    const params = new URLSearchParams(u.query || undefined);
    const sid = params.get("sid");
    console.log("SID: ", sid);
    props.onAuthTokenChanged(sid || "");
    params.delete("sid");
    u.search = params.toString() ? `?${params.toString()}` : "";
    window.history.replaceState(null, "", url.format(u));
  }, []);
  return <Fragment></Fragment>;
}
