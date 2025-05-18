// import React from "react";
// import type { ReactNode } from "react";
// import "./App.css";

// type Props = {
//   children: ReactNode;
// };

// const Layout: React.FC<Props> = ({ children }) => {
//   return (
//     <div className="app">
//       <header className="app-header">
//         <h1>My Stylish App</h1>
//         <nav>
//           <a href="#">Home</a>
//           <a href="#">Users</a>
//           <a href="#">About</a>
//         </nav>
//       </header>

//       <main className="app-main">{children}</main>

//       <footer className="app-footer">
//         <p>© {new Date().getFullYear()} My Stylish App. All rights reserved.</p>
//       </footer>
//     </div>
//   );
// };

// export default Layout;

import { Routes, Route } from "react-router-dom";
import Layout from "./layout";
import UserList from "./pages/user";

function App() {
  return (
    <Layout>
      <Routes>
        <Route path="/" element={<UserList />} />
        <Route path="/user" element={<UserList />} />
      </Routes>
    </Layout>
  );
}

export default App;
