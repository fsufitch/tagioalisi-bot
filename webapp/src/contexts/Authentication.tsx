import React, { ReactNode } from 'react';
import { APIConfigurationContext } from './APIConfiguration';
import { StorageContext } from './Storage';

const AUTH_STORAGE_KEY = 'tagioalisi/auth';

export interface Authentication {
    jwt?: string;
    id?: string;
    fullname?: string;
    avatarUrl?: string;
    error?: string;
}

interface AuthenticationContextValue {
    authentication: Authentication,
    login: () => void,
    logout: () => void,
}

export const AuthenticationContext = React.createContext<AuthenticationContextValue>({
    authentication: {},
    login: () => { },
    logout: () => { },
});

export default (props: { children: ReactNode }) => {
    const [urlJWT, setURLJWT] = React.useState<string>('');
    React.useEffect(() => {
        const jwt = extractJWT();
        console.log('extract jwt from url', jwt);
        if (!jwt) {
            console.log("No JWT data in URL");
            return;
        }
        console.log("Found JWT in URL: ", jwt);
        setURLJWT(jwt);
    }, []);

    const { configuration } = React.useContext(APIConfigurationContext);
    const { state } = React.useContext(StorageContext);    
    const [ authentication, setAuthentication, clearAuthentication ] = state?.useJSON<Authentication>('auth') ?? [{}, () => {}, () => {}];
    React.useEffect(() => {
        if (!configuration || !urlJWT ) {
            return;
        }
        console.log('Running authentication');
        fetch(`${configuration?.baseURL}/whoami`, {
            headers: { Authorization: `Bearer ${urlJWT}` },
        })
            .then(async response => {
                const { id, fullname, avatar_url: avatarUrl } = await response.json();
                console.log("Logged in:", { jwt: urlJWT, id, fullname, avatarUrl });
                setAuthentication({ jwt: urlJWT, id, fullname, avatarUrl });
            })
            .catch(err => {
                console.error(`Authentication error`, err);
            })
    }, [configuration, urlJWT]);


    const login = () => {
        console.log('config in login', configuration)
        if (!configuration?.baseURL) {
            console.error("No base URL configured; cannot log in");
            return;
        }
        const url = new URL(configuration.baseURL);
        url.searchParams.set('return_url', window.location.href);
        url.pathname = '/login';
        console.log('redirect would be', url.toString());
        window.location.href = url.toString();
    }

    const logout = () => {
        clearAuthentication();
    }

    return <AuthenticationContext.Provider value={{ authentication: authentication || {}, login, logout }}>
        {props.children}
    </AuthenticationContext.Provider>
}

const extractJWT = () => {
    const u = new URL(document.location.href);
    const jwt = u.searchParams.get("jwt") ?? '';
    if (!jwt) { return '' };
    u.searchParams.delete("jwt");
    window.history.replaceState(null, "", u);
    return jwt;
}
