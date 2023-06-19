import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import { RootState } from '@/store';

interface InitialState {
  status: boolean;
}

const initialState = { status: false } as InitialState;

const modalSlice = createSlice({
  name: 'modalSlice',
  initialState,
  reducers: {
    setStatus: (state, action: PayloadAction<boolean>) => {
      state.status = action.payload;
    },
  },
});

export const { setStatus } = modalSlice.actions;
export const selectModalStatus = (state: RootState) => state.modal;
export default modalSlice.reducer;
