import React from 'react';
import {grpc} from '@improbable-eng/grpc-web';

import { GreeterClientImpl } from '@tagioalisi/proto/hello';
import { SockpuppetClientImpl } from '@tagioalisi/proto/sockpuppet';
import { AuthenticationContext } from '@tagioalisi/contexts/Authentication';
import { APIConfigurationContext } from '@tagioalisi/contexts/APIConfiguration';

export const useGreeterClient = () => {
    const { configuration } = React.useContext(APIConfigurationContext)
    const { authentication } = React.useContext(AuthenticationContext);
    

    const c = grpc.client(GreeterClientImpl, {host: configuration.baseURL})
    const transport = grpc.WebsocketTransport();
    const client = React.useMemo(() => new GreeterClientImpl(configuration.baseURL ?? '', {transport}), [configuration, authentication]);
    
    return client;
}

export const useSockpuppetClient = () => {
    const { configuration } = React.useContext(APIConfigurationContext)
    const { authentication } = React.useContext(AuthenticationContext);
    
    const transport = grpc.WebsocketTransport();
    const client = React.useMemo(() => new SockpuppetClient(configuration.baseURL ?? '', {transport}), [configuration, authentication]);
    
    return client;

}
