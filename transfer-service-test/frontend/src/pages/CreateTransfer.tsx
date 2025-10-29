import { useApi } from "../hooks/useApi";

export default function CreateTransfer() {
  const api = useApi();
  const [error, setError] = useState("");

  const submit = async () => {
    try {
      await api.post("/transfers", form, { headers: { "Idempotency-Key": Date.now().toString() } });
      alert("Transfer created!");
    } catch (err: any) {
      setError(err.response?.data || "Error");
    }
  };

  return (
    <div>
      <h2>Create Transfer</h2>
      <input placeholder="Pallet ID" onChange={e => setForm({ ...form, pallet_id: e.target.value })} />
      <input placeholder="From Location" onChange={e => setForm({ ...form, from_location: e.target.value })} />
      <input placeholder="To Location" onChange={e => setForm({ ...form, to_location: e.target.value })} />
      <input placeholder="Note" onChange={e => setForm({ ...form, note: e.target.value })} />
      <button onClick={submit}>Create</button>
      {error && <p style={{ color: "red" }}>{error}</p>}
    </div>
  );
}
