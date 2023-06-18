interface Checkbox {
  checked: boolean;
}

export default function Checkbox({ checked }: Checkbox) {
  if (checked) {
    return (
      <div className="w-[20px] h-[20px] bg-blue-600 hover:bg-blue-400">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          strokeWidth={1.5}
          stroke="currentColor"
          className="w-5 h-5 text-white"
        >
          <path
            strokeLinecap="round"
            strokeLinejoin="round"
            d="M4.5 12.75l6 6 9-13.5"
          />
        </svg>
      </div>
    );
  } else {
    return (
      <div className="h-[20px] w-[20px] bg-slate-200 hover:bg-slate-400"></div>
    );
  }
}
