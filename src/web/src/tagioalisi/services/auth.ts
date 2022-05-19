import { useEffect } from 'react';
import * as url from "url";

import { useAPIEndpoint } from 'tagioalisi/services/api';
import { useLocalStorage } from 'tagioalisi/services/localStorage';

type AuthenticationToken = string;

const AUTH_LOCALSTORAGE_JWT_KEY = 'tagioalisi.auth.jwt';

export function useAuthentication() {
  const [endpoint] = useAPIEndpoint();
  const [jwt, setJWT] = useLocalStorage<AuthenticationToken>(AUTH_LOCALSTORAGE_JWT_KEY, '');

  const login = () => {
    window.location.href = `${endpoint}/login?return_url=${encodeURIComponent(window.location.href)}`;
  }

  const logout = () => {
    fetch(`${endpoint}/logout`);
    setJWT('');
  }

  return { jwt, setJWT, login, logout };
}

export const useOnLoadAuthenticationEffect = () => {
  const { setJWT } = useAuthentication();
  useEffect(() => {
    const u = url.parse(document.location.href);
    const params = new URLSearchParams(u.query ?? undefined);
    const jwt = params?.get("jwt") ?? '';
    if (!jwt) {
      return;
    }
    console.log("Found JWT in URL: ", jwt);
    setJWT(jwt);

    params.delete("jwt");
    u.search = params.toString() ? `?${params.toString()}` : "";
    window.history.replaceState(null, "", url.format(u));
  }, []);
}

interface UserData {
  authenticated: boolean;
  authPending: boolean;
  error?: string;
  id?: string;
  fullname?: string;
  avatarUrl?: string;
}

const AUTH_LOCALSTORAGE_USER_DATA_KEY = 'tagioalisi.auth.user-data';

export function useAuthenticatedUserData(): [UserData, (val: UserData) => void] {
  const [userData, setUserData] = useLocalStorage<UserData>(AUTH_LOCALSTORAGE_USER_DATA_KEY, { authenticated: false, authPending: false });
  return [userData, setUserData];
}

export function useUpdateAuthenticatedUserDataEffect() {
  const { jwt } = useAuthentication();
  const [, setUserData] = useAuthenticatedUserData();
  const [endpoint] = useAPIEndpoint();
  console.log('fff');

  useEffect(() => {
    console.log('omg');
    if (!jwt || !endpoint) {
      setUserData({ authenticated: false, authPending: false });
      return;
    }

    setUserData({ authPending: true, authenticated: false });
    new Promise(async () => {
      let response: Response;
      try {
        response = await fetch(`${endpoint}/whoami`, {
          headers: {
            Authorization: `Bearer ${jwt}`
          }
        });
      } catch (e) {
        setUserData({
          authenticated: false,
          authPending: false,
          error: `error running user data fetch: ${e}`,
        })
        return;
      }

      if (response.status >= 200 && response.status < 300) {
        const data = await response.json();
        setUserData({
          authenticated: true,
          authPending: false,
          id: data.id,
          fullname: data.fullname,
          avatarUrl: data.avatar_url,
        })
      } else {
        setUserData({
          authenticated: false,
          authPending: false,
          error: `Error (${response.status}): ${await response.text()}`,
        });
      }
    });
    console.log('wtf');
  }, [jwt, endpoint]);
}