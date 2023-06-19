'use client';
import { useState } from 'react';
import Button from './button';

import { useAppDispatch } from '@/hooks';
import { setStatus } from '@/slices/modalslice';
import { changeType } from '@/slices/filterSlice';

export default function Body() {
  const dispatch = useAppDispatch();

  const handlerClick = () => {
    dispatch(setStatus(true));
  };

  const handlerSwitch = (type: string) => {
    dispatch(changeType(type));
  };

  return (
    <div className="flex justify-between">
      <Button
        onClick={() => handlerClick()}
        hovercolor="hover:bg-blue-400"
        text="Add Task"
        bgcolor="bg-blue-700"
        textcolor="text-white"
      />
      <Filter
        options={['ALL', 'Complete', 'Incomplete']}
        onSwitch={handlerSwitch}
      />
    </div>
  );
}

interface FilterParams {
  options: string[];
  onSwitch: (params: any) => any;
}

function Filter({ options, onSwitch }: FilterParams) {
  const [activeIndex, setActiveIndex] = useState<number>(0);

  function handler() {
    if (activeIndex != 2) {
      setActiveIndex((past) => {
        return past + 1;
      });
      onSwitch(options[activeIndex + 1]);
    } else {
      setActiveIndex(0);
      onSwitch(options[0]);
    }
  }

  return (
    <div
      onClick={handler}
      className={
        'bg-gray-300 text-gray-500 w-28 h-8 rounded-md flex justify-center font-lato cursor-pointer select-none hover:bg-gray-200'
      }
    >
      <span className="my-auto">{options[activeIndex]}</span>
    </div>
  );
}
