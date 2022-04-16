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

  fetch("/rooms", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(value),
  })
    .then((res) => res.json())
    .then((res) => redirect(`/room/${res.id}`));
}

function redirect(route) {
  window.location.href = route;
}
