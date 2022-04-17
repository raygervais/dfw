function formatJson(event) {
  event.preventDefault();
  const data = new FormData(event.target);

  const value = Object.fromEntries(data.entries());

  fetch("/rooms", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(value),
  }).then((res) => {
    console.log("Request complete! response:", res);
  });
}

function formatJsonForPromptSubmission(event) {
  event.preventDefault();
  const data = new FormData(event.target);

  value = Object.fromEntries(data.entries());

  value.prompt = {
    text: value.prompt,
    active: true,
  };

  value.timeToLive = 60;

  fetch("/room", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(value),
  })
    .then((res) => res.json())
    .then((res) => redirect(`/room/${res.id}`))
    .catch((err) => console.log(err));
}

function redirect(route) {
  window.location.href = route;
}

function formatJsonForCommentSubmission(event, id) {
  event.preventDefault();
  const data = new FormData(event.target);

  value = Object.fromEntries(data.entries());

  fetch(`/room/${id}/comment`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(value),
  })
    .then((res) => res.json())
    .catch((err) => console.log(err));
}

function getRoomIdForRedirect(event) {
  event.preventDefault();
  const data = new FormData(event.target);

  value = Object.fromEntries(data.entries());
  console.log(value)
  redirect(`room/${value.roomId}`)
}

