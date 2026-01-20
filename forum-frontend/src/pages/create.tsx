import { useState, ChangeEvent } from "react";
import { useNavigate } from "react-router-dom";

export default function Create() {
  const [username, setUsername] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [message, setMessage] = useState("");
  const nav = useNavigate();

  const handleCreate = async (e: React.FormEvent) => {
      e.preventDefault();

      try {
        const res = await fetch("http://localhost:8080/users", {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify({ username, password, email }),
        });

        const data = await res.json();

        if (res.ok) {
          setMessage("Account created!");
          nav("/")
        } else {
          setMessage("Failed: " + (data.error || res.statusText));
        }
      } catch (err) {
        setMessage("Error connecting to backend");
      }
  };

  return (
      <div style={{ maxWidth: 400, margin: "50px auto" }}>
      <h2>Create Account</h2>
        <label>Username</label>
        <input
          type="username"
          value={username}
          onChange={(e: ChangeEvent<HTMLInputElement>) => setUsername(e.target.value)}
          required
          style={{ width: "100%", marginBottom: 10 }}
        />

        <label>Email</label>
        <input 
          type="email"
          value={email}
          onChange={(e: ChangeEvent<HTMLInputElement>) => setEmail(e.target.value)}
          required
          style={{ width: "100%", marginBottom: 10 }}
        />

        <label>Password</label>
        <input
          type="password"
          value={password}
          onChange={(e: ChangeEvent<HTMLInputElement>) => setPassword(e.target.value)}
          required
          style={{ width: "100%", marginBottom: 10 }}
        />

      <button onClick={handleCreate}>Create</button>
      <button onClick={() => nav("/")}>Already have an account? Login</button>
        

      {message && (
        <div style={{ marginTop: 20, padding: 10, border: "1px solid gray" }}>
          {message}
        </div>
      )}
    </div>
  )

  }