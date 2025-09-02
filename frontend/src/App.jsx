import { useState, useEffect } from "react";

const env =
  import.meta.env.VITE_BACKEND_URL?.includes("app-qa")
    ? "QA"
    : import.meta.env.VITE_BACKEND_URL?.includes("app-prod")
    ? "PROD"
    : "";

function App() {
  const [name, setName] = useState("");
  const [users, setUsers] = useState([]);
  const [msg, setMsg] = useState("");

  useEffect(() => {
    document.title = env ? `Front ${env}` : "Front";
  }, []);

  const loadUsers = async () => {
    try {
      const res = await fetch("/api/users");
      const data = await res.json();
      setUsers(data);
    } catch (err) {
      setMsg("Error cargando usuarios: " + err.message);
    }
  };

  const addUser = async (e) => {
    e.preventDefault();
    setMsg("");
    try {
      const res = await fetch("/api/users", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ name }),
      });
      if (!res.ok) throw new Error("Error al insertar");
      setName("");
      setMsg("✔ Usuario insertado");
      loadUsers();
    } catch (err) {
      setMsg("✖ " + err.message);
    }
  };

  useEffect(() => {
    loadUsers();
  }, []);

  return (
    <div style={{ maxWidth: 600, margin: "40px auto", fontFamily: "sans-serif" }}>
      <h1>{env}</h1>
      <h1>Usuarios</h1>
      <form onSubmit={addUser} style={{ display: "flex", gap: "8px", marginBottom: "16px" }}>
        <input
          value={name}
          onChange={(e) => setName(e.target.value)}
          placeholder="Nombre..."
          required
          style={{ flex: 1, padding: "8px" }}
        />
        <button type="submit">Agregar</button>
      </form>

      {msg && <div>{msg}</div>}

      <h2>Lista</h2>
      <ul>
        {users.map((u) => (
          <li key={u.id}>
            <strong>#{u.id}</strong> {u.name}
          </li>
        ))}
      </ul>
    </div>
  );
}

export default App;
