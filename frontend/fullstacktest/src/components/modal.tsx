import Button from './button';

export default function Modal() {
  return (
    <>
      <div className="bg-black fixed opacity-[50%] w-full h-full"></div>
      <div className="w-[400px] fixed h-[200px] bg-white opacity-100 top-[100px] rounded-md p-3 flex flex-col px-6 justify-between">
        <div>
          <p className=" font-lato text-slate-700 ">Add task</p>
          <p className="text-slate-600 text-sm  mt-3">Title</p>
          <input
            className="w-full bg-white h-10 border-[2px] border-gray-500 pl-3"
            value={'2'}
          ></input>
        </div>
        <div className="flex justify-between">
          <Button
            text="Add task"
            bgcolor="bg-blue-600"
            textcolor="text-white"
          />
          <Button
            text="Cancel"
            bgcolor="bg-slate-200"
            textcolor="text-gray-400"
          />
        </div>
      </div>
    </>
  );
}
