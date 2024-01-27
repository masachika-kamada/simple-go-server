import React, { useState } from "react";
import "./App.css";

function App() {
  const [message, setMessage] = useState("");
  const [response, setResponse] = useState("");

  const handleSubmit = async (event) => {
    event.preventDefault();
    const response = await fetch("http://localhost:8080/contact", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ message }),
    });
    const responseData = await response.text();
    setResponse(responseData);
  };

  return (
    <div className="App">
      <header className="App-header">
        <form onSubmit={handleSubmit}>
          <input
            type="text"
            value={message}
            onChange={(e) => setMessage(e.target.value)}
          />
          <button type="submit">Send</button>
        </form>
        <p>Response from Go server: {response}</p>
      </header>
    </div>
  );
}

export default App;
