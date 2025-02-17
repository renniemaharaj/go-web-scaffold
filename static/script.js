// Initiate WebSocket connection
const socket = new WebSocket("ws://localhost:8080/ws");

socket.onopen = function () {
  socket.onmessage = function (event) {
    const [command, path, content] = event.data.split(":");
    console.log(`Received message: ${command} ${path}`);
    if (command === "reload") {
      console.log("Reloading page...");
      window.location.reload(true); // true forces reload from server, bypassing cache
    }
  };
  socket.onclose = function (event) {
    if (event.wasClean) {
      console.log(
        `[close] Connection closed cleanly, code=${event.code} reason=${event.reason}`
      );
    } else {
      console.error("[close] Connection died");
    }
  };
};

// DOM manipulation
document.addEventListener("DOMContentLoaded", function () {
  const title = document.getElementById("title");
  const changeTextBtn = document.getElementById("changeTextBtn");

  changeTextBtn.addEventListener("click", function () {
    title.textContent = "Golang Rocks!";
  });
});
