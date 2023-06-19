import { configureStore } from '@reduxjs/toolkit';

import modalslice from './slices/modalslice';
import tododataslice from './slices/tododataslice';

export const store = configureStore({
  reducer: {
    modal: modalslice,
    todo: tododataslice,
  },
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
