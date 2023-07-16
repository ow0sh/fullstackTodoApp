'use client';
import { useAppSelector, useAppDispatch } from '@/hooks';
import { Provider } from 'react-redux';
import { store } from '@/store';

import Header from '@/components/header';
import Body from '@/components/body';
import TodoList from '@/components/todofield';

import Modal from '@/components/modal';
import { useEffect } from 'react';
import { fetchData } from '@/slices/tododataslice';

export default function Page() {
  return (
    <>
      <Provider store={store}>
        <Home />
      </Provider>
    </>
  );
}

const Home = () => {
  const dispatch = useAppDispatch();

  const modal = useAppSelector((state) => state.modal.status);
  const todos = useAppSelector((state) => state.todo.data);

  useEffect(() => {
    const tmp = async () => {
      const responce = await fetch('http://localhost:3001/api/gettodos');
      const result = await responce.json();
      dispatch(fetchData(result));
    };
    tmp();
  }, [dispatch]);

  return (
    <>
      {modal && <Modal />}
      <div className="w-[600px] flex flex-col ">
        <Header />
        <Body />
        <TodoList todos={todos} />
      </div>
    </>
  );
};
