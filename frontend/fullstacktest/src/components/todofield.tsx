'use client';
import { Todo } from '@/interfaces';
import Checkbox from './checkbox';
import { useAppDispatch, useAppSelector } from '@/hooks';
import { deletetodo, switchStatus, updateText } from '@/slices/tododataslice';
import { KeyboardEvent, useEffect, useRef, useState } from 'react';

interface TodoListParams {
  todos: Todo[];
}

export default function TodoList({ todos }: TodoListParams) {
  const [copyTodos, setCopyTodos] = useState<Todo[]>(todos);
  const type = useAppSelector((state) => state.filter.type);

  useEffect(() => {
    switch (type) {
      case 'ALL':
        setCopyTodos(todos);
        break;
      case 'Complete':
        setCopyTodos(
          todos.filter((element) => {
            return element.Status;
          })
        );
        break;
      case 'Incomplete':
        setCopyTodos(
          todos.filter((element) => {
            return !element.Status;
          })
        );
        break;
    }
  }, [type, todos]);

  if (!copyTodos || copyTodos.length == 0) {
    return (
      <div className="mt-[20px] bg-gray-200 rounded-md min-h-[80px] flex justify-center">
        <div className="bg-stone-400 text-gray-700 w-32 h-8 flex justify-center rounded-lg my-auto">
          <span className="my-auto">No Todo Found</span>
        </div>
      </div>
    );
  }
  return (
    <div className="mt-[20px] bg-gray-200 rounded-md min-h-[70px] pt-5">
      {copyTodos.map((todo) => {
        return (
          <Todo
            key={todo.Id}
            Text={todo.Text}
            Status={todo.Status}
            Id={todo.Id}
          />
        );
      })}
    </div>
  );
}

function Todo({ Text, Status, Id }: Todo) {
  const textareaRef = useRef<HTMLTextAreaElement>(null);

  const dispatch = useAppDispatch();

  const UpdateTextHandler = () => {
    if (!textareaRef.current) {
      return;
    }
    textareaRef.current.disabled = false;
    textareaRef.current.focus();
  };

  const onKeyDownHandler = (e: KeyboardEvent) => {
    if (e.key == 'Enter' && textareaRef.current) {
      dispatch(updateText({ ID: Id, Text: textareaRef.current.value }));
      textareaRef.current.disabled = true;
    }
  };

  return (
    <div className="px-5 pb-5">
      <div className="bg-white rounded-sm h-10 select-none flex justify-between">
        <div className="flex">
          <div
            className="my-auto ml-2 mr-2"
            onClick={() => {
              dispatch(switchStatus(Id));
            }}
          >
            <Checkbox checked={Status} />
          </div>
          {Status ? (
            <textarea
              onKeyDown={(e) => onKeyDownHandler(e)}
              ref={textareaRef}
              disabled={true}
              defaultValue={Text}
              className="my-auto line-through h-5 resize-none overflow-hidden outline-none"
            />
          ) : (
            <textarea
              onKeyDown={(e) => onKeyDownHandler(e)}
              defaultValue={Text}
              ref={textareaRef}
              disabled={true}
              className="my-auto resize-none h-5 overflow-y-hidden outline-none"
            />
          )}
        </div>
        <div className="flex my-auto mr-2">
          <div className="w-[25px] h-[25px]">
            <svg
              onClick={() => {
                dispatch(deletetodo(Id));
              }}
              fill="currentColor"
              viewBox="0 0 16 16"
              className="bg-gray-300 hover:bg-gray-200 text-gray-500 p-1"
            >
              <path d="M2.5 1a1 1 0 00-1 1v1a1 1 0 001 1H3v9a2 2 0 002 2h6a2 2 0 002-2V4h.5a1 1 0 001-1V2a1 1 0 00-1-1H10a1 1 0 00-1-1H7a1 1 0 00-1 1H2.5zm3 4a.5.5 0 01.5.5v7a.5.5 0 01-1 0v-7a.5.5 0 01.5-.5zM8 5a.5.5 0 01.5.5v7a.5.5 0 01-1 0v-7A.5.5 0 018 5zm3 .5v7a.5.5 0 01-1 0v-7a.5.5 0 011 0z" />
            </svg>
          </div>
          <div
            className="w-[25px] h-[25px] ml-[10px]"
            onClick={() => UpdateTextHandler()}
          >
            <svg
              viewBox="0 0 24 24"
              fill="currentColor"
              className="bg-gray-300 hover:bg-gray-200 text-gray-500 p-[2px]"
            >
              <path d="M20.71 7.04c.39-.39.39-1.04 0-1.41l-2.34-2.34c-.37-.39-1.02-.39-1.41 0l-1.84 1.83 3.75 3.75M3 17.25V21h3.75L17.81 9.93l-3.75-3.75L3 17.25z" />
            </svg>
          </div>
        </div>
      </div>
    </div>
  );
}
