import { useEffect } from 'react';
export type Serializer<T> = (data: T) => string
export type Deserializer<T> = (raw: string) => T


class SynchronizedState {
    private keyUpdateCallbacks: {[key: string]: Set<() => void>} = {}

    constructor(public readonly storageKeyPrefix: string) {
        window.addEventListener('storage', (ev) => this.handleStorageEvent(ev));
    }

    private handleStorageEvent(event: StorageEvent) {
        if (!event.key?.startsWith(this.storageKeyPrefix)) {
            // Only handle events actually relevant to our interests
            return;
        }
        const key = event.key.slice(this.storageKeyPrefix.length);
        if (!key) {
            // Weird edge case
            console.error('Empty storage event key after prefix strip', event);
            return;
        }
    }

    private notifyKey(key: string) {
        // Call the callbacks for this key; async, since who knows how heavy they might be
        const callbacks = this.keyUpdateCallbacks[key] ?? [];
        for (const cb of callbacks) {
            Promise.resolve().then(cb);
        }
    }

    useSynchronizedState<T>(key: string, defaultValue: T, serialize: Serializer<T>, deserialize: Deserializer<T>) {
        const storageKey = this.storageKeyPrefix + key;
        let value: T;

        // Initialize the localStorage entry if necessary
        const existingValue = localStorage.getItem(storageKey)
        if (existingValue == null) {
            value = defaultValue;
            localStorage.setItem(storageKey, serialize(value));
        } else {
            value = deserialize(existingValue);
        }

        const setter = (newValue: T) => {
            value = newValue;
            localStorage.setItem(storageKey, serialize(newValue));
            this.notifyKey(key);
        }

        const notifyKeyCallback = () => {
            const rawValue = localStorage.getItem(storageKey);
            if (rawValue == null) {
                console.warn(`Received update for key '${key}' but value was missing; using default`);
                value = defaultValue;
            } else {
                value = deserialize(rawValue);
            }
        }

        useEffect(() => {
            // Register the update notify callback, and a cleanup
            this.keyUpdateCallbacks[key] ??= new Set();
            this.keyUpdateCallbacks[key].add(notifyKeyCallback);
            return () => {
                this.keyUpdateCallbacks[key].delete(notifyKeyCallback);
            }
        }, []);

        return [value, setter] as [T, (newValue: T) => void];
    }

    // Static below

    private static _instance?: SynchronizedState;
    static get instance() {
        this._instance = this._instance ?? new SynchronizedState('sync-state::');
        return this._instance;
    }
}


export const useSynchronizedState = <T>(key: string, defaultValue: T, serialize: Serializer<T>, deserialize: Deserializer<T>) => 
    SynchronizedState.instance.useSynchronizedState(key, defaultValue, serialize, deserialize);

// Convenience JSON wrapepr

export const useSynchronizedJSONState = <T>(key: string, defaultValue: T) =>
    useSynchronizedState<T>(key, defaultValue, JSON.stringify, JSON.parse);