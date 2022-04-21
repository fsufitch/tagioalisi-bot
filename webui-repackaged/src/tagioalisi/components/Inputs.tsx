import React from "react";
import styles from "tagioalisi/styles";

interface BotInput {
  name: string;
  value: string;
  setter?: (val: string) => void;
}

export function Inputs(props: { fields: BotInput[] }) {
  return (
    <ul className={styles.inputs}>
      {props.fields.map((input, i) => (
        <li key={i}>
          <label htmlFor={input.name}>{input.name}</label>
          {!!input.setter ?
            <input
              type="text"
              id={input.name}
              value={input.value}
              onChange={e => !!input.setter ? input.setter(e.target.value) : void (0)}
            />
            :
            <input
              type="text"
              id={input.name}
              value={input.value}
              disabled={true}
            />
          }
        </li>
      ))}
    </ul>
  );
}
