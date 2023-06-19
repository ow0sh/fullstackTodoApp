interface ButtonParams {
  text: string;
  bgcolor: string;
  textcolor: string;
  hovercolor: string;
  onClick?: (params: any) => any;
}

export default function Button({
  text,
  bgcolor,
  textcolor,
  hovercolor,
  onClick,
}: ButtonParams) {
  return (
    <div
      onClick={onClick}
      className={
        bgcolor +
        ' ' +
        textcolor +
        ' w-24 h-8 rounded-md flex justify-center font-lato cursor-pointer select-none ' +
        hovercolor
      }
    >
      <span className="my-auto">{text}</span>
    </div>
  );
}
