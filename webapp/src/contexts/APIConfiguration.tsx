import React, { ReactNode } from 'react';

export const getDefaultBaseURL = () => {
    const url = __BOT_BASE_URL__ || 'http://localhost:8091'; // XXX: HARDCODING
    return url;
}

export const getDefaultGrpcBaseURL = () => {
    const url = __BOT_GRPC_BASE_URL__ || 'https://localhost:8092'; // XXX: HARDCODING
    return url;

}

export interface APIConfiguration {
    baseURL?: string;
    grpcBaseURL?: string;
}

interface APIConfigurationContextValue {
    configuration: APIConfiguration;
    setBaseURL: (url: string) => void;
    setGrpcBaseURL: (url: string) => void;
}

export const APIConfigurationContext = React.createContext<APIConfigurationContextValue>({
    configuration: {},
    setBaseURL: (url: string) => {},
    setGrpcBaseURL: (url: string) => {},
});

export default (props: {children: ReactNode}) => {
    const [configuration, setConfiguration] = React.useState<APIConfiguration>({
        baseURL: getDefaultBaseURL(),
        grpcBaseURL: getDefaultGrpcBaseURL(),
    });

    const setBaseURL = (baseURL: string) => setConfiguration({...configuration, baseURL});

    const setGrpcBaseURL = (grpcBaseURL: string) => setConfiguration({...configuration, grpcBaseURL})

    React.useEffect(() => {
        console.log('config', configuration);
    }, [configuration]);

    return <APIConfigurationContext.Provider value={{configuration, setBaseURL, setGrpcBaseURL}}>
        {props.children}
    </APIConfigurationContext.Provider>

}
