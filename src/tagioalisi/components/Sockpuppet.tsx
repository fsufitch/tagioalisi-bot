import React, { useState, Fragment } from "react";

import styles from "tagioalisi/styles";

import { Inputs } from "./Inputs";
import { useAPIEndpoint } from "tagioalisi/services/api";
import { useAuthentication } from "tagioalisi/services/auth";

export function Sockpuppet() {
  return (
    <div>
      <h2> Sockpuppet Module </h2>
      <p>
        Tagioailisi supports a system for sending custom messages to any of the
        channels it is in. Sending a message requires &quot;Manage Messages&quot; permissions
        in the target channel (be sure to authenticate).
      </p>
      <SendMessageSection />
    </div>
  );
}

function SendMessageSection() {
  const [message, setMessage] = useState("");
  const [resultMessage, setResultMessage] = useState("");
  const [loading, setLoading] = useState(false);
  const [channelID, setChannelID] = useState("");
  const [endpoint] = useAPIEndpoint();
  const [jwt] = useAuthentication();

  const send = () => {
    const url = `${endpoint}${
      endpoint.endsWith("/") ? "" : "/"
    }sockpuppet`;

    setLoading(true);
    fetch(url, {
      method: "POST",
      body: JSON.stringify({ channelID, message }),
      headers: {
        "Content-Type": "application/json",
        "Authorization": `Bearer ${jwt}`,
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
            value: endpoint,
          },
          {
            name: "Auth Token",
            value: jwt,
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
