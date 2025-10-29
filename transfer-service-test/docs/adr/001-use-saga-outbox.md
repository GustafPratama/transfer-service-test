# ADR 001: Use Outbox Pattern for Reliable Event Delivery

**Context:** Transfer completion should trigger an event that may be consumed by another system.  
To ensure reliability even if the broker is down, events are stored in an **outbox** table first.

**Decision:** Implement Outbox Pattern with periodic `/dev/flush-outbox` endpoint.

**Consequences:**  
- Event delivery becomes reliable.  
- Slight complexity added in database layer.
