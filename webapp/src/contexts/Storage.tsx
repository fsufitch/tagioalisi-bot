
import React, { ReactNode } from 'react';

const PREFIX = 'tagioalisi::';
const STORAGE = window.sessionStorage;

type UseJSONOutput<T> = [
    T | null,  // value
    (newValue: T) => void, // setter
    () => void,  // clearer
];

class CrossDocumentState {
    constructor(private prefix: string, private storage: Storage) {}

    parse<T>(data: string, default_: T): T {
        try {
            return JSON.parse(data);
        } catch (err) {
            console.error(`Invalid JSON: ${data}`);
            return default_;
        }
    }

    useJSON<T>(key: string): UseJSONOutput<T>{
        const storageKey = `${this.prefix}${key}`;

        const initialStorageValue = this.storage.getItem(storageKey);
        console.log("initial local value in storage", storageKey, initialStorageValue);
        const defaultStateValue = initialStorageValue == null ? null : this.parse<T | null>(initialStorageValue, null);

        const [value, setValue] = React.useState<T | null>(defaultStateValue);
        // Async load value from storage events
        React.useEffect(() => {
            const cb = (ev: StorageEvent) => {
                if (ev.key != storageKey) {
                    return;
                }
                let newValueRaw = this.storage.getItem(storageKey);
                const newValue = newValueRaw == null ? null : this.parse<T | null>(newValueRaw, null);
                setValue(newValue);
                console.log("update local value from storage", storageKey, newValue);
            };
            window.addEventListener('storage', cb);
            // Deregister the callback when the user component is unloaded
            return () => {
                window.removeEventListener('storage', cb);
            };
        }, []);

        const [writeValue, setWriteValue] = React.useState<T | null>(null);
        React.useEffect(() => {
            // Write any non-null value in writeValue to storage
            if (writeValue === null) {
                return;
            }
            setWriteValue(null);
            this.storage.setItem(storageKey, JSON.stringify(writeValue));
            setValue(writeValue);
            console.log("update local value from writeValue", storageKey, writeValue);
        }, [writeValue])
        const save = (newValue: T) => setWriteValue(newValue);

        const [removeValue, setRemoveValue] = React.useState<boolean>(false);
        React.useEffect(() => {
            if (!removeValue) {
                return;
            }
            setRemoveValue(false);
            this.storage.removeItem(storageKey);
            setValue(null); // Need to also manually update it here, since storage events only happen in other documents
            console.log("clear local value", storageKey);
        }, [removeValue]);
        const clear = () => setRemoveValue(true);

        React.useEffect(() => {
            console.log("value updated", value);
        }, [value]);

        return [value, save, clear];
    }
}

interface StorageContextValue {
    state?: CrossDocumentState,
}

export const StorageContext = React.createContext<StorageContextValue>({ });

export default (props: {prefix?: string, storage?: Storage, children: ReactNode}) => {
    const prefix = props.prefix ?? PREFIX;
    const storage = props.storage ?? STORAGE;
    const state = new CrossDocumentState(prefix, storage);

    return <StorageContext.Provider value={{state}}>{props.children}</StorageContext.Provider>;
}