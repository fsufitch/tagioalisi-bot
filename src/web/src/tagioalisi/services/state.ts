import { useEffect, useState } from 'react';
export type Serializer<T> = (data: T) => string
export type Deserializer<T> = (raw: string) => T


class SynchronizedState {
    constructor(public readonly storageKeyPrefix: string, public readonly storage: Storage) {}

    use<T>(key: string, defaultValue: T, serialize: Serializer<T>, deserialize: Deserializer<T>) {
        const storageKey = this.storageKeyPrefix + key;

        // Initialize the localStorage entry if necessary
        let defaultStateValue: T = defaultValue;
        const initialStorageValue = localStorage.getItem(storageKey);
        if (initialStorageValue != null) {
            try {
                defaultStateValue = deserialize(initialStorageValue);
            } catch (err) {
                console.error(`Could not load initial value for key '${key}' from storage`, err);
                localStorage.setItem(key, serialize(defaultValue));
            }
        }
        
        const [value, setValue] = useState<T>(defaultStateValue);
        // Async load value from storage events
        useEffect(() => {
            const cb = (ev: StorageEvent) => {
                if (ev.key != storageKey) {
                    return;
                }
                let newValue = localStorage.getItem(storageKey);
                if (newValue == null) {
                    console.log(`After storage update, key '${key}' contained null`);
                    setValue(defaultValue);
                    return;
                }
                try {
                    setValue(deserialize(newValue));
                } catch (ex) {
                    console.error(`After update, key ${key} had invalid value`, newValue, ex);
                    setValue(defaultValue);
                }
            };
            window.addEventListener('storage', cb);
            // Deregister the callback when the user component is unloaded
            return () => {
                window.removeEventListener('storage', cb);
            };
        }, []);

        // writeValue ephemerally contains (and triggers) writing values to localstorage
        const [writeValue, setWriteValue] = useState<T | null>(null);
        useEffect(() => {
            if (writeValue == null) return;
            setWriteValue(null);
            localStorage.setItem(storageKey, serialize(writeValue));
            setValue(writeValue); // Need to also manually update it here, since storage events only happen in other documents
        }, [writeValue]);

        // removeValue ephemerally contains a trigger to delete stuff from localstorage
        const [removeValue, setRemoveValue] = useState<boolean>(false);
        useEffect(() => {
            if (!removeValue) return;
            setRemoveValue(false);
            localStorage.removeItem(storageKey);
            setValue(defaultValue); // Need to also manually update it here, since storage events only happen in other documents
        }, [writeValue]);


        return [value, setWriteValue, () => setRemoveValue(true)] as [T, (newValue: T) => void, () => void];


    }

    // Static below

    private static _instance?: SynchronizedState;
    static get instance() {
        this._instance = this._instance ?? new SynchronizedState('sync-state::', sessionStorage);
        return this._instance;
    }
}


export const useSynchronizedState = <T>(key: string, defaultValue: T, serialize: Serializer<T>, deserialize: Deserializer<T>) => 
    SynchronizedState.instance.use(key, defaultValue, serialize, deserialize);

// Convenience JSON wrapepr

export const useSynchronizedJSONState = <T>(key: string, defaultValue: T) =>
    useSynchronizedState<T>(key, defaultValue, JSON.stringify, JSON.parse);