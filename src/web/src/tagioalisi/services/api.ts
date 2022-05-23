import { useCookies } from 'react-cookie';
import { useLocalStorage, SetterFunc } from './localStorage';
import { useSynchronizedJSONState } from './state';


export const useDefaultAPIEndpoint = () => {
    const [cookies] = useCookies(['BOT_EXTERNAL_BASE_URL']);
    const defaultBaseUrl = `${cookies.BOT_EXTERNAL_BASE_URL}`;
    return defaultBaseUrl;
}

export interface APIConnection {
    baseUrl: string;
}


export const useAPIConnection = () => 
    useSynchronizedJSONState<APIConnection>('tagioalisi/api', {
        baseUrl: useDefaultAPIEndpoint(),
    });
