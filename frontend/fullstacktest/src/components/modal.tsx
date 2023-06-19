import { useState } from 'react';
import Button from './button';
import { useAppDispatch, useAppSelector } from '@/hooks';

import { setStatus } from '@/slices/modalslice';
import { addTodo } from '@/slices/tododataslice';

import { Todo } from '@/interfaces';

export default function Modal() {
  const [value, setValue] = useState<string>('');

  const lastid = useAppSelector((state) => state.id.lastId);
  const dispatch = useAppDispatch();

  const handlerClick = (type: number) => {
    switch (type) {
      case 0:
        dispatch(setStatus(false));
        break;
      case 1:
        if (value) {
          const tmp = async () => {
            const responce = await fetch('http://localhost:3001/api/getlastid');
            const result = await responce.json();

            dispatch(
              addTodo({
                Text: value,
                Status: false,
                ID: result + 1,
              } as Todo)
            );
          };

          tmp();
          dispatch(setStatus(false));
          break;
        }
    }
  };

  return (
    <>
      <div className="bg-black fixed opacity-[50%] w-full h-full"></div>
      <div className="w-[400px] fixed h-[200px] bg-white opacity-100 top-[100px] rounded-md p-3 flex flex-col px-6 justify-between">
        <div>
          <p className=" font-lato text-slate-700 ">Add task</p>
          <p className="text-slate-600 text-sm  mt-3">Title</p>
          <input
            onChange={(e) => {
              setValue(e.target.value);
            }}
            className="w-full bg-white h-10 border-[2px] border-gray-500 pl-3"
            value={value}
          ></input>
        </div>
        <div className="flex justify-between">
          <Button
            onClick={() => handlerClick(1)}
            hovercolor="hover:bg-blue-400"
            text="Add task"
            bgcolor="bg-blue-600"
            textcolor="text-white"
          />
          <Button
            onClick={() => handlerClick(0)}
            hovercolor="hover:bg-red-500"
            text="Cancel"
            bgcolor="bg-slate-200"
            textcolor="text-gray-400"
          />
        </div>
      </div>
    </>
  );
}
