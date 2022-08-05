import { atom } from 'recoil';

export const wsState = atom<WebSocket | null>({key: 'ws', default: null});