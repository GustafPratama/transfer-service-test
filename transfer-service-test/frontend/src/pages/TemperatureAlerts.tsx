import React, { useEffect, useState } from "react";
import alerts from "../mock/alerts.json";

export default function TemperatureAlerts() {
  const [data, setData] = useState<any[]>([]);
  useEffect(() => setData(alerts), []);
  return (
    <div>
      <h2>Temperature Alerts</h2>
      <ul>
        {data.map(a => (
          <li key={a.id}>
            {a.location}: {a.temp}°C{" "}
            {a.temp < -20 || a.temp > -16 ? <span style={{ color: "red" }}>Abnormal</span> : <span>Normal</span>}
          </li>
        ))}
      </ul>
    </div>
  );
}
