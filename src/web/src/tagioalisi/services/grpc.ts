import React from 'react';
import { useCookies } from 'react-cookie';

import { useSynchronizedJSONState } from './state';
import { useAuthentication } from './auth';

import { GreeterClient } from 'tagioalisi/proto/hello_pb_service';

export const useDefaultGRPCEndpoint = () => {
    const [cookies] = useCookies(['BOT_EXTERNAL_GRPC_URL']);
    const defaultBaseUrl = `${cookies.BOT_EXTERNAL_GRPC_URL}`;
    return defaultBaseUrl;
}

export interface GRPCConfiguration {
    url: string;
}

export const useGRPCConfiguration = () => 
    useSynchronizedJSONState<GRPCConfiguration>('tagioalisi/api', {
        url: useDefaultGRPCEndpoint(),
    });


export const useGreeterClient = () => {
    const [grpcConfig] = useGRPCConfiguration();
    const [authData] = useAuthentication();
    
    const client = React.useMemo(() => new GreeterClient(grpcConfig.url), [grpcConfig, authData]);
    
    return client;
}
