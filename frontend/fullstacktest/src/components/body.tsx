'use client';
import { useState } from 'react';
import Button from './button';

export default function Body() {
  return (
    <div className="flex justify-between">
      <Button text="Add Task" bgcolor="bg-blue-700" textcolor="text-white" />
      <Filter options={['ALL', 'Complete', 'Incomplete']} />
    </div>
  );
}

interface FilterParams {
  options: string[];
}

function Filter({ options }: FilterParams) {
  const [activeIndex, setActiveIndex] = useState<number>(0);

  function handler() {
    setActiveIndex((past) => {
      if (past == 2) {
        return 0;
      }
      return past + 1;
    });
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