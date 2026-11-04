import LoginForm from "./LoginForm.jsx";

export default function App() {
  const handleSubmit = (data) => {
    console.log("Отправлено:", data);
  };

  return <LoginForm onSubmit={handleSubmit} />;
}
