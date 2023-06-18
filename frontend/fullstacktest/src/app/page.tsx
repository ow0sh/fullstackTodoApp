'use client';
import { useState } from 'react';

import Header from '@/components/header';
import Body from '@/components/body';
import TodoList from '@/components/todofield';

import Modal from '@/components/modal';

export default function Home() {
  const [modal, setModal] = useState<boolean>(true);
  return (
    <>
      {modal && <Modal />}
      <div className="w-[600px] flex flex-col ">
        <Header />
        <Body />
        <TodoList todos={[{ text: 'test', status: false, id: 0 }]} />
      </div>
    </>
  );
}
