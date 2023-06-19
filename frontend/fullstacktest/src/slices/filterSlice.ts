import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import { RootState } from '@/store';

interface InitialState {
  type: string;
}

const initialState = { type: 'ALL' } as InitialState;

const filterSlice = createSlice({
  name: 'filterSlice',
  initialState,
  reducers: {
    changeType: (state, action: PayloadAction<string>) => {
      console.log(action.payload);
      state.type = action.payload;
    },
  },
});

export const { changeType } = filterSlice.actions;
export const selectFilterType = (state: RootState) => state.filter;
export default filterSlice.reducer;
