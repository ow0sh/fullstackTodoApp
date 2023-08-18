export const addTodo = async (ctx) => {
  const responce = await fetch('http://backend:3001/api/getlastid');
  const lastId = await responce.json();

  let todo = {
    Text: '',
    Status: false,
    ID: lastId + 1,
  };

  todo.Text = ctx.message.text;

  fetch('http://backend:3001/api/inserttodo', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(todo),
  });

  ctx.reply('TODO успішно збережений');
};

export const deleteTodo = (ctx) => {
  let idStr = ctx.message.text;
  let id = Number(idStr);
  if (isNaN(id) || id == 0) {
    ctx.reply('Вводи тільки числа після /delete');
    return;
  }

  fetch('http://backend:3001/api/deletetodo', {
    method: 'DELETE',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(id),
  }).then(() => ctx.reply('TODO успішно видалений'));
};

export const switchTodo = (ctx) => {
  let idStr = ctx.message.text;
  let id = Number(idStr);
  if (isNaN(id) || id == 0) {
    ctx.reply('Вводи тільки числа після /complete');
    return;
  }

  fetch('http://backend:3001/api/switchstatus', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(id),
  }).then(() => ctx.reply('TODO успішно switched'));
};

export const updateTodo = (ctx, id) => {
  const request = { Text: ctx.message.text, ID: id };

  fetch('http://backend:3001/api/updatetext', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(request),
  }).then(() => ctx.reply('TODO updated'));
};
