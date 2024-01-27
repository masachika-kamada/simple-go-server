import React, { useState } from "react";
import "./App.css";

function App() {
  const [inputMessage, setInputMessage] = useState("");
  const [responseMessage, setResponseMessage] = useState("");

  const handleSubmit = async (event) => {
    event.preventDefault();
    const response = await fetch("http://localhost:8080/", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ message: inputMessage }),
    });
    const responseData = await response.text();
    setResponseMessage(responseData);
  };

  return (
    <div className="App">
      <header className="App-header">
        <p>Enter a message, and we'll convert it to uppercase</p>
        <form onSubmit={handleSubmit}>
          <input
            type="text"
            value={inputMessage}
            onChange={(e) => setInputMessage(e.target.value)}
            style={{ fontSize: "30px", margin: "10px" }}
          />
          <button type="submit" style={{ fontSize: "30px" }}>Submit</button>
        </form>
        <p style={{ textAlign: "left" }}>
          <strong>Sent message:</strong> {inputMessage}<br />
          <strong>Response from Go server:</strong> {responseMessage}
        </p>
      </header>
    </div>
  );
}

export default App;
