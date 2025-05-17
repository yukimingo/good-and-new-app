import React, { useEffect, useState } from "react";
import { fetchUsers } from "../api/user";

type User = {
  id: number;
  name: string;
  email: string;
};

function UserList() {
  const [users, setUsers] = useState<User[]>([]); // データ保存用
  const [error, setError] = useState(null); // エラー表示用
  const [loading, setLoading] = useState(true); // ローディング中

  useEffect(() => {
    fetchUsers()
      .then((data) => {
        setUsers(data); // 取得したデータを保存
        setLoading(false);
      })
      .catch((err) => {
        setError(err.message);
        setLoading(false);
      });
  }, []); // 初回のみ実行

  if (loading) return <div>読み込み中...</div>;
  if (error) return <div>エラー: {error}</div>;

  return (
    <div>
      <h1>ユーザー一覧</h1>
      {/* <div>{users}</div> */}
      <ul>
        {users.map((user) => (
          <li key={user.id}>
            {user.name} {user.email}
          </li>
        ))}
      </ul>
    </div>
  );
}

export default UserList;
