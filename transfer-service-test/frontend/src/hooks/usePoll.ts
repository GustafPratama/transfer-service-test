import { useEffect } from "react";
export const usePoll = (fn: () => void, interval: number) => {
  useEffect(() => {
    fn();
    return () => clearInterval(id);
  }, [fn, interval]);
};

cat > transfer-service-test/frontend/src/mock/alerts.json <<'EOF'
[
  {"id": 3, "location": "Freezer-3", "temp": -15}
]
