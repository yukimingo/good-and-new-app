import type { ReactNode } from "react";
import "./layout.css";

type Props = {
  children: ReactNode;
};

const Layout = ({ children }: Props) => {
  return (
    <div className="app">
      <header className="app-header">
        <h1>Good And New</h1>
        <nav>
          <a href="/">Home</a>
          <a href="/user">Users</a>
          <a href="/category">Categories</a>
        </nav>
      </header>

      <main className="app-main">{children}</main>

      <footer className="app-footer">
        <p>© {new Date().getFullYear()} Good And New</p>
      </footer>
    </div>
  );
};

export default Layout;
