import { useEffect, useState } from "react";
import { fetchUsers } from "../api/user";
import "./user.css";

type User = {
  id: number;
  name: string;
  email: string;
};

function UserList() {
  const [users, setUsers] = useState<User[]>([]);
  const [error, setError] = useState(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetchUsers()
      .then((data) => {
        setUsers(data);
        setLoading(false);
      })
      .catch((err) => {
        setError(err.message);
        setLoading(false);
      });
  }, []);

  if (loading) return <div>読み込み中...</div>;
  if (error) return <div>エラー: {error}</div>;

  return (
    <div className="user-list-container">
      <h2 className="user-list-title">ユーザーリスト</h2>
      <div className="user-grid">
        {users.map((user) => (
          <div className="user-card" key={user.id}>
            <div className="user-info">
              <p>-----------------------------------</p>
              <h3>{user.name}</h3>
              <p>{user.email}</p>
              <p>-----------------------------------</p>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}

export default UserList;
