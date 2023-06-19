import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import { RootState } from '@/store';

import { Todo } from '@/interfaces';

interface InitialState {
  data: Todo[];
}

const initialState = { data: [] } as InitialState;

export const todoDataSlice = createSlice({
  name: 'tododataslice',
  initialState,
  reducers: {
    addTodo: (state, action: PayloadAction<Todo>) => {
      fetch('http://localhost:3001/api/inserttodo', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(action.payload),
      });

      if (state.data) {
        let tmp = state.data.slice();
        tmp.push(action.payload);
        state.data = tmp;
        return;
      }
      let tmp = [] as Todo[];
      tmp.push(action.payload);
      state.data = tmp;
    },
    deletetodo: (state, action: PayloadAction<number>) => {
      fetch('http://localhost:3001/api/deletetodo', {
        method: 'DELETE',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(action.payload),
      });

      let tmp = state.data.slice();
      state.data = tmp.filter((element) => element.ID != action.payload);
    },
    fetchData: (state, action: PayloadAction<Todo[]>) => {
      state.data = action.payload;
    },
  },
});

export const { addTodo, deletetodo, fetchData } = todoDataSlice.actions;
export const selectTodoData = (state: RootState) => state.todo;
export default todoDataSlice.reducer;
