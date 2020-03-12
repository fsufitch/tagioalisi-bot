import React, { useState, Fragment } from "react";

import styles from "tagioalisi/styles";

import { Inputs } from "./Inputs";

export function Sockpuppet(props: {
  endpoint: string;
  onEndpointChanged: (endpoint: string) => void;
  authToken: string;
  onAuthTokenChanged: (token: string) => void;
}) {
  return (
    <div>
      <h2> Sockpuppet Module </h2>
      <p>
        Tagioailisi supports a system for sending custom messages to any of the
        channels it is in. Sending a message requires &quot;Manage Messages&quot; permissions
        in the target channel (be sure to authenticate).
      </p>
      <SendMessageSection
        endpoint={props.endpoint}
        onEndpointChanged={props.onEndpointChanged}
        authToken={props.authToken}
        onAuthTokenChanged={props.onAuthTokenChanged}
      />
    </div>
  );
}

function SendMessageSection(props: {
  endpoint: string;
  onEndpointChanged: (endpoint: string) => void;
  authToken: string;
  onAuthTokenChanged: (authToken: string) => void;
}) {
  const [message, setMessage] = useState("");
  const [resultMessage, setResultMessage] = useState("");
  const [loading, setLoading] = useState(false);
  const [channelID, setChannelID] = useState("");

  const send = () => {
    const url = `${props.endpoint}${
      props.endpoint.endsWith("/") ? "" : "/"
    }sockpuppet`;

    setLoading(true);
    fetch(url, {
      method: "POST",
      body: JSON.stringify({ channelID, message }),
      headers: {
        "Content-Type": "application/json",
        "Authorization": `Bearer ${props.authToken}`,
      }
    })
      .then(
        r =>
          r.text().then(text => ({
            ok: r.ok,
            text: r.ok ? "Message sent successfully" : text
          })),
        reason => ({
          ok: false,
          text: `error: ${reason}`
        })
      )
      .then(r => {
        if (r.ok) {
          setMessage("");
        }
        setResultMessage(r.text);
        setLoading(false);
      });
  };

  return (
    <Fragment>
      <h3> Send Message to Channel </h3>
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
          },
          {
            name: "Channel ID",
            value: channelID,
            setter: setChannelID
          },
          {
            name: "Message",
            value: message,
            setter: setMessage
          }
        ]}
      />
      {loading ? (
        <p> Loading... </p>
      ) : (
        <Fragment>
          <p>
            <button className={`${styles.btn} ${styles["btn-primary"]}`} onClick={() => send()}> Send Message </button>
          </p>
          {!!resultMessage ? (
            <p>
              <strong>Result:</strong>
              {resultMessage}
            </p>
          ) : (
            ""
          )}
        </Fragment>
      )}
    </Fragment>
  );
}
