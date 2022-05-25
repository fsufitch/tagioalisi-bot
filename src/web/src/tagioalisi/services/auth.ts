import { useAPIConnection } from './api';
import { useSynchronizedJSONState } from './state';
import { usePromiseEffect } from './async';
import { useState, useEffect } from 'react';

type AuthenticationToken = string;

const AUTH_LOCALSTORAGE_JWT_KEY = 'tagioalisi.auth.jwt';

interface AuthData {
  jwt?: string;
  id?: string;
  fullname?: string;
  avatarUrl?: string;
  error?: string;
}


export function useAuthentication() {
  const [endpoint] = useAPIConnection();

  const [authData, setAuthData] = useSynchronizedJSONState<AuthData>(AUTH_LOCALSTORAGE_JWT_KEY, {});
  usePromiseEffect(async () => {
    const jwt = extractJWT();
    if (!jwt) {
      return;
    }
    console.log("Found JWT in URL: ", jwt);

    try {
      const whoamiResponse = await fetch(`${endpoint.baseUrl}/whoami`, {
        headers: { Authorization: `Bearer ${jwt}` },
      });
      const { id, fullname, avatar_url: avatarUrl } = await whoamiResponse.json();
      console.log("Logged in:", { jwt, id, fullname, avatarUrl });
      setAuthData({ jwt, id, fullname, avatarUrl });
    }
    catch (err) {
      setAuthData({ error: `${err}` });
    }
  }, []);

  const login = () => {
    window.location.href = `${endpoint.baseUrl}/login?return_url=${encodeURIComponent(window.location.href)}`;
  }

  // Logout processing requires an effect to update state properly. Yucky, but React is React.
  const [triggerLogout, setTriggerLogout] = useState<boolean>(false);
  useEffect(() => {
    if (triggerLogout) {
      fetch(`${endpoint.baseUrl}/logout`)
      setAuthData({});
      setTriggerLogout(false);
    }
  })

  const logout = () => setTriggerLogout(true);

  return [authData, login, logout] as [AuthData, () => void, () => void];
}


const extractJWT = () => {
  const u = new URL(document.location.href);
  const jwt = u.searchParams.get("jwt") ?? '';
  if (!jwt) { return '' };
  u.searchParams.delete("jwt");
  window.history.replaceState(null, "", u);
  return jwt;
}
