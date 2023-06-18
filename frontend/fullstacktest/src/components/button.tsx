interface ButtonParams {
  text: string;
  bgcolor: string;
  textcolor: string;
  onClick?: (params: any) => any;
}

export default function Button({
  text,
  bgcolor,
  textcolor,
  onClick,
}: ButtonParams) {
  return (
    <div
      onClick={() => onclick}
      className={
        bgcolor +
        ' ' +
        textcolor +
        ' w-24 h-8 rounded-md flex justify-center font-lato cursor-pointer select-none hover:bg-blue-500'
      }
    >
      <span className="my-auto">{text}</span>
    </div>
  );
}
