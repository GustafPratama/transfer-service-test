import React, { useState } from "react";
import { useApi } from "../hooks/useApi";
import { usePoll } from "../hooks/usePoll";

export default function TransferDetail() {
  const api = useApi();
  const [transfer, setTransfer] = useState<any>(null);
  const [id, setId] = useState("");
  const [role] = useState("supervisor");

  const load = async () => {
    if (id) {
      const res = await api.get(`/transfers/${id}`);
      setTransfer(res.data);
    }
  };

  usePoll(load, 5000);

  return (
    <div>
      <h2>Transfer Detail</h2>
      <input placeholder="Transfer ID" onChange={e => setId(e.target.value)} />
      {transfer && (
        <div>
          <p>Status: {transfer.status}</p>
          {role === "supervisor" && (
            <>
              <button onClick={() => api.post(`/transfers/${id}/accept`)}>Accept</button>
              <button onClick={() => api.post(`/transfers/${id}/complete`)}>Complete</button>
            </>
          )}
        </div>
      )}
    </div>
  );
}
