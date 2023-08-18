import { Telegraf } from 'telegraf';
import { addTodo, deleteTodo, switchTodo, updateTodo } from './handlers.js';

const bot = new Telegraf('5765769733:AAF7jPi4pdv9jccd3utixGZI00V-c4xinEQ');

let status = '';
let statusId = 0;

bot.start((ctx) => {
  if (ctx.from.id != 533108461) {
    return;
  }
  ctx.reply('/get - отримати список todo \n/link - ссилка на сайт');
});

bot.command('link', (ctx) => {
  if (ctx.from.id != 533108461) {
    return;
  }
  ctx.reply(
    'Запусти сайт в /mine/fullstackgo, пак перейди на http://localhost:3000/'
  );
});

bot.command('get', async (ctx) => {
  if (ctx.from.id != 533108461) {
    return;
  }
  console.log("it's here");
  const responce = await fetch('http://backend:3001/api/gettodos');
  const todos = await responce.json();

  console.log(todos);

  let parsedStr = [];

  todos.forEach((todo) => {
    parsedStr.push(
      `[${todo.Id}] - ` + todo.Text + ' ' + (todo.Status ? '✅' : '❌')
    );
  });

  ctx.reply(parsedStr.join('\n'));
});

bot.command('add', (ctx) => {
  status = 'add';
  ctx.reply('Введи текст нового TODO');
});

bot.command('delete', (ctx) => {
  status = 'delete';
  ctx.reply('Введи ID TODO для видалення');
});

bot.command('switch', (ctx) => {
  status = 'switch';
  ctx.reply('Введи ID TODO для зміни статусу');
});

bot.command('update', (ctx) => {
  status = 'update';
  ctx.reply('Введи ID TODO для зміни');
});

bot.on('text', (ctx) => {
  switch (status) {
    case 'add':
      addTodo(ctx);
      status = '';
      break;
    case 'delete':
      deleteTodo(ctx);
      status = '';
      break;
    case 'switch':
      switchTodo(ctx);
      status = '';
      break;
    case 'update':
      let idStr = ctx.message.text;
      let id = Number(idStr);
      if (isNaN(id) || id == 0) {
        ctx.reply('Вводи тільки числа після /delete');
        return;
      }

      statusId = id;
      status = 'update2';
      ctx.reply('Тепер введи текст для зміни');
      break;
    case 'update2':
      updateTodo(ctx, statusId);
      break;
    default:
      return;
  }
});

console.log('Bot has been started');
bot.launch();
