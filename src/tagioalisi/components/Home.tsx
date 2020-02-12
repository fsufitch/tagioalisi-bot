import React, { useState, Fragment } from "react";

import Logo from "tagioalisi/resources/logo.png";
import styles from "tagioalisi/styles";
import { Inputs } from "./Inputs";

interface HealthResponse {
  response?: Response;
  error: string;
}

function queryBotHealth(
  endpoint: string,
  setLoading: (loading: boolean) => void,
  setResponse: (response: HealthResponse) => void
) {
  setLoading(true);
  const url = endpoint + (endpoint.endsWith("/") ? "/" : "") + "";
  fetch(url).then(
    r => {
      console.log(r);
      setResponse({ response: r, error: "" });
      setLoading(false);
    },
    err => {
      setResponse({ error: `${err}` });
      setLoading(false);
    }
  );
}

function ResponseRender(props: { response: HealthResponse }) {
  return (
    <Fragment>
      {props.response.error ? (
        <p>
          <strong>Query error:</strong> {props.response.error}
        </p>
      ) : (
        ""
      )}

      {props.response.response ? (
        <Fragment>
          <p>
            {props.response.response.status} -{" "}
            {props.response.response.statusText}
          </p>
          <p>
            {props.response.response.ok ? "Bot is OK! :)" : "Bot is not OK :("}
          </p>
        </Fragment>
      ) : (
        ""
      )}
    </Fragment>
  );
}

export function Home(props: {
  endpoint: string;
  onEndpointChanged: (endpoint: string) => void;
}) {
  const [loading, setLoading] = useState(false);
  const [response, setResponse] = useState<HealthResponse>({ error: "" });

  return (
    <div className={styles.home}>
      <img src={Logo} />
      <h2> Hello, world! </h2>
      <p>
        If you found this website, it means you were likely told about the
        Tagioalisi bot. Tagioalisi is a custom Discord bot, and this is its web
        control panel.
      </p>
      <div className={styles.clearfix}></div>

      <h3> Check Bot Health </h3>
      <Inputs
        fields={[
          {
            name: "Endpoint",
            value: props.endpoint,
            setter: props.onEndpointChanged
          }
        ]}
      />
      {loading ? (
        <p> Loading... </p>
      ) : (
        <Fragment>
          <p>
            <button
              onClick={() =>
                queryBotHealth(props.endpoint, setLoading, setResponse)
              }
            >
              {" "}
              Check Bot Health{" "}
            </button>
          </p>
          {!!response ? <ResponseRender response={response} /> : ""}
        </Fragment>
      )}
    </div>
  );
}
