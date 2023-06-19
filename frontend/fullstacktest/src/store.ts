import { configureStore } from '@reduxjs/toolkit';

import modalslice from './slices/modalslice';
import tododataslice from './slices/tododataslice';
import filterSlice from './slices/filterSlice';

export const store = configureStore({
  reducer: {
    modal: modalslice,
    todo: tododataslice,
    filter: filterSlice,
  },
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
