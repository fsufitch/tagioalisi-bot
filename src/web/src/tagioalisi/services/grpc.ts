import React from 'react';
import { useAuthentication } from './auth';
import {grpc} from '@improbable-eng/grpc-web';

import { GreeterClient } from 'tagioalisi/proto/hello_pb_service';
import { useAPIConnection } from './api';

export const useGreeterClient = () => {
    const [api] = useAPIConnection();
    const [authData] = useAuthentication();
    
    const transport = grpc.WebsocketTransport();
    const client = React.useMemo(() => new GreeterClient(api.baseUrl, {transport}), [api, authData]);
    
    return client;
}
